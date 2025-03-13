package awswisdom


// Source configuration for managed resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedSourceConfigurationProperty := &ManagedSourceConfigurationProperty{
//   	WebCrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   		UrlConfiguration: &UrlConfigurationProperty{
//   			SeedUrls: []interface{}{
//   				&SeedUrlProperty{
//   					Url: jsii.String("url"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		CrawlerLimits: &CrawlerLimitsProperty{
//   			RateLimit: jsii.Number(123),
//   		},
//   		ExclusionFilters: []*string{
//   			jsii.String("exclusionFilters"),
//   		},
//   		InclusionFilters: []*string{
//   			jsii.String("inclusionFilters"),
//   		},
//   		Scope: jsii.String("scope"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-managedsourceconfiguration.html
//
type CfnKnowledgeBase_ManagedSourceConfigurationProperty struct {
	// Configuration data for web crawler data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-managedsourceconfiguration.html#cfn-wisdom-knowledgebase-managedsourceconfiguration-webcrawlerconfiguration
	//
	WebCrawlerConfiguration interface{} `field:"required" json:"webCrawlerConfiguration" yaml:"webCrawlerConfiguration"`
}

