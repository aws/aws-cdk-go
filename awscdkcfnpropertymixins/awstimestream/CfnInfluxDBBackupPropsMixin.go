package awstimestream

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awstimestream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for AWS::Timestream::InfluxDBBackup.
//
// An on-demand backup of an Amazon Timestream for InfluxDB DB instance or DB cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnInfluxDBBackupPropsMixin := awscdkcfnpropertymixins.Aws_timestream.NewCfnInfluxDBBackupPropsMixin(&CfnInfluxDBBackupMixinProps{
//   	DbResourceId: jsii.String("dbResourceId"),
//   	Name: jsii.String("name"),
//   	RetentionDays: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbbackup.html
//
type CfnInfluxDBBackupPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnInfluxDBBackupMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInfluxDBBackupPropsMixin
type jsiiProxy_CfnInfluxDBBackupPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnInfluxDBBackupPropsMixin) Props() *CfnInfluxDBBackupMixinProps {
	var returns *CfnInfluxDBBackupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInfluxDBBackupPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Timestream::InfluxDBBackup`.
func NewCfnInfluxDBBackupPropsMixin(props *CfnInfluxDBBackupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnInfluxDBBackupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInfluxDBBackupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInfluxDBBackupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBBackupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Timestream::InfluxDBBackup`.
func NewCfnInfluxDBBackupPropsMixin_Override(c CfnInfluxDBBackupPropsMixin, props *CfnInfluxDBBackupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBBackupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnInfluxDBBackupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInfluxDBBackupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBBackupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInfluxDBBackupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_timestream.CfnInfluxDBBackupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInfluxDBBackupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInfluxDBBackupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

