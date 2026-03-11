package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGateway::RestApi` resource creates a REST API.
//
// For more information, see [restapi:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateRestApi.html) in the *Amazon API Gateway REST API Reference* .
//
// > On January 1, 2016, the Swagger Specification was donated to the [OpenAPI initiative](https://docs.aws.amazon.com/https://www.openapis.org/) , becoming the foundation of the OpenAPI Specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var body interface{}
//   var mergeStrategy IMergeStrategy
//   var policy interface{}
//
//   cfnRestApiPropsMixin := awscdkcfnpropertymixins.Aws_apigateway.NewCfnRestApiPropsMixin(&CfnRestApiMixinProps{
//   	ApiKeySourceType: jsii.String("apiKeySourceType"),
//   	BinaryMediaTypes: []*string{
//   		jsii.String("binaryMediaTypes"),
//   	},
//   	Body: body,
//   	BodyS3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		ETag: jsii.String("eTag"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   	CloneFrom: jsii.String("cloneFrom"),
//   	Description: jsii.String("description"),
//   	DisableExecuteApiEndpoint: jsii.Boolean(false),
//   	EndpointAccessMode: jsii.String("endpointAccessMode"),
//   	EndpointConfiguration: &EndpointConfigurationProperty{
//   		IpAddressType: jsii.String("ipAddressType"),
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   		VpcEndpointIds: []interface{}{
//   			jsii.String("vpcEndpointIds"),
//   		},
//   	},
//   	FailOnWarnings: jsii.Boolean(false),
//   	MinimumCompressionSize: jsii.Number(123),
//   	Mode: jsii.String("mode"),
//   	Name: jsii.String("name"),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Policy: policy,
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
//
type CfnRestApiPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRestApiMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRestApiPropsMixin
type jsiiProxy_CfnRestApiPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRestApiPropsMixin) Props() *CfnRestApiMixinProps {
	var returns *CfnRestApiMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRestApiPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApiGateway::RestApi`.
func NewCfnRestApiPropsMixin(props *CfnRestApiMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRestApiPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRestApiPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRestApiPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apigateway.CfnRestApiPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApiGateway::RestApi`.
func NewCfnRestApiPropsMixin_Override(c CfnRestApiPropsMixin, props *CfnRestApiMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apigateway.CfnRestApiPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRestApiPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRestApiPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_apigateway.CfnRestApiPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRestApiPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_apigateway.CfnRestApiPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRestApiPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRestApiPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

