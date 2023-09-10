package tag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	require.Equal(t, "<a></a>", A().String())
}

func TestAbbr(t *testing.T) {
	require.Equal(t, "<abbr></abbr>", Abbr().String())
}

func TestAcronym(t *testing.T) {
	require.Equal(t, "<acronym></acronym>", Acronym().String())
}

func TestAddress(t *testing.T) {
	require.Equal(t, "<address></address>", Address().String())
}

func TestApplet(t *testing.T) {
	require.Equal(t, "<applet></applet>", Applet().String())
}

func TestArticle(t *testing.T) {
	require.Equal(t, "<article></article>", Article().String())
}

func TestAside(t *testing.T) {
	require.Equal(t, "<aside></aside>", Aside().String())
}

func TestAudio(t *testing.T) {
	require.Equal(t, "<audio></audio>", Audio().String())
}

func TestBDI(t *testing.T) {
	require.Equal(t, "<bdi></bdi>", BDI().String())
}

func TestBDO(t *testing.T) {
	require.Equal(t, "<bdo></bdo>", BDO().String())
}

func TestBig(t *testing.T) {
	require.Equal(t, "<big></big>", Big().String())
}

func TestBlockQuote(t *testing.T) {
	require.Equal(t, "<blockquote></blockquote>", BlockQuote().String())
}

func TestBody(t *testing.T) {
	require.Equal(t, "<body></body>", Body().String())
}

func TestButton(t *testing.T) {
	require.Equal(t, "<button></button>", Button().String())
}

func TestCanvas(t *testing.T) {
	require.Equal(t, "<canvas></canvas>", Canvas().String())
}

func TestCaption(t *testing.T) {
	require.Equal(t, "<caption></caption>", Caption().String())
}

func TestCenter(t *testing.T) {
	require.Equal(t, "<center></center>", Center().String())
}

func TestCite(t *testing.T) {
	require.Equal(t, "<cite></cite>", Cite().String())
}

func TestCode(t *testing.T) {
	require.Equal(t, "<code></code>", Code().String())
}

func TestColGroup(t *testing.T) {
	require.Equal(t, "<colgroup></colgroup>", ColGroup().String())
}

func TestDataList(t *testing.T) {
	require.Equal(t, "<datalist></datalist>", DataList().String())
}

func TestDD(t *testing.T) {
	require.Equal(t, "<dd></dd>", DD().String())
}

func TestDel(t *testing.T) {
	require.Equal(t, "<del></del>", Del().String())
}

func TestDetails(t *testing.T) {
	require.Equal(t, "<details></details>", Details().String())
}

func TestDfn(t *testing.T) {
	require.Equal(t, "<dfn></dfn>", Dfn().String())
}

func TestDialog(t *testing.T) {
	require.Equal(t, "<dialog></dialog>", Dialog().String())
}

func TestDiv(t *testing.T) {
	require.Equal(t, "<div></div>", Div().String())
}

func TestDL(t *testing.T) {
	require.Equal(t, "<dl></dl>", DL().String())
}

func TestDT(t *testing.T) {
	require.Equal(t, "<dt></dt>", DT().String())
}

func TestEm(t *testing.T) {
	require.Equal(t, "<em></em>", Em().String())
}

func TestFieldSet(t *testing.T) {
	require.Equal(t, "<fieldset></fieldset>", FieldSet().String())
}

func TestFigCaption(t *testing.T) {
	require.Equal(t, "<figcaption></figcaption>", FigCaption().String())
}

func TestFigure(t *testing.T) {
	require.Equal(t, "<figure></figure>", Figure().String())
}

func TestFont(t *testing.T) {
	require.Equal(t, "<font></font>", Font().String())
}

func TestFooter(t *testing.T) {
	require.Equal(t, "<footer></footer>", Footer().String())
}

func TestForm(t *testing.T) {
	require.Equal(t, "<form></form>", Form().String())
}

func TestFrame(t *testing.T) {
	require.Equal(t, "<frame></frame>", Frame().String())
}

func TestFrameSet(t *testing.T) {
	require.Equal(t, "<frameset></frameset>", FrameSet().String())
}

func TestH1(t *testing.T) {
	require.Equal(t, "<h1></h1>", H1().String())
}

func TestH2(t *testing.T) {
	require.Equal(t, "<h2></h2>", H2().String())
}

func TestH3(t *testing.T) {
	require.Equal(t, "<h3></h3>", H3().String())
}

func TestH4(t *testing.T) {
	require.Equal(t, "<h4></h4>", H4().String())
}

func TestH5(t *testing.T) {
	require.Equal(t, "<h5></h5>", H5().String())
}

func TestHead(t *testing.T) {
	require.Equal(t, "<head></head>", Head().String())
}

func TestHeader(t *testing.T) {
	require.Equal(t, "<header></header>", Header().String())
}

func TestHGroup(t *testing.T) {
	require.Equal(t, "<hgroup></hgroup>", HGroup().String())
}

func TestHTML(t *testing.T) {
	require.Equal(t, "<html></html>", HTML().String())
}

func TestIFrame(t *testing.T) {
	require.Equal(t, "<iframe></iframe>", IFrame().String())
}

func TestIns(t *testing.T) {
	require.Equal(t, "<ins></ins>", Ins().String())
}

func TestKbd(t *testing.T) {
	require.Equal(t, "<kbd></kbd>", Kbd().String())
}

func TestLabel(t *testing.T) {
	require.Equal(t, "<label></label>", Label().String())
}

func TestLegend(t *testing.T) {
	require.Equal(t, "<legend></legend>", Legend().String())
}

func TestLI(t *testing.T) {
	require.Equal(t, "<li></li>", LI().String())
}

func TestMain(t *testing.T) {
	require.Equal(t, "<main></main>", Main().String())
}

func TestMap(t *testing.T) {
	require.Equal(t, "<map></map>", Map().String())
}

func TestMark(t *testing.T) {
	require.Equal(t, "<mark></mark>", Mark().String())
}

func TestMenu(t *testing.T) {
	require.Equal(t, "<menu></menu>", Menu().String())
}

func TestMenuItem(t *testing.T) {
	require.Equal(t, "<menuitem></menuitem>", MenuItem().String())
}

func TestMeter(t *testing.T) {
	require.Equal(t, "<meter></meter>", Meter().String())
}

func TestNav(t *testing.T) {
	require.Equal(t, "<nav></nav>", Nav().String())
}

func TestNoFrames(t *testing.T) {
	require.Equal(t, "<noframes></noframes>", NoFrames().String())
}

func TestNoScript(t *testing.T) {
	require.Equal(t, "<noscript></noscript>", NoScript().String())
}

func TestObject(t *testing.T) {
	require.Equal(t, "<object></object>", Object().String())
}

func TestOL(t *testing.T) {
	require.Equal(t, "<ol></ol>", OL().String())
}

func TestOptGroup(t *testing.T) {
	require.Equal(t, "<optgroup></optgroup>", OptGroup().String())
}

func TestOption(t *testing.T) {
	require.Equal(t, "<option></option>", Option().String())
}

func TestOutput(t *testing.T) {
	require.Equal(t, "<output></output>", Output().String())
}

func TestP(t *testing.T) {
	require.Equal(t, "<p></p>", P().String())
}

func TestPicture(t *testing.T) {
	require.Equal(t, "<picture></picture>", Picture().String())
}

func TestPre(t *testing.T) {
	require.Equal(t, "<pre></pre>", Pre().String())
}

func TestProgress(t *testing.T) {
	require.Equal(t, "<progress></progress>", Progress().String())
}

func TestRP(t *testing.T) {
	require.Equal(t, "<rp></rp>", RP().String())
}

func TestRT(t *testing.T) {
	require.Equal(t, "<rt></rt>", RT().String())
}

func TestRuby(t *testing.T) {
	require.Equal(t, "<ruby></ruby>", Ruby().String())
}

func TestSamp(t *testing.T) {
	require.Equal(t, "<samp></samp>", Samp().String())
}

func TestScript(t *testing.T) {
	require.Equal(t, "<script></script>", Script().String())
}

func TestSection(t *testing.T) {
	require.Equal(t, "<section></section>", Section().String())
}

func TestSelect(t *testing.T) {
	require.Equal(t, "<select></select>", Select().String())
}

func TestSlot(t *testing.T) {
	require.Equal(t, "<slot></slot>", Slot().String())
}

func TestSmall(t *testing.T) {
	require.Equal(t, "<small></small>", Small().String())
}

func TestSpan(t *testing.T) {
	require.Equal(t, "<span></span>", Span().String())
}

func TestStrike(t *testing.T) {
	require.Equal(t, "<strike></strike>", Strike().String())
}

func TestStrong(t *testing.T) {
	require.Equal(t, "<strong></strong>", Strong().String())
}

func TestSub(t *testing.T) {
	require.Equal(t, "<sub></sub>", Sub().String())
}

func TestSummary(t *testing.T) {
	require.Equal(t, "<summary></summary>", Summary().String())
}

func TestSup(t *testing.T) {
	require.Equal(t, "<sup></sup>", Sup().String())
}

func TestSVG(t *testing.T) {
	require.Equal(t, "<svg></svg>", SVG().String())
}

func TestTable(t *testing.T) {
	require.Equal(t, "<table></table>", Table().String())
}

func TestTBody(t *testing.T) {
	require.Equal(t, "<tbody></tbody>", TBody().String())
}

func TestTD(t *testing.T) {
	require.Equal(t, "<td></td>", TD().String())
}

func TestTemplate(t *testing.T) {
	require.Equal(t, "<template></template>", Template().String())
}

func TestTextArea(t *testing.T) {
	require.Equal(t, "<textarea></textarea>", TextArea().String())
}

func TestTFoot(t *testing.T) {
	require.Equal(t, "<tfoot></tfoot>", TFoot().String())
}

func TestTH(t *testing.T) {
	require.Equal(t, "<th></th>", TH().String())
}

func TestTHead(t *testing.T) {
	require.Equal(t, "<thead></thead>", THead().String())
}

func TestTime(t *testing.T) {
	require.Equal(t, "<time></time>", Time().String())
}

func TestTitle(t *testing.T) {
	require.Equal(t, "<title></title>", Title().String())
}

func TestTR(t *testing.T) {
	require.Equal(t, "<tr></tr>", TR().String())
}

func TestTT(t *testing.T) {
	require.Equal(t, "<tt></tt>", TT().String())
}

func TestUL(t *testing.T) {
	require.Equal(t, "<ul></ul>", UL().String())
}

func TestVar(t *testing.T) {
	require.Equal(t, "<var></var>", Var().String())
}

func TestVideo(t *testing.T) {
	require.Equal(t, "<video></video>", Video().String())
}

func TestWBr(t *testing.T) {
	require.Equal(t, "<wbr></wbr>", WBr().String())
}
