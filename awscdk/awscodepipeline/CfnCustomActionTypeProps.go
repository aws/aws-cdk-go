package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCustomActionType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomActionTypeProps := &CfnCustomActionTypeProps{
//   	Category: jsii.String("category"),
//   	InputArtifactDetails: &ArtifactDetailsProperty{
//   		MaximumCount: jsii.Number(123),
//   		MinimumCount: jsii.Number(123),
//   	},
//   	OutputArtifactDetails: &ArtifactDetailsProperty{
//   		MaximumCount: jsii.Number(123),
//   		MinimumCount: jsii.Number(123),
//   	},
//   	Provider: jsii.String("provider"),
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	ConfigurationProperties: []interface{}{
//   		&ConfigurationPropertiesProperty{
//   			Key: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			Required: jsii.Boolean(false),
//   			Secret: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			Queryable: jsii.Boolean(false),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Settings: &SettingsProperty{
//   		EntityUrlTemplate: jsii.String("entityUrlTemplate"),
//   		ExecutionUrlTemplate: jsii.String("executionUrlTemplate"),
//   		RevisionUrlTemplate: jsii.String("revisionUrlTemplate"),
//   		ThirdPartyConfigurationUrl: jsii.String("thirdPartyConfigurationUrl"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html
//
type CfnCustomActionTypeProps struct {
	// The category of the custom action, such as a build action or a test action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-category
	//
	Category *string `field:"required" json:"category" yaml:"category"`
	// The details of the input artifact for the action, such as its commit ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-inputartifactdetails
	//
	InputArtifactDetails interface{} `field:"required" json:"inputArtifactDetails" yaml:"inputArtifactDetails"`
	// The details of the output artifact of the action, such as its commit ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-outputartifactdetails
	//
	OutputArtifactDetails interface{} `field:"required" json:"outputArtifactDetails" yaml:"outputArtifactDetails"`
	// The provider of the service used in the custom action, such as CodeDeploy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-provider
	//
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// The version identifier of the custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
	// The configuration properties for the custom action.
	//
	// > You can refer to a name in the configuration properties of the custom action within the URL templates by following the format of {Config:name}, as long as the configuration property is both required and not secret. For more information, see [Create a Custom Action for a Pipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-configurationproperties
	//
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// URLs that provide users information about this custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// The tags for the custom action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

