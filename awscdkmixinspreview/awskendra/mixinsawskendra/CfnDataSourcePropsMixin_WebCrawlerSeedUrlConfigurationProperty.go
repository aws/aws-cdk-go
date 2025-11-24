package mixinsawskendra


// Provides the configuration information of the seed or starting point URLs to crawl.
//
// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index your own webpages, or webpages that you have authorization to index.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   webCrawlerSeedUrlConfigurationProperty := &WebCrawlerSeedUrlConfigurationProperty{
//   	SeedUrls: []*string{
//   		jsii.String("seedUrls"),
//   	},
//   	WebCrawlerMode: jsii.String("webCrawlerMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerseedurlconfiguration.html
//
type CfnDataSourcePropsMixin_WebCrawlerSeedUrlConfigurationProperty struct {
	// The list of seed or starting point URLs of the websites you want to crawl.
	//
	// The list can include a maximum of 100 seed URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerseedurlconfiguration.html#cfn-kendra-datasource-webcrawlerseedurlconfiguration-seedurls
	//
	SeedUrls *[]*string `field:"optional" json:"seedUrls" yaml:"seedUrls"`
	// You can choose one of the following modes:.
	//
	// - `HOST_ONLY` —crawl only the website host names. For example, if the seed URL is "abc.example.com", then only URLs with host name "abc.example.com" are crawled.
	// - `SUBDOMAINS` —crawl the website host names with subdomains. For example, if the seed URL is "abc.example.com", then "a.abc.example.com" and "b.abc.example.com" are also crawled.
	// - `EVERYTHING` —crawl the website host names with subdomains and other domains that the web pages link to.
	//
	// The default mode is set to `HOST_ONLY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-webcrawlerseedurlconfiguration.html#cfn-kendra-datasource-webcrawlerseedurlconfiguration-webcrawlermode
	//
	WebCrawlerMode *string `field:"optional" json:"webCrawlerMode" yaml:"webCrawlerMode"`
}

