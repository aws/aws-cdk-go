package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   environmentVariableProperty := &EnvironmentVariableProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-environmentvariable.html
//
type CfnMicrovmImagePropsMixin_EnvironmentVariableProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-environmentvariable.html#cfn-lambda-microvmimage-environmentvariable-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-environmentvariable.html#cfn-lambda-microvmimage-environmentvariable-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

