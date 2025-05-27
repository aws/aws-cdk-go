package awsbedrock


// Inline code config strucuture, contains code configs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inlineCodeFlowNodeConfigurationProperty := &InlineCodeFlowNodeConfigurationProperty{
//   	Code: jsii.String("code"),
//   	Language: jsii.String("language"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-inlinecodeflownodeconfiguration.html
//
type CfnFlow_InlineCodeFlowNodeConfigurationProperty struct {
	// The inline code entered by customers.
	//
	// max size is 5MB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-inlinecodeflownodeconfiguration.html#cfn-bedrock-flow-inlinecodeflownodeconfiguration-code
	//
	Code *string `field:"required" json:"code" yaml:"code"`
	// Enum encodes the supported language type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-inlinecodeflownodeconfiguration.html#cfn-bedrock-flow-inlinecodeflownodeconfiguration-language
	//
	Language *string `field:"required" json:"language" yaml:"language"`
}

