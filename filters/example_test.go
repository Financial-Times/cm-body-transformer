package filters_test

import (
	"fmt"

	"github.com/Financial-Times/cm-body-transformer/filters"
)

func ExampleApply() {
	body := `
<body><p><a href="http://www.ft.com/fake-blog/files/2017/02/Fake_blog_post_title-line_chart-ft-web-themelarge-600x397.1234567890.png"><img alt="" height="398" src="http://www.ft.com/fake-blog/files/2017/02/Fake_blog_post_title-line_chart-ft-web-themelarge-600x397.1234567890.png" width="600"/></a></p>
<p>Aliquam sagittis ipsum non tortor placerat scelerisque.</p>
<p>Maecenas lobortis purus ut cursus tempor. Vestibulum lacus neque, auctor et euismod in, ultricies dictum sem. Fusce finibus erat quis ipsum pharetra, quis vehicula urna varius. Donec consequat pellentesque erat nec porta.</p>
<p>Praesent vel leo feugiat, rhoncus quam quis, ullamcorper augue. Pellentesque quis nisi nec sapien accumsan efficitur. Quisque commodo mollis metus.</p>
<p><a href="http://www.ft.com/fake-blog/files/2017/02/fake-image.png"><img alt="" height="382" src="http://www.ft.com/fake-blog/files/2017/02/fake-image.png" width="733"/></a></p>
<p>Aliquam eros tellus, pharetra non orci eu, dictum semper enim. Donec vel dapibus mi, vel fermentum sapien.</p>
<p>Ut nec nibh ex. Proin dignissim ipsum at lacus condimentum efficitur. Donec at felis felis. Etiam sagittis condimentum maximus.</p>
<p><em>Donec id faucibus erat </em></p>
</body>
`
	result := filters.Apply(body, filters.DefaultContentFilters()...)
	fmt.Println(result)
	// Output: Aliquam sagittis ipsum non tortor placerat scelerisque. Maecenas lobortis purus ut cursus tempor. Vestibulum lacus neque, auctor et euismod in, ultricies dictum sem. Fusce finibus erat quis ipsum pharetra, quis vehicula urna varius. Donec consequat pellentesque erat nec porta. Praesent vel leo feugiat, rhoncus quam quis, ullamcorper augue. Pellentesque quis nisi nec sapien accumsan efficitur. Quisque commodo mollis metus. Aliquam eros tellus, pharetra non orci eu, dictum semper enim. Donec vel dapibus mi, vel fermentum sapien. Ut nec nibh ex. Proin dignissim ipsum at lacus condimentum efficitur. Donec at felis felis. Etiam sagittis condimentum maximus. Donec id faucibus erat
}

func ExampleDedupSpaces() {
	body := "testing\n\tdedup"
	result := filters.Apply(body, filters.DedupSpaces)
	fmt.Println(result)
	// Output: testing	dedup
}
