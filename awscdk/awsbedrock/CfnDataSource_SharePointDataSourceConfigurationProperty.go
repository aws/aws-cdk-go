package awsbedrock


// The configuration information to connect to SharePoint as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharePointDataSourceConfigurationProperty := &SharePointDataSourceConfigurationProperty{
//   	SourceConfiguration: &SharePointSourceConfigurationProperty{
//   		AuthType: jsii.String("authType"),
//   		CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   		Domain: jsii.String("domain"),
//   		HostType: jsii.String("hostType"),
//   		SiteUrls: []*string{
//   			jsii.String("siteUrls"),
//   		},
//
//   		// the properties below are optional
//   		TenantId: jsii.String("tenantId"),
//   	},
//
//   	// the properties below are optional
//   	CrawlerConfiguration: &SharePointCrawlerConfigurationProperty{
//   		FilterConfiguration: &CrawlFilterConfigurationProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   				Filters: []interface{}{
//   					&PatternObjectFilterProperty{
//   						ObjectType: jsii.String("objectType"),
//
//   						// the properties below are optional
//   						ExclusionFilters: []*string{
//   							jsii.String("exclusionFilters"),
//   						},
//   						InclusionFilters: []*string{
//   							jsii.String("inclusionFilters"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointdatasourceconfiguration.html
//
type CfnDataSource_SharePointDataSourceConfigurationProperty struct {
	// The endpoint information to connect to your SharePoint data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointdatasourceconfiguration.html#cfn-bedrock-datasource-sharepointdatasourceconfiguration-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"required" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The configuration of the SharePoint content.
	//
	// For example, configuring specific types of SharePoint content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointdatasourceconfiguration.html#cfn-bedrock-datasource-sharepointdatasourceconfiguration-crawlerconfiguration
	//
	CrawlerConfiguration interface{} `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
}

