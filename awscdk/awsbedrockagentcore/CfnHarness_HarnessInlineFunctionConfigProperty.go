package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var inputSchema interface{}
//
//   harnessInlineFunctionConfigProperty := &HarnessInlineFunctionConfigProperty{
//   	Description: jsii.String("description"),
//   	InputSchema: inputSchema,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessinlinefunctionconfig.html
//
type CfnHarness_HarnessInlineFunctionConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessinlinefunctionconfig.html#cfn-bedrockagentcore-harness-harnessinlinefunctionconfig-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// JSON Schema describing the tool's input parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessinlinefunctionconfig.html#cfn-bedrockagentcore-harness-harnessinlinefunctionconfig-inputschema
	//
	InputSchema interface{} `field:"required" json:"inputSchema" yaml:"inputSchema"`
}

