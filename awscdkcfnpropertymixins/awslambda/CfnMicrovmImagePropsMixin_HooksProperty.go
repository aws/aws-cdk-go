package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   hooksProperty := &HooksProperty{
//   	MicrovmHooks: &MicrovmHooksProperty{
//   		Resume: jsii.String("resume"),
//   		ResumeTimeoutInSeconds: jsii.Number(123),
//   		Run: jsii.String("run"),
//   		RunTimeoutInSeconds: jsii.Number(123),
//   		Suspend: jsii.String("suspend"),
//   		SuspendTimeoutInSeconds: jsii.Number(123),
//   		Terminate: jsii.String("terminate"),
//   		TerminateTimeoutInSeconds: jsii.Number(123),
//   	},
//   	MicrovmImageHooks: &MicrovmImageHooksProperty{
//   		Ready: jsii.String("ready"),
//   		ReadyTimeoutInSeconds: jsii.Number(123),
//   		Validate: jsii.String("validate"),
//   		ValidateTimeoutInSeconds: jsii.Number(123),
//   	},
//   	Port: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-hooks.html
//
type CfnMicrovmImagePropsMixin_HooksProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-hooks.html#cfn-lambda-microvmimage-hooks-microvmhooks
	//
	MicrovmHooks interface{} `field:"optional" json:"microvmHooks" yaml:"microvmHooks"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-hooks.html#cfn-lambda-microvmimage-hooks-microvmimagehooks
	//
	MicrovmImageHooks interface{} `field:"optional" json:"microvmImageHooks" yaml:"microvmImageHooks"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-hooks.html#cfn-lambda-microvmimage-hooks-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

