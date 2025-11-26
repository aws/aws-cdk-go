package previewawsrdsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrds/previewawsrdsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::RDS::DBProxyTargetGroup` resource represents a set of RDS DB instances, Aurora DB clusters, or both that a proxy can connect to.
//
// Currently, each target group is associated with exactly one RDS DB instance or Aurora DB cluster.
//
// This data type is used as a response element in the `DescribeDBProxyTargetGroups` action.
//
// For information about RDS Proxy for Amazon RDS, see [Managing Connections with Amazon RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html) in the *Amazon RDS User Guide* .
//
// For information about RDS Proxy for Amazon Aurora, see [Managing Connections with Amazon RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-proxy.html) in the *Amazon Aurora User Guide* .
//
// For a sample template that creates a DB proxy and registers a DB instance, see [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html#aws-resource-rds-dbproxy--examples) in AWS::RDS::DBProxy.
//
// > Limitations apply to RDS Proxy, including DB engine version limitations and AWS Region limitations.
// >
// > For information about limitations that apply to RDS Proxy for Amazon RDS, see [Limitations for RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html#rds-proxy.limitations) in the *Amazon RDS User Guide* .
// >
// > For information about that apply to RDS Proxy for Amazon Aurora, see [Limitations for RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-proxy.html#rds-proxy.limitations) in the *Amazon Aurora User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBProxyTargetGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnDBProxyTargetGroupPropsMixin(&CfnDBProxyTargetGroupMixinProps{
//   	ConnectionPoolConfigurationInfo: &ConnectionPoolConfigurationInfoFormatProperty{
//   		ConnectionBorrowTimeout: jsii.Number(123),
//   		InitQuery: jsii.String("initQuery"),
//   		MaxConnectionsPercent: jsii.Number(123),
//   		MaxIdleConnectionsPercent: jsii.Number(123),
//   		SessionPinningFilters: []*string{
//   			jsii.String("sessionPinningFilters"),
//   		},
//   	},
//   	DbClusterIdentifiers: []*string{
//   		jsii.String("dbClusterIdentifiers"),
//   	},
//   	DbInstanceIdentifiers: []*string{
//   		jsii.String("dbInstanceIdentifiers"),
//   	},
//   	DbProxyName: jsii.String("dbProxyName"),
//   	TargetGroupName: jsii.String("targetGroupName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html
//
type CfnDBProxyTargetGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDBProxyTargetGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDBProxyTargetGroupPropsMixin
type jsiiProxy_CfnDBProxyTargetGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDBProxyTargetGroupPropsMixin) Props() *CfnDBProxyTargetGroupMixinProps {
	var returns *CfnDBProxyTargetGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyTargetGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RDS::DBProxyTargetGroup`.
func NewCfnDBProxyTargetGroupPropsMixin(props *CfnDBProxyTargetGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDBProxyTargetGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDBProxyTargetGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBProxyTargetGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBProxyTargetGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RDS::DBProxyTargetGroup`.
func NewCfnDBProxyTargetGroupPropsMixin_Override(c CfnDBProxyTargetGroupPropsMixin, props *CfnDBProxyTargetGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBProxyTargetGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDBProxyTargetGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBProxyTargetGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBProxyTargetGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBProxyTargetGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBProxyTargetGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBProxyTargetGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDBProxyTargetGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

