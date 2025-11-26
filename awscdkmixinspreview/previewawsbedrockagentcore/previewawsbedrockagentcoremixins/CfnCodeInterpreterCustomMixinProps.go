package previewawsbedrockagentcoremixins


// Properties for CfnCodeInterpreterCustomPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomMixinProps := &CfnCodeInterpreterCustomMixinProps{
//   	Description: jsii.String("description"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Name: jsii.String("name"),
//   	NetworkConfiguration: &CodeInterpreterNetworkConfigurationProperty{
//   		NetworkMode: jsii.String("networkMode"),
//   		VpcConfig: &VpcConfigProperty{
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html
//
type CfnCodeInterpreterCustomMixinProps struct {
	// The code interpreter description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html#cfn-bedrockagentcore-codeinterpretercustom-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the execution role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html#cfn-bedrockagentcore-codeinterpretercustom-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the code interpreter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html#cfn-bedrockagentcore-codeinterpretercustom-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The network configuration for a code interpreter.
	//
	// This structure defines how the code interpreter connects to the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html#cfn-bedrockagentcore-codeinterpretercustom-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// The tags for the code interpreter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-codeinterpretercustom.html#cfn-bedrockagentcore-codeinterpretercustom-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

