package memo

//Func is the type of the function to memoize
type Func func(key string)(interface{}, error)

//A result is the result of calling a Func
type result struct{
	value interface{}
	err error
}

//A request is a message requesting that the Func be applied to key
//key调用Func型函数
type request struct{
	key string
	response chan<- result		//the client wants a single result
}

type Memo struct{requests chan request}

//New returns a memoizatoin of f.
//Clients must subsequently call Close.
func New(f Func) *Memo{
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func(memo *Memo)Get(key string)(interface{}, error){
	response := make(chan result)								//创建，结果通道
	memo.requests <- request{key, response}		//分装，请求。并发送
	res := <- response											//阻塞，直到消息
	return res.value, res.err
}

func(memo *Memo)Close(){
	close(memo.requests)
}

func(memo *Memo)server(f Func){
	cache := make(map[string]*entry)
	for req := range memo.requests{								//读取，请求
		e := cache[req.key]										//查询，请求
		if e == nil{											//查询，失败，处理......
			//this is the first request for this key.
			e = &entry{ready: make(chan struct{})}				//生成新条目entry
			cache[req.key] = e									//缓存新条目
			go e.call(f,req.key) 		//call f(key)			//条目调用call方法，传入函数f，key，生成结果
		}
		go e.deliver(req.response)								//条目调用deliver方法,向客户请求的结果通道（req.response）发送结果
	}
}

//结果条目，封装了1,结果值和2,结果值的状态。
type entry struct{
	res result
	ready chan struct{}					//closed when res is raady
}

func (e *entry)call(f Func, key string){
	e.res.value, e.res.err = f(key)		//Evaluate the function
	close(e.ready)						//Broadcast the ready condition
}

func (e *entry)deliver(response chan<- result){
	<-e.ready							//Wait for the ready condition
	response <- e.res					//Send the result to the client
}

