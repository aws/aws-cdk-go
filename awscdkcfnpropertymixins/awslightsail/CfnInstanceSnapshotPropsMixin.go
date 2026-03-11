package awslightsail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes an instance snapshot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnInstanceSnapshotPropsMixin := awscdkcfnpropertymixins.Aws_lightsail.NewCfnInstanceSnapshotPropsMixin(&CfnInstanceSnapshotMixinProps{
//   	InstanceName: jsii.String("instanceName"),
//   	InstanceSnapshotName: jsii.String("instanceSnapshotName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instancesnapshot.html
//
type CfnInstanceSnapshotPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnInstanceSnapshotMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInstanceSnapshotPropsMixin
type jsiiProxy_CfnInstanceSnapshotPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnInstanceSnapshotPropsMixin) Props() *CfnInstanceSnapshotMixinProps {
	var returns *CfnInstanceSnapshotMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceSnapshotPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lightsail::InstanceSnapshot`.
func NewCfnInstanceSnapshotPropsMixin(props *CfnInstanceSnapshotMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnInstanceSnapshotPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInstanceSnapshotPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstanceSnapshotPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstanceSnapshotPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lightsail::InstanceSnapshot`.
func NewCfnInstanceSnapshotPropsMixin_Override(c CfnInstanceSnapshotPropsMixin, props *CfnInstanceSnapshotMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstanceSnapshotPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnInstanceSnapshotPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstanceSnapshotPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstanceSnapshotPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstanceSnapshotPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_lightsail.CfnInstanceSnapshotPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstanceSnapshotPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInstanceSnapshotPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

