package mixinsawsbedrock


// The specific filters applied to your data source content.
//
// You can filter out or include certain content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   patternObjectFilterProperty := &PatternObjectFilterProperty{
//   	ExclusionFilters: []*string{
//   		jsii.String("exclusionFilters"),
//   	},
//   	InclusionFilters: []*string{
//   		jsii.String("inclusionFilters"),
//   	},
//   	ObjectType: jsii.String("objectType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-patternobjectfilter.html
//
type CfnDataSourcePropsMixin_PatternObjectFilterProperty struct {
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
	// The supported object type or content type of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-patternobjectfilter.html#cfn-bedrock-datasource-patternobjectfilter-objecttype
	//
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
}

