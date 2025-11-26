package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconnect/previewawsconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The storage configuration for the instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceStorageConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnInstanceStorageConfigPropsMixin(&CfnInstanceStorageConfigMixinProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	KinesisFirehoseConfig: &KinesisFirehoseConfigProperty{
//   		FirehoseArn: jsii.String("firehoseArn"),
//   	},
//   	KinesisStreamConfig: &KinesisStreamConfigProperty{
//   		StreamArn: jsii.String("streamArn"),
//   	},
//   	KinesisVideoStreamConfig: &KinesisVideoStreamConfigProperty{
//   		EncryptionConfig: &EncryptionConfigProperty{
//   			EncryptionType: jsii.String("encryptionType"),
//   			KeyId: jsii.String("keyId"),
//   		},
//   		Prefix: jsii.String("prefix"),
//   		RetentionPeriodHours: jsii.Number(123),
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	S3Config: &S3ConfigProperty{
//   		BucketName: jsii.String("bucketName"),
//   		BucketPrefix: jsii.String("bucketPrefix"),
//   		EncryptionConfig: &EncryptionConfigProperty{
//   			EncryptionType: jsii.String("encryptionType"),
//   			KeyId: jsii.String("keyId"),
//   		},
//   	},
//   	StorageType: jsii.String("storageType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html
//
type CfnInstanceStorageConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInstanceStorageConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInstanceStorageConfigPropsMixin
type jsiiProxy_CfnInstanceStorageConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInstanceStorageConfigPropsMixin) Props() *CfnInstanceStorageConfigMixinProps {
	var returns *CfnInstanceStorageConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceStorageConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::InstanceStorageConfig`.
func NewCfnInstanceStorageConfigPropsMixin(props *CfnInstanceStorageConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInstanceStorageConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstanceStorageConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstanceStorageConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceStorageConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::InstanceStorageConfig`.
func NewCfnInstanceStorageConfigPropsMixin_Override(c CfnInstanceStorageConfigPropsMixin, props *CfnInstanceStorageConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceStorageConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInstanceStorageConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstanceStorageConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceStorageConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstanceStorageConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceStorageConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstanceStorageConfigPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnInstanceStorageConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

