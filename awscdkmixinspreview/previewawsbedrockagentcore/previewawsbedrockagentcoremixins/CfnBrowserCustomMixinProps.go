package previewawsbedrockagentcoremixins


// Properties for CfnBrowserCustomPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserCustomMixinProps := &CfnBrowserCustomMixinProps{
//   	BrowserSigning: &BrowserSigningProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Name: jsii.String("name"),
//   	NetworkConfiguration: &BrowserNetworkConfigurationProperty{
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
//   	RecordingConfig: &RecordingConfigProperty{
//   		Enabled: jsii.Boolean(false),
//   		S3Location: &S3LocationProperty{
//   			Bucket: jsii.String("bucket"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browsercustom.html
//
type CfnBrowserCustomMixinProps struct {
	// Browser signing configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browsercustom.html#cfn-bedrockagentcore-browsercustom-browsersigning
	//
	BrowserSigning interface{} `field:"optional" json:"browserSigning" yaml:"browserSigning"`
	// The custom browser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browsercustom.html#cfn-bedrockagentcore-browsercustom-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the execution role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browsercustom.html#cfn-bedrockagentcore-browsercustom-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the custom browser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browsercustom.html#cfn-bedrockagentcore-browsercustom-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The network configuration for a code interpreter.
	//
	// This structure defines how the code interpreter connects to the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browsercustom.html#cfn-bedrockagentcore-browsercustom-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// THe custom browser configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browsercustom.html#cfn-bedrockagentcore-browsercustom-recordingconfig
	//
	RecordingConfig interface{} `field:"optional" json:"recordingConfig" yaml:"recordingConfig"`
	// The tags for the custom browser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-browsercustom.html#cfn-bedrockagentcore-browsercustom-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

