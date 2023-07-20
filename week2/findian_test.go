package main

import "testing"

func TestStartswithCharI(t *testing.T) {
	testString1 := "Indianapolis"
	wanted1 := true
	testString2 := "iris"
	wanted2 := true
	testString3 := "tesla"
	wanted3 := false

	if startswithCharI(testString1) != wanted1 {
		t.Errorf("startswithCharI(%s) should be %t", testString1, wanted1)
	}
	if startswithCharI(testString2) != wanted2 {
		t.Errorf("startswithCharI(%s) should be %t", testString2, wanted2)
	}
	if startswithCharI(testString3) != wanted3 {
		t.Errorf("startswithCharI(%s) should be %t", testString3, wanted3)
	}
}

func TestEndswithCharN(t *testing.T) {
	testString1 := "Clinton"
	wanted1 := true
	testString2 := "visit BOSTON"
	wanted2 := true
	testString3 := "bill123"
	wanted3 := false

	if endswithCharN(testString1) != wanted1 {
		t.Errorf("endswithCharN(%s) should be %t", testString1, wanted1)
	}
	if endswithCharN(testString2) != wanted2 {
		t.Errorf("endswithCharN(%s) should be %t", testString2, wanted2)
	}
	if endswithCharN(testString3) != wanted3 {
		t.Errorf("endswithCharN(%s) should be %t", testString3, wanted3)
	}
}

func TestContainsCharA(t *testing.T) {
	testString1 := "America"
	wanted1 := true
	testString2 := "What a string "
	wanted2 := true
	testString3 := "$ _ $"
	wanted3 := false

	if containsCharA(testString1) != wanted1 {
		t.Errorf("containsCharA(%s) should be %t", testString1, wanted1)
	}
	if containsCharA(testString2) != wanted2 {
		t.Errorf("containsCharA(%s) should be %t", testString2, wanted2)
	}
	if containsCharA(testString3) != wanted3 {
		t.Errorf("containsCharA(%s) should be %t", testString3, wanted3)
	}
}

func TestFindianPass(t *testing.T) {
	tests := map[string]struct {
		given    string
		expected string
	}{
		"test1": {
			given:    "ian",
			expected: "Found!"},
		"test2": {
			given:    "Ian",
			expected: "Found!"},
		"test3": {
			given:    "iuiygaygn",
			expected: "Found!"},
		"test4": {
			given:    "I d skd a efju N",
			expected: "Found!"},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			if findian(test.given) != test.expected {
				t.Errorf("findian(%s) should be %s", test.given, test.expected)
			}
		})
	}
}

func TestFindianFail(t *testing.T) {
	tests := map[string]struct {
		given    string
		expected string
	}{
		"test1": {
			given:    "ihhhhhn",
			expected: "Not Found!"},
		"test2": {
			given:    "ina",
			expected: "Not Found!"},
		"test3": {
			given:    "xian",
			expected: "Not Found!"},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			if findian(test.given) != test.expected {
				t.Errorf("findian(%s) should be %s", test.given, test.expected)
			}
		})
	}
}
