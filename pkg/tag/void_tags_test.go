package tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArea(t *testing.T) {
    require.Equal(t, "<area>", Area().String())
}

func TestBase(t *testing.T) {
    require.Equal(t, "<base>", Base().String())
}

func TestBaseFont(t *testing.T) {
    require.Equal(t, "<basefont>", BaseFont().String())
}

func TestBr(t *testing.T) {
    require.Equal(t, "<br>", Br().String())
}

func TestCol(t *testing.T) {
    require.Equal(t, "<col>", Col().String())
}

func TestEmbed(t *testing.T) {
    require.Equal(t, "<embed>", Embed().String())
}

func TestHr(t *testing.T) {
    require.Equal(t, "<hr>", Hr().String())
}

func TestImg(t *testing.T) {
    require.Equal(t, "<img>", Img().String())
}

func TestInput(t *testing.T) {
    require.Equal(t, "<input>", Input().String())
}

func TestLink(t *testing.T) {
    require.Equal(t, "<link>", Link().String())
}

func TestMeta(t *testing.T) {
    require.Equal(t, "<meta>", Meta().String())
}

func TestParam(t *testing.T) {
    require.Equal(t, "<param>", Param().String())
}

func TestSource(t *testing.T) {
    require.Equal(t, "<source>", Source().String())
}

func TestTrack(t *testing.T) {
    require.Equal(t, "<track>", Track().String())
}
