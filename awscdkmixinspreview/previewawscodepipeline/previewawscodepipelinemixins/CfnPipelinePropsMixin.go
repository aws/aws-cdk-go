package previewawscodepipelinemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodepipeline/previewawscodepipelinemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CodePipeline::Pipeline` resource creates a CodePipeline pipeline that describes how software changes go through a release process.
//
// For more information, see [What Is CodePipeline?](https://docs.aws.amazon.com/codepipeline/latest/userguide/welcome.html) in the *CodePipeline User Guide* .
//
// For an example in YAML and JSON that contains the parameters in this reference, see [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#aws-resource-codepipeline-pipeline--examples) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var configuration interface{}
//
//   cfnPipelinePropsMixin := awscdkmixinspreview.Mixins.NewCfnPipelinePropsMixin(&CfnPipelineMixinProps{
//   	ArtifactStore: &ArtifactStoreProperty{
//   		EncryptionKey: &EncryptionKeyProperty{
//   			Id: jsii.String("id"),
//   			Type: jsii.String("type"),
//   		},
//   		Location: jsii.String("location"),
//   		Type: jsii.String("type"),
//   	},
//   	ArtifactStores: []interface{}{
//   		&ArtifactStoreMapProperty{
//   			ArtifactStore: &ArtifactStoreProperty{
//   				EncryptionKey: &EncryptionKeyProperty{
//   					Id: jsii.String("id"),
//   					Type: jsii.String("type"),
//   				},
//   				Location: jsii.String("location"),
//   				Type: jsii.String("type"),
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
//   					Commands: []*string{
//   						jsii.String("commands"),
//   					},
//   					Configuration: configuration,
//   					EnvironmentVariables: []interface{}{
//   						&EnvironmentVariableProperty{
//   							Name: jsii.String("name"),
//   							Type: jsii.String("type"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					InputArtifacts: []interface{}{
//   						&InputArtifactProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					Name: jsii.String("name"),
//   					Namespace: jsii.String("namespace"),
//   					OutputArtifacts: []interface{}{
//   						&OutputArtifactProperty{
//   							Files: []*string{
//   								jsii.String("files"),
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					OutputVariables: []*string{
//   						jsii.String("outputVariables"),
//   					},
//   					Region: jsii.String("region"),
//   					RoleArn: jsii.String("roleArn"),
//   					RunOrder: jsii.Number(123),
//   					TimeoutInMinutes: jsii.Number(123),
//   				},
//   			},
//   			BeforeEntry: &BeforeEntryConditionsProperty{
//   				Conditions: []interface{}{
//   					&ConditionProperty{
//   						Result: jsii.String("result"),
//   						Rules: []interface{}{
//   							&RuleDeclarationProperty{
//   								Commands: []*string{
//   									jsii.String("commands"),
//   								},
//   								Configuration: configuration,
//   								InputArtifacts: []interface{}{
//   									&InputArtifactProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Name: jsii.String("name"),
//   								Region: jsii.String("region"),
//   								RoleArn: jsii.String("roleArn"),
//   								RuleTypeId: &RuleTypeIdProperty{
//   									Category: jsii.String("category"),
//   									Owner: jsii.String("owner"),
//   									Provider: jsii.String("provider"),
//   									Version: jsii.String("version"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   			Blockers: []interface{}{
//   				&BlockerDeclarationProperty{
//   					Name: jsii.String("name"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			OnFailure: &FailureConditionsProperty{
//   				Conditions: []interface{}{
//   					&ConditionProperty{
//   						Result: jsii.String("result"),
//   						Rules: []interface{}{
//   							&RuleDeclarationProperty{
//   								Commands: []*string{
//   									jsii.String("commands"),
//   								},
//   								Configuration: configuration,
//   								InputArtifacts: []interface{}{
//   									&InputArtifactProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Name: jsii.String("name"),
//   								Region: jsii.String("region"),
//   								RoleArn: jsii.String("roleArn"),
//   								RuleTypeId: &RuleTypeIdProperty{
//   									Category: jsii.String("category"),
//   									Owner: jsii.String("owner"),
//   									Provider: jsii.String("provider"),
//   									Version: jsii.String("version"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				Result: jsii.String("result"),
//   				RetryConfiguration: &RetryConfigurationProperty{
//   					RetryMode: jsii.String("retryMode"),
//   				},
//   			},
//   			OnSuccess: &SuccessConditionsProperty{
//   				Conditions: []interface{}{
//   					&ConditionProperty{
//   						Result: jsii.String("result"),
//   						Rules: []interface{}{
//   							&RuleDeclarationProperty{
//   								Commands: []*string{
//   									jsii.String("commands"),
//   								},
//   								Configuration: configuration,
//   								InputArtifacts: []interface{}{
//   									&InputArtifactProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   								Name: jsii.String("name"),
//   								Region: jsii.String("region"),
//   								RoleArn: jsii.String("roleArn"),
//   								RuleTypeId: &RuleTypeIdProperty{
//   									Category: jsii.String("category"),
//   									Owner: jsii.String("owner"),
//   									Provider: jsii.String("provider"),
//   									Version: jsii.String("version"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Triggers: []interface{}{
//   		&PipelineTriggerDeclarationProperty{
//   			GitConfiguration: &GitConfigurationProperty{
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
//   				SourceActionName: jsii.String("sourceActionName"),
//   			},
//   			ProviderType: jsii.String("providerType"),
//   		},
//   	},
//   	Variables: []interface{}{
//   		&VariableDeclarationProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html
//
type CfnPipelinePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPipelineMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPipelinePropsMixin
type jsiiProxy_CfnPipelinePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPipelinePropsMixin) Props() *CfnPipelineMixinProps {
	var returns *CfnPipelineMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipelinePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodePipeline::Pipeline`.
func NewCfnPipelinePropsMixin(props *CfnPipelineMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPipelinePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPipelinePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPipelinePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodePipeline::Pipeline`.
func NewCfnPipelinePropsMixin_Override(c CfnPipelinePropsMixin, props *CfnPipelineMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPipelinePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipelinePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipelinePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codepipeline.mixins.CfnPipelinePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipelinePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPipelinePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

