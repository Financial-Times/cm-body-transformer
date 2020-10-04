# cm-body-transformer
Go library that transforms content body in format expected by external consumers

The end goal of this library is to do body transformations so that the article body is the same as the one returned by the public content API or enriched content API. The logic for these transformations is spread across [public-content-read](https://github.com/Financial-Times/content-public-read) and [api-policy-component](https://github.com/Financial-Times/api-policy-component). There are small transformation done by [enriched-content-read-api](https://github.com/Financial-Times/enriched-content-read-api). The actual library for parsing and manipulating the xml body is [content-body-processing](https://github.com/Financial-Times/content-body-processing). 

**Please note that the current version of this library does NOT fully replicate the logic for content body transformation implemented for public content API so the transformed body is not always the same string as the one returned by public content API. However the differences should be only cosmetic. No semantic transformation rule should be skipped.**

The integration tests are meant only for manual execution as they make calls to our public and internal APIs. In order to run the integration tests use:
```shell script
go test -tags=manualintegration -v --cover --count=1 --apiKey XXXX --basicAuthUser XXXX --basicAuthPassword XXXX
```
