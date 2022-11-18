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
//   cfnPipelineProps := &cfnPipelineProps{
//   	roleArn: jsii.String("roleArn"),
//   	stages: []interface{}{
//   		&stageDeclarationProperty{
//   			actions: []interface{}{
//   				&actionDeclarationProperty{
//   					actionTypeId: &actionTypeIdProperty{
//   						category: jsii.String("category"),
//   						owner: jsii.String("owner"),
//   						provider: jsii.String("provider"),
//   						version: jsii.String("version"),
//   					},
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					configuration: configuration,
//   					inputArtifacts: []interface{}{
//   						&inputArtifactProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					namespace: jsii.String("namespace"),
//   					outputArtifacts: []interface{}{
//   						&outputArtifactProperty{
//   							name: jsii.String("name"),
//   						},
//   					},
//   					region: jsii.String("region"),
//   					roleArn: jsii.String("roleArn"),
//   					runOrder: jsii.Number(123),
//   				},
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			blockers: []interface{}{
//   				&blockerDeclarationProperty{
//   					name: jsii.String("name"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	artifactStore: &artifactStoreProperty{
//   		location: jsii.String("location"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		encryptionKey: &encryptionKeyProperty{
//   			id: jsii.String("id"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	artifactStores: []interface{}{
//   		&artifactStoreMapProperty{
//   			artifactStore: &artifactStoreProperty{
//   				location: jsii.String("location"),
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				encryptionKey: &encryptionKeyProperty{
//   					id: jsii.String("id"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			region: jsii.String("region"),
//   		},
//   	},
//   	disableInboundStageTransitions: []interface{}{
//   		&stageTransitionProperty{
//   			reason: jsii.String("reason"),
//   			stageName: jsii.String("stageName"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	restartExecutionOnUpdate: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

