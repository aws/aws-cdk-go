package previewawsroute53recoveryreadinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53recoveryreadiness/previewawsroute53recoveryreadinessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a resource set in Amazon Route 53 Application Recovery Controller.
//
// A resource set is a set of resources of one type, such as Network Load Balancers, that span multiple cells. You can associate a resource set with a readiness check to have Route 53 ARC continually monitor the resources in the set for failover readiness.
//
// You typically create a resource set and a readiness check for each supported type of AWS resource in your application.
//
// For more information, see [Readiness checks, resource sets, and readiness scopes](https://docs.aws.amazon.com/r53recovery/latest/dg/recovery-readiness.recovery-groups.readiness-scope.html) in the Amazon Route 53 Application Recovery Controller Developer Guide.
//
// Route 53 ARC Readiness supports us-east-1 and us-west-2 AWS Regions only.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnResourceSetPropsMixin(&CfnResourceSetMixinProps{
//   	Resources: []interface{}{
//   		&ResourceProperty{
//   			ComponentId: jsii.String("componentId"),
//   			DnsTargetResource: &DNSTargetResourceProperty{
//   				DomainName: jsii.String("domainName"),
//   				HostedZoneArn: jsii.String("hostedZoneArn"),
//   				RecordSetId: jsii.String("recordSetId"),
//   				RecordType: jsii.String("recordType"),
//   				TargetResource: &TargetResourceProperty{
//   					NlbResource: &NLBResourceProperty{
//   						Arn: jsii.String("arn"),
//   					},
//   					R53Resource: &R53ResourceRecordProperty{
//   						DomainName: jsii.String("domainName"),
//   						RecordSetId: jsii.String("recordSetId"),
//   					},
//   				},
//   			},
//   			ReadinessScopes: []*string{
//   				jsii.String("readinessScopes"),
//   			},
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   	},
//   	ResourceSetName: jsii.String("resourceSetName"),
//   	ResourceSetType: jsii.String("resourceSetType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoveryreadiness-resourceset.html
//
type CfnResourceSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResourceSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResourceSetPropsMixin
type jsiiProxy_CfnResourceSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResourceSetPropsMixin) Props() *CfnResourceSetMixinProps {
	var returns *CfnResourceSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53RecoveryReadiness::ResourceSet`.
func NewCfnResourceSetPropsMixin(props *CfnResourceSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResourceSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResourceSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResourceSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.mixins.CfnResourceSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53RecoveryReadiness::ResourceSet`.
func NewCfnResourceSetPropsMixin_Override(c CfnResourceSetPropsMixin, props *CfnResourceSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.mixins.CfnResourceSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResourceSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResourceSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.mixins.CfnResourceSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.mixins.CfnResourceSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceSetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnResourceSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

