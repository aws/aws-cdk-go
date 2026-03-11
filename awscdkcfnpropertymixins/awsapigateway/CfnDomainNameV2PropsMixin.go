package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGateway::DomainNameV2` resource specifies a custom domain name for your private APIs in API Gateway.
//
// You can use a private custom domain name to provide a URL for your private API that's more intuitive and easier to recall.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var policy interface{}
//
//   cfnDomainNameV2PropsMixin := awscdkcfnpropertymixins.Aws_apigateway.NewCfnDomainNameV2PropsMixin(&CfnDomainNameV2MixinProps{
//   	CertificateArn: jsii.String("certificateArn"),
//   	DomainName: jsii.String("domainName"),
//   	EndpointAccessMode: jsii.String("endpointAccessMode"),
//   	EndpointConfiguration: &EndpointConfigurationProperty{
//   		IpAddressType: jsii.String("ipAddressType"),
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   	Policy: policy,
//   	RoutingMode: jsii.String("routingMode"),
//   	SecurityPolicy: jsii.String("securityPolicy"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainnamev2.html
//
type CfnDomainNameV2PropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDomainNameV2MixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDomainNameV2PropsMixin
type jsiiProxy_CfnDomainNameV2PropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDomainNameV2PropsMixin) Props() *CfnDomainNameV2MixinProps {
	var returns *CfnDomainNameV2MixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainNameV2PropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGateway::DomainNameV2`.
func NewCfnDomainNameV2PropsMixin(props *CfnDomainNameV2MixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDomainNameV2PropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDomainNameV2PropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDomainNameV2PropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apigateway.CfnDomainNameV2PropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGateway::DomainNameV2`.
func NewCfnDomainNameV2PropsMixin_Override(c CfnDomainNameV2PropsMixin, props *CfnDomainNameV2MixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apigateway.CfnDomainNameV2PropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDomainNameV2PropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomainNameV2PropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_apigateway.CfnDomainNameV2PropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainNameV2PropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_apigateway.CfnDomainNameV2PropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainNameV2PropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDomainNameV2PropsMixin) Supports(construct constructs.IConstruct) *bool {
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

