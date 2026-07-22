package awsconnect


// The Entry Point associated with the test case.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   entryPointProperty := &EntryPointProperty{
//   	ChatEntryPointParameters: &ChatEntryPointParametersProperty{
//   		FlowId: jsii.String("flowId"),
//   	},
//   	Type: jsii.String("type"),
//   	VoiceCallEntryPointParameters: &VoiceCallEntryPointParametersProperty{
//   		DestinationPhoneNumber: jsii.String("destinationPhoneNumber"),
//   		FlowId: jsii.String("flowId"),
//   		SourcePhoneNumber: jsii.String("sourcePhoneNumber"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-testcase-entrypoint.html
//
type CfnTestCasePropsMixin_EntryPointProperty struct {
	// The chat entry point parameters for the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-testcase-entrypoint.html#cfn-connect-testcase-entrypoint-chatentrypointparameters
	//
	ChatEntryPointParameters interface{} `field:"optional" json:"chatEntryPointParameters" yaml:"chatEntryPointParameters"`
	// The type of the Entry Point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-testcase-entrypoint.html#cfn-connect-testcase-entrypoint-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The voice call entry point parameters for the test case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-testcase-entrypoint.html#cfn-connect-testcase-entrypoint-voicecallentrypointparameters
	//
	VoiceCallEntryPointParameters interface{} `field:"optional" json:"voiceCallEntryPointParameters" yaml:"voiceCallEntryPointParameters"`
}

