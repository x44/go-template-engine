# go-template-engine
Text template processing library

## Install
`
go get "github.com/x44/go-template-engine"
`

## Usage
`
import temple "github.com/x44/go-template-engine"
`

## Example
```go
in := []string{
	...
}
out, err := temple.New().
		SetInputStrings(in).
		// or
		// SetInputBytes(...).
		// SetInputString(...).
		// SetInputFile("filename.txt").
		SetOutputEndOfLine(eol.LF).
		Filter("var1", true).
		Filter("var2", false).
		Replace("rep1", "replacement1").
		Replace("rep2", "replacement2").
		Replace("empty", "").
		Replace("multi", "multi_line_1\nmulti_line_2\n").
		Process()
```

## Filter Syntax
```go
//?var1
This line is output if var1 is set to true via temple.Filter("var1", true)
//-

//!var2
This line is output if var2 is set to false via temple.Filter("var2", false) or if var2 is not set
//-

//?var1
This line is output if var1 is true
	//?var2
	This line is output if var1 is true and var2 is true
	//-
//-
```

## Replacer Syntax
```go
To replace ${rep1} with "replacement1" call temple.Replace("rep1", "replacement1")
```