package attr

import (
	"testing"

	"github.com/jeffswenson/sanity/pkg/tag"
	"github.com/stretchr/testify/require"
)

func TestAllowFullScreen(t *testing.T) {
    require.Equal(t, 
        "<div allowfullscreen></div>", 
        tag.Div(AllowFullScreen()).String())
}

func TestAsync(t *testing.T) {
    require.Equal(t, 
        "<div async></div>", 
        tag.Div(Async()).String())
}

func TestAutoFocus(t *testing.T) {
    require.Equal(t, 
        "<div autofocus></div>", 
        tag.Div(AutoFocus()).String())
}

func TestAutoPlay(t *testing.T) {
    require.Equal(t, 
        "<div autoplay></div>", 
        tag.Div(AutoPlay()).String())
}

func TestChecked(t *testing.T) {
    require.Equal(t, 
        "<div checked></div>", 
        tag.Div(Checked()).String())
}

func TestControls(t *testing.T) {
    require.Equal(t, 
        "<div controls></div>", 
        tag.Div(Controls()).String())
}

func TestDefault(t *testing.T) {
    require.Equal(t, 
        "<div default></div>", 
        tag.Div(Default()).String())
}

func TestDefer(t *testing.T) {
    require.Equal(t, 
        "<div defer></div>", 
        tag.Div(Defer()).String())
}

func TestDisabled(t *testing.T) {
    require.Equal(t, 
        "<div disabled></div>", 
        tag.Div(Disabled()).String())
}

func TestFormNoValidate(t *testing.T) {
    require.Equal(t, 
        "<div formnovalidate></div>", 
        tag.Div(FormNoValidate()).String())
}

func TestHidden(t *testing.T) {
    require.Equal(t, 
        "<div hidden></div>", 
        tag.Div(Hidden()).String())
}

func TestInert(t *testing.T) {
    require.Equal(t, 
        "<div inert></div>", 
        tag.Div(Inert()).String())
}

func TestIsMap(t *testing.T) {
    require.Equal(t, 
        "<div ismap></div>", 
        tag.Div(IsMap()).String())
}

func TestItemScope(t *testing.T) {
    require.Equal(t, 
        "<div itemscope></div>", 
        tag.Div(ItemScope()).String())
}

func TestLoop(t *testing.T) {
    require.Equal(t, 
        "<div loop></div>", 
        tag.Div(Loop()).String())
}

func TestMultiple(t *testing.T) {
    require.Equal(t, 
        "<div multiple></div>", 
        tag.Div(Multiple()).String())
}

func TestMuted(t *testing.T) {
    require.Equal(t, 
        "<div muted></div>", 
        tag.Div(Muted()).String())
}

func TestNoModule(t *testing.T) {
    require.Equal(t, 
        "<div nomodule></div>", 
        tag.Div(NoModule()).String())
}

func TestNoValidate(t *testing.T) {
    require.Equal(t, 
        "<div novalidate></div>", 
        tag.Div(NoValidate()).String())
}

func TestOpen(t *testing.T) {
    require.Equal(t, 
        "<div open></div>", 
        tag.Div(Open()).String())
}

func TestPlaysInline(t *testing.T) {
    require.Equal(t, 
        "<div playsinline></div>", 
        tag.Div(PlaysInline()).String())
}

func TestReadOnly(t *testing.T) {
    require.Equal(t, 
        "<div readonly></div>", 
        tag.Div(ReadOnly()).String())
}

func TestRequired(t *testing.T) {
    require.Equal(t, 
        "<div required></div>", 
        tag.Div(Required()).String())
}

func TestReversed(t *testing.T) {
    require.Equal(t, 
        "<div reversed></div>", 
        tag.Div(Reversed()).String())
}

func TestSelected(t *testing.T) {
    require.Equal(t, 
        "<div selected></div>", 
        tag.Div(Selected()).String())
}
