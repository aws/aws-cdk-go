package awsstoragegateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsstoragegateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::StorageGateway::NFSFileShare.
//
// An NFS file share exposes an Amazon S3 bucket through a Storage Gateway file gateway, so clients read and write objects as files over NFS. The share's variant is fixed by its resource type rather than by a property: the service constructs the two variants with different APIs, and a property carrying the choice would belong to neither API's input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnNFSFileSharePropsMixin := awscdkcfnpropertymixins.Aws_storagegateway.NewCfnNFSFileSharePropsMixin(&CfnNFSFileShareMixinProps{
//   	AuditDestinationArn: jsii.String("auditDestinationArn"),
//   	BucketRegion: jsii.String("bucketRegion"),
//   	CacheAttributes: &CacheAttributesProperty{
//   		CacheStaleTimeoutInSeconds: jsii.Number(123),
//   	},
//   	ClientList: []*string{
//   		jsii.String("clientList"),
//   	},
//   	DefaultStorageClass: jsii.String("defaultStorageClass"),
//   	EncryptionType: jsii.String("encryptionType"),
//   	FileShareName: jsii.String("fileShareName"),
//   	GatewayArn: jsii.String("gatewayArn"),
//   	GuessMimeTypeEnabled: jsii.Boolean(false),
//   	KmsEncrypted: jsii.Boolean(false),
//   	KmsKey: jsii.String("kmsKey"),
//   	LocationArn: jsii.String("locationArn"),
//   	NfsFileShareDefaults: &NFSFileShareDefaultsProperty{
//   		DirectoryMode: jsii.String("directoryMode"),
//   		FileMode: jsii.String("fileMode"),
//   		GroupId: jsii.Number(123),
//   		OwnerId: jsii.Number(123),
//   	},
//   	NotificationPolicy: jsii.String("notificationPolicy"),
//   	ObjectAcl: jsii.String("objectAcl"),
//   	ReadOnly: jsii.Boolean(false),
//   	RequesterPays: jsii.Boolean(false),
//   	Role: jsii.String("role"),
//   	Squash: jsii.String("squash"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcEndpointDnsName: jsii.String("vpcEndpointDnsName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-nfsfileshare.html
//
type CfnNFSFileSharePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnNFSFileShareMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNFSFileSharePropsMixin
type jsiiProxy_CfnNFSFileSharePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnNFSFileSharePropsMixin) Props() *CfnNFSFileShareMixinProps {
	var returns *CfnNFSFileShareMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNFSFileSharePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::StorageGateway::NFSFileShare`.
func NewCfnNFSFileSharePropsMixin(props *CfnNFSFileShareMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnNFSFileSharePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNFSFileSharePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNFSFileSharePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_storagegateway.CfnNFSFileSharePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::StorageGateway::NFSFileShare`.
func NewCfnNFSFileSharePropsMixin_Override(c CfnNFSFileSharePropsMixin, props *CfnNFSFileShareMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_storagegateway.CfnNFSFileSharePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnNFSFileSharePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNFSFileSharePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_storagegateway.CfnNFSFileSharePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNFSFileSharePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_storagegateway.CfnNFSFileSharePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNFSFileSharePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnNFSFileSharePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

