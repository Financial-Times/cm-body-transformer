package filters

import (
	"fmt"
	"testing"
)

func TestTextTransform(t *testing.T) {
	inputText := `<body><content data-embedded="true" id="aae9611e-f66c-4fe4-a6c6-2e2bdea69060" type="http://www.ft.com/ontology/content/ImageSet"></content>
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris scelerisque, nunc vel consectetur sagittis, purus ex ultrices metus, in consectetur nisl lacus congue nulla. Integer fermentum molestie dui at accumsan.</p>
<p>Nam <content id="396d9102-9845-4ce2-8783-49b73f8f1302" type="http://www.ft.com/ontology/content/Article">scelerisque luctus</content> tristique. Aliquam orci massa, hendrerit non pulvinar a, tristique vitae enim. Pellentesque laoreet condimentum nulla sed tempor. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculous mus. Quisque euismod euismod porta. Praesent id sapien et magna porta malesuada. Proin sit amet justo vel augue sollicitudin volutpat sodales id turpis.</p>
<p>Sed posuere vestibulum metus non cursus. Fusce ac blandit erat. Fusce turpis turpis, vehicula et condimentum quis, dapibus eget odio. Vivamus lobortis vulputate sapien quis ultrices. </p>
<p>Morbi laoreet, sem at bibendum rutrum, ligula erat rhoncus est, eget hendrerit leo diam sit amet mauris. Curabitur cursus dictum mi id eleifend. Pellentesque sed massa sit amet massa ornare accumsan. Nulla eget lobortis velit. </p>
<p>Cras vel libero ut arcu hendrerit accumsan. “Vivamus ligula lectus”, vestibulum at nisi id, imperdiet “ornare libero”.</p>
<pull-quote>
    <pull-quote-text><p>Maecenas ac ipsum in elit aliquam consectetur. Proin felis metus, efficitur et nulla eu, interdum malesuada diam.</p></pull-quote-text><pull-quote-image><content data-embedded="true" id="77c8a5b5-c9e3-4df2-ad5f-3ef35fe1d9d4" type="http://www.ft.com/ontology/content/ImageSet"></content></pull-quote-image><pull-quote-source>Pellentesque habitant, morbi tristique</pull-quote-source>
</pull-quote>
<p>Donec id faucibus erat. Suspendisse tempor laoreet lorem, sit amet vehicula massa facilisis at. Nulla quis feugiat massa. Praesent viverra non lectus ut ullamcorper. Phasellus <content id="c71efed9-fe5a-488d-9f47-20c15d177153" type="http://www.ft.com/ontology/content/Article">porttitor neque</content> at volutpat pulvinar.</p>
<p>“Curabitur fermentum, dolor vel interdum varius, tellus justo dapibus velit, interdum sollicitudin dolor nibh varius velit.”</p>
</body>`

	transformedText := Apply(inputText, DefaultContentFilters()...)

	expectedText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris scelerisque, nunc vel consectetur sagittis, " +
		"purus ex ultrices metus, in consectetur nisl lacus congue nulla. Integer fermentum molestie dui at accumsan. Nam scelerisque luctus tristique. " +
		"Aliquam orci massa, hendrerit non pulvinar a, tristique vitae enim. Pellentesque laoreet condimentum nulla sed tempor. " +
		"Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculous mus. Quisque euismod euismod porta. " +
		"Praesent id sapien et magna porta malesuada. Proin sit amet justo vel augue sollicitudin volutpat sodales id turpis. " +
		"Sed posuere vestibulum metus non cursus. Fusce ac blandit erat. Fusce turpis turpis, vehicula et condimentum quis, dapibus eget odio. " +
		"Vivamus lobortis vulputate sapien quis ultrices. Morbi laoreet, sem at bibendum rutrum, ligula erat rhoncus est, eget hendrerit leo diam sit amet mauris. " +
		"Curabitur cursus dictum mi id eleifend. Pellentesque sed massa sit amet massa ornare accumsan. Nulla eget lobortis velit. " +
		"Cras vel libero ut arcu hendrerit accumsan. “Vivamus ligula lectus”, vestibulum at nisi id, imperdiet “ornare libero”. " +
		"Donec id faucibus erat. Suspendisse tempor laoreet lorem, sit amet vehicula massa facilisis at. Nulla quis feugiat massa. " +
		"Praesent viverra non lectus ut ullamcorper. Phasellus porttitor neque at volutpat pulvinar. “Curabitur fermentum, dolor vel interdum varius, " +
		"tellus justo dapibus velit, interdum sollicitudin dolor nibh varius velit.”"

	equal(t, expectedText, transformedText, fmt.Sprintf("Expected text %s differs from actual text %s ", transformedText, expectedText))
}

func TestBlogTransform(t *testing.T) {
	inputText := `<body><p><a href="http://www.ft.com/fake-blog/files/2017/02/Fake_blog_post_title-line_chart-ft-web-themelarge-600x397.1234567890.png"><img alt="" height="398" src="http://www.ft.com/fake-blog/files/2017/02/Fake_blog_post_title-line_chart-ft-web-themelarge-600x397.1234567890.png" width="600"/></a></p>
<p>Aliquam sagittis ipsum non tortor placerat scelerisque.</p>
<p>Maecenas lobortis purus ut cursus tempor. Vestibulum lacus neque, auctor et euismod in, ultricies dictum sem. Fusce finibus erat quis ipsum pharetra, quis vehicula urna varius. Donec consequat pellentesque erat nec porta.</p>
<p>Praesent vel leo feugiat, rhoncus quam quis, ullamcorper augue. Pellentesque quis nisi nec sapien accumsan efficitur. Quisque commodo mollis metus.</p>
<p><a href="http://www.ft.com/fake-blog/files/2017/02/fake-image.png"><img alt="" height="382" src="http://www.ft.com/fake-blog/files/2017/02/fake-image.png" width="733"/></a></p>
<p>Aliquam eros tellus, pharetra non orci eu, dictum semper enim. Donec vel dapibus mi, vel fermentum sapien.</p>
<p>Ut nec nibh ex. Proin dignissim ipsum at lacus condimentum efficitur. Donec at felis felis. Etiam sagittis condimentum maximus.</p>
<p><em>Donec id faucibus erat </em></p>
</body>`

	transformedText := Apply(inputText, DefaultContentFilters()...)

	expectedText := "Aliquam sagittis ipsum non tortor placerat scelerisque. Maecenas lobortis purus ut cursus tempor. " +
		"Vestibulum lacus neque, auctor et euismod in, ultricies dictum sem. Fusce finibus erat quis ipsum pharetra, " +
		"quis vehicula urna varius. Donec consequat pellentesque erat nec porta. Praesent vel leo feugiat, rhoncus quam quis, " +
		"ullamcorper augue. Pellentesque quis nisi nec sapien accumsan efficitur. Quisque commodo mollis metus. Aliquam eros " +
		"tellus, pharetra non orci eu, dictum semper enim. Donec vel dapibus mi, vel fermentum sapien. Ut nec nibh ex. " +
		"Proin dignissim ipsum at lacus condimentum efficitur. Donec at felis felis. Etiam sagittis condimentum maximus. Donec id faucibus erat"

	equal(t, expectedText, transformedText, fmt.Sprintf("Expected text %s differs from actual text %s ", transformedText, expectedText))
}

func TestPullTagTransformer(t *testing.T) {
	equal(t, "this is a test followed by another test", RemovePullQuoteTag("this is a test<pull-quote>pull quote</pull-quote> followed by another test<pull-quote>\npull quote\n</pull-quote>"), "Pull tags not transformed properly")
}

func TestWebPullTagTransformer(t *testing.T) {
	equal(t, "this is a test followed by another test", RemoveWebPullQuoteTag("this is a test<web-pull-quote>web-pull quote</web-pull-quote> followed by another test<web-pull-quote>\nweb-pull quote\n</web-pull-quote>"), "Web-pull tags not transformed properly")
}

func TestTableTagTransformer(t *testing.T) {
	equal(t, "this is a test followed by another test", RemoveTableTag("this is a test<table style=\"width:100%\">\n	<tr><th>Firstname</th><th>Lastname</th><th>Age</th></tr>\n<tr>	<td>Jill</td><td>Smith</td>	<td>50</td>	</tr>\n	<tr><td>Eve</td><td>Jackson</td>	<td>94</td></tr>\t	</table> followed by another test<table>\nempty table\n</table>"), "Table tags not transformed properly")
}

func TestPromoBoxTagTransformer(t *testing.T) {
	equal(t, "this is a test followed by another test", RemovePromoBoxTag("this is a test<promo-box>promo-box stuff</promo-box> followed by another test<promo-box>\npromo-box stuff\n</promo-box>"), "Promobox tags not transformed properly")
}

func TestWebInlinePictureTagTransformer(t *testing.T) {
	equal(t, "this is a test followed by another test", RemoveWebInlinePictureTag("this is a test<web-inline-picture>web-inline-picture stuff</web-inline-picture> followed by another test<web-inline-picture>\nweb-inline-picture stuff\n</web-inline-picture>"), "web-inline-picture tags not transformed properly")
}

func TestHtmlEntityTransformer(t *testing.T) {
	equal(t, "test ‑£& >&", RemoveHTMLEntity("test &#8209;&pound;&amp;&nbsp;&gt;&"), "Entities not transformed properly")
}

func TestTagsRemover(t *testing.T) {
	equal(t, "this is a simple test for tag removal", RemoveGenericTags("this is a <b>simple </b>test<br> for <span attr=\"val\">tag </span>removal"), "Tags not transformed properly")
}

func TestDedupSpaces(t *testing.T) {
	equal(t, " lots of\tspace\nbut\nno room", DedupSpaces(" lots  of\t\tspace\r\nbut \t\nno room"), "Whitespace not transformed properly")
}

func TestTagsRemover2(t *testing.T) {
	input := `<experimental><div data-layout-name=\"card\" class=\"n-content-layout\" data-layout-width=\"fullWidth\"><div class=\"n-content-layout__container\"><h3>Recommended newsletters for you</h3><div class=\"n-content-layout__slot\" data-slot-width=\"true\"><p><strong>#fintechFT</strong> — The biggest themes in the digital disruption of financial services. Sign up <a href=\"https://ep.ft.com/newsletters/subscribe?newsletterIds=575981ede74eb90300a44d8e\">here</a></p><p><strong>Martin Sandbu’s Free Lunch</strong> — Your guide to the global economic policy debate. Sign up <a href=\"https://ep.ft.com/newsletters/subscribe?newsletterIds=56388465e4b0c3d64132e189\">here</a></p></div></div></div></experimental>`
	expected := `Recommended newsletters for you #fintechFT — The biggest themes in the digital disruption of financial services. Sign up here Martin Sandbu’s Free Lunch — Your guide to the global economic policy debate. Sign up here`
	result := RemoveGenericTags(input)
	equal(t, expected, result, "")
}

func equal(t *testing.T, expected, actual string, msg string) {
	t.Helper()
	if expected != actual {
		if msg != "" {
			t.Fatal(msg)
		} else {
			t.Fatalf("missmatch\nexpected: '%s'\nactual: '%s'\n", expected, actual)
		}
	}
}
