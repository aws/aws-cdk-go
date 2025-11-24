package mixinsawsbedrock


// The configuration of the Salesforce content.
//
// For example, configuring specific types of Salesforce content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   salesforceCrawlerConfigurationProperty := &SalesforceCrawlerConfigurationProperty{
//   	FilterConfiguration: &CrawlFilterConfigurationProperty{
//   		PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   			Filters: []interface{}{
//   				&PatternObjectFilterProperty{
//   					ExclusionFilters: []*string{
//   						jsii.String("exclusionFilters"),
//   					},
//   					InclusionFilters: []*string{
//   						jsii.String("inclusionFilters"),
//   					},
//   					ObjectType: jsii.String("objectType"),
//   				},
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcecrawlerconfiguration.html
//
type CfnDataSourcePropsMixin_SalesforceCrawlerConfigurationProperty struct {
	// The configuration of filtering the Salesforce content.
	//
	// For example, configuring regular expression patterns to include or exclude certain content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-salesforcecrawlerconfiguration.html#cfn-bedrock-datasource-salesforcecrawlerconfiguration-filterconfiguration
	//
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
}

