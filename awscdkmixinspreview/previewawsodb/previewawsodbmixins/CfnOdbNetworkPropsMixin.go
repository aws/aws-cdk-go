package previewawsodbmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsodb/previewawsodbmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ODB::OdbNetwork` resource creates an ODB network.
//
// An ODB network provides the networking foundation for Oracle Database resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOdbNetworkPropsMixin := awscdkmixinspreview.Mixins.NewCfnOdbNetworkPropsMixin(&CfnOdbNetworkMixinProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	BackupSubnetCidr: jsii.String("backupSubnetCidr"),
//   	ClientSubnetCidr: jsii.String("clientSubnetCidr"),
//   	CustomDomainName: jsii.String("customDomainName"),
//   	DefaultDnsPrefix: jsii.String("defaultDnsPrefix"),
//   	DeleteAssociatedResources: jsii.Boolean(false),
//   	DisplayName: jsii.String("displayName"),
//   	S3Access: jsii.String("s3Access"),
//   	S3PolicyDocument: jsii.String("s3PolicyDocument"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ZeroEtlAccess: jsii.String("zeroEtlAccess"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html
//
type CfnOdbNetworkPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOdbNetworkMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOdbNetworkPropsMixin
type jsiiProxy_CfnOdbNetworkPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOdbNetworkPropsMixin) Props() *CfnOdbNetworkMixinProps {
	var returns *CfnOdbNetworkMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetworkPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ODB::OdbNetwork`.
func NewCfnOdbNetworkPropsMixin(props *CfnOdbNetworkMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOdbNetworkPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOdbNetworkPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOdbNetworkPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ODB::OdbNetwork`.
func NewCfnOdbNetworkPropsMixin_Override(c CfnOdbNetworkPropsMixin, props *CfnOdbNetworkMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOdbNetworkPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOdbNetworkPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOdbNetworkPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOdbNetworkPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnOdbNetworkPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

