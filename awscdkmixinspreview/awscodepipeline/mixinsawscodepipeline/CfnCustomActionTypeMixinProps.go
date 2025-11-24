package mixinsawscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCustomActionTypePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCustomActionTypeMixinProps := &CfnCustomActionTypeMixinProps{
//   	Category: jsii.String("category"),
//   	ConfigurationProperties: []interface{}{
//   		&ConfigurationPropertiesProperty{
//   			Description: jsii.String("description"),
//   			Key: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			Queryable: jsii.Boolean(false),
//   			Required: jsii.Boolean(false),
//   			Secret: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	InputArtifactDetails: &ArtifactDetailsProperty{
//   		MaximumCount: jsii.Number(123),
//   		MinimumCount: jsii.Number(123),
//   	},
//   	OutputArtifactDetails: &ArtifactDetailsProperty{
//   		MaximumCount: jsii.Number(123),
//   		MinimumCount: jsii.Number(123),
//   	},
//   	Provider: jsii.String("provider"),
//   	Settings: &SettingsProperty{
//   		EntityUrlTemplate: jsii.String("entityUrlTemplate"),
//   		ExecutionUrlTemplate: jsii.String("executionUrlTemplate"),
//   		RevisionUrlTemplate: jsii.String("revisionUrlTemplate"),
//   		ThirdPartyConfigurationUrl: jsii.String("thirdPartyConfigurationUrl"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html
//
type CfnCustomActionTypeMixinProps struct {
	// The category of the custom action, such as a build action or a test action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The configuration properties for the custom action.
	//
	// > You can refer to a name in the configuration properties of the custom action within the URL templates by following the format of {Config:name}, as long as the configuration property is both required and not secret. For more information, see [Create a Custom Action for a Pipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-configurationproperties
	//
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// The details of the input artifact for the action, such as its commit ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-inputartifactdetails
	//
	InputArtifactDetails interface{} `field:"optional" json:"inputArtifactDetails" yaml:"inputArtifactDetails"`
	// The details of the output artifact of the action, such as its commit ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-outputartifactdetails
	//
	OutputArtifactDetails interface{} `field:"optional" json:"outputArtifactDetails" yaml:"outputArtifactDetails"`
	// The provider of the service used in the custom action, such as CodeDeploy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-provider
	//
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// URLs that provide users information about this custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// The tags for the custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The version identifier of the custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

