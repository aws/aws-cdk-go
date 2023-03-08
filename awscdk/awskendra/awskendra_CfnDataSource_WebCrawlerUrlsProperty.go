package awskendra


// Specifies the seed or starting point URLs of the websites or the sitemap URLs of the websites you want to crawl.
//
// You can include website subdomains. You can list up to 100 seed URLs and up to three sitemap URLs.
//
// You can only crawl websites that use the secure communication protocol, Hypertext Transfer Protocol Secure (HTTPS). If you receive an error when crawling a website, it could be that the website is blocked from crawling.
//
// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index your own webpages, or webpages that you have authorization to index.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerUrlsProperty := &webCrawlerUrlsProperty{
//   	seedUrlConfiguration: &webCrawlerSeedUrlConfigurationProperty{
//   		seedUrls: []*string{
//   			jsii.String("seedUrls"),
//   		},
//
//   		// the properties below are optional
//   		webCrawlerMode: jsii.String("webCrawlerMode"),
//   	},
//   	siteMapsConfiguration: &webCrawlerSiteMapsConfigurationProperty{
//   		siteMaps: []*string{
//   			jsii.String("siteMaps"),
//   		},
//   	},
//   }
//
type CfnDataSource_WebCrawlerUrlsProperty struct {
	// Configuration of the seed or starting point URLs of the websites you want to crawl.
	//
	// You can choose to crawl only the website host names, or the website host names with subdomains, or the website host names with subdomains and other domains that the webpages link to.
	//
	// You can list up to 100 seed URLs.
	SeedUrlConfiguration interface{} `field:"optional" json:"seedUrlConfiguration" yaml:"seedUrlConfiguration"`
	// Configuration of the sitemap URLs of the websites you want to crawl.
	//
	// Only URLs belonging to the same website host names are crawled. You can list up to three sitemap URLs.
	SiteMapsConfiguration interface{} `field:"optional" json:"siteMapsConfiguration" yaml:"siteMapsConfiguration"`
}

