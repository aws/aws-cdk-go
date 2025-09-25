package awsbedrockagentcore


// Network configuration for code interpreter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeInterpreterNetworkConfigurationProperty := &CodeInterpreterNetworkConfigurationProperty{
//   	NetworkMode: jsii.String("networkMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-codeinterpretercustom-codeinterpreternetworkconfiguration.html
//
type CfnCodeInterpreterCustom_CodeInterpreterNetworkConfigurationProperty struct {
	// Network modes supported by code interpreter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-codeinterpretercustom-codeinterpreternetworkconfiguration.html#cfn-bedrockagentcore-codeinterpretercustom-codeinterpreternetworkconfiguration-networkmode
	//
	NetworkMode *string `field:"required" json:"networkMode" yaml:"networkMode"`
}

