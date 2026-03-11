package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource to create a traffic policy for a Mail Manager ingress endpoint which contains policy statements used to evaluate whether incoming emails should be allowed or denied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMailManagerTrafficPolicyPropsMixin := awscdkcfnpropertymixins.Aws_ses.NewCfnMailManagerTrafficPolicyPropsMixin(&CfnMailManagerTrafficPolicyMixinProps{
//   	DefaultAction: jsii.String("defaultAction"),
//   	MaxMessageSizeBytes: jsii.Number(123),
//   	PolicyStatements: []interface{}{
//   		&PolicyStatementProperty{
//   			Action: jsii.String("action"),
//   			Conditions: []interface{}{
//   				&PolicyConditionProperty{
//   					BooleanExpression: &IngressBooleanExpressionProperty{
//   						Evaluate: &IngressBooleanToEvaluateProperty{
//   							Analysis: &IngressAnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							IsInAddressList: &IngressIsInAddressListProperty{
//   								AddressLists: []*string{
//   									jsii.String("addressLists"),
//   								},
//   								Attribute: jsii.String("attribute"),
//   							},
//   						},
//   						Operator: jsii.String("operator"),
//   					},
//   					IpExpression: &IngressIpv4ExpressionProperty{
//   						Evaluate: &IngressIpToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					Ipv6Expression: &IngressIpv6ExpressionProperty{
//   						Evaluate: &IngressIpv6ToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					StringExpression: &IngressStringExpressionProperty{
//   						Evaluate: &IngressStringToEvaluateProperty{
//   							Analysis: &IngressAnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					TlsExpression: &IngressTlsProtocolExpressionProperty{
//   						Evaluate: &IngressTlsProtocolToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrafficPolicyName: jsii.String("trafficPolicyName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagertrafficpolicy.html
//
type CfnMailManagerTrafficPolicyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMailManagerTrafficPolicyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMailManagerTrafficPolicyPropsMixin
type jsiiProxy_CfnMailManagerTrafficPolicyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMailManagerTrafficPolicyPropsMixin) Props() *CfnMailManagerTrafficPolicyMixinProps {
	var returns *CfnMailManagerTrafficPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMailManagerTrafficPolicyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SES::MailManagerTrafficPolicy`.
func NewCfnMailManagerTrafficPolicyPropsMixin(props *CfnMailManagerTrafficPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMailManagerTrafficPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMailManagerTrafficPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMailManagerTrafficPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SES::MailManagerTrafficPolicy`.
func NewCfnMailManagerTrafficPolicyPropsMixin_Override(c CfnMailManagerTrafficPolicyPropsMixin, props *CfnMailManagerTrafficPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMailManagerTrafficPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMailManagerTrafficPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMailManagerTrafficPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ses.CfnMailManagerTrafficPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMailManagerTrafficPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMailManagerTrafficPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

