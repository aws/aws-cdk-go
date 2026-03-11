package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds the specified revocation contents to the specified trust store.
//
// You must specify `TrustStoreArn` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTrustStoreRevocationPropsMixin := awscdkcfnpropertymixins.Aws_elasticloadbalancingv2.NewCfnTrustStoreRevocationPropsMixin(&CfnTrustStoreRevocationMixinProps{
//   	RevocationContents: []interface{}{
//   		&RevocationContentProperty{
//   			RevocationType: jsii.String("revocationType"),
//   			S3Bucket: jsii.String("s3Bucket"),
//   			S3Key: jsii.String("s3Key"),
//   			S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		},
//   	},
//   	TrustStoreArn: jsii.String("trustStoreArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-truststorerevocation.html
//
type CfnTrustStoreRevocationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTrustStoreRevocationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTrustStoreRevocationPropsMixin
type jsiiProxy_CfnTrustStoreRevocationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTrustStoreRevocationPropsMixin) Props() *CfnTrustStoreRevocationMixinProps {
	var returns *CfnTrustStoreRevocationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrustStoreRevocationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElasticLoadBalancingV2::TrustStoreRevocation`.
func NewCfnTrustStoreRevocationPropsMixin(props *CfnTrustStoreRevocationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTrustStoreRevocationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTrustStoreRevocationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTrustStoreRevocationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTrustStoreRevocationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElasticLoadBalancingV2::TrustStoreRevocation`.
func NewCfnTrustStoreRevocationPropsMixin_Override(c CfnTrustStoreRevocationPropsMixin, props *CfnTrustStoreRevocationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTrustStoreRevocationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTrustStoreRevocationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrustStoreRevocationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTrustStoreRevocationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrustStoreRevocationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_elasticloadbalancingv2.CfnTrustStoreRevocationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTrustStoreRevocationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTrustStoreRevocationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

