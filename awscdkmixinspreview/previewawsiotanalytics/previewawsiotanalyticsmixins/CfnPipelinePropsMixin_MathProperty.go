package previewawsiotanalyticsmixins


// An activity that computes an arithmetic expression using the message's attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mathProperty := &MathProperty{
//   	Attribute: jsii.String("attribute"),
//   	Math: jsii.String("math"),
//   	Name: jsii.String("name"),
//   	Next: jsii.String("next"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html
//
type CfnPipelinePropsMixin_MathProperty struct {
	// The name of the attribute that contains the result of the math operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html#cfn-iotanalytics-pipeline-math-attribute
	//
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// An expression that uses one or more existing attributes and must return an integer value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html#cfn-iotanalytics-pipeline-math-math
	//
	Math *string `field:"optional" json:"math" yaml:"math"`
	// The name of the 'math' activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html#cfn-iotanalytics-pipeline-math-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-math.html#cfn-iotanalytics-pipeline-math-next
	//
	Next *string `field:"optional" json:"next" yaml:"next"`
}

