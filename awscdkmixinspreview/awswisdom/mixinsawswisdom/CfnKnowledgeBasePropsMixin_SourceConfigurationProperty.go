package mixinsawswisdom


// Configuration information about the external data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceConfigurationProperty := &SourceConfigurationProperty{
//   	AppIntegrations: &AppIntegrationsConfigurationProperty{
//   		AppIntegrationArn: jsii.String("appIntegrationArn"),
//   		ObjectFields: []*string{
//   			jsii.String("objectFields"),
//   		},
//   	},
//   	ManagedSourceConfiguration: &ManagedSourceConfigurationProperty{
//   		WebCrawlerConfiguration: &WebCrawlerConfigurationProperty{
//   			CrawlerLimits: &CrawlerLimitsProperty{
//   				RateLimit: jsii.Number(123),
//   			},
//   			ExclusionFilters: []*string{
//   				jsii.String("exclusionFilters"),
//   			},
//   			InclusionFilters: []*string{
//   				jsii.String("inclusionFilters"),
//   			},
//   			Scope: jsii.String("scope"),
//   			UrlConfiguration: &UrlConfigurationProperty{
//   				SeedUrls: []interface{}{
//   					&SeedUrlProperty{
//   						Url: jsii.String("url"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-sourceconfiguration.html
//
type CfnKnowledgeBasePropsMixin_SourceConfigurationProperty struct {
	// Configuration information for Amazon AppIntegrations to automatically ingest content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-sourceconfiguration.html#cfn-wisdom-knowledgebase-sourceconfiguration-appintegrations
	//
	AppIntegrations interface{} `field:"optional" json:"appIntegrations" yaml:"appIntegrations"`
	// Source configuration for managed resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-sourceconfiguration.html#cfn-wisdom-knowledgebase-sourceconfiguration-managedsourceconfiguration
	//
	ManagedSourceConfiguration interface{} `field:"optional" json:"managedSourceConfiguration" yaml:"managedSourceConfiguration"`
}

