// Integrations for AWS APIGateway V2
package awscdkapigatewayv2integrationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/internal"
)

// The Application Load Balancer integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpAlbIntegration interface {
	awscdkapigatewayv2alpha.HttpRouteIntegration
	ConnectionType() awscdkapigatewayv2alpha.HttpConnectionType
	SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType)
	HttpMethod() awscdkapigatewayv2alpha.HttpMethod
	SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod)
	IntegrationType() awscdkapigatewayv2alpha.HttpIntegrationType
	SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType)
	PayloadFormatVersion() awscdkapigatewayv2alpha.PayloadFormatVersion
	SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion)
	Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
}

// The jsii proxy struct for HttpAlbIntegration
type jsiiProxy_HttpAlbIntegration struct {
	internal.Type__awscdkapigatewayv2alphaHttpRouteIntegration
}

func (j *jsiiProxy_HttpAlbIntegration) ConnectionType() awscdkapigatewayv2alpha.HttpConnectionType {
	var returns awscdkapigatewayv2alpha.HttpConnectionType
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpAlbIntegration) HttpMethod() awscdkapigatewayv2alpha.HttpMethod {
	var returns awscdkapigatewayv2alpha.HttpMethod
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpAlbIntegration) IntegrationType() awscdkapigatewayv2alpha.HttpIntegrationType {
	var returns awscdkapigatewayv2alpha.HttpIntegrationType
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpAlbIntegration) PayloadFormatVersion() awscdkapigatewayv2alpha.PayloadFormatVersion {
	var returns awscdkapigatewayv2alpha.PayloadFormatVersion
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
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpAlbIntegration",
		[]interface{}{id, listener, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpAlbIntegration_Override(h HttpAlbIntegration, id *string, listener awselasticloadbalancingv2.IApplicationListener, props *HttpAlbIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpAlbIntegration",
		[]interface{}{id, listener, props},
		h,
	)
}

func (j *jsiiProxy_HttpAlbIntegration) SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_HttpAlbIntegration) SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_HttpAlbIntegration) SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType) {
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_HttpAlbIntegration) SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion) {
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

// (experimental) Bind this integration to the route.
// Experimental.
func (h *jsiiProxy_HttpAlbIntegration) Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
	var returns *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig

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
	Method awscdkapigatewayv2alpha.HttpMethod `json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awscdkapigatewayv2alpha.IVpcLink `json:"vpcLink" yaml:"vpcLink"`
}

// The Lambda Proxy integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpLambdaIntegration interface {
	awscdkapigatewayv2alpha.HttpRouteIntegration
	Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
}

// The jsii proxy struct for HttpLambdaIntegration
type jsiiProxy_HttpLambdaIntegration struct {
	internal.Type__awscdkapigatewayv2alphaHttpRouteIntegration
}

// Experimental.
func NewHttpLambdaIntegration(id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) HttpLambdaIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpLambdaIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpLambdaIntegration",
		[]interface{}{id, handler, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpLambdaIntegration_Override(h HttpLambdaIntegration, id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpLambdaIntegration",
		[]interface{}{id, handler, props},
		h,
	)
}

// (experimental) Bind this integration to the route.
// Experimental.
func (h *jsiiProxy_HttpLambdaIntegration) Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
	var returns *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig

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
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping" yaml:"parameterMapping"`
	// Version of the payload sent to the lambda handler.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Experimental.
	PayloadFormatVersion awscdkapigatewayv2alpha.PayloadFormatVersion `json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
}

// The Network Load Balancer integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpNlbIntegration interface {
	awscdkapigatewayv2alpha.HttpRouteIntegration
	ConnectionType() awscdkapigatewayv2alpha.HttpConnectionType
	SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType)
	HttpMethod() awscdkapigatewayv2alpha.HttpMethod
	SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod)
	IntegrationType() awscdkapigatewayv2alpha.HttpIntegrationType
	SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType)
	PayloadFormatVersion() awscdkapigatewayv2alpha.PayloadFormatVersion
	SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion)
	Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
}

// The jsii proxy struct for HttpNlbIntegration
type jsiiProxy_HttpNlbIntegration struct {
	internal.Type__awscdkapigatewayv2alphaHttpRouteIntegration
}

func (j *jsiiProxy_HttpNlbIntegration) ConnectionType() awscdkapigatewayv2alpha.HttpConnectionType {
	var returns awscdkapigatewayv2alpha.HttpConnectionType
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNlbIntegration) HttpMethod() awscdkapigatewayv2alpha.HttpMethod {
	var returns awscdkapigatewayv2alpha.HttpMethod
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNlbIntegration) IntegrationType() awscdkapigatewayv2alpha.HttpIntegrationType {
	var returns awscdkapigatewayv2alpha.HttpIntegrationType
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNlbIntegration) PayloadFormatVersion() awscdkapigatewayv2alpha.PayloadFormatVersion {
	var returns awscdkapigatewayv2alpha.PayloadFormatVersion
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
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpNlbIntegration",
		[]interface{}{id, listener, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpNlbIntegration_Override(h HttpNlbIntegration, id *string, listener awselasticloadbalancingv2.INetworkListener, props *HttpNlbIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpNlbIntegration",
		[]interface{}{id, listener, props},
		h,
	)
}

func (j *jsiiProxy_HttpNlbIntegration) SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_HttpNlbIntegration) SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_HttpNlbIntegration) SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType) {
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_HttpNlbIntegration) SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion) {
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

// (experimental) Bind this integration to the route.
// Experimental.
func (h *jsiiProxy_HttpNlbIntegration) Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
	var returns *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig

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
	Method awscdkapigatewayv2alpha.HttpMethod `json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awscdkapigatewayv2alpha.IVpcLink `json:"vpcLink" yaml:"vpcLink"`
}

// Base options for private integration.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpPrivateIntegrationOptions struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awscdkapigatewayv2alpha.HttpMethod `json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awscdkapigatewayv2alpha.IVpcLink `json:"vpcLink" yaml:"vpcLink"`
}

// The Service Discovery integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpServiceDiscoveryIntegration interface {
	awscdkapigatewayv2alpha.HttpRouteIntegration
	ConnectionType() awscdkapigatewayv2alpha.HttpConnectionType
	SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType)
	HttpMethod() awscdkapigatewayv2alpha.HttpMethod
	SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod)
	IntegrationType() awscdkapigatewayv2alpha.HttpIntegrationType
	SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType)
	PayloadFormatVersion() awscdkapigatewayv2alpha.PayloadFormatVersion
	SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion)
	Bind(_arg *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
}

// The jsii proxy struct for HttpServiceDiscoveryIntegration
type jsiiProxy_HttpServiceDiscoveryIntegration struct {
	internal.Type__awscdkapigatewayv2alphaHttpRouteIntegration
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) ConnectionType() awscdkapigatewayv2alpha.HttpConnectionType {
	var returns awscdkapigatewayv2alpha.HttpConnectionType
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) HttpMethod() awscdkapigatewayv2alpha.HttpMethod {
	var returns awscdkapigatewayv2alpha.HttpMethod
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) IntegrationType() awscdkapigatewayv2alpha.HttpIntegrationType {
	var returns awscdkapigatewayv2alpha.HttpIntegrationType
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) PayloadFormatVersion() awscdkapigatewayv2alpha.PayloadFormatVersion {
	var returns awscdkapigatewayv2alpha.PayloadFormatVersion
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
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpServiceDiscoveryIntegration",
		[]interface{}{id, service, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpServiceDiscoveryIntegration_Override(h HttpServiceDiscoveryIntegration, id *string, service awsservicediscovery.IService, props *HttpServiceDiscoveryIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpServiceDiscoveryIntegration",
		[]interface{}{id, service, props},
		h,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType) {
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration) SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion) {
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

// (experimental) Bind this integration to the route.
// Experimental.
func (h *jsiiProxy_HttpServiceDiscoveryIntegration) Bind(_arg *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
	var returns *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig

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
	Method awscdkapigatewayv2alpha.HttpMethod `json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awscdkapigatewayv2alpha.IVpcLink `json:"vpcLink" yaml:"vpcLink"`
}

// The HTTP Proxy integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpUrlIntegration interface {
	awscdkapigatewayv2alpha.HttpRouteIntegration
	Bind(_arg *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
}

// The jsii proxy struct for HttpUrlIntegration
type jsiiProxy_HttpUrlIntegration struct {
	internal.Type__awscdkapigatewayv2alphaHttpRouteIntegration
}

// Experimental.
func NewHttpUrlIntegration(id *string, url *string, props *HttpUrlIntegrationProps) HttpUrlIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpUrlIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpUrlIntegration",
		[]interface{}{id, url, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpUrlIntegration_Override(h HttpUrlIntegration, id *string, url *string, props *HttpUrlIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpUrlIntegration",
		[]interface{}{id, url, props},
		h,
	)
}

// (experimental) Bind this integration to the route.
// Experimental.
func (h *jsiiProxy_HttpUrlIntegration) Bind(_arg *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
	var returns *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig

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
	Method awscdkapigatewayv2alpha.HttpMethod `json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping" yaml:"parameterMapping"`
}

// Lambda WebSocket Integration.
//
// TODO: EXAMPLE
//
// Experimental.
type WebSocketLambdaIntegration interface {
	awscdkapigatewayv2alpha.WebSocketRouteIntegration
	Bind(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketLambdaIntegration
type jsiiProxy_WebSocketLambdaIntegration struct {
	internal.Type__awscdkapigatewayv2alphaWebSocketRouteIntegration
}

// Experimental.
func NewWebSocketLambdaIntegration(id *string, handler awslambda.IFunction) WebSocketLambdaIntegration {
	_init_.Initialize()

	j := jsiiProxy_WebSocketLambdaIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.WebSocketLambdaIntegration",
		[]interface{}{id, handler},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketLambdaIntegration_Override(w WebSocketLambdaIntegration, id *string, handler awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.WebSocketLambdaIntegration",
		[]interface{}{id, handler},
		w,
	)
}

// (experimental) Bind this integration to the route.
// Experimental.
func (w *jsiiProxy_WebSocketLambdaIntegration) Bind(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig {
	var returns *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig

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
// TODO: EXAMPLE
//
// Experimental.
type WebSocketMockIntegration interface {
	awscdkapigatewayv2alpha.WebSocketRouteIntegration
	Bind(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketMockIntegration
type jsiiProxy_WebSocketMockIntegration struct {
	internal.Type__awscdkapigatewayv2alphaWebSocketRouteIntegration
}

// Experimental.
func NewWebSocketMockIntegration(id *string) WebSocketMockIntegration {
	_init_.Initialize()

	j := jsiiProxy_WebSocketMockIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.WebSocketMockIntegration",
		[]interface{}{id},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketMockIntegration_Override(w WebSocketMockIntegration, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.WebSocketMockIntegration",
		[]interface{}{id},
		w,
	)
}

// (experimental) Bind this integration to the route.
// Experimental.
func (w *jsiiProxy_WebSocketMockIntegration) Bind(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig {
	var returns *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

