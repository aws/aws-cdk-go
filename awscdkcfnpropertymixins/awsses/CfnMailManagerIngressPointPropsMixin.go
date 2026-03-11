package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource to provision an ingress endpoint for receiving email.
//
// An ingress endpoint serves as the entry point for incoming emails, allowing you to define how emails are received and processed within your AWS environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMailManagerIngressPointPropsMixin := awscdkcfnpropertymixins.Aws_ses.NewCfnMailManagerIngressPointPropsMixin(&CfnMailManagerIngressPointMixinProps{
//   	IngressPointConfiguration: &IngressPointConfigurationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   		SmtpPassword: jsii.String("smtpPassword"),
//   	},
//   	IngressPointName: jsii.String("ingressPointName"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		PrivateNetworkConfiguration: &PrivateNetworkConfigurationProperty{
//   			VpcEndpointId: jsii.String("vpcEndpointId"),
//   		},
//   		PublicNetworkConfiguration: &PublicNetworkConfigurationProperty{
//   			IpType: jsii.String("ipType"),
//   		},
//   	},
//   	RuleSetId: jsii.String("ruleSetId"),
//   	StatusToUpdate: jsii.String("statusToUpdate"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrafficPolicyId: jsii.String("trafficPolicyId"),
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html
//
type CfnMailManagerIngressPointPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMailManagerIngressPointMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMailManagerIngressPointPropsMixin
type jsiiProxy_CfnMailManagerIngressPointPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMailManagerIngressPointPropsMixin) Props() *CfnMailManagerIngressPointMixinProps {
	var returns *CfnMailManagerIngressPointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMailManagerIngressPointPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SES::MailManagerIngressPoint`.
func NewCfnMailManagerIngressPointPropsMixin(props *CfnMailManagerIngressPointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMailManagerIngressPointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMailManagerIngressPointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMailManagerIngressPointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerIngressPointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SES::MailManagerIngressPoint`.
func NewCfnMailManagerIngressPointPropsMixin_Override(c CfnMailManagerIngressPointPropsMixin, props *CfnMailManagerIngressPointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerIngressPointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMailManagerIngressPointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMailManagerIngressPointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerIngressPointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMailManagerIngressPointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerIngressPointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMailManagerIngressPointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMailManagerIngressPointPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

