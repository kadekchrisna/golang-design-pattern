package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

// HTMLElement ...
type HTMLElement struct {
	Name, Text string
	Elements   []HTMLElement
}

// HTMLBuilder ...
type HTMLBuilder struct {
	RootName string
	Root     HTMLElement
}

// NewHTMLBuilder ...
func NewHTMLBuilder(rootName string) *HTMLBuilder {
	return &HTMLBuilder{
		RootName: rootName,
		Root: HTMLElement{
			Name: rootName, Text: "", Elements: []HTMLElement{},
		},
	}
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.Name))
	if len(e.Text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize*(indent+1)))
		sb.WriteString(e.Text)
		sb.WriteString("\n")
	}

	for _, el := range e.Elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n",
		i, e.Name))
	return sb.String()
}

func (b *HTMLBuilder) String() string {
	return b.Root.String()
}

// AddChild ...
func (b *HTMLBuilder) AddChild(childName, childText string) *HTMLBuilder {
	e := HTMLElement{childName, childText, []HTMLElement{}}
	b.Root.Elements = append(b.Root.Elements, e)
	return b
}

func main() {
	b := NewHTMLBuilder("ul")
	b.AddChild("li", "hello")

	fmt.Println(b.String())

}
