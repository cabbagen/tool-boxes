package utils

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Ustring struct {
	ustring   string
}

func NewUstring(str string) Ustring {
	return Ustring {str}
}

func (us Ustring) Value() string {
	return us.ustring
}

func (us Ustring) Length() int {
	return utf8.RuneCountInString(us.ustring)
}

func (us Ustring) Len() int {
	return len(us.ustring)
}

func (us Ustring) ToByte() []byte {
	return []byte(us.ustring)
}

func (us Ustring) Concat(strs ...string) Ustring {
	rstring := us.ustring
	
	for _, str := range strs {
		rstring += str
	}
	return Ustring {rstring}
}

func (us Ustring) Substring(indexStart, indexEnd int) Ustring {
	rrune := []rune(us.ustring)[indexStart:indexEnd]
	
	return Ustring { string(rrune) }
}

func (us Ustring) EndsWith(searchString string, position int) bool {
	substr := us.Substring(position, us.Length())
	
	return strings.HasSuffix(substr.ustring, searchString)
}

func (us Ustring) Includes(searchString string, position int) bool {
	substr := us.Substring(position, us.Length())
	
	return strings.Contains(substr.ustring, searchString)
}

func (us Ustring) IndexOf(searchString string, formIndex int) int {
	substr := us.Substring(formIndex, us.Length())
	
	return strings.Index(substr.ustring, searchString)
}

func (us Ustring) LastIndexOf(searchString string, formIndex int) int {
	substr := us.Substring(formIndex, us.Length())
	
	return strings.LastIndex(substr.ustring, searchString)
}

func (us Ustring) Match(reg *regexp.Regexp) []string {
	return reg.FindStringSubmatch(us.ustring)
}

func (us Ustring) PadEnd(targetLength int, padString string) Ustring {
	padLength := targetLength - us.Length()
	
	if padLength < 0 {
		return us
	}
	
	padUString := Ustring {padString}
	
	groups, surplus := padLength / padUString.Length(), padLength % padUString.Length()
	
	for i := 0; i < groups; i++ {
		us = us.Concat(padString)
	}
	
	if surplus > 0 {
		us = us.Concat(padUString.Substring(0, surplus).ustring)
	}
	
	return us
}

func (us Ustring) PadStart(targetLength int, padString string) Ustring {
	padLength := targetLength - us.Length()
	
	if padLength < 0 {
		return us
	}
	
	padUString := Ustring {padString}
	
	groups, surplus := padLength / padUString.Length(), padLength % padUString.Length()
	
	prefix := Ustring {""}
	
	for i := 0; i < groups; i++ {
		prefix = prefix.Concat(padString)
	}
	
	if surplus > 0 {
		prefix = prefix.Concat(padUString.Substring(0, surplus).ustring)
	}
	
	return prefix.Concat(us.ustring)
}

func (us Ustring) Repeat(count int) Ustring {
	return Ustring { strings.Repeat(us.ustring, count) }
}

func (us Ustring) SReplace(substr, newSubstr string) Ustring {
	return Ustring { strings.Replace(us.ustring, substr, newSubstr, -1) }
}

func (us Ustring) RReplace(reg *regexp.Regexp, f func(s string) string) Ustring {
	matchAllStrings := reg.FindAllString(us.ustring, -1)
	
	for _, value := range matchAllStrings {
		us = us.SReplace(value, f(value))
	}
	
	return us
}

func (us Ustring) Search(reg *regexp.Regexp) int {
	loc := reg.FindStringIndex(us.ustring)
	
	if loc == nil {
		return -1
	}
	
	return loc[0]
}


func (us Ustring) Split(separator string) []Ustring {
	ustrings := []Ustring{}
	splitedString := strings.Split(us.ustring, separator)
	
	for _, value := range splitedString {
		ustrings = append(ustrings, Ustring {value})
	}
	
	return ustrings
}

func (us Ustring) StartsWill(searchString string) bool {
	return strings.HasPrefix(us.ustring, searchString)
}

func (us Ustring) ToLowerCase() Ustring {
	return Ustring { strings.ToLower(us.ustring) }
}

func (us Ustring) ToUpperCase() Ustring {
	return Ustring { strings.ToUpper(us.ustring) }
}

func (us Ustring) Trim() Ustring {
	return Ustring { strings.TrimSpace(us.ustring) }
}

func (us Ustring) TrimRight() Ustring {
	return Ustring { strings.TrimRightFunc(us.ustring, unicode.IsSpace) }
}

func (us Ustring) TrimLeft() Ustring {
	return Ustring { strings.TrimLeftFunc(us.ustring, unicode.IsSpace) }
}









