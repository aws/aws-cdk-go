package awsstoragegateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsstoragegateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::StorageGateway::SMBFileShare.
//
// An SMB file share exposes an Amazon S3 bucket through a Storage Gateway file gateway, so clients read and write objects as files over SMB. The share's variant is fixed by its resource type rather than by a property: the service constructs the two variants with different APIs, and a property carrying the choice would belong to neither API's input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSMBFileSharePropsMixin := awscdkcfnpropertymixins.Aws_storagegateway.NewCfnSMBFileSharePropsMixin(&CfnSMBFileShareMixinProps{
//   	AccessBasedEnumeration: jsii.Boolean(false),
//   	AdminUserList: []*string{
//   		jsii.String("adminUserList"),
//   	},
//   	AuditDestinationArn: jsii.String("auditDestinationArn"),
//   	Authentication: jsii.String("authentication"),
//   	BucketRegion: jsii.String("bucketRegion"),
//   	CacheAttributes: &CacheAttributesProperty{
//   		CacheStaleTimeoutInSeconds: jsii.Number(123),
//   	},
//   	CaseSensitivity: jsii.String("caseSensitivity"),
//   	DefaultStorageClass: jsii.String("defaultStorageClass"),
//   	EncryptionType: jsii.String("encryptionType"),
//   	FileShareName: jsii.String("fileShareName"),
//   	GatewayArn: jsii.String("gatewayArn"),
//   	GuessMimeTypeEnabled: jsii.Boolean(false),
//   	InvalidUserList: []*string{
//   		jsii.String("invalidUserList"),
//   	},
//   	KmsEncrypted: jsii.Boolean(false),
//   	KmsKey: jsii.String("kmsKey"),
//   	LocationArn: jsii.String("locationArn"),
//   	NotificationPolicy: jsii.String("notificationPolicy"),
//   	ObjectAcl: jsii.String("objectAcl"),
//   	OplocksEnabled: jsii.Boolean(false),
//   	ReadOnly: jsii.Boolean(false),
//   	RequesterPays: jsii.Boolean(false),
//   	Role: jsii.String("role"),
//   	SmbaclEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ValidUserList: []*string{
//   		jsii.String("validUserList"),
//   	},
//   	VpcEndpointDnsName: jsii.String("vpcEndpointDnsName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html
//
type CfnSMBFileSharePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSMBFileShareMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSMBFileSharePropsMixin
type jsiiProxy_CfnSMBFileSharePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSMBFileSharePropsMixin) Props() *CfnSMBFileShareMixinProps {
	var returns *CfnSMBFileShareMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileSharePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::StorageGateway::SMBFileShare`.
func NewCfnSMBFileSharePropsMixin(props *CfnSMBFileShareMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSMBFileSharePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSMBFileSharePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSMBFileSharePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_storagegateway.CfnSMBFileSharePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::StorageGateway::SMBFileShare`.
func NewCfnSMBFileSharePropsMixin_Override(c CfnSMBFileSharePropsMixin, props *CfnSMBFileShareMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_storagegateway.CfnSMBFileSharePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSMBFileSharePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSMBFileSharePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_storagegateway.CfnSMBFileSharePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSMBFileSharePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_storagegateway.CfnSMBFileSharePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSMBFileSharePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSMBFileSharePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

