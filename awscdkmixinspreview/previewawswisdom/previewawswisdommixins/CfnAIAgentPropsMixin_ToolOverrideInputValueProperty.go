package previewawswisdommixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnAIAgentPropsMixin_ToolOverrideInputValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloverrideinputvalue.html#cfn-wisdom-aiagent-tooloverrideinputvalue-jsonpath
	//
	JsonPath *string `field:"optional" json:"jsonPath" yaml:"jsonPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloverrideinputvalue.html#cfn-wisdom-aiagent-tooloverrideinputvalue-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

