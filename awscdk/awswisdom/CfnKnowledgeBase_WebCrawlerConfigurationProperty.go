package awswisdom


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html#cfn-wisdom-knowledgebase-webcrawlerconfiguration-urlconfiguration
	//
	UrlConfiguration interface{} `field:"required" json:"urlConfiguration" yaml:"urlConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html#cfn-wisdom-knowledgebase-webcrawlerconfiguration-crawlerlimits
	//
	CrawlerLimits interface{} `field:"optional" json:"crawlerLimits" yaml:"crawlerLimits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html#cfn-wisdom-knowledgebase-webcrawlerconfiguration-exclusionfilters
	//
	ExclusionFilters *[]*string `field:"optional" json:"exclusionFilters" yaml:"exclusionFilters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html#cfn-wisdom-knowledgebase-webcrawlerconfiguration-inclusionfilters
	//
	InclusionFilters *[]*string `field:"optional" json:"inclusionFilters" yaml:"inclusionFilters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-webcrawlerconfiguration.html#cfn-wisdom-knowledgebase-webcrawlerconfiguration-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

