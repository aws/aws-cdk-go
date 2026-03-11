package awsopensearchserverless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an encryption or network policy to be used by one or more OpenSearch Serverless collections.
//
// Network policies specify access to a collection and its OpenSearch Dashboards endpoint from public networks or specific VPC endpoints. For more information, see [Network access for Amazon OpenSearch Serverless](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-network.html) .
//
// Encryption policies specify a KMS encryption key to assign to particular collections. For more information, see [Encryption at rest for Amazon OpenSearch Serverless](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-encryption.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSecurityPolicyPropsMixin := awscdkcfnpropertymixins.Aws_opensearchserverless.NewCfnSecurityPolicyPropsMixin(&CfnSecurityPolicyMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Policy: jsii.String("policy"),
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-securitypolicy.html
//
type CfnSecurityPolicyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSecurityPolicyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSecurityPolicyPropsMixin
type jsiiProxy_CfnSecurityPolicyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSecurityPolicyPropsMixin) Props() *CfnSecurityPolicyMixinProps {
	var returns *CfnSecurityPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityPolicyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OpenSearchServerless::SecurityPolicy`.
func NewCfnSecurityPolicyPropsMixin(props *CfnSecurityPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSecurityPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSecurityPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSecurityPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_opensearchserverless.CfnSecurityPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OpenSearchServerless::SecurityPolicy`.
func NewCfnSecurityPolicyPropsMixin_Override(c CfnSecurityPolicyPropsMixin, props *CfnSecurityPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_opensearchserverless.CfnSecurityPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSecurityPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSecurityPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_opensearchserverless.CfnSecurityPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_opensearchserverless.CfnSecurityPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSecurityPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSecurityPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

