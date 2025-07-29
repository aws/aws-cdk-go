package awswisdom


// The configuration details for the web data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webCrawlerConfigurationProperty := &WebCrawlerConfigurationProperty{
//   	UrlConfiguration: &UrlConfigurationProperty{
//   		SeedUrls: []interface{}{
//   			&SeedUrlProperty{
//   				Url: jsii.String("url"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	CrawlerLimits: &CrawlerLimitsProperty{
//   		RateLimit: jsii.Number(123),
//   	},
//   	ExclusionFilters: []*string{
//   		jsii.String("exclusionFilters"),
//   	},
//   	InclusionFilters: []*string{
//   		jsii.String("inclusionFilters"),
//   	},
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html
//
type CfnKnowledgeBase_WebCrawlerConfigurationProperty struct {
	// The configuration of the URL/URLs for the web content that you want to crawl.
	//
	// You should be authorized to crawl the URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html#cfn-wisdom-knowledgebase-webcrawlerconfiguration-urlconfiguration
	//
	UrlConfiguration interface{} `field:"required" json:"urlConfiguration" yaml:"urlConfiguration"`
	// The configuration of crawl limits for the web URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html#cfn-wisdom-knowledgebase-webcrawlerconfiguration-crawlerlimits
	//
	CrawlerLimits interface{} `field:"optional" json:"crawlerLimits" yaml:"crawlerLimits"`
	// A list of one or more exclusion regular expression patterns to exclude certain URLs.
	//
	// If you specify an inclusion and exclusion filter/pattern and both match a URL, the exclusion filter takes precedence and the web content of the URL isn’t crawled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html#cfn-wisdom-knowledgebase-webcrawlerconfiguration-exclusionfilters
	//
	ExclusionFilters *[]*string `field:"optional" json:"exclusionFilters" yaml:"exclusionFilters"`
	// A list of one or more inclusion regular expression patterns to include certain URLs.
	//
	// If you specify an inclusion and exclusion filter/pattern and both match a URL, the exclusion filter takes precedence and the web content of the URL isn’t crawled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html#cfn-wisdom-knowledgebase-webcrawlerconfiguration-inclusionfilters
	//
	InclusionFilters *[]*string `field:"optional" json:"inclusionFilters" yaml:"inclusionFilters"`
	// The scope of what is crawled for your URLs.
	//
	// You can choose to crawl only web pages that belong to the same host or primary domain. For example, only web pages that contain the seed URL `https://docs.aws.amazon.com/bedrock/latest/userguide/` and no other domains. You can choose to include sub domains in addition to the host or primary domain. For example, web pages that contain `aws.amazon.com` can also include sub domain `docs.aws.amazon.com` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html#cfn-wisdom-knowledgebase-webcrawlerconfiguration-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

