package benchmark

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"

	"github.com/jeffswenson/sanity/pkg/attr"
	"github.com/jeffswenson/sanity/pkg/html"
	"github.com/jeffswenson/sanity/pkg/tag"
	"github.com/stretchr/testify/require"
)

type simpleListModel struct {
	Items []string
}

func generateModel(length int) simpleListModel {
	var result simpleListModel
	for i := 0; i < length; i++ {
		result.Items = append(result.Items, fmt.Sprintf("list item %d", i))
	}
	return result
}

func listView(model simpleListModel) html.Node {
	return tag.UL(
		attr.Class("list"),
		html.ForEach(model.Items, func(m string) html.Node {
			return tag.LI(attr.Class("item"), html.InnerText(m))
		}),
	)
}

func BenchmarkSanityList1000(b *testing.B) {
	list := generateModel(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := listView(list).Render()
		require.NotEmpty(b, result)
	}
}

func BenchmarkSanityListRenderOnly1000(b *testing.B) {
	list := generateModel(1000)
	element := listView(list)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		require.NotEmpty(b, element.Render())
	}
}

func BenchmarkGoTemplate1000(b *testing.B) {
	listTemplate, err := template.New("list").Parse(`<ul class="list">{{range .Items}}<li class="item">{{.}}</li>{{end}}</ul>`)
	require.NoError(b, err)
	list := generateModel(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		require.NoError(b, listTemplate.Execute(&buffer, list))
	}
}
