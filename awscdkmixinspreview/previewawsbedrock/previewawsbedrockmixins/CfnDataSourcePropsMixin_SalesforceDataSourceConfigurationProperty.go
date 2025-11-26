package previewawsbedrockmixins


// The configuration information to connect to Salesforce as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   salesforceDataSourceConfigurationProperty := &SalesforceDataSourceConfigurationProperty{
//   	CrawlerConfiguration: &SalesforceCrawlerConfigurationProperty{
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
//   	SourceConfiguration: &SalesforceSourceConfigurationProperty{
//   		AuthType: jsii.String("authType"),
//   		CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   		HostUrl: jsii.String("hostUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcedatasourceconfiguration.html
//
type CfnDataSourcePropsMixin_SalesforceDataSourceConfigurationProperty struct {
	// The configuration of the Salesforce content.
	//
	// For example, configuring specific types of Salesforce content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcedatasourceconfiguration.html#cfn-bedrock-datasource-salesforcedatasourceconfiguration-crawlerconfiguration
	//
	CrawlerConfiguration interface{} `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
	// The endpoint information to connect to your Salesforce data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcedatasourceconfiguration.html#cfn-bedrock-datasource-salesforcedatasourceconfiguration-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

