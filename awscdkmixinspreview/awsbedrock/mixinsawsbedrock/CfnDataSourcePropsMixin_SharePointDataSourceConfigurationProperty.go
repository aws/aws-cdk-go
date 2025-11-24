package mixinsawsbedrock


// The configuration information to connect to SharePoint as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sharePointDataSourceConfigurationProperty := &SharePointDataSourceConfigurationProperty{
//   	CrawlerConfiguration: &SharePointCrawlerConfigurationProperty{
//   		FilterConfiguration: &CrawlFilterConfigurationProperty{
//   			PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   				Filters: []interface{}{
//   					&PatternObjectFilterProperty{
//   						ExclusionFilters: []*string{
//   							jsii.String("exclusionFilters"),
//   						},
//   						InclusionFilters: []*string{
//   							jsii.String("inclusionFilters"),
//   						},
//   						ObjectType: jsii.String("objectType"),
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SourceConfiguration: &SharePointSourceConfigurationProperty{
//   		AuthType: jsii.String("authType"),
//   		CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   		Domain: jsii.String("domain"),
//   		HostType: jsii.String("hostType"),
//   		SiteUrls: []*string{
//   			jsii.String("siteUrls"),
//   		},
//   		TenantId: jsii.String("tenantId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointdatasourceconfiguration.html
//
type CfnDataSourcePropsMixin_SharePointDataSourceConfigurationProperty struct {
	// The configuration of the SharePoint content.
	//
	// For example, configuring specific types of SharePoint content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointdatasourceconfiguration.html#cfn-bedrock-datasource-sharepointdatasourceconfiguration-crawlerconfiguration
	//
	CrawlerConfiguration interface{} `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
	// The endpoint information to connect to your SharePoint data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointdatasourceconfiguration.html#cfn-bedrock-datasource-sharepointdatasourceconfiguration-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

