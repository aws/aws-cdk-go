package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessAgentCoreCodeInterpreterConfigProperty := &HarnessAgentCoreCodeInterpreterConfigProperty{
//   	CodeInterpreterArn: jsii.String("codeInterpreterArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorecodeinterpreterconfig.html
//
type CfnHarnessPropsMixin_HarnessAgentCoreCodeInterpreterConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessagentcorecodeinterpreterconfig.html#cfn-bedrockagentcore-harness-harnessagentcorecodeinterpreterconfig-codeinterpreterarn
	//
	CodeInterpreterArn *string `field:"optional" json:"codeInterpreterArn" yaml:"codeInterpreterArn"`
}

