package awsconnect


// The voice call entry point parameters for the test case.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   voiceCallEntryPointParametersProperty := &VoiceCallEntryPointParametersProperty{
//   	DestinationPhoneNumber: jsii.String("destinationPhoneNumber"),
//   	FlowId: jsii.String("flowId"),
//   	SourcePhoneNumber: jsii.String("sourcePhoneNumber"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-testcase-voicecallentrypointparameters.html
//
type CfnTestCase_VoiceCallEntryPointParametersProperty struct {
	// The destination phonenumber of the EntryPoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-testcase-voicecallentrypointparameters.html#cfn-connect-testcase-voicecallentrypointparameters-destinationphonenumber
	//
	DestinationPhoneNumber *string `field:"optional" json:"destinationPhoneNumber" yaml:"destinationPhoneNumber"`
	// The flow id used for the TestCase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-testcase-voicecallentrypointparameters.html#cfn-connect-testcase-voicecallentrypointparameters-flowid
	//
	FlowId *string `field:"optional" json:"flowId" yaml:"flowId"`
	// The source phonenumber of the EntryPoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-testcase-voicecallentrypointparameters.html#cfn-connect-testcase-voicecallentrypointparameters-sourcephonenumber
	//
	SourcePhoneNumber *string `field:"optional" json:"sourcePhoneNumber" yaml:"sourcePhoneNumber"`
}

