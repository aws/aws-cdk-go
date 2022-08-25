package awskendra


// Provides the configuration information required for Amazon Kendra Web Crawler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerConfigurationProperty := &webCrawlerConfigurationProperty{
//   	urls: &webCrawlerUrlsProperty{
//   		seedUrlConfiguration: &webCrawlerSeedUrlConfigurationProperty{
//   			seedUrls: []*string{
//   				jsii.String("seedUrls"),
//   			},
//
//   			// the properties below are optional
//   			webCrawlerMode: jsii.String("webCrawlerMode"),
//   		},
//   		siteMapsConfiguration: &webCrawlerSiteMapsConfigurationProperty{
//   			siteMaps: []*string{
//   				jsii.String("siteMaps"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	authenticationConfiguration: &webCrawlerAuthenticationConfigurationProperty{
//   		basicAuthentication: []interface{}{
//   			&webCrawlerBasicAuthenticationProperty{
//   				credentials: jsii.String("credentials"),
//   				host: jsii.String("host"),
//   				port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	crawlDepth: jsii.Number(123),
//   	maxContentSizePerPageInMegaBytes: jsii.Number(123),
//   	maxLinksPerPage: jsii.Number(123),
//   	maxUrlsPerMinuteCrawlRate: jsii.Number(123),
//   	proxyConfiguration: &proxyConfigurationProperty{
//   		host: jsii.String("host"),
//   		port: jsii.Number(123),
//
//   		// the properties below are optional
//   		credentials: jsii.String("credentials"),
//   	},
//   	urlExclusionPatterns: []*string{
//   		jsii.String("urlExclusionPatterns"),
//   	},
//   	urlInclusionPatterns: []*string{
//   		jsii.String("urlInclusionPatterns"),
//   	},
//   }
//
type CfnDataSource_WebCrawlerConfigurationProperty struct {
	// Specifies the seed or starting point URLs of the websites or the sitemap URLs of the websites you want to crawl.
	//
	// You can include website subdomains. You can list up to 100 seed URLs and up to three sitemap URLs.
	//
	// You can only crawl websites that use the secure communication protocol, Hypertext Transfer Protocol Secure (HTTPS). If you receive an error when crawling a website, it could be that the website is blocked from crawling.
	//
	// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use Amazon Kendra Web Crawler to index your own webpages, or webpages that you have authorization to index.*
	Urls interface{} `field:"required" json:"urls" yaml:"urls"`
	// Configuration information required to connect to websites using authentication.
	//
	// You can connect to websites using basic authentication of user name and password.
	//
	// You must provide the website host name and port number. For example, the host name of https://a.example.com/page1.html is "a.example.com" and the port is 443, the standard port for HTTPS. You use a secret in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) to store your authentication credentials.
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Specifies the number of levels in a website that you want to crawl.
	//
	// The first level begins from the website seed or starting point URL. For example, if a website has 3 levels – index level (i.e. seed in this example), sections level, and subsections level – and you are only interested in crawling information up to the sections level (i.e. levels 0-1), you can set your depth to 1.
	//
	// The default crawl depth is set to 2.
	CrawlDepth *float64 `field:"optional" json:"crawlDepth" yaml:"crawlDepth"`
	// The maximum size (in MB) of a webpage or attachment to crawl.
	//
	// Files larger than this size (in MB) are skipped/not crawled.
	//
	// The default maximum size of a webpage or attachment is set to 50 MB.
	MaxContentSizePerPageInMegaBytes *float64 `field:"optional" json:"maxContentSizePerPageInMegaBytes" yaml:"maxContentSizePerPageInMegaBytes"`
	// The maximum number of URLs on a webpage to include when crawling a website. This number is per webpage.
	//
	// As a website’s webpages are crawled, any URLs the webpages link to are also crawled. URLs on a webpage are crawled in order of appearance.
	//
	// The default maximum links per page is 100.
	MaxLinksPerPage *float64 `field:"optional" json:"maxLinksPerPage" yaml:"maxLinksPerPage"`
	// The maximum number of URLs crawled per website host per minute.
	//
	// A minimum of one URL is required.
	//
	// The default maximum number of URLs crawled per website host per minute is 300.
	MaxUrlsPerMinuteCrawlRate *float64 `field:"optional" json:"maxUrlsPerMinuteCrawlRate" yaml:"maxUrlsPerMinuteCrawlRate"`
	// Configuration information required to connect to your internal websites via a web proxy.
	//
	// You must provide the website host name and port number. For example, the host name of https://a.example.com/page1.html is "a.example.com" and the port is 443, the standard port for HTTPS.
	//
	// Web proxy credentials are optional and you can use them to connect to a web proxy server that requires basic authentication. To store web proxy credentials, you use a secret in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) .
	ProxyConfiguration interface{} `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// A list of regular expression patterns to exclude certain URLs to crawl.
	//
	// URLs that match the patterns are excluded from the index. URLs that don't match the patterns are included in the index. If a URL matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the URL file isn't included in the index.
	UrlExclusionPatterns *[]*string `field:"optional" json:"urlExclusionPatterns" yaml:"urlExclusionPatterns"`
	// A list of regular expression patterns to include certain URLs to crawl.
	//
	// URLs that match the patterns are included in the index. URLs that don't match the patterns are excluded from the index. If a URL matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the URL file isn't included in the index.
	UrlInclusionPatterns *[]*string `field:"optional" json:"urlInclusionPatterns" yaml:"urlInclusionPatterns"`
}

