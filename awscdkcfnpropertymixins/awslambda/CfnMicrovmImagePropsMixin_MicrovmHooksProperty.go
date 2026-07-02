package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   microvmHooksProperty := &MicrovmHooksProperty{
//   	Resume: jsii.String("resume"),
//   	ResumeTimeoutInSeconds: jsii.Number(123),
//   	Run: jsii.String("run"),
//   	RunTimeoutInSeconds: jsii.Number(123),
//   	Suspend: jsii.String("suspend"),
//   	SuspendTimeoutInSeconds: jsii.Number(123),
//   	Terminate: jsii.String("terminate"),
//   	TerminateTimeoutInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmhooks.html
//
type CfnMicrovmImagePropsMixin_MicrovmHooksProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmhooks.html#cfn-lambda-microvmimage-microvmhooks-resume
	//
	Resume *string `field:"optional" json:"resume" yaml:"resume"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmhooks.html#cfn-lambda-microvmimage-microvmhooks-resumetimeoutinseconds
	//
	ResumeTimeoutInSeconds *float64 `field:"optional" json:"resumeTimeoutInSeconds" yaml:"resumeTimeoutInSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmhooks.html#cfn-lambda-microvmimage-microvmhooks-run
	//
	Run *string `field:"optional" json:"run" yaml:"run"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmhooks.html#cfn-lambda-microvmimage-microvmhooks-runtimeoutinseconds
	//
	RunTimeoutInSeconds *float64 `field:"optional" json:"runTimeoutInSeconds" yaml:"runTimeoutInSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmhooks.html#cfn-lambda-microvmimage-microvmhooks-suspend
	//
	Suspend *string `field:"optional" json:"suspend" yaml:"suspend"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmhooks.html#cfn-lambda-microvmimage-microvmhooks-suspendtimeoutinseconds
	//
	SuspendTimeoutInSeconds *float64 `field:"optional" json:"suspendTimeoutInSeconds" yaml:"suspendTimeoutInSeconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmhooks.html#cfn-lambda-microvmimage-microvmhooks-terminate
	//
	Terminate *string `field:"optional" json:"terminate" yaml:"terminate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-microvmhooks.html#cfn-lambda-microvmimage-microvmhooks-terminatetimeoutinseconds
	//
	TerminateTimeoutInSeconds *float64 `field:"optional" json:"terminateTimeoutInSeconds" yaml:"terminateTimeoutInSeconds"`
}

