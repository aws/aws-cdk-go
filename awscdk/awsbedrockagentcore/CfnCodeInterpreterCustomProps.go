package awsbedrockagentcore


// Properties for defining a `CfnCodeInterpreterCustom`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeInterpreterCustomProps := &CfnCodeInterpreterCustomProps{
//   	Name: jsii.String("name"),
//   	NetworkConfiguration: &CodeInterpreterNetworkConfigurationProperty{
//   		NetworkMode: jsii.String("networkMode"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html
//
type CfnCodeInterpreterCustomProps struct {
	// The name of the sandbox.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html#cfn-bedrockagentcore-codeinterpretercustom-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Network configuration for code interpreter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html#cfn-bedrockagentcore-codeinterpretercustom-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"required" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Description of the code interpreter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html#cfn-bedrockagentcore-codeinterpretercustom-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the IAM role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html#cfn-bedrockagentcore-codeinterpretercustom-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// A map of tag keys and values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html#cfn-bedrockagentcore-codeinterpretercustom-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

