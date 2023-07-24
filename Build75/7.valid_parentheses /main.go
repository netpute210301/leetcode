package main

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '(') || (v == '{') || (v == '[') {
			stack = append(stack, v)
		} else {
			if len(stack) <= 0 {
				return false
			}
			if v == ')' {
				chat := stack[len(stack)-1]
				if chat == '(' {
					stack = stack[:len(stack)-1]
					continue
				}
				return false
			}
			if v == '}' {
				chat := stack[len(stack)-1]
				if chat == '{' {
					stack = stack[:len(stack)-1]
					continue
				}
				return false
			}
			if v == ']' {
				chat := stack[len(stack)-1]
				if chat == '[' {
					stack = stack[:len(stack)-1]
					continue
				}
				return false
			}
		}
	}
	return len(stack) == 0
	// 判斷邊緣 case 是不是存在
	// for s 被走訪完成
	// 如果是 ( 或 { 或 [ 則加入堆疊
	// 如果是 ) 或 } 或 [ 先判斷堆疊是否為零
	//	為零 return false
	//	不為零
	//	如果是 ） 的話，從堆疊，最後一個取出來，看看是不是( 如果是繼續比對直到否則return fasle
	//	如果是 } 的話，從堆疊，最後一個取出來，看看是不是{ 如果是繼續比對直到否則return fasle
	//如果是 ] 的話，從堆疊，最後一個取出來，看看是不是[ 如果是繼續比對直到否則return fasle
}
