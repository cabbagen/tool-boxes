package utils

import (
	"regexp"
	"testing"
)

func Test_String_Value(t *testing.T) {
	if NewUstring("hello").Value() != "hello" {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_Length(t *testing.T) {
	if NewUstring("你好").Length() != 2 {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_Len(t *testing.T) {
	if NewUstring("你好").Len() != 6 {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_ToByte(t *testing.T) {
	greet := NewUstring("hello").ToByte()
	
	if string(greet) != "hello" {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_Concat(t *testing.T) {	
	rstring := NewUstring("hello").Concat(" ", "world", " !")
	
	if rstring.Value() != "hello world !" {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_Substring(t *testing.T) {
	rstring := NewUstring("hello").Substring(1, 3)
	
	if rstring.Value() != "el" {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_EndsWith(t *testing.T) {
	if !NewUstring("hello").EndsWith("lo", 0) {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_Includes(t *testing.T) {
	if NewUstring("hello").EndsWith("alo", 0) {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_PadEnd(t *testing.T) {
	if NewUstring("hello").PadEnd(10, "dacb").Value() != "hellodacbd" {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_PadStart(t *testing.T) {
	if NewUstring("hello").PadStart(10, "dacb").Value() != "dacbdhello" {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_Repeat(t *testing.T) {
	if NewUstring("hello").Repeat(2).Value() != "hellohello" {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_SReplace(t *testing.T) {
	if NewUstring("hello").SReplace("ll", "ab").Value() != "heabo" {
		t.Error("no, failed !")
		return
	}
	t.Log("yeah, passed !")
}

func Test_String_RReplace(t *testing.T) {
	reg := regexp.MustCompile(`l+`)
	
	newUNewUstring := NewUstring("helloll").RReplace(reg, func(value string) string {
		return "ss"
	})
	
	if newUNewUstring.Value() != "hessoss" {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_String_Search(t *testing.T) {
	reg := regexp.MustCompile(`\d+`)
	
	if NewUstring("ab765").Search(reg) != 2 {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_String_Split(t *testing.T) {
	stringSlice := NewUstring("a,b,c").Split(",")
	
	if stringSlice[0].Value() != "a" {
		t.Error("no, failed !")
		return
	}
	
	if stringSlice[1].Value() != "b" {
		t.Error("no, failed !")
		return
	}
	
	if stringSlice[2].Value() != "c" {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_String_StartsWill(t *testing.T) {
	if !NewUstring("hello").StartsWill("hell") {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_String_ToLowerCase(t *testing.T) {
	if NewUstring("HELLOE").ToLowerCase().Value() != "helloe" {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_String_ToUpperCase(t *testing.T) {
	if NewUstring("hello").ToUpperCase().Value() != "HELLO" {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_String_Trim(t *testing.T) {
	if NewUstring("  hello  ").Trim().Value() != "hello" {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_String_TrimRight(t *testing.T) {
	if NewUstring("hello  ").TrimRight().Value() != "hello" {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_String_TrimLeft(t *testing.T) {
	if NewUstring("  hello").TrimLeft().Value() != "hello" {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}



































