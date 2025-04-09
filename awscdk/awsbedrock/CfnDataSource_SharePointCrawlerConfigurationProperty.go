package awsbedrock


// The configuration of the SharePoint content.
//
// For example, configuring specific types of SharePoint content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharePointCrawlerConfigurationProperty := &SharePointCrawlerConfigurationProperty{
//   	FilterConfiguration: &CrawlFilterConfigurationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   			Filters: []interface{}{
//   				&PatternObjectFilterProperty{
//   					ObjectType: jsii.String("objectType"),
//
//   					// the properties below are optional
//   					ExclusionFilters: []*string{
//   						jsii.String("exclusionFilters"),
//   					},
//   					InclusionFilters: []*string{
//   						jsii.String("inclusionFilters"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointcrawlerconfiguration.html
//
type CfnDataSource_SharePointCrawlerConfigurationProperty struct {
	// The configuration of filtering the SharePoint content.
	//
	// For example, configuring regular expression patterns to include or exclude certain content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointcrawlerconfiguration.html#cfn-bedrock-datasource-sharepointcrawlerconfiguration-filterconfiguration
	//
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
}

