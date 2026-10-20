package awsbedrockalphamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2"
	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2/awsbedrockalphamixins/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Configures a knowledge base as the given type.
//
// Applies the type-specific configuration, permissions and dependencies that the
// `type` property of `KnowledgeBase` applies, so a `CfnKnowledgeBase` can be
// configured from a `KnowledgeBaseType` without applying its individual mixins.
//
// Example:
//   var collection ICollectionRef
//   var role Role
//
//
//   awscdk.Aws_bedrock.NewCfnKnowledgeBase(this, jsii.String("KnowledgeBase"), &CfnKnowledgeBaseProps{
//   	Name: jsii.String("product-documentation"),
//   	RoleArn: role.roleArn,
//   	KnowledgeBaseConfiguration: &KnowledgeBaseConfigurationProperty{
//   		Type: jsii.String("VECTOR"),
//   	},
//   }).With(mixins.NewKnowledgeBaseTypeConfiguration(bedrock.KnowledgeBaseType_Vector(&VectorKnowledgeBaseTypeProps{
//   	EmbeddingsModel: bedrock.BedrockFoundationModel_TITAN_EMBED_TEXT_V2_1024(),
//   	VectorStore: bedrock.VectorStore_OpenSearchServerless(&OpenSearchServerlessVectorStoreProps{
//   		Collection: *Collection,
//   		VectorIndexName: jsii.String("bedrock-knowledge-base-default-index"),
//   		VectorField: jsii.String("bedrock-knowledge-base-default-vector"),
//   		TextField: jsii.String("AMAZON_BEDROCK_TEXT_CHUNK"),
//   		MetadataField: jsii.String("AMAZON_BEDROCK_METADATA"),
//   	}),
//   })))
//
// Experimental.
type KnowledgeBaseTypeConfiguration interface {
	awscdk.Mixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for KnowledgeBaseTypeConfiguration
type jsiiProxy_KnowledgeBaseTypeConfiguration struct {
	internal.Type__awscdkMixin
}

// Experimental.
func NewKnowledgeBaseTypeConfiguration(knowledgeBaseType awsbedrockalpha.KnowledgeBaseType) KnowledgeBaseTypeConfiguration {
	_init_.Initialize()

	if err := validateNewKnowledgeBaseTypeConfigurationParameters(knowledgeBaseType); err != nil {
		panic(err)
	}
	j := jsiiProxy_KnowledgeBaseTypeConfiguration{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseTypeConfiguration",
		[]interface{}{knowledgeBaseType},
		&j,
	)

	return &j
}

// Experimental.
func NewKnowledgeBaseTypeConfiguration_Override(k KnowledgeBaseTypeConfiguration, knowledgeBaseType awsbedrockalpha.KnowledgeBaseType) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseTypeConfiguration",
		[]interface{}{knowledgeBaseType},
		k,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func KnowledgeBaseTypeConfiguration_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKnowledgeBaseTypeConfiguration_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseTypeConfiguration",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBaseTypeConfiguration) ApplyTo(construct constructs.IConstruct) {
	if err := k.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"applyTo",
		[]interface{}{construct},
	)
}

func (k *jsiiProxy_KnowledgeBaseTypeConfiguration) Supports(construct constructs.IConstruct) *bool {
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

