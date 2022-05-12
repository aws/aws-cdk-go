package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2integrations/internal"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery"
)

// The Application Load Balancer integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var lb applicationLoadBalancer
//
//   listener := lb.addListener(jsii.String("listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addApplicationTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdk.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener, &httpAlbIntegrationProps{
//   		parameterMapping: apigwv2.NewParameterMapping().custom(jsii.String("myKey"), jsii.String("myValue")),
//   	}),
//   })
//
// Experimental.
type HttpAlbIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	// Experimental.
	ConnectionType() awsapigatewayv2.HttpConnectionType
	// Experimental.
	SetConnectionType(val awsapigatewayv2.HttpConnectionType)
	// Experimental.
	HttpMethod() awsapigatewayv2.HttpMethod
	// Experimental.
	SetHttpMethod(val awsapigatewayv2.HttpMethod)
	// Experimental.
	IntegrationType() awsapigatewayv2.HttpIntegrationType
	// Experimental.
	SetIntegrationType(val awsapigatewayv2.HttpIntegrationType)
	// Experimental.
	PayloadFormatVersion() awsapigatewayv2.PayloadFormatVersion
	// Experimental.
	SetPayloadFormatVersion(val awsapigatewayv2.PayloadFormatVersion)
	// Bind this integration to the route.
	// Experimental.
	Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Experimental.
	CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpAlbIntegration
type jsiiProxy_HttpAlbIntegration struct {
	internal.Type__awsapigatewayv2HttpRouteIntegration
}

func (j *jsiiProxy_HttpAlbIntegration) ConnectionType() awsapigatewayv2.HttpConnectionType {
	var returns awsapigatewayv2.HttpConnectionType
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpAlbIntegration) HttpMethod() awsapigatewayv2.HttpMethod {
	var returns awsapigatewayv2.HttpMethod
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpAlbIntegration) IntegrationType() awsapigatewayv2.HttpIntegrationType {
	var returns awsapigatewayv2.HttpIntegrationType
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpAlbIntegration) PayloadFormatVersion() awsapigatewayv2.PayloadFormatVersion {
	var returns awsapigatewayv2.PayloadFormatVersion
	_jsii_.Get(
		j,
		"payloadFormatVersion",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpAlbIntegration(id *string, listener awselasticloadbalancingv2.IApplicationListener, props *HttpAlbIntegrationProps) HttpAlbIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpAlbIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpAlbIntegration",
		[]interface{}{id, listener, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpAlbIntegration_Override(h HttpAlbIntegration, id *string, listener awselasticloadbalancingv2.IApplicationListener, props *HttpAlbIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpAlbIntegration",
		[]interface{}{id, listener, props},
		h,
	)
}

func (j *jsiiProxy_HttpAlbIntegration) SetConnectionType(val awsapigatewayv2.HttpConnectionType) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_HttpAlbIntegration) SetHttpMethod(val awsapigatewayv2.HttpMethod) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_HttpAlbIntegration) SetIntegrationType(val awsapigatewayv2.HttpIntegrationType) {
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_HttpAlbIntegration) SetPayloadFormatVersion(val awsapigatewayv2.PayloadFormatVersion) {
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (h *jsiiProxy_HttpAlbIntegration) Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpAlbIntegration) CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

// Properties to initialize `HttpAlbIntegration`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var lb applicationLoadBalancer
//
//   listener := lb.addListener(jsii.String("listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addApplicationTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdk.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener, &httpAlbIntegrationProps{
//   		parameterMapping: apigwv2.NewParameterMapping().appendHeader(jsii.String("header2"), apigwv2.mappingValue.requestHeader(jsii.String("header1"))).removeHeader(jsii.String("header1")),
//   	}),
//   })
//
// Experimental.
type HttpAlbIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

// The Lambda Proxy integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var booksDefaultFn function
//
//   booksIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), booksDefaultFn)
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   httpApi.addRoutes(&addRoutesOptions{
//   	path: jsii.String("/books"),
//   	methods: []httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   	integration: booksIntegration,
//   })
//
// Experimental.
type HttpLambdaIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	// Bind this integration to the route.
	// Experimental.
	Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Experimental.
	CompleteBind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpLambdaIntegration
type jsiiProxy_HttpLambdaIntegration struct {
	internal.Type__awsapigatewayv2HttpRouteIntegration
}

// Experimental.
func NewHttpLambdaIntegration(id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) HttpLambdaIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpLambdaIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpLambdaIntegration",
		[]interface{}{id, handler, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpLambdaIntegration_Override(h HttpLambdaIntegration, id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpLambdaIntegration",
		[]interface{}{id, handler, props},
		h,
	)
}

func (h *jsiiProxy_HttpLambdaIntegration) Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpLambdaIntegration) CompleteBind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{options},
	)
}

// Lambda Proxy integration properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterMapping parameterMapping
//   var payloadFormatVersion payloadFormatVersion
//
//   httpLambdaIntegrationProps := &httpLambdaIntegrationProps{
//   	parameterMapping: parameterMapping,
//   	payloadFormatVersion: payloadFormatVersion,
//   }
//
// Experimental.
type HttpLambdaIntegrationProps struct {
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Version of the payload sent to the lambda handler.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Experimental.
	PayloadFormatVersion awsapigatewayv2.PayloadFormatVersion `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
}

// The Network Load Balancer integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   listener := lb.addListener(jsii.String("listener"), &baseNetworkListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addNetworkTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdk.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
// Experimental.
type HttpNlbIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	// Experimental.
	ConnectionType() awsapigatewayv2.HttpConnectionType
	// Experimental.
	SetConnectionType(val awsapigatewayv2.HttpConnectionType)
	// Experimental.
	HttpMethod() awsapigatewayv2.HttpMethod
	// Experimental.
	SetHttpMethod(val awsapigatewayv2.HttpMethod)
	// Experimental.
	IntegrationType() awsapigatewayv2.HttpIntegrationType
	// Experimental.
	SetIntegrationType(val awsapigatewayv2.HttpIntegrationType)
	// Experimental.
	PayloadFormatVersion() awsapigatewayv2.PayloadFormatVersion
	// Experimental.
	SetPayloadFormatVersion(val awsapigatewayv2.PayloadFormatVersion)
	// Bind this integration to the route.
	// Experimental.
	Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Experimental.
	CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpNlbIntegration
type jsiiProxy_HttpNlbIntegration struct {
	internal.Type__awsapigatewayv2HttpRouteIntegration
}

func (j *jsiiProxy_HttpNlbIntegration) ConnectionType() awsapigatewayv2.HttpConnectionType {
	var returns awsapigatewayv2.HttpConnectionType
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNlbIntegration) HttpMethod() awsapigatewayv2.HttpMethod {
	var returns awsapigatewayv2.HttpMethod
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNlbIntegration) IntegrationType() awsapigatewayv2.HttpIntegrationType {
	var returns awsapigatewayv2.HttpIntegrationType
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNlbIntegration) PayloadFormatVersion() awsapigatewayv2.PayloadFormatVersion {
	var returns awsapigatewayv2.PayloadFormatVersion
	_jsii_.Get(
		j,
		"payloadFormatVersion",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpNlbIntegration(id *string, listener awselasticloadbalancingv2.INetworkListener, props *HttpNlbIntegrationProps) HttpNlbIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpNlbIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpNlbIntegration",
		[]interface{}{id, listener, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpNlbIntegration_Override(h HttpNlbIntegration, id *string, listener awselasticloadbalancingv2.INetworkListener, props *HttpNlbIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpNlbIntegration",
		[]interface{}{id, listener, props},
		h,
	)
}

func (j *jsiiProxy_HttpNlbIntegration) SetConnectionType(val awsapigatewayv2.HttpConnectionType) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_HttpNlbIntegration) SetHttpMethod(val awsapigatewayv2.HttpMethod) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_HttpNlbIntegration) SetIntegrationType(val awsapigatewayv2.HttpIntegrationType) {
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_HttpNlbIntegration) SetPayloadFormatVersion(val awsapigatewayv2.PayloadFormatVersion) {
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (h *jsiiProxy_HttpNlbIntegration) Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpNlbIntegration) CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

// Properties to initialize `HttpNlbIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterMapping parameterMapping
//   var vpcLink vpcLink
//
//   httpNlbIntegrationProps := &httpNlbIntegrationProps{
//   	method: awscdk.Aws_apigatewayv2.httpMethod_ANY,
//   	parameterMapping: parameterMapping,
//   	secureServerName: jsii.String("secureServerName"),
//   	vpcLink: vpcLink,
//   }
//
// Experimental.
type HttpNlbIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

// Base options for private integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterMapping parameterMapping
//   var vpcLink vpcLink
//
//   httpPrivateIntegrationOptions := &httpPrivateIntegrationOptions{
//   	method: awscdk.Aws_apigatewayv2.httpMethod_ANY,
//   	parameterMapping: parameterMapping,
//   	secureServerName: jsii.String("secureServerName"),
//   	vpcLink: vpcLink,
//   }
//
// Experimental.
type HttpPrivateIntegrationOptions struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

// The Service Discovery integration resource for HTTP API.
//
// Example:
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   vpcLink := apigwv2.NewVpcLink(this, jsii.String("VpcLink"), &vpcLinkProps{
//   	vpc: vpc,
//   })
//   namespace := servicediscovery.NewPrivateDnsNamespace(this, jsii.String("Namespace"), &privateDnsNamespaceProps{
//   	name: jsii.String("boobar.com"),
//   	vpc: vpc,
//   })
//   service := namespace.createService(jsii.String("Service"))
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdk.NewHttpServiceDiscoveryIntegration(jsii.String("DefaultIntegration"), service, &httpServiceDiscoveryIntegrationProps{
//   		vpcLink: vpcLink,
//   	}),
//   })
//
// Experimental.
type HttpServiceDiscoveryIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	// Experimental.
	ConnectionType() awsapigatewayv2.HttpConnectionType
	// Experimental.
	SetConnectionType(val awsapigatewayv2.HttpConnectionType)
	// Experimental.
	HttpMethod() awsapigatewayv2.HttpMethod
	// Experimental.
	SetHttpMethod(val awsapigatewayv2.HttpMethod)
	// Experimental.
	IntegrationType() awsapigatewayv2.HttpIntegrationType
	// Experimental.
	SetIntegrationType(val awsapigatewayv2.HttpIntegrationType)
	// Experimental.
	PayloadFormatVersion() awsapigatewayv2.PayloadFormatVersion
	// Experimental.
	SetPayloadFormatVersion(val awsapigatewayv2.PayloadFormatVersion)
	// Bind this integration to the route.
	// Experimental.
	Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Experimental.
	CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpServiceDiscoveryIntegration
type jsiiProxy_HttpServiceDiscoveryIntegration struct {
	internal.Type__awsapigatewayv2HttpRouteIntegration
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) ConnectionType() awsapigatewayv2.HttpConnectionType {
	var returns awsapigatewayv2.HttpConnectionType
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) HttpMethod() awsapigatewayv2.HttpMethod {
	var returns awsapigatewayv2.HttpMethod
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) IntegrationType() awsapigatewayv2.HttpIntegrationType {
	var returns awsapigatewayv2.HttpIntegrationType
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) PayloadFormatVersion() awsapigatewayv2.PayloadFormatVersion {
	var returns awsapigatewayv2.PayloadFormatVersion
	_jsii_.Get(
		j,
		"payloadFormatVersion",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpServiceDiscoveryIntegration(id *string, service awsservicediscovery.IService, props *HttpServiceDiscoveryIntegrationProps) HttpServiceDiscoveryIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpServiceDiscoveryIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpServiceDiscoveryIntegration",
		[]interface{}{id, service, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpServiceDiscoveryIntegration_Override(h HttpServiceDiscoveryIntegration, id *string, service awsservicediscovery.IService, props *HttpServiceDiscoveryIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpServiceDiscoveryIntegration",
		[]interface{}{id, service, props},
		h,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) SetConnectionType(val awsapigatewayv2.HttpConnectionType) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) SetHttpMethod(val awsapigatewayv2.HttpMethod) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) SetIntegrationType(val awsapigatewayv2.HttpIntegrationType) {
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) SetPayloadFormatVersion(val awsapigatewayv2.PayloadFormatVersion) {
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

// Properties to initialize `HttpServiceDiscoveryIntegration`.
//
// Example:
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   vpcLink := apigwv2.NewVpcLink(this, jsii.String("VpcLink"), &vpcLinkProps{
//   	vpc: vpc,
//   })
//   namespace := servicediscovery.NewPrivateDnsNamespace(this, jsii.String("Namespace"), &privateDnsNamespaceProps{
//   	name: jsii.String("boobar.com"),
//   	vpc: vpc,
//   })
//   service := namespace.createService(jsii.String("Service"))
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdk.NewHttpServiceDiscoveryIntegration(jsii.String("DefaultIntegration"), service, &httpServiceDiscoveryIntegrationProps{
//   		vpcLink: vpcLink,
//   	}),
//   })
//
// Experimental.
type HttpServiceDiscoveryIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

// The HTTP Proxy integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // This function handles your auth logic
//   var authHandler function
//
//
//   authorizer := awscdk.NewHttpLambdaAuthorizer(jsii.String("BooksAuthorizer"), authHandler, &httpLambdaAuthorizerProps{
//   	responseTypes: []httpLambdaResponseType{
//   		awscdk.HttpLambdaResponseType_SIMPLE,
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	authorizer: authorizer,
//   })
//
// Experimental.
type HttpUrlIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	// Bind this integration to the route.
	// Experimental.
	Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Experimental.
	CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpUrlIntegration
type jsiiProxy_HttpUrlIntegration struct {
	internal.Type__awsapigatewayv2HttpRouteIntegration
}

// Experimental.
func NewHttpUrlIntegration(id *string, url *string, props *HttpUrlIntegrationProps) HttpUrlIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpUrlIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpUrlIntegration",
		[]interface{}{id, url, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpUrlIntegration_Override(h HttpUrlIntegration, id *string, url *string, props *HttpUrlIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpUrlIntegration",
		[]interface{}{id, url, props},
		h,
	)
}

func (h *jsiiProxy_HttpUrlIntegration) Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpUrlIntegration) CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

// Properties to initialize a new `HttpProxyIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterMapping parameterMapping
//
//   httpUrlIntegrationProps := &httpUrlIntegrationProps{
//   	method: awscdk.Aws_apigatewayv2.httpMethod_ANY,
//   	parameterMapping: parameterMapping,
//   }
//
// Experimental.
type HttpUrlIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
}

// Lambda WebSocket Integration.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var messageHandler function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &webSocketStageProps{
//   	webSocketApi: webSocketApi,
//   	stageName: jsii.String("dev"),
//   	autoDeploy: jsii.Boolean(true),
//   })
//   webSocketApi.addRoute(jsii.String("sendmessage"), &webSocketRouteOptions{
//   	integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
//   })
//
// Experimental.
type WebSocketLambdaIntegration interface {
	awsapigatewayv2.WebSocketRouteIntegration
	// Bind this integration to the route.
	// Experimental.
	Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketLambdaIntegration
type jsiiProxy_WebSocketLambdaIntegration struct {
	internal.Type__awsapigatewayv2WebSocketRouteIntegration
}

// Experimental.
func NewWebSocketLambdaIntegration(id *string, handler awslambda.IFunction) WebSocketLambdaIntegration {
	_init_.Initialize()

	j := jsiiProxy_WebSocketLambdaIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.WebSocketLambdaIntegration",
		[]interface{}{id, handler},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketLambdaIntegration_Override(w WebSocketLambdaIntegration, id *string, handler awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.WebSocketLambdaIntegration",
		[]interface{}{id, handler},
		w,
	)
}

func (w *jsiiProxy_WebSocketLambdaIntegration) Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig {
	var returns *awsapigatewayv2.WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Mock WebSocket Integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSocketMockIntegration := awscdk.Aws_apigatewayv2_integrations.NewWebSocketMockIntegration(jsii.String("id"))
//
// Experimental.
type WebSocketMockIntegration interface {
	awsapigatewayv2.WebSocketRouteIntegration
	// Bind this integration to the route.
	// Experimental.
	Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketMockIntegration
type jsiiProxy_WebSocketMockIntegration struct {
	internal.Type__awsapigatewayv2WebSocketRouteIntegration
}

// Experimental.
func NewWebSocketMockIntegration(id *string) WebSocketMockIntegration {
	_init_.Initialize()

	j := jsiiProxy_WebSocketMockIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.WebSocketMockIntegration",
		[]interface{}{id},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketMockIntegration_Override(w WebSocketMockIntegration, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.WebSocketMockIntegration",
		[]interface{}{id},
		w,
	)
}

func (w *jsiiProxy_WebSocketMockIntegration) Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig {
	var returns *awsapigatewayv2.WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

