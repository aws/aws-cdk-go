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
// TODO: EXAMPLE
//
// Experimental.
type HttpAlbIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	ConnectionType() awsapigatewayv2.HttpConnectionType
	SetConnectionType(val awsapigatewayv2.HttpConnectionType)
	HttpMethod() awsapigatewayv2.HttpMethod
	SetHttpMethod(val awsapigatewayv2.HttpMethod)
	IntegrationType() awsapigatewayv2.HttpIntegrationType
	SetIntegrationType(val awsapigatewayv2.HttpIntegrationType)
	PayloadFormatVersion() awsapigatewayv2.PayloadFormatVersion
	SetPayloadFormatVersion(val awsapigatewayv2.PayloadFormatVersion)
	Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
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

// Bind this integration to the route.
// Experimental.
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

// Properties to initialize `HttpAlbIntegration`.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpAlbIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `json:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `json:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `json:"vpcLink"`
}

// The Lambda Proxy integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpLambdaIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
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

// Bind this integration to the route.
// Experimental.
func (h *jsiiProxy_HttpLambdaIntegration) Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Lambda Proxy integration properties.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpLambdaIntegrationProps struct {
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `json:"parameterMapping"`
	// Version of the payload sent to the lambda handler.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Experimental.
	PayloadFormatVersion awsapigatewayv2.PayloadFormatVersion `json:"payloadFormatVersion"`
}

// The Network Load Balancer integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpNlbIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	ConnectionType() awsapigatewayv2.HttpConnectionType
	SetConnectionType(val awsapigatewayv2.HttpConnectionType)
	HttpMethod() awsapigatewayv2.HttpMethod
	SetHttpMethod(val awsapigatewayv2.HttpMethod)
	IntegrationType() awsapigatewayv2.HttpIntegrationType
	SetIntegrationType(val awsapigatewayv2.HttpIntegrationType)
	PayloadFormatVersion() awsapigatewayv2.PayloadFormatVersion
	SetPayloadFormatVersion(val awsapigatewayv2.PayloadFormatVersion)
	Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
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

// Bind this integration to the route.
// Experimental.
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

// Properties to initialize `HttpNlbIntegration`.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpNlbIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `json:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `json:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `json:"vpcLink"`
}

// Base options for private integration.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpPrivateIntegrationOptions struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `json:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `json:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `json:"vpcLink"`
}

// The Service Discovery integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpServiceDiscoveryIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	ConnectionType() awsapigatewayv2.HttpConnectionType
	SetConnectionType(val awsapigatewayv2.HttpConnectionType)
	HttpMethod() awsapigatewayv2.HttpMethod
	SetHttpMethod(val awsapigatewayv2.HttpMethod)
	IntegrationType() awsapigatewayv2.HttpIntegrationType
	SetIntegrationType(val awsapigatewayv2.HttpIntegrationType)
	PayloadFormatVersion() awsapigatewayv2.PayloadFormatVersion
	SetPayloadFormatVersion(val awsapigatewayv2.PayloadFormatVersion)
	Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
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

// Bind this integration to the route.
// Experimental.
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

// Properties to initialize `HttpServiceDiscoveryIntegration`.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpServiceDiscoveryIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `json:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `json:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `json:"vpcLink"`
}

// The HTTP Proxy integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpUrlIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
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

// Bind this integration to the route.
// Experimental.
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

// Properties to initialize a new `HttpProxyIntegration`.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpUrlIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `json:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `json:"parameterMapping"`
}

// Lambda WebSocket Integration.
//
// TODO: EXAMPLE
//
// Experimental.
type WebSocketLambdaIntegration interface {
	awsapigatewayv2.WebSocketRouteIntegration
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

// Bind this integration to the route.
// Experimental.
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

