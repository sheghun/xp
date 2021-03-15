package lib

import "fmt"

func FilterStrings(strArr []string) (ret []string) {
	for _, s := range strArr {
		if s != "" {
			ret = append(ret, s)
		}
	}
	return
}

func IsSorted(strArr []string) bool {
	sorted := false
	for i := 0; i < len(strArr); i++ {

		if i+1 == len(strArr) {
			break
		}

		if strArr[i] < strArr[i+1] || strArr[i] == strArr[i+1] {
			sorted = true
		} else {
			fmt.Println(strArr[i])
			fmt.Println(strArr[i+1])
			return false
		}
	}

	return sorted
}
