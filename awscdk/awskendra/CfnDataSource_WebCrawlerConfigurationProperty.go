package awskendra


// Provides the configuration information required for Amazon Kendra Web Crawler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerConfigurationProperty := &WebCrawlerConfigurationProperty{
//   	Urls: &WebCrawlerUrlsProperty{
//   		SeedUrlConfiguration: &WebCrawlerSeedUrlConfigurationProperty{
//   			SeedUrls: []*string{
//   				jsii.String("seedUrls"),
//   			},
//
//   			// the properties below are optional
//   			WebCrawlerMode: jsii.String("webCrawlerMode"),
//   		},
//   		SiteMapsConfiguration: &WebCrawlerSiteMapsConfigurationProperty{
//   			SiteMaps: []*string{
//   				jsii.String("siteMaps"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	AuthenticationConfiguration: &WebCrawlerAuthenticationConfigurationProperty{
//   		BasicAuthentication: []interface{}{
//   			&WebCrawlerBasicAuthenticationProperty{
//   				Credentials: jsii.String("credentials"),
//   				Host: jsii.String("host"),
//   				Port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	CrawlDepth: jsii.Number(123),
//   	MaxContentSizePerPageInMegaBytes: jsii.Number(123),
//   	MaxLinksPerPage: jsii.Number(123),
//   	MaxUrlsPerMinuteCrawlRate: jsii.Number(123),
//   	ProxyConfiguration: &ProxyConfigurationProperty{
//   		Host: jsii.String("host"),
//   		Port: jsii.Number(123),
//
//   		// the properties below are optional
//   		Credentials: jsii.String("credentials"),
//   	},
//   	UrlExclusionPatterns: []*string{
//   		jsii.String("urlExclusionPatterns"),
//   	},
//   	UrlInclusionPatterns: []*string{
//   		jsii.String("urlInclusionPatterns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerconfiguration.html
//
type CfnDataSource_WebCrawlerConfigurationProperty struct {
	// Specifies the seed or starting point URLs of the websites or the sitemap URLs of the websites you want to crawl.
	//
	// You can include website subdomains. You can list up to 100 seed URLs and up to three sitemap URLs.
	//
	// You can only crawl websites that use the secure communication protocol, Hypertext Transfer Protocol Secure (HTTPS). If you receive an error when crawling a website, it could be that the website is blocked from crawling.
	//
	// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use Amazon Kendra Web Crawler to index your own webpages, or webpages that you have authorization to index.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerconfiguration.html#cfn-kendra-datasource-webcrawlerconfiguration-urls
	//
	Urls interface{} `field:"required" json:"urls" yaml:"urls"`
	// Configuration information required to connect to websites using authentication.
	//
	// You can connect to websites using basic authentication of user name and password. You use a secret in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) to store your authentication credentials.
	//
	// You must provide the website host name and port number. For example, the host name of https://a.example.com/page1.html is "a.example.com" and the port is 443, the standard port for HTTPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerconfiguration.html#cfn-kendra-datasource-webcrawlerconfiguration-authenticationconfiguration
	//
	AuthenticationConfiguration interface{} `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// The 'depth' or number of levels from the seed level to crawl.
	//
	// For example, the seed URL page is depth 1 and any hyperlinks on this page that are also crawled are depth 2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerconfiguration.html#cfn-kendra-datasource-webcrawlerconfiguration-crawldepth
	//
	CrawlDepth *float64 `field:"optional" json:"crawlDepth" yaml:"crawlDepth"`
	// The maximum size (in MB) of a web page or attachment to crawl.
	//
	// Files larger than this size (in MB) are skipped/not crawled.
	//
	// The default maximum size of a web page or attachment is set to 50 MB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerconfiguration.html#cfn-kendra-datasource-webcrawlerconfiguration-maxcontentsizeperpageinmegabytes
	//
	MaxContentSizePerPageInMegaBytes *float64 `field:"optional" json:"maxContentSizePerPageInMegaBytes" yaml:"maxContentSizePerPageInMegaBytes"`
	// The maximum number of URLs on a web page to include when crawling a website.
	//
	// This number is per web page.
	//
	// As a website’s web pages are crawled, any URLs the web pages link to are also crawled. URLs on a web page are crawled in order of appearance.
	//
	// The default maximum links per page is 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerconfiguration.html#cfn-kendra-datasource-webcrawlerconfiguration-maxlinksperpage
	//
	MaxLinksPerPage *float64 `field:"optional" json:"maxLinksPerPage" yaml:"maxLinksPerPage"`
	// The maximum number of URLs crawled per website host per minute.
	//
	// A minimum of one URL is required.
	//
	// The default maximum number of URLs crawled per website host per minute is 300.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerconfiguration.html#cfn-kendra-datasource-webcrawlerconfiguration-maxurlsperminutecrawlrate
	//
	MaxUrlsPerMinuteCrawlRate *float64 `field:"optional" json:"maxUrlsPerMinuteCrawlRate" yaml:"maxUrlsPerMinuteCrawlRate"`
	// Configuration information required to connect to your internal websites via a web proxy.
	//
	// You must provide the website host name and port number. For example, the host name of https://a.example.com/page1.html is "a.example.com" and the port is 443, the standard port for HTTPS.
	//
	// Web proxy credentials are optional and you can use them to connect to a web proxy server that requires basic authentication. To store web proxy credentials, you use a secret in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerconfiguration.html#cfn-kendra-datasource-webcrawlerconfiguration-proxyconfiguration
	//
	ProxyConfiguration interface{} `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// A list of regular expression patterns to exclude certain URLs to crawl.
	//
	// URLs that match the patterns are excluded from the index. URLs that don't match the patterns are included in the index. If a URL matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the URL file isn't included in the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerconfiguration.html#cfn-kendra-datasource-webcrawlerconfiguration-urlexclusionpatterns
	//
	UrlExclusionPatterns *[]*string `field:"optional" json:"urlExclusionPatterns" yaml:"urlExclusionPatterns"`
	// A list of regular expression patterns to include certain URLs to crawl.
	//
	// URLs that match the patterns are included in the index. URLs that don't match the patterns are excluded from the index. If a URL matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the URL file isn't included in the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerconfiguration.html#cfn-kendra-datasource-webcrawlerconfiguration-urlinclusionpatterns
	//
	UrlInclusionPatterns *[]*string `field:"optional" json:"urlInclusionPatterns" yaml:"urlInclusionPatterns"`
}

