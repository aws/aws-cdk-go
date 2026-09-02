package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::SageMaker::Workforce.
//
// Use to create a private workforce that you can use to label your training data using Amazon SageMaker Ground Truth.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWorkforcePropsMixin := awscdkcfnpropertymixins.Aws_sagemaker.NewCfnWorkforcePropsMixin(&CfnWorkforceMixinProps{
//   	CognitoConfig: &CognitoConfigProperty{
//   		ClientId: jsii.String("clientId"),
//   		UserPool: jsii.String("userPool"),
//   	},
//   	IpAddressType: jsii.String("ipAddressType"),
//   	OidcConfig: &OidcConfigProperty{
//   		AuthenticationRequestExtraParams: map[string]*string{
//   			"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   		},
//   		AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   		ClientId: jsii.String("clientId"),
//   		ClientSecret: jsii.String("clientSecret"),
//   		Issuer: jsii.String("issuer"),
//   		JwksUri: jsii.String("jwksUri"),
//   		LogoutEndpoint: jsii.String("logoutEndpoint"),
//   		Scope: jsii.String("scope"),
//   		TokenEndpoint: jsii.String("tokenEndpoint"),
//   		UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   	},
//   	SourceIpConfig: &SourceIpConfigProperty{
//   		Cidrs: []*string{
//   			jsii.String("cidrs"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkforceName: jsii.String("workforceName"),
//   	WorkforceVpcConfig: &WorkforceVpcConfigRequestProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workforce.html
//
type CfnWorkforcePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWorkforceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkforcePropsMixin
type jsiiProxy_CfnWorkforcePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnWorkforcePropsMixin) Props() *CfnWorkforceMixinProps {
	var returns *CfnWorkforceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkforcePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::Workforce`.
func NewCfnWorkforcePropsMixin(props *CfnWorkforceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWorkforcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkforcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkforcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnWorkforcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::Workforce`.
func NewCfnWorkforcePropsMixin_Override(c CfnWorkforcePropsMixin, props *CfnWorkforceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnWorkforcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWorkforcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkforcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnWorkforcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkforcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_sagemaker.CfnWorkforcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkforcePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWorkforcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

