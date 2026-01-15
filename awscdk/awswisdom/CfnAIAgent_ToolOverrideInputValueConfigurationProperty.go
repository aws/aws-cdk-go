package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   toolOverrideInputValueConfigurationProperty := &ToolOverrideInputValueConfigurationProperty{
//   	Constant: &ToolOverrideConstantInputValueProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloverrideinputvalueconfiguration.html
//
type CfnAIAgent_ToolOverrideInputValueConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloverrideinputvalueconfiguration.html#cfn-wisdom-aiagent-tooloverrideinputvalueconfiguration-constant
	//
	Constant interface{} `field:"required" json:"constant" yaml:"constant"`
}

