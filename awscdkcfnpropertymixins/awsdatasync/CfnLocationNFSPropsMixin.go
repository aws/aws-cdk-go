package awsdatasync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DataSync::LocationNFS` resource specifies a Network File System (NFS) file server that AWS DataSync can use as a transfer source or destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnLocationNFSPropsMixin := awscdkcfnpropertymixins.Aws_datasync.NewCfnLocationNFSPropsMixin(&CfnLocationNFSMixinProps{
//   	MountOptions: &MountOptionsProperty{
//   		Version: jsii.String("version"),
//   	},
//   	OnPremConfig: &OnPremConfigProperty{
//   		AgentArns: []interface{}{
//   			jsii.String("agentArns"),
//   		},
//   	},
//   	ServerHostname: jsii.String("serverHostname"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locationnfs.html
//
type CfnLocationNFSPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLocationNFSMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLocationNFSPropsMixin
type jsiiProxy_CfnLocationNFSPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLocationNFSPropsMixin) Props() *CfnLocationNFSMixinProps {
	var returns *CfnLocationNFSMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLocationNFSPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataSync::LocationNFS`.
func NewCfnLocationNFSPropsMixin(props *CfnLocationNFSMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLocationNFSPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLocationNFSPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLocationNFSPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationNFSPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataSync::LocationNFS`.
func NewCfnLocationNFSPropsMixin_Override(c CfnLocationNFSPropsMixin, props *CfnLocationNFSMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationNFSPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLocationNFSPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLocationNFSPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationNFSPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLocationNFSPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationNFSPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLocationNFSPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLocationNFSPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

