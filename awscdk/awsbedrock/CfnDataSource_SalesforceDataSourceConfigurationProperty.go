package awsbedrock


// The configuration information to connect to Salesforce as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceDataSourceConfigurationProperty := &SalesforceDataSourceConfigurationProperty{
//   	SourceConfiguration: &SalesforceSourceConfigurationProperty{
//   		AuthType: jsii.String("authType"),
//   		CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   		HostUrl: jsii.String("hostUrl"),
//   	},
//
//   	// the properties below are optional
//   	CrawlerConfiguration: &SalesforceCrawlerConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcedatasourceconfiguration.html
//
type CfnDataSource_SalesforceDataSourceConfigurationProperty struct {
	// The endpoint information to connect to your Salesforce data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcedatasourceconfiguration.html#cfn-bedrock-datasource-salesforcedatasourceconfiguration-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"required" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The configuration of the Salesforce content.
	//
	// For example, configuring specific types of Salesforce content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcedatasourceconfiguration.html#cfn-bedrock-datasource-salesforcedatasourceconfiguration-crawlerconfiguration
	//
	CrawlerConfiguration interface{} `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
}

