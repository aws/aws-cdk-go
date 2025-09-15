package awsbedrock


// The configuration of filtering the data source content.
//
// For example, configuring regular expression patterns to include or exclude certain content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crawlFilterConfigurationProperty := &CrawlFilterConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   		Filters: []interface{}{
//   			&PatternObjectFilterProperty{
//   				ObjectType: jsii.String("objectType"),
//
//   				// the properties below are optional
//   				ExclusionFilters: []*string{
//   					jsii.String("exclusionFilters"),
//   				},
//   				InclusionFilters: []*string{
//   					jsii.String("inclusionFilters"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-crawlfilterconfiguration.html
//
type CfnDataSource_CrawlFilterConfigurationProperty struct {
	// The type of filtering that you want to apply to certain objects or content of the data source.
	//
	// For example, the `PATTERN` type is regular expression patterns you can apply to filter your content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-crawlfilterconfiguration.html#cfn-bedrock-datasource-crawlfilterconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration of filtering certain objects or content types of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-crawlfilterconfiguration.html#cfn-bedrock-datasource-crawlfilterconfiguration-patternobjectfilter
	//
	PatternObjectFilter interface{} `field:"optional" json:"patternObjectFilter" yaml:"patternObjectFilter"`
}

