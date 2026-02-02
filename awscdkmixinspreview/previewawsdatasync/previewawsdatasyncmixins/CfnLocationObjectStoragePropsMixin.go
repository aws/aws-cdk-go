package previewawsdatasyncmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdatasync/previewawsdatasyncmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DataSync::LocationObjectStorage` resource specifies an endpoint for a self-managed object storage bucket.
//
// For more information about self-managed object storage locations, see [Creating a Location for Object Storage](https://docs.aws.amazon.com/datasync/latest/userguide/create-object-location.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLocationObjectStoragePropsMixin := awscdkmixinspreview.Mixins.NewCfnLocationObjectStoragePropsMixin(&CfnLocationObjectStorageMixinProps{
//   	AccessKey: jsii.String("accessKey"),
//   	AgentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	BucketName: jsii.String("bucketName"),
//   	CmkSecretConfig: &CmkSecretConfigProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	CustomSecretConfig: &CustomSecretConfigProperty{
//   		SecretAccessRoleArn: jsii.String("secretAccessRoleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	SecretKey: jsii.String("secretKey"),
//   	ServerCertificate: jsii.String("serverCertificate"),
//   	ServerHostname: jsii.String("serverHostname"),
//   	ServerPort: jsii.Number(123),
//   	ServerProtocol: jsii.String("serverProtocol"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationobjectstorage.html
//
type CfnLocationObjectStoragePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLocationObjectStorageMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLocationObjectStoragePropsMixin
type jsiiProxy_CfnLocationObjectStoragePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLocationObjectStoragePropsMixin) Props() *CfnLocationObjectStorageMixinProps {
	var returns *CfnLocationObjectStorageMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationObjectStoragePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataSync::LocationObjectStorage`.
func NewCfnLocationObjectStoragePropsMixin(props *CfnLocationObjectStorageMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLocationObjectStoragePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLocationObjectStoragePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLocationObjectStoragePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationObjectStoragePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataSync::LocationObjectStorage`.
func NewCfnLocationObjectStoragePropsMixin_Override(c CfnLocationObjectStoragePropsMixin, props *CfnLocationObjectStorageMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationObjectStoragePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLocationObjectStoragePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLocationObjectStoragePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationObjectStoragePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationObjectStoragePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationObjectStoragePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLocationObjectStoragePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLocationObjectStoragePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

