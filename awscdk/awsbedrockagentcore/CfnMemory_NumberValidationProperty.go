package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberValidationProperty := &NumberValidationProperty{
//   	MaxValue: jsii.Number(123),
//   	MinValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-numbervalidation.html
//
type CfnMemory_NumberValidationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-numbervalidation.html#cfn-bedrockagentcore-memory-numbervalidation-maxvalue
	//
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-numbervalidation.html#cfn-bedrockagentcore-memory-numbervalidation-minvalue
	//
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
}

