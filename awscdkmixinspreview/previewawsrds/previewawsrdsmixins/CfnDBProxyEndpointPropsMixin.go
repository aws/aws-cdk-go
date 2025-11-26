package previewawsrdsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrds/previewawsrdsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::RDS::DBProxyEndpoint` resource creates or updates a DB proxy endpoint.
//
// You can use custom proxy endpoints to access a proxy through a different VPC than the proxy's default VPC.
//
// For more information about RDS Proxy, see [AWS::RDS::DBProxy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBProxyEndpointPropsMixin := awscdkmixinspreview.Mixins.NewCfnDBProxyEndpointPropsMixin(&CfnDBProxyEndpointMixinProps{
//   	DbProxyEndpointName: jsii.String("dbProxyEndpointName"),
//   	DbProxyName: jsii.String("dbProxyName"),
//   	EndpointNetworkType: jsii.String("endpointNetworkType"),
//   	Tags: []TagFormatProperty{
//   		&TagFormatProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetRole: jsii.String("targetRole"),
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   	VpcSubnetIds: []*string{
//   		jsii.String("vpcSubnetIds"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html
//
type CfnDBProxyEndpointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDBProxyEndpointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDBProxyEndpointPropsMixin
type jsiiProxy_CfnDBProxyEndpointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDBProxyEndpointPropsMixin) Props() *CfnDBProxyEndpointMixinProps {
	var returns *CfnDBProxyEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBProxyEndpointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RDS::DBProxyEndpoint`.
func NewCfnDBProxyEndpointPropsMixin(props *CfnDBProxyEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDBProxyEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDBProxyEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBProxyEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBProxyEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RDS::DBProxyEndpoint`.
func NewCfnDBProxyEndpointPropsMixin_Override(c CfnDBProxyEndpointPropsMixin, props *CfnDBProxyEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBProxyEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDBProxyEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBProxyEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBProxyEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBProxyEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBProxyEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBProxyEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDBProxyEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

