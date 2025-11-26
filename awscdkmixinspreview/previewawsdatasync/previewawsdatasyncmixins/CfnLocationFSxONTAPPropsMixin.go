package previewawsdatasyncmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdatasync/previewawsdatasyncmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DataSync::LocationFSxONTAP` resource creates an endpoint for an Amazon FSx for NetApp ONTAP file system.
//
// AWS DataSync can access this endpoint as a source or destination location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLocationFSxONTAPPropsMixin := awscdkmixinspreview.Mixins.NewCfnLocationFSxONTAPPropsMixin(&CfnLocationFSxONTAPMixinProps{
//   	Protocol: &ProtocolProperty{
//   		Nfs: &NFSProperty{
//   			MountOptions: &NfsMountOptionsProperty{
//   				Version: jsii.String("version"),
//   			},
//   		},
//   		Smb: &SMBProperty{
//   			Domain: jsii.String("domain"),
//   			MountOptions: &SmbMountOptionsProperty{
//   				Version: jsii.String("version"),
//   			},
//   			Password: jsii.String("password"),
//   			User: jsii.String("user"),
//   		},
//   	},
//   	SecurityGroupArns: []*string{
//   		jsii.String("securityGroupArns"),
//   	},
//   	StorageVirtualMachineArn: jsii.String("storageVirtualMachineArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationfsxontap.html
//
type CfnLocationFSxONTAPPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLocationFSxONTAPMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLocationFSxONTAPPropsMixin
type jsiiProxy_CfnLocationFSxONTAPPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLocationFSxONTAPPropsMixin) Props() *CfnLocationFSxONTAPMixinProps {
	var returns *CfnLocationFSxONTAPMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationFSxONTAPPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataSync::LocationFSxONTAP`.
func NewCfnLocationFSxONTAPPropsMixin(props *CfnLocationFSxONTAPMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLocationFSxONTAPPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLocationFSxONTAPPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLocationFSxONTAPPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataSync::LocationFSxONTAP`.
func NewCfnLocationFSxONTAPPropsMixin_Override(c CfnLocationFSxONTAPPropsMixin, props *CfnLocationFSxONTAPMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLocationFSxONTAPPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLocationFSxONTAPPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationFSxONTAPPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLocationFSxONTAPPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLocationFSxONTAPPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

