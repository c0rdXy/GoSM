package myUtils

const (
	E = "XPATH syntax error"
	B = "You are in"
)

var (
	KeyList = []string{}
	NumList = []int{}
)

func init() {
	func() {
		// 初始化字符与数字集合
		for i := 33; i <= 126; i++ {
			KeyList = append(KeyList, string(rune(i)))
			NumList = append(NumList, i)
		}
	}()
}
