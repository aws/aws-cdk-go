package awsconnect


// The chat entry point parameters for the test case.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   chatEntryPointParametersProperty := &ChatEntryPointParametersProperty{
//   	FlowId: jsii.String("flowId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-testcase-chatentrypointparameters.html
//
type CfnTestCase_ChatEntryPointParametersProperty struct {
	// The flow id used for the TestCase.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-testcase-chatentrypointparameters.html#cfn-connect-testcase-chatentrypointparameters-flowid
	//
	FlowId *string `field:"optional" json:"flowId" yaml:"flowId"`
}

