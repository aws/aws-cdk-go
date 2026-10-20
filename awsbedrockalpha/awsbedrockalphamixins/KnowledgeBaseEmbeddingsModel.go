package awsbedrockalphamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2/awsbedrockalphamixins/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Configures the embeddings model of a vector knowledge base.
//
// Sets `EmbeddingModelArn`, `Dimensions` (only for models with a configurable
// dimension; CloudFormation rejects it otherwise) and `EmbeddingDataType`.
// Grants the role referenced by the knowledge base's `roleArn` `bedrock:InvokeModel*`
// and `bedrock:GetFoundationModel` on the model and makes the knowledge base depend
// on that policy.
//
// Example:
//   var role Role
//
//
//   awscdk.Aws_bedrock.NewCfnKnowledgeBase(this, jsii.String("KnowledgeBase"), &CfnKnowledgeBaseProps{
//   	Name: jsii.String("product-documentation"),
//   	RoleArn: role.roleArn,
//   	KnowledgeBaseConfiguration: &KnowledgeBaseConfigurationProperty{
//   		Type: jsii.String("VECTOR"),
//   	},
//   }).With(mixins.NewKnowledgeBaseEmbeddingsModel(&KnowledgeBaseEmbeddingsModelProps{
//   	EmbeddingsModel: bedrock.BedrockFoundationModel_TITAN_EMBED_TEXT_V2_1024(),
//   }))
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-permissions.html#kb-permissions-fm
//
// Experimental.
type KnowledgeBaseEmbeddingsModel interface {
	awscdk.Mixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for KnowledgeBaseEmbeddingsModel
type jsiiProxy_KnowledgeBaseEmbeddingsModel struct {
	internal.Type__awscdkMixin
}

// Experimental.
func NewKnowledgeBaseEmbeddingsModel(props *KnowledgeBaseEmbeddingsModelProps) KnowledgeBaseEmbeddingsModel {
	_init_.Initialize()

	if err := validateNewKnowledgeBaseEmbeddingsModelParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KnowledgeBaseEmbeddingsModel{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseEmbeddingsModel",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewKnowledgeBaseEmbeddingsModel_Override(k KnowledgeBaseEmbeddingsModel, props *KnowledgeBaseEmbeddingsModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseEmbeddingsModel",
		[]interface{}{props},
		k,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func KnowledgeBaseEmbeddingsModel_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKnowledgeBaseEmbeddingsModel_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseEmbeddingsModel",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBaseEmbeddingsModel) ApplyTo(construct constructs.IConstruct) {
	if err := k.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"applyTo",
		[]interface{}{construct},
	)
}

func (k *jsiiProxy_KnowledgeBaseEmbeddingsModel) Supports(construct constructs.IConstruct) *bool {
	if err := k.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		k,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

