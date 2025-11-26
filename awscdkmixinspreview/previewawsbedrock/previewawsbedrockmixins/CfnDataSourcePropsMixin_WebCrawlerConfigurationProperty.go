package previewawsbedrockmixins


// The configuration of web URLs that you want to crawl.
//
// You should be authorized to crawl the URLs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   webCrawlerConfigurationProperty := &WebCrawlerConfigurationProperty{
//   	CrawlerLimits: &WebCrawlerLimitsProperty{
//   		MaxPages: jsii.Number(123),
//   		RateLimit: jsii.Number(123),
//   	},
//   	ExclusionFilters: []*string{
//   		jsii.String("exclusionFilters"),
//   	},
//   	InclusionFilters: []*string{
//   		jsii.String("inclusionFilters"),
//   	},
//   	Scope: jsii.String("scope"),
//   	UserAgent: jsii.String("userAgent"),
//   	UserAgentHeader: jsii.String("userAgentHeader"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webcrawlerconfiguration.html
//
type CfnDataSourcePropsMixin_WebCrawlerConfigurationProperty struct {
	// The configuration of crawl limits for the web URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webcrawlerconfiguration.html#cfn-bedrock-datasource-webcrawlerconfiguration-crawlerlimits
	//
	CrawlerLimits interface{} `field:"optional" json:"crawlerLimits" yaml:"crawlerLimits"`
	// A list of one or more exclusion regular expression patterns to exclude certain URLs.
	//
	// If you specify an inclusion and exclusion filter/pattern and both match a URL, the exclusion filter takes precedence and the web content of the URL isn’t crawled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webcrawlerconfiguration.html#cfn-bedrock-datasource-webcrawlerconfiguration-exclusionfilters
	//
	ExclusionFilters *[]*string `field:"optional" json:"exclusionFilters" yaml:"exclusionFilters"`
	// A list of one or more inclusion regular expression patterns to include certain URLs.
	//
	// If you specify an inclusion and exclusion filter/pattern and both match a URL, the exclusion filter takes precedence and the web content of the URL isn’t crawled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webcrawlerconfiguration.html#cfn-bedrock-datasource-webcrawlerconfiguration-inclusionfilters
	//
	InclusionFilters *[]*string `field:"optional" json:"inclusionFilters" yaml:"inclusionFilters"`
	// The scope of what is crawled for your URLs.
	//
	// You can choose to crawl only web pages that belong to the same host or primary domain. For example, only web pages that contain the seed URL "https://docs.aws.amazon.com/bedrock/latest/userguide/" and no other domains. You can choose to include sub domains in addition to the host or primary domain. For example, web pages that contain "aws.amazon.com" can also include sub domain "docs.aws.amazon.com".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webcrawlerconfiguration.html#cfn-bedrock-datasource-webcrawlerconfiguration-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Returns the user agent suffix for your web crawler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webcrawlerconfiguration.html#cfn-bedrock-datasource-webcrawlerconfiguration-useragent
	//
	UserAgent *string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// A string used for identifying the crawler or bot when it accesses a web server.
	//
	// The user agent header value consists of the `bedrockbot` , UUID, and a user agent suffix for your crawler (if one is provided). By default, it is set to `bedrockbot_UUID` . You can optionally append a custom suffix to `bedrockbot_UUID` to allowlist a specific user agent permitted to access your source URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-webcrawlerconfiguration.html#cfn-bedrock-datasource-webcrawlerconfiguration-useragentheader
	//
	UserAgentHeader *string `field:"optional" json:"userAgentHeader" yaml:"userAgentHeader"`
}

