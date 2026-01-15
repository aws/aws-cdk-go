package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   toolOverrideInputValueProperty := &ToolOverrideInputValueProperty{
//   	JsonPath: jsii.String("jsonPath"),
//   	Value: &ToolOverrideInputValueConfigurationProperty{
//   		Constant: &ToolOverrideConstantInputValueProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloverrideinputvalue.html
//
type CfnAIAgent_ToolOverrideInputValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloverrideinputvalue.html#cfn-wisdom-aiagent-tooloverrideinputvalue-jsonpath
	//
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloverrideinputvalue.html#cfn-wisdom-aiagent-tooloverrideinputvalue-value
	//
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

