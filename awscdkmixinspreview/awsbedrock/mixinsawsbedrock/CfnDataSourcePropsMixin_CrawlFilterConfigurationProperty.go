package mixinsawsbedrock


// The configuration of filtering the data source content.
//
// For example, configuring regular expression patterns to include or exclude certain content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   crawlFilterConfigurationProperty := &CrawlFilterConfigurationProperty{
//   	PatternObjectFilter: &PatternObjectFilterConfigurationProperty{
//   		Filters: []interface{}{
//   			&PatternObjectFilterProperty{
//   				ExclusionFilters: []*string{
//   					jsii.String("exclusionFilters"),
//   				},
//   				InclusionFilters: []*string{
//   					jsii.String("inclusionFilters"),
//   				},
//   				ObjectType: jsii.String("objectType"),
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-crawlfilterconfiguration.html
//
type CfnDataSourcePropsMixin_CrawlFilterConfigurationProperty struct {
	// The configuration of filtering certain objects or content types of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-crawlfilterconfiguration.html#cfn-bedrock-datasource-crawlfilterconfiguration-patternobjectfilter
	//
	PatternObjectFilter interface{} `field:"optional" json:"patternObjectFilter" yaml:"patternObjectFilter"`
	// The type of filtering that you want to apply to certain objects or content of the data source.
	//
	// For example, the `PATTERN` type is regular expression patterns you can apply to filter your content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-crawlfilterconfiguration.html#cfn-bedrock-datasource-crawlfilterconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

