package wildcard

import (
	"os"
	"fmt"
	"regexp"
	"gmv/option"
)

type PathElement struct {
	charType CharType
	content string
	match string
	referenceNumbers []int
}
func (this PathElement) GetMatch() (string, error) {
	return this.match, nil
}
func (this PathElement) GetMatchReferenceNumber(n int) string {
	for _, i := range this.referenceNumbers {
		if n == i {
			return this.match
		}
	}
	return ""
}

//enum
type CharType int
const (
	_ CharType = iota
	Literal
	Star
	Question
	ClassBegin
	ClassEnd
	ReferenceBegin
	ReferenceEnd
)
// Char は、第一段階でsrc文字列の解析結果を格納するstructです。
type Char struct {
	charType CharType
	char string
}
// parseChars は、src文字列の第一段階の解析です。各文字を分類しエスケープ文字を処理します。
func parseChars(options option.Option, src string) []Char {
	chars := []Char{}
	strs := src
	inClass := false
	for len(strs) != 0 {
		n := 1
		s := string([]rune(strs)[0:1])
		switch s {
		case "?":
			ct := Question
			if inClass {
				ct = Literal
			}
			chars = append(chars, Char{charType: ct, char: s})
			break
		case "*":
			ct := Star
			if inClass {
				ct = Literal
			}
			chars = append(chars, Char{charType: ct, char: s})
			break
		case "[":
			chars = append(chars, Char{charType: ClassBegin, char: s})
			inClass = true
			break
		case "]":
			inClass = false
			chars = append(chars, Char{charType: ClassEnd, char: s})
			break
		case "\\":
			s = string([]rune(strs)[1:2]) 
			chars = append(chars, Char{charType: Literal, char: s})
			n += 1
			break
		case "(":
			chars = append(chars, Char{charType: ReferenceBegin, char: s})
			break
		case ")":
			chars = append(chars, Char{charType: ReferenceEnd, char: s})
			break
		default:
			chars = append(chars, Char{charType: Literal, char: s})
		}
		strs = string([]rune(strs)[n:])
	}
	return chars
}
func referenceNumbering(options option.Option, elements []PathElement) []PathElement {
	if ! *options.Opt_w && ! *options.Opt_W {
		return elements
	}
	nextReferenceNumber := 1
	newElements := []PathElement{}
	classReferenceNumber := 0
	for _, e := range elements {
		switch e.charType {
		case Star, Question:
			if classReferenceNumber > 0 {
				//クラス指定中ならクラスの参照番号を指定
				e.referenceNumbers = []int{classReferenceNumber}
			} else {
				e.referenceNumbers = []int{nextReferenceNumber}
				nextReferenceNumber += 1
			}
			break
		case ClassBegin:
			classReferenceNumber = nextReferenceNumber
			nextReferenceNumber += 1
			break
		case ClassEnd:
			classReferenceNumber = 0
			break
		default:
			if classReferenceNumber > 0 {
				//クラス指定中ならクラスの参照番号を指定
				e.referenceNumbers = []int{classReferenceNumber}
			} else {
				e.referenceNumbers = []int{}
			}
			break
		}
		newElements = append(newElements, e)
	}
	fmt.Println(newElements)
	return newElements
}
func Parse(options option.Option, src string) ([]PathElement, error) {

	elements := []PathElement{}

	chars := parseChars(options, src)

	nextReferenceNumber := 1
	referenceLevels := []int{}

	for len(chars) != 0 {
		c := chars[0]
		buf := c.char

		CharSwich: switch c.charType {
		case ReferenceBegin:
			referenceLevels = append(referenceLevels, nextReferenceNumber)
			nextReferenceNumber += 1
			break
		case ReferenceEnd:
			referenceLevels = referenceLevels[0:len(referenceLevels) - 1]
			break
		case Literal, Star:
			if len(chars) > 1 {
				for _, next := range chars[1:] {
					if next.charType != c.charType {
						elements = append(elements, PathElement{charType: c.charType, content: buf, referenceNumbers: referenceLevels})
						break CharSwich
					}
					//次の要素も同じcharTypeなら続けて取得する。
					buf += next.char
				}
				elements = append(elements, PathElement{charType: c.charType, content: buf, referenceNumbers: referenceLevels})
			} else { // len(chars) == 1
				elements = append(elements, PathElement{charType: c.charType, content: buf, referenceNumbers: referenceLevels})
			}
			break
		default:
			elements = append(elements, PathElement{charType: c.charType, content: buf, referenceNumbers: referenceLevels})
		}
		chars = chars[len(buf):]
	}
	if len(referenceLevels) != 0 {
		return nil, fmt.Errorf("unmatch reference braces.")
	}
	//Option w W のために、参照の設定を行なう
	elements = referenceNumbering(options, elements)

	return elements, nil
}

func GetSearchPath(elements []PathElement) string {
	path := ""
	for _, e := range elements {
		switch e.charType {
		case ReferenceBegin, ReferenceEnd:
			break
		default:
			path += e.content
		}
	}
	return path
}

func findReferences(elements []PathElement) map[int]string {

	referenceSet := make(map[int]string)
	for _, e := range elements {
		for _, r := range e.referenceNumbers {
			referenceSet[r] = ""
		}
	}
	for k, _ := range referenceSet {
		inClass := false
		s := ""
		for _, e := range elements {
			switch e.charType {
			case ClassBegin:
				inClass = true
				break
			case ClassEnd:
				inClass = false
				break
			default:
				if !inClass {
					s += e.GetMatchReferenceNumber(k)
				}
			}
		}
		referenceSet[k] = s
	}
	return referenceSet
}


func GetDestPath(options option.Option, elements []PathElement, realPath, dest string) (string, error) {
	regexString := "^"
	for _, e := range elements {
		switch e.charType {
		case Literal:
			regexString += e.content
			break
		case Question, Star:
			regexString += "(.*)"
			break
		case ClassBegin:
			regexString += "(" + e.content
			break
		case ClassEnd:
			regexString += e.content + ")"
			break
		}

	}
	regexString += "$"
	re := regexp.MustCompile(regexString)
	matched := re.FindStringSubmatch(realPath)
	if len(matched) == 0 {
		fmt.Fprintln(os.Stderr, "regexp failed. " + regexString)
		return "", fmt.Errorf("regexp failed.")
	}

	//マッチした部分を対応するPathElementに保存する。
	referenceNumber := 1
	resultElements := []PathElement{}
	for _, e := range elements {
		switch e.charType {
		case Literal:
			e.match = e.content
			break
		case Question, Star:
			e.match = matched[referenceNumber]
			referenceNumber += 1
			break
		case ClassBegin:
			e.match = matched[referenceNumber]
			referenceNumber += 1
		}
		resultElements = append(resultElements, e)
	}
	fmt.Println(resultElements)

	// dest を生成する。
	fmt.Println(realPath)
	if !*options.Opt_W {
		//dest の * ** ? を参照の$i に変換する。
	}
	references := findReferences(resultElements)
	for k, v := range references {
		re := regexp.MustCompile(fmt.Sprintf("\\$%d", k))
		dest = re.ReplaceAllLiteralString(dest, v)
	}

	return dest, nil
}

