package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   microvmImageHooksProperty := &MicrovmImageHooksProperty{
//   	Ready: jsii.String("ready"),
//   	ReadyTimeoutInSeconds: jsii.Number(123),
//   	Validate: jsii.String("validate"),
//   	ValidateTimeoutInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmimagehooks.html
//
type CfnMicrovmImage_MicrovmImageHooksProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmimagehooks.html#cfn-lambda-microvmimage-microvmimagehooks-ready
	//
	Ready *string `field:"optional" json:"ready" yaml:"ready"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmimagehooks.html#cfn-lambda-microvmimage-microvmimagehooks-readytimeoutinseconds
	//
	ReadyTimeoutInSeconds *float64 `field:"optional" json:"readyTimeoutInSeconds" yaml:"readyTimeoutInSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmimagehooks.html#cfn-lambda-microvmimage-microvmimagehooks-validate
	//
	Validate *string `field:"optional" json:"validate" yaml:"validate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmimagehooks.html#cfn-lambda-microvmimage-microvmimagehooks-validatetimeoutinseconds
	//
	ValidateTimeoutInSeconds *float64 `field:"optional" json:"validateTimeoutInSeconds" yaml:"validateTimeoutInSeconds"`
}

