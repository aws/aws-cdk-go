package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   toolOutputConfigurationProperty := &ToolOutputConfigurationProperty{
//   	OutputVariableNameOverride: jsii.String("outputVariableNameOverride"),
//   	SessionDataNamespace: jsii.String("sessionDataNamespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloutputconfiguration.html
//
type CfnAIAgent_ToolOutputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloutputconfiguration.html#cfn-wisdom-aiagent-tooloutputconfiguration-outputvariablenameoverride
	//
	OutputVariableNameOverride *string `field:"optional" json:"outputVariableNameOverride" yaml:"outputVariableNameOverride"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloutputconfiguration.html#cfn-wisdom-aiagent-tooloutputconfiguration-sessiondatanamespace
	//
	SessionDataNamespace *string `field:"optional" json:"sessionDataNamespace" yaml:"sessionDataNamespace"`
}

