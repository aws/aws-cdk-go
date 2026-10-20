package awsbedrockalphamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2/awsbedrockalphamixins/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Stores a knowledge base's vectors in an existing Amazon OpenSearch Serverless vector index.
//
// Sets `StorageConfiguration` to `OPENSEARCH_SERVERLESS` with the given
// collection, index and field mapping. Grants the role referenced by the knowledge
// base's `roleArn` `aoss:APIAccessAll` on the collection and makes the knowledge
// base depend on that policy.
//
// The index must already exist and the collection's data access policy must
// allow the service role to describe, read and write it; neither is created here.
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
//   		VectorKnowledgeBaseConfiguration: &VectorKnowledgeBaseConfigurationProperty{
//   			EmbeddingModelArn: bedrock.BedrockFoundationModel_TITAN_EMBED_TEXT_V2_1024().ModelArn,
//   		},
//   	},
//   }).With(mixins.NewKnowledgeBaseOpenSearchServerlessStorage(&KnowledgeBaseOpenSearchServerlessStorageProps{
//   	Collection: Collection,
//   	VectorIndexName: jsii.String("bedrock-knowledge-base-default-index"),
//   	VectorField: jsii.String("bedrock-knowledge-base-default-vector"),
//   	TextField: jsii.String("AMAZON_BEDROCK_TEXT_CHUNK"),
//   	MetadataField: jsii.String("AMAZON_BEDROCK_METADATA"),
//   }))
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-permissions.html#kb-permissions-oss
//
// Experimental.
type KnowledgeBaseOpenSearchServerlessStorage interface {
	awscdk.Mixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for KnowledgeBaseOpenSearchServerlessStorage
type jsiiProxy_KnowledgeBaseOpenSearchServerlessStorage struct {
	internal.Type__awscdkMixin
}

// Experimental.
func NewKnowledgeBaseOpenSearchServerlessStorage(props *KnowledgeBaseOpenSearchServerlessStorageProps) KnowledgeBaseOpenSearchServerlessStorage {
	_init_.Initialize()

	if err := validateNewKnowledgeBaseOpenSearchServerlessStorageParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KnowledgeBaseOpenSearchServerlessStorage{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseOpenSearchServerlessStorage",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewKnowledgeBaseOpenSearchServerlessStorage_Override(k KnowledgeBaseOpenSearchServerlessStorage, props *KnowledgeBaseOpenSearchServerlessStorageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseOpenSearchServerlessStorage",
		[]interface{}{props},
		k,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func KnowledgeBaseOpenSearchServerlessStorage_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKnowledgeBaseOpenSearchServerlessStorage_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseOpenSearchServerlessStorage",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBaseOpenSearchServerlessStorage) ApplyTo(construct constructs.IConstruct) {
	if err := k.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"applyTo",
		[]interface{}{construct},
	)
}

func (k *jsiiProxy_KnowledgeBaseOpenSearchServerlessStorage) Supports(construct constructs.IConstruct) *bool {
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

