package linter

import (
	"fmt"
	"testing"
)

func TestEOFLint_TextWithSingleNewLine(t *testing.T) {
	got, fix := NewEndOfFileRule(autofixTestConf).Lint([]byte(textWithSingleNewLine))

	if fix != nil {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(textWithSingleNewLine):\n\tExpected nil, got:\n%v", string(fix))
	}

	if got != true {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(textWithSingleNewLine):\n\tExpected %v, got %v", true, got)
	}
}

func TestEOFLint_ShortTextWithSingleNewLine(t *testing.T) {
	got, fix := NewEndOfFileRule(autofixTestConf).Lint([]byte(shortTextWithSingleNewLine))

	if fix != nil {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(shortTextWithSingleNewLine):\n\tExpected nil, got:\n%v", string(fix))
	}

	if got != true {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(shortTextWithSingleNewLine):\n\tExpected %v, got %v", true, got)
	}
}

func TestEOFLint_TextWithTwoNewLines(t *testing.T) {
	got, fixed := NewEndOfFileRule(autofixTestConf).Lint([]byte(textWithTwoNewLines))

	if got != false {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(textWithTwoNewLines):\n\tExpected %v, got %v", false, got)
	}

	if string(fixed) != string(textWithSingleNewLine) {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(textWithTwoNewLines): autofix did not work\n\tExpected:\n%q\n\tGot:\n%q", textWithSingleNewLine, string(fixed))
	}
}

func TestEOFLint_TextWithoutNewLine(t *testing.T) {
	got, fixed := NewEndOfFileRule(autofixTestConf).Lint([]byte(textWithoutNewLine))

	if string(fixed) != string(textWithSingleNewLine) {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(textWithoutNewLine): autofix did not work\n\tExpected:\n%q\n\tGot:\n%q", textWithSingleNewLine, string(fixed))
	}

	if got != false {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(textWithoutNewLine):\n\tExpected %v, got %v", false, got)
	}
}

func TestEOFLint_EmptyString(t *testing.T) {

	// empty files are valid
	got, _ := NewEndOfFileRule(autofixTestConf).Lint([]byte(""))

	if got != true {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(emptyFileText):\n\tExpected %v, got %v", true, got)
	}
}

func TestEOFLint_StringWithOneNewline(t *testing.T) {
	// files with a single newline char are also valid
	got, _ := NewEndOfFileRule(autofixTestConf).Lint([]byte(fmt.Sprintf("\n")))

	if got != true {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(emptyFileText):\n\tExpected %v, got %v", true, got)
	}
}

func TestEOFLint_StringWithTwoNewlines(t *testing.T) {
	// files with a two newlines should be reduced to one newline if singleNewLineRule is set
	got, fixed := NewEndOfFileRule(autofixTestConf).Lint([]byte(fmt.Sprintf("\n\n")))

	if string(fixed) != string(fmt.Sprintf("\n")) {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(textWithoutNewLine): autofix did not work\n\tExpected:\n%q\n\tGot:\n%q", fmt.Sprintf("\n\n"), string(fixed))
	}

	if got != false {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(emptyFileText):\n\tExpected %v, got %v", false, got)
	}
}

func TestEOFLint_NotTextFile(t *testing.T) {
	// the 0xFFFD UTF-8 control character should be ignored, because the Lint method
	// does not check if the input is a valid Text file or not 'IsText' check fail
	got, _ := NewEndOfFileRule(autofixTestConf).Lint([]byte(string([]rune{0xFFFD, '????'})))

	if got != false {
		t.Errorf("NewEndOfFileRule(autofixTestConf).Lint(textNotText):\n\tExpected %v, got %v", false, got)
	}
}
