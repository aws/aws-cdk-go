package awsbedrock


// The configuration of filtering certain objects or content types of the data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   patternObjectFilterConfigurationProperty := &PatternObjectFilterConfigurationProperty{
//   	Filters: []interface{}{
//   		&PatternObjectFilterProperty{
//   			ObjectType: jsii.String("objectType"),
//
//   			// the properties below are optional
//   			ExclusionFilters: []*string{
//   				jsii.String("exclusionFilters"),
//   			},
//   			InclusionFilters: []*string{
//   				jsii.String("inclusionFilters"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-patternobjectfilterconfiguration.html
//
type CfnDataSource_PatternObjectFilterConfigurationProperty struct {
	// The configuration of specific filters applied to your data source content.
	//
	// You can filter out or include certain content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-patternobjectfilterconfiguration.html#cfn-bedrock-datasource-patternobjectfilterconfiguration-filters
	//
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
}

