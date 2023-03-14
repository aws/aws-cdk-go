package awskendra


// Provides the configuration information of the seed or starting point URLs to crawl.
//
// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index your own webpages, or webpages that you have authorization to index.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerSeedUrlConfigurationProperty := &webCrawlerSeedUrlConfigurationProperty{
//   	seedUrls: []*string{
//   		jsii.String("seedUrls"),
//   	},
//
//   	// the properties below are optional
//   	webCrawlerMode: jsii.String("webCrawlerMode"),
//   }
//
type CfnDataSource_WebCrawlerSeedUrlConfigurationProperty struct {
	// The list of seed or starting point URLs of the websites you want to crawl.
	//
	// The list can include a maximum of 100 seed URLs.
	SeedUrls *[]*string `field:"required" json:"seedUrls" yaml:"seedUrls"`
	// You can choose one of the following modes:.
	//
	// - `HOST_ONLY` – crawl only the website host names. For example, if the seed URL is "abc.example.com", then only URLs with host name "abc.example.com" are crawled.
	// - `SUBDOMAINS` – crawl the website host names with subdomains. For example, if the seed URL is "abc.example.com", then "a.abc.example.com" and "b.abc.example.com" are also crawled.
	// - `EVERYTHING` – crawl the website host names with subdomains and other domains that the webpages link to.
	//
	// The default mode is set to `HOST_ONLY` .
	WebCrawlerMode *string `field:"optional" json:"webCrawlerMode" yaml:"webCrawlerMode"`
}

