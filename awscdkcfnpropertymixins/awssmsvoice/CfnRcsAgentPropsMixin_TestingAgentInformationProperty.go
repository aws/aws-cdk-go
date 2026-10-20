package awssmsvoice


// Information about the testing agent (RCS for Business ID) associated with the RCS agent.
//
// A testing agent is created by submitting a testing registration, and allows sending to registered test devices without carrier approval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   testingAgentInformationProperty := &TestingAgentInformationProperty{
//   	RegistrationId: jsii.String("registrationId"),
//   	TestingAgentId: jsii.String("testingAgentId"),
//   	TestingAgentStatus: jsii.String("testingAgentStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-rcsagent-testingagentinformation.html
//
type CfnRcsAgentPropsMixin_TestingAgentInformationProperty struct {
	// The unique identifier of the testing registration that created the testing agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-rcsagent-testingagentinformation.html#cfn-smsvoice-rcsagent-testingagentinformation-registrationid
	//
	RegistrationId *string `field:"optional" json:"registrationId" yaml:"registrationId"`
	// The identifier of the testing agent assigned by the RCS infrastructure provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-rcsagent-testingagentinformation.html#cfn-smsvoice-rcsagent-testingagentinformation-testingagentid
	//
	TestingAgentId *string `field:"optional" json:"testingAgentId" yaml:"testingAgentId"`
	// The status of the testing agent.
	//
	// Named distinctly from the resource-level Status property, which has its own larger set of values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-rcsagent-testingagentinformation.html#cfn-smsvoice-rcsagent-testingagentinformation-testingagentstatus
	//
	TestingAgentStatus *string `field:"optional" json:"testingAgentStatus" yaml:"testingAgentStatus"`
}

