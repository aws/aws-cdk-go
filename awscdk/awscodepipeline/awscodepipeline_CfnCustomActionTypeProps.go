package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCustomActionType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomActionTypeProps := &cfnCustomActionTypeProps{
//   	category: jsii.String("category"),
//   	inputArtifactDetails: &artifactDetailsProperty{
//   		maximumCount: jsii.Number(123),
//   		minimumCount: jsii.Number(123),
//   	},
//   	outputArtifactDetails: &artifactDetailsProperty{
//   		maximumCount: jsii.Number(123),
//   		minimumCount: jsii.Number(123),
//   	},
//   	provider: jsii.String("provider"),
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	configurationProperties: []interface{}{
//   		&configurationPropertiesProperty{
//   			key: jsii.Boolean(false),
//   			name: jsii.String("name"),
//   			required: jsii.Boolean(false),
//   			secret: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			queryable: jsii.Boolean(false),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	settings: &settingsProperty{
//   		entityUrlTemplate: jsii.String("entityUrlTemplate"),
//   		executionUrlTemplate: jsii.String("executionUrlTemplate"),
//   		revisionUrlTemplate: jsii.String("revisionUrlTemplate"),
//   		thirdPartyConfigurationUrl: jsii.String("thirdPartyConfigurationUrl"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCustomActionTypeProps struct {
	// The category of the custom action, such as a build action or a test action.
	Category *string `field:"required" json:"category" yaml:"category"`
	// The details of the input artifact for the action, such as its commit ID.
	InputArtifactDetails interface{} `field:"required" json:"inputArtifactDetails" yaml:"inputArtifactDetails"`
	// The details of the output artifact of the action, such as its commit ID.
	OutputArtifactDetails interface{} `field:"required" json:"outputArtifactDetails" yaml:"outputArtifactDetails"`
	// The provider of the service used in the custom action, such as CodeDeploy.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// The version identifier of the custom action.
	Version *string `field:"required" json:"version" yaml:"version"`
	// The configuration properties for the custom action.
	//
	// > You can refer to a name in the configuration properties of the custom action within the URL templates by following the format of {Config:name}, as long as the configuration property is both required and not secret. For more information, see [Create a Custom Action for a Pipeline](https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html) .
	ConfigurationProperties interface{} `field:"optional" json:"configurationProperties" yaml:"configurationProperties"`
	// URLs that provide users information about this custom action.
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// The tags for the custom action.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

