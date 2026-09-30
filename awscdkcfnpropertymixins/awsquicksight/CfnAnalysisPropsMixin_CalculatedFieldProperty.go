package awsquicksight


// The calculated field of an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   calculatedFieldProperty := &CalculatedFieldProperty{
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	Expression: jsii.String("expression"),
//   	Name: jsii.String("name"),
//   	TopicIdentifier: jsii.String("topicIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-calculatedfield.html
//
type CfnAnalysisPropsMixin_CalculatedFieldProperty struct {
	// The data set that is used in this calculated field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-calculatedfield.html#cfn-quicksight-analysis-calculatedfield-datasetidentifier
	//
	// Default: - "".
	//
	DataSetIdentifier *string `field:"optional" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// The expression of the calculated field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-calculatedfield.html#cfn-quicksight-analysis-calculatedfield-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The name of the calculated field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-calculatedfield.html#cfn-quicksight-analysis-calculatedfield-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-calculatedfield.html#cfn-quicksight-analysis-calculatedfield-topicidentifier
	//
	TopicIdentifier *string `field:"optional" json:"topicIdentifier" yaml:"topicIdentifier"`
}

