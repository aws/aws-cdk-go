package awsbedrockalphamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockalpha/v2/awsbedrockalphamixins/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Stores the images, audio and video a vector knowledge base extracts from multimodal documents in an Amazon S3 bucket.
//
// Merges `SupplementalDataStorageConfiguration` into the existing vector
// configuration, replacing any location already set. Grants the role referenced
// by the knowledge base's `roleArn` `s3:ListBucket` on the bucket and
// `s3:GetObject`, `s3:PutObject` and `s3:DeleteObject` on its objects, and makes
// the knowledge base depend on that policy.
//
// Example:
//   var knowledgeBase CfnKnowledgeBase
//   var bucket IBucket
//
//
//   knowledgeBase.With(mixins.NewKnowledgeBaseSupplementalDataStorage(&KnowledgeBaseSupplementalDataStorageProps{
//   	Bucket: Bucket,
//   }))
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-permissions.html#kb-permissions-multimodal
//
// Experimental.
type KnowledgeBaseSupplementalDataStorage interface {
	awscdk.Mixin
	// Applies the mixin functionality to the target construct.
	// Experimental.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	// Experimental.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for KnowledgeBaseSupplementalDataStorage
type jsiiProxy_KnowledgeBaseSupplementalDataStorage struct {
	internal.Type__awscdkMixin
}

// Experimental.
func NewKnowledgeBaseSupplementalDataStorage(props *KnowledgeBaseSupplementalDataStorageProps) KnowledgeBaseSupplementalDataStorage {
	_init_.Initialize()

	if err := validateNewKnowledgeBaseSupplementalDataStorageParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KnowledgeBaseSupplementalDataStorage{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseSupplementalDataStorage",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewKnowledgeBaseSupplementalDataStorage_Override(k KnowledgeBaseSupplementalDataStorage, props *KnowledgeBaseSupplementalDataStorageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseSupplementalDataStorage",
		[]interface{}{props},
		k,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func KnowledgeBaseSupplementalDataStorage_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKnowledgeBaseSupplementalDataStorage_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseSupplementalDataStorage",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBaseSupplementalDataStorage) ApplyTo(construct constructs.IConstruct) {
	if err := k.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"applyTo",
		[]interface{}{construct},
	)
}

func (k *jsiiProxy_KnowledgeBaseSupplementalDataStorage) Supports(construct constructs.IConstruct) *bool {
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

