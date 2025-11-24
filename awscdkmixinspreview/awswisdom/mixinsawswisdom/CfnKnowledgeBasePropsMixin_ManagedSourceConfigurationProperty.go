package mixinsawswisdom


// Source configuration for managed resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedSourceConfigurationProperty := &ManagedSourceConfigurationProperty{
//   	WebCrawlerConfiguration: &WebCrawlerConfigurationProperty{
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
//   		UrlConfiguration: &UrlConfigurationProperty{
//   			SeedUrls: []interface{}{
//   				&SeedUrlProperty{
//   					Url: jsii.String("url"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-managedsourceconfiguration.html
//
type CfnKnowledgeBasePropsMixin_ManagedSourceConfigurationProperty struct {
	// Configuration data for web crawler data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-managedsourceconfiguration.html#cfn-wisdom-knowledgebase-managedsourceconfiguration-webcrawlerconfiguration
	//
	WebCrawlerConfiguration interface{} `field:"optional" json:"webCrawlerConfiguration" yaml:"webCrawlerConfiguration"`
}

