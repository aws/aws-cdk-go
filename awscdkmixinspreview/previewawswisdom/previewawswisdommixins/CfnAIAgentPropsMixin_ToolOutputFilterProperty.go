package previewawswisdommixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   toolOutputFilterProperty := &ToolOutputFilterProperty{
//   	JsonPath: jsii.String("jsonPath"),
//   	OutputConfiguration: &ToolOutputConfigurationProperty{
//   		OutputVariableNameOverride: jsii.String("outputVariableNameOverride"),
//   		SessionDataNamespace: jsii.String("sessionDataNamespace"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloutputfilter.html
//
type CfnAIAgentPropsMixin_ToolOutputFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloutputfilter.html#cfn-wisdom-aiagent-tooloutputfilter-jsonpath
	//
	JsonPath *string `field:"optional" json:"jsonPath" yaml:"jsonPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloutputfilter.html#cfn-wisdom-aiagent-tooloutputfilter-outputconfiguration
	//
	OutputConfiguration interface{} `field:"optional" json:"outputConfiguration" yaml:"outputConfiguration"`
}

