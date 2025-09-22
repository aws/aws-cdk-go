package awsbedrock


// The specific filters applied to your data source content.
//
// You can filter out or include certain content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   patternObjectFilterProperty := &PatternObjectFilterProperty{
//   	ObjectType: jsii.String("objectType"),
//
//   	// the properties below are optional
//   	ExclusionFilters: []*string{
//   		jsii.String("exclusionFilters"),
//   	},
//   	InclusionFilters: []*string{
//   		jsii.String("inclusionFilters"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-patternobjectfilter.html
//
type CfnDataSource_PatternObjectFilterProperty struct {
	// The supported object type or content type of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-patternobjectfilter.html#cfn-bedrock-datasource-patternobjectfilter-objecttype
	//
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// A list of one or more exclusion regular expression patterns to exclude certain object types that adhere to the pattern.
	//
	// If you specify an inclusion and exclusion filter/pattern and both match a document, the exclusion filter takes precedence and the document isn’t crawled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-patternobjectfilter.html#cfn-bedrock-datasource-patternobjectfilter-exclusionfilters
	//
	ExclusionFilters *[]*string `field:"optional" json:"exclusionFilters" yaml:"exclusionFilters"`
	// A list of one or more inclusion regular expression patterns to include certain object types that adhere to the pattern.
	//
	// If you specify an inclusion and exclusion filter/pattern and both match a document, the exclusion filter takes precedence and the document isn’t crawled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-patternobjectfilter.html#cfn-bedrock-datasource-patternobjectfilter-inclusionfilters
	//
	InclusionFilters *[]*string `field:"optional" json:"inclusionFilters" yaml:"inclusionFilters"`
}

