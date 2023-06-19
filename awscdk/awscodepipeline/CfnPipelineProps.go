package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   cfnPipelineProps := &CfnPipelineProps{
//   	RoleArn: jsii.String("roleArn"),
//   	Stages: []interface{}{
//   		&StageDeclarationProperty{
//   			Actions: []interface{}{
//   				&ActionDeclarationProperty{
//   					ActionTypeId: &ActionTypeIdProperty{
//   						Category: jsii.String("category"),
//   						Owner: jsii.String("owner"),
//   						Provider: jsii.String("provider"),
//   						Version: jsii.String("version"),
//   					},
//   					Name: jsii.String("name"),
//
//   					// the properties below are optional
//   					Configuration: configuration,
//   					InputArtifacts: []interface{}{
//   						&InputArtifactProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Namespace: jsii.String("namespace"),
//   					OutputArtifacts: []interface{}{
//   						&OutputArtifactProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Region: jsii.String("region"),
//   					RoleArn: jsii.String("roleArn"),
//   					RunOrder: jsii.Number(123),
//   				},
//   			},
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Blockers: []interface{}{
//   				&BlockerDeclarationProperty{
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ArtifactStore: &ArtifactStoreProperty{
//   		Location: jsii.String("location"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		EncryptionKey: &EncryptionKeyProperty{
//   			Id: jsii.String("id"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ArtifactStores: []interface{}{
//   		&ArtifactStoreMapProperty{
//   			ArtifactStore: &ArtifactStoreProperty{
//   				Location: jsii.String("location"),
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				EncryptionKey: &EncryptionKeyProperty{
//   					Id: jsii.String("id"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Region: jsii.String("region"),
//   		},
//   	},
//   	DisableInboundStageTransitions: []interface{}{
//   		&StageTransitionProperty{
//   			Reason: jsii.String("reason"),
//   			StageName: jsii.String("stageName"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RestartExecutionOnUpdate: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipelineProps struct {
	// The Amazon Resource Name (ARN) for CodePipeline to use to either perform actions with no `actionRoleArn` , or to use to assume roles for actions with an `actionRoleArn` .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Represents information about a stage and its definition.
	Stages interface{} `field:"required" json:"stages" yaml:"stages"`
	// The S3 bucket where artifacts for the pipeline are stored.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStore interface{} `field:"optional" json:"artifactStore" yaml:"artifactStore"`
	// A mapping of `artifactStore` objects and their corresponding AWS Regions.
	//
	// There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	ArtifactStores interface{} `field:"optional" json:"artifactStores" yaml:"artifactStores"`
	// Represents the input of a `DisableStageTransition` action.
	DisableInboundStageTransitions interface{} `field:"optional" json:"disableInboundStageTransitions" yaml:"disableInboundStageTransitions"`
	// The name of the pipeline.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates whether to rerun the CodePipeline pipeline after you update it.
	RestartExecutionOnUpdate interface{} `field:"optional" json:"restartExecutionOnUpdate" yaml:"restartExecutionOnUpdate"`
	// Specifies the tags applied to the pipeline.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

