package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentVariableProperty := &EnvironmentVariableProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-environmentvariable.html
//
type CfnMicrovmImage_EnvironmentVariableProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-environmentvariable.html#cfn-lambda-microvmimage-environmentvariable-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-environmentvariable.html#cfn-lambda-microvmimage-environmentvariable-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

