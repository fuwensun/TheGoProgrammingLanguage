package word

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func TestIsPalindrome(t *testing.T){
	var tests = []struct{
		input string
		want bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	//这种表格驱动的测试在Go语言中很常见。我们可以很容易地向表格添加新的测试数据，
	//并且后面的测试逻辑也没有冗余，这样我们可以有更多的精力地完善错误信息
	for _, test := range tests{
		if got := IsPalindrome(test.input); got != test.want{
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++{
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

//示例函数是以Example为函数名前缀的函数，提供一个由编译器保证正确性的示例文档
func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
}

func randomPalindrome(rng *rand.Rand)string{
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n + 1)/2; i++{
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n - 1 - i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T){
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++{
		p := randomPalindrome(rng)
		if !IsPalindrome(p){
			t.Errorf("IsPalindrome(%q) = false", p)
		}

	}
}
