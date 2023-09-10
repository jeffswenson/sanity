package attr

import (
	"testing"

	"github.com/jeffswenson/sanity/pkg/tag"
	"github.com/stretchr/testify/require"
)

func TestAccept(t *testing.T) {
    require.Equal(t, 
        "<div accept=\"bar\"></div>", 
        tag.Div(Accept("bar")).String())
}

func TestAcceptCharSet(t *testing.T) {
    require.Equal(t, 
        "<div accept-charset=\"bar\"></div>", 
        tag.Div(AcceptCharSet("bar")).String())
}

func TestAccessKey(t *testing.T) {
    require.Equal(t, 
        "<div accesskey=\"bar\"></div>", 
        tag.Div(AccessKey("bar")).String())
}

func TestAllow(t *testing.T) {
    require.Equal(t, 
        "<div allow=\"bar\"></div>", 
        tag.Div(Allow("bar")).String())
}

func TestAlt(t *testing.T) {
    require.Equal(t, 
        "<div alt=\"bar\"></div>", 
        tag.Div(Alt("bar")).String())
}

func TestAs(t *testing.T) {
    require.Equal(t, 
        "<div as=\"bar\"></div>", 
        tag.Div(As("bar")).String())
}

func TestAutoCapitalize(t *testing.T) {
    require.Equal(t, 
        "<div autocapitalize=\"bar\"></div>", 
        tag.Div(AutoCapitalize("bar")).String())
}

func TestAutoComplete(t *testing.T) {
    require.Equal(t, 
        "<div autocomplete=\"bar\"></div>", 
        tag.Div(AutoComplete("bar")).String())
}

func TestBlocking(t *testing.T) {
    require.Equal(t, 
        "<div blocking=\"bar\"></div>", 
        tag.Div(Blocking("bar")).String())
}

func TestCharSet(t *testing.T) {
    require.Equal(t, 
        "<div charset=\"bar\"></div>", 
        tag.Div(CharSet("bar")).String())
}

func TestClass(t *testing.T) {
    require.Equal(t, 
        "<div class=\"bar\"></div>", 
        tag.Div(Class("bar")).String())
}

func TestColor(t *testing.T) {
    require.Equal(t, 
        "<div color=\"bar\"></div>", 
        tag.Div(Color("bar")).String())
}

func TestCols(t *testing.T) {
    require.Equal(t, 
        "<div cols=\"bar\"></div>", 
        tag.Div(Cols("bar")).String())
}

func TestColSpan(t *testing.T) {
    require.Equal(t, 
        "<div colspan=\"bar\"></div>", 
        tag.Div(ColSpan("bar")).String())
}

func TestContent(t *testing.T) {
    require.Equal(t, 
        "<div content=\"bar\"></div>", 
        tag.Div(Content("bar")).String())
}

func TestContentEditable(t *testing.T) {
    require.Equal(t, 
        "<div contenteditable=\"bar\"></div>", 
        tag.Div(ContentEditable("bar")).String())
}

func TestCoords(t *testing.T) {
    require.Equal(t, 
        "<div coords=\"bar\"></div>", 
        tag.Div(Coords("bar")).String())
}

func TestCrossOrigin(t *testing.T) {
    require.Equal(t, 
        "<div crossorigin=\"bar\"></div>", 
        tag.Div(CrossOrigin("bar")).String())
}

func TestDateTime(t *testing.T) {
    require.Equal(t, 
        "<div datetime=\"bar\"></div>", 
        tag.Div(DateTime("bar")).String())
}

func TestDecoding(t *testing.T) {
    require.Equal(t, 
        "<div decoding=\"bar\"></div>", 
        tag.Div(Decoding("bar")).String())
}

func TestDir(t *testing.T) {
    require.Equal(t, 
        "<div dir=\"bar\"></div>", 
        tag.Div(Dir("bar")).String())
}

func TestDirName(t *testing.T) {
    require.Equal(t, 
        "<div dirname=\"bar\"></div>", 
        tag.Div(DirName("bar")).String())
}

func TestDraggable(t *testing.T) {
    require.Equal(t, 
        "<div draggable=\"bar\"></div>", 
        tag.Div(Draggable("bar")).String())
}

func TestEnctype(t *testing.T) {
    require.Equal(t, 
        "<div enctype=\"bar\"></div>", 
        tag.Div(Enctype("bar")).String())
}

func TestEnterKeyHint(t *testing.T) {
    require.Equal(t, 
        "<div enterkeyhint=\"bar\"></div>", 
        tag.Div(EnterKeyHint("bar")).String())
}

func TestFetchPriority(t *testing.T) {
    require.Equal(t, 
        "<div fetchpriority=\"bar\"></div>", 
        tag.Div(FetchPriority("bar")).String())
}

func TestFor(t *testing.T) {
    require.Equal(t, 
        "<div for=\"bar\"></div>", 
        tag.Div(For("bar")).String())
}

func TestFormAction(t *testing.T) {
    require.Equal(t, 
        "<div formaction=\"bar\"></div>", 
        tag.Div(FormAction("bar")).String())
}

func TestFormEncType(t *testing.T) {
    require.Equal(t, 
        "<div formenctype=\"bar\"></div>", 
        tag.Div(FormEncType("bar")).String())
}

func TestFormMethod(t *testing.T) {
    require.Equal(t, 
        "<div formmethod=\"bar\"></div>", 
        tag.Div(FormMethod("bar")).String())
}

func TestFormTarget(t *testing.T) {
    require.Equal(t, 
        "<div formtarget=\"bar\"></div>", 
        tag.Div(FormTarget("bar")).String())
}

func TestHeaders(t *testing.T) {
    require.Equal(t, 
        "<div headers=\"bar\"></div>", 
        tag.Div(Headers("bar")).String())
}

func TestHeight(t *testing.T) {
    require.Equal(t, 
        "<div height=\"bar\"></div>", 
        tag.Div(Height("bar")).String())
}

func TestHigh(t *testing.T) {
    require.Equal(t, 
        "<div high=\"bar\"></div>", 
        tag.Div(High("bar")).String())
}

func TestHRef(t *testing.T) {
    require.Equal(t, 
        "<div href=\"bar\"></div>", 
        tag.Div(HRef("bar")).String())
}

func TestHRefLang(t *testing.T) {
    require.Equal(t, 
        "<div hreflang=\"bar\"></div>", 
        tag.Div(HRefLang("bar")).String())
}

func TestHttpEquiv(t *testing.T) {
    require.Equal(t, 
        "<div http-equiv=\"bar\"></div>", 
        tag.Div(HttpEquiv("bar")).String())
}

func TestId(t *testing.T) {
    require.Equal(t, 
        "<div id=\"bar\"></div>", 
        tag.Div(Id("bar")).String())
}

func TestInputMode(t *testing.T) {
    require.Equal(t, 
        "<div inputmode=\"bar\"></div>", 
        tag.Div(InputMode("bar")).String())
}

func TestIntegrity(t *testing.T) {
    require.Equal(t, 
        "<div integrity=\"bar\"></div>", 
        tag.Div(Integrity("bar")).String())
}

func TestIs(t *testing.T) {
    require.Equal(t, 
        "<div is=\"bar\"></div>", 
        tag.Div(Is("bar")).String())
}

func TestItemId(t *testing.T) {
    require.Equal(t, 
        "<div itemid=\"bar\"></div>", 
        tag.Div(ItemId("bar")).String())
}

func TestItemProp(t *testing.T) {
    require.Equal(t, 
        "<div itemprop=\"bar\"></div>", 
        tag.Div(ItemProp("bar")).String())
}

func TestItemRef(t *testing.T) {
    require.Equal(t, 
        "<div itemref=\"bar\"></div>", 
        tag.Div(ItemRef("bar")).String())
}

func TestItemType(t *testing.T) {
    require.Equal(t, 
        "<div itemtype=\"bar\"></div>", 
        tag.Div(ItemType("bar")).String())
}

func TestKind(t *testing.T) {
    require.Equal(t, 
        "<div kind=\"bar\"></div>", 
        tag.Div(Kind("bar")).String())
}

func TestList(t *testing.T) {
    require.Equal(t, 
        "<div list=\"bar\"></div>", 
        tag.Div(List("bar")).String())
}

func TestLoading(t *testing.T) {
    require.Equal(t, 
        "<div loading=\"bar\"></div>", 
        tag.Div(Loading("bar")).String())
}

func TestLow(t *testing.T) {
    require.Equal(t, 
        "<div low=\"bar\"></div>", 
        tag.Div(Low("bar")).String())
}

func TestMax(t *testing.T) {
    require.Equal(t, 
        "<div max=\"bar\"></div>", 
        tag.Div(Max("bar")).String())
}

func TestMaxLength(t *testing.T) {
    require.Equal(t, 
        "<div maxlength=\"bar\"></div>", 
        tag.Div(MaxLength("bar")).String())
}

func TestMedia(t *testing.T) {
    require.Equal(t, 
        "<div media=\"bar\"></div>", 
        tag.Div(Media("bar")).String())
}

func TestMethod(t *testing.T) {
    require.Equal(t, 
        "<div method=\"bar\"></div>", 
        tag.Div(Method("bar")).String())
}

func TestMin(t *testing.T) {
    require.Equal(t, 
        "<div min=\"bar\"></div>", 
        tag.Div(Min("bar")).String())
}

func TestMinLength(t *testing.T) {
    require.Equal(t, 
        "<div minlength=\"bar\"></div>", 
        tag.Div(MinLength("bar")).String())
}

func TestName(t *testing.T) {
    require.Equal(t, 
        "<div name=\"bar\"></div>", 
        tag.Div(Name("bar")).String())
}

func TestNonce(t *testing.T) {
    require.Equal(t, 
        "<div nonce=\"bar\"></div>", 
        tag.Div(Nonce("bar")).String())
}

func TestOptimum(t *testing.T) {
    require.Equal(t, 
        "<div optimum=\"bar\"></div>", 
        tag.Div(Optimum("bar")).String())
}

func TestPattern(t *testing.T) {
    require.Equal(t, 
        "<div pattern=\"bar\"></div>", 
        tag.Div(Pattern("bar")).String())
}

func TestPing(t *testing.T) {
    require.Equal(t, 
        "<div ping=\"bar\"></div>", 
        tag.Div(Ping("bar")).String())
}

func TestPlaceholder(t *testing.T) {
    require.Equal(t, 
        "<div placeholder=\"bar\"></div>", 
        tag.Div(Placeholder("bar")).String())
}

func TestPopOver(t *testing.T) {
    require.Equal(t, 
        "<div popover=\"bar\"></div>", 
        tag.Div(PopOver("bar")).String())
}

func TestPopOverTarget(t *testing.T) {
    require.Equal(t, 
        "<div popovertarget=\"bar\"></div>", 
        tag.Div(PopOverTarget("bar")).String())
}

func TestPopOverTargetAction(t *testing.T) {
    require.Equal(t, 
        "<div popovertargetaction=\"bar\"></div>", 
        tag.Div(PopOverTargetAction("bar")).String())
}

func TestPoster(t *testing.T) {
    require.Equal(t, 
        "<div poster=\"bar\"></div>", 
        tag.Div(Poster("bar")).String())
}

func TestPreLoad(t *testing.T) {
    require.Equal(t, 
        "<div preload=\"bar\"></div>", 
        tag.Div(PreLoad("bar")).String())
}

func TestReferrerPolicy(t *testing.T) {
    require.Equal(t, 
        "<div referrerpolicy=\"bar\"></div>", 
        tag.Div(ReferrerPolicy("bar")).String())
}

func TestRel(t *testing.T) {
    require.Equal(t, 
        "<div rel=\"bar\"></div>", 
        tag.Div(Rel("bar")).String())
}

func TestRows(t *testing.T) {
    require.Equal(t, 
        "<div rows=\"bar\"></div>", 
        tag.Div(Rows("bar")).String())
}

func TestRowSpan(t *testing.T) {
    require.Equal(t, 
        "<div rowspan=\"bar\"></div>", 
        tag.Div(RowSpan("bar")).String())
}

func TestSandbox(t *testing.T) {
    require.Equal(t, 
        "<div sandbox=\"bar\"></div>", 
        tag.Div(Sandbox("bar")).String())
}

func TestScope(t *testing.T) {
    require.Equal(t, 
        "<div scope=\"bar\"></div>", 
        tag.Div(Scope("bar")).String())
}

func TestSize(t *testing.T) {
    require.Equal(t, 
        "<div size=\"bar\"></div>", 
        tag.Div(Size("bar")).String())
}

func TestSizes(t *testing.T) {
    require.Equal(t, 
        "<div sizes=\"bar\"></div>", 
        tag.Div(Sizes("bar")).String())
}

func TestSpellCheck(t *testing.T) {
    require.Equal(t, 
        "<div spellcheck=\"bar\"></div>", 
        tag.Div(SpellCheck("bar")).String())
}

func TestSrc(t *testing.T) {
    require.Equal(t, 
        "<div src=\"bar\"></div>", 
        tag.Div(Src("bar")).String())
}

func TestSrcDoc(t *testing.T) {
    require.Equal(t, 
        "<div srcdoc=\"bar\"></div>", 
        tag.Div(SrcDoc("bar")).String())
}

func TestSrcLang(t *testing.T) {
    require.Equal(t, 
        "<div srclang=\"bar\"></div>", 
        tag.Div(SrcLang("bar")).String())
}

func TestSrcSet(t *testing.T) {
    require.Equal(t, 
        "<div srcset=\"bar\"></div>", 
        tag.Div(SrcSet("bar")).String())
}

func TestStart(t *testing.T) {
    require.Equal(t, 
        "<div start=\"bar\"></div>", 
        tag.Div(Start("bar")).String())
}

func TestStep(t *testing.T) {
    require.Equal(t, 
        "<div step=\"bar\"></div>", 
        tag.Div(Step("bar")).String())
}

func TestStyle(t *testing.T) {
    require.Equal(t, 
        "<div style=\"bar\"></div>", 
        tag.Div(Style("bar")).String())
}

func TestTabIndex(t *testing.T) {
    require.Equal(t, 
        "<div tabindex=\"bar\"></div>", 
        tag.Div(TabIndex("bar")).String())
}

func TestTarget(t *testing.T) {
    require.Equal(t, 
        "<div target=\"bar\"></div>", 
        tag.Div(Target("bar")).String())
}

func TestTranslate(t *testing.T) {
    require.Equal(t, 
        "<div translate=\"bar\"></div>", 
        tag.Div(Translate("bar")).String())
}

func TestType(t *testing.T) {
    require.Equal(t, 
        "<div type=\"bar\"></div>", 
        tag.Div(Type("bar")).String())
}

func TestUseMap(t *testing.T) {
    require.Equal(t, 
        "<div usemap=\"bar\"></div>", 
        tag.Div(UseMap("bar")).String())
}

func TestValue(t *testing.T) {
    require.Equal(t, 
        "<div value=\"bar\"></div>", 
        tag.Div(Value("bar")).String())
}

func TestWidth(t *testing.T) {
    require.Equal(t, 
        "<div width=\"bar\"></div>", 
        tag.Div(Width("bar")).String())
}

func TestWrap(t *testing.T) {
    require.Equal(t, 
        "<div wrap=\"bar\"></div>", 
        tag.Div(Wrap("bar")).String())
}
