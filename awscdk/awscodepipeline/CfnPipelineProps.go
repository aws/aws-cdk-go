package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   					TimeoutInMinutes: jsii.Number(123),
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
//   	ExecutionMode: jsii.String("executionMode"),
//   	Name: jsii.String("name"),
//   	PipelineType: jsii.String("pipelineType"),
//   	RestartExecutionOnUpdate: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Triggers: []interface{}{
//   		&PipelineTriggerDeclarationProperty{
//   			ProviderType: jsii.String("providerType"),
//
//   			// the properties below are optional
//   			GitConfiguration: &GitConfigurationProperty{
//   				SourceActionName: jsii.String("sourceActionName"),
//
//   				// the properties below are optional
//   				PullRequest: []interface{}{
//   					&GitPullRequestFilterProperty{
//   						Branches: &GitBranchFilterCriteriaProperty{
//   							Excludes: []*string{
//   								jsii.String("excludes"),
//   							},
//   							Includes: []*string{
//   								jsii.String("includes"),
//   							},
//   						},
//   						Events: []*string{
//   							jsii.String("events"),
//   						},
//   						FilePaths: &GitFilePathFilterCriteriaProperty{
//   							Excludes: []*string{
//   								jsii.String("excludes"),
//   							},
//   							Includes: []*string{
//   								jsii.String("includes"),
//   							},
//   						},
//   					},
//   				},
//   				Push: []interface{}{
//   					&GitPushFilterProperty{
//   						Branches: &GitBranchFilterCriteriaProperty{
//   							Excludes: []*string{
//   								jsii.String("excludes"),
//   							},
//   							Includes: []*string{
//   								jsii.String("includes"),
//   							},
//   						},
//   						FilePaths: &GitFilePathFilterCriteriaProperty{
//   							Excludes: []*string{
//   								jsii.String("excludes"),
//   							},
//   							Includes: []*string{
//   								jsii.String("includes"),
//   							},
//   						},
//   						Tags: &GitTagFilterCriteriaProperty{
//   							Excludes: []*string{
//   								jsii.String("excludes"),
//   							},
//   							Includes: []*string{
//   								jsii.String("includes"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Variables: []interface{}{
//   		&VariableDeclarationProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			DefaultValue: jsii.String("defaultValue"),
//   			Description: jsii.String("description"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html
//
type CfnPipelineProps struct {
	// The Amazon Resource Name (ARN) for CodePipeline to use to either perform actions with no `actionRoleArn` , or to use to assume roles for actions with an `actionRoleArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Represents information about a stage and its definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-stages
	//
	Stages interface{} `field:"required" json:"stages" yaml:"stages"`
	// The S3 bucket where artifacts for the pipeline are stored.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-artifactstore
	//
	ArtifactStore interface{} `field:"optional" json:"artifactStore" yaml:"artifactStore"`
	// A mapping of `artifactStore` objects and their corresponding AWS Regions.
	//
	// There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.
	//
	// > You must include either `artifactStore` or `artifactStores` in your pipeline, but you cannot use both. If you create a cross-region action in your pipeline, you must use `artifactStores` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-artifactstores
	//
	ArtifactStores interface{} `field:"optional" json:"artifactStores" yaml:"artifactStores"`
	// Represents the input of a `DisableStageTransition` action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-disableinboundstagetransitions
	//
	DisableInboundStageTransitions interface{} `field:"optional" json:"disableInboundStageTransitions" yaml:"disableInboundStageTransitions"`
	// The method that the pipeline will use to handle multiple executions.
	//
	// The default mode is SUPERSEDED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-executionmode
	//
	ExecutionMode *string `field:"optional" json:"executionMode" yaml:"executionMode"`
	// The name of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// CodePipeline provides the following pipeline types, which differ in characteristics and price, so that you can tailor your pipeline features and cost to the needs of your applications.
	//
	// - V1 type pipelines have a JSON structure that contains standard pipeline, stage, and action-level parameters.
	// - V2 type pipelines have the same structure as a V1 type, along with additional parameters for release safety and trigger configuration.
	//
	// > Including V2 parameters, such as triggers on Git tags, in the pipeline JSON when creating or updating a pipeline will result in the pipeline having the V2 type of pipeline and the associated costs.
	//
	// For information about pricing for CodePipeline, see [Pricing](https://docs.aws.amazon.com/codepipeline/pricing/) .
	//
	// For information about which type of pipeline to choose, see [What type of pipeline is right for me?](https://docs.aws.amazon.com/codepipeline/latest/userguide/pipeline-types-planning.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-pipelinetype
	//
	PipelineType *string `field:"optional" json:"pipelineType" yaml:"pipelineType"`
	// Indicates whether to rerun the CodePipeline pipeline after you update it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-restartexecutiononupdate
	//
	RestartExecutionOnUpdate interface{} `field:"optional" json:"restartExecutionOnUpdate" yaml:"restartExecutionOnUpdate"`
	// Specifies the tags applied to the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The trigger configuration specifying a type of event, such as Git tags, that starts the pipeline.
	//
	// > When a trigger configuration is specified, default change detection for repository and branch commits is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-triggers
	//
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
	// A list that defines the pipeline variables for a pipeline resource.
	//
	// Variable names can have alphanumeric and underscore characters, and the values must match `[A-Za-z0-9@\-_]+` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

