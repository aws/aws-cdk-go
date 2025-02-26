package awsbedrock


// The configuration information to connect to Confluence as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceDataSourceConfigurationProperty := &ConfluenceDataSourceConfigurationProperty{
//   	SourceConfiguration: &ConfluenceSourceConfigurationProperty{
//   		AuthType: jsii.String("authType"),
//   		CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   		HostType: jsii.String("hostType"),
//   		HostUrl: jsii.String("hostUrl"),
//   	},
//
//   	// the properties below are optional
//   	CrawlerConfiguration: &ConfluenceCrawlerConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-confluencedatasourceconfiguration.html
//
type CfnDataSource_ConfluenceDataSourceConfigurationProperty struct {
	// The endpoint information to connect to your Confluence data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-confluencedatasourceconfiguration.html#cfn-bedrock-datasource-confluencedatasourceconfiguration-sourceconfiguration
	//
	SourceConfiguration interface{} `field:"required" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// The configuration of the Confluence content.
	//
	// For example, configuring specific types of Confluence content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-confluencedatasourceconfiguration.html#cfn-bedrock-datasource-confluencedatasourceconfiguration-crawlerconfiguration
	//
	CrawlerConfiguration interface{} `field:"optional" json:"crawlerConfiguration" yaml:"crawlerConfiguration"`
}

