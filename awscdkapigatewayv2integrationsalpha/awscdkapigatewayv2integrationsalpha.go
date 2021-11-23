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
	awscdkapigatewayv2alpha.IHttpRouteIntegration
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
	internal.Type__awscdkapigatewayv2alphaIHttpRouteIntegration
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
func NewHttpAlbIntegration(props *HttpAlbIntegrationProps) HttpAlbIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpAlbIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpAlbIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpAlbIntegration_Override(h HttpAlbIntegration, props *HttpAlbIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpAlbIntegration",
		[]interface{}{props},
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
	Method awscdkapigatewayv2alpha.HttpMethod `json:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awscdkapigatewayv2alpha.IVpcLink `json:"vpcLink"`
	// The listener to the application load balancer used for the integration.
	// Experimental.
	Listener awselasticloadbalancingv2.IApplicationListener `json:"listener"`
}

// The Network Load Balancer integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpNlbIntegration interface {
	awscdkapigatewayv2alpha.IHttpRouteIntegration
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
	internal.Type__awscdkapigatewayv2alphaIHttpRouteIntegration
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
func NewHttpNlbIntegration(props *HttpNlbIntegrationProps) HttpNlbIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpNlbIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpNlbIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpNlbIntegration_Override(h HttpNlbIntegration, props *HttpNlbIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpNlbIntegration",
		[]interface{}{props},
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
	Method awscdkapigatewayv2alpha.HttpMethod `json:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awscdkapigatewayv2alpha.IVpcLink `json:"vpcLink"`
	// The listener to the network load balancer used for the integration.
	// Experimental.
	Listener awselasticloadbalancingv2.INetworkListener `json:"listener"`
}

// Base options for private integration.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpPrivateIntegrationOptions struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awscdkapigatewayv2alpha.HttpMethod `json:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awscdkapigatewayv2alpha.IVpcLink `json:"vpcLink"`
}

// The HTTP Proxy integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpProxyIntegration interface {
	awscdkapigatewayv2alpha.IHttpRouteIntegration
	Bind(_arg *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
}

// The jsii proxy struct for HttpProxyIntegration
type jsiiProxy_HttpProxyIntegration struct {
	internal.Type__awscdkapigatewayv2alphaIHttpRouteIntegration
}

// Experimental.
func NewHttpProxyIntegration(props *HttpProxyIntegrationProps) HttpProxyIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpProxyIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpProxyIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpProxyIntegration_Override(h HttpProxyIntegration, props *HttpProxyIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpProxyIntegration",
		[]interface{}{props},
		h,
	)
}

// (experimental) Bind this integration to the route.
// Experimental.
func (h *jsiiProxy_HttpProxyIntegration) Bind(_arg *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
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
type HttpProxyIntegrationProps struct {
	// The full-qualified HTTP URL for the HTTP integration.
	// Experimental.
	Url *string `json:"url"`
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awscdkapigatewayv2alpha.HttpMethod `json:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping"`
}

// The Service Discovery integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type HttpServiceDiscoveryIntegration interface {
	awscdkapigatewayv2alpha.IHttpRouteIntegration
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
	internal.Type__awscdkapigatewayv2alphaIHttpRouteIntegration
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
func NewHttpServiceDiscoveryIntegration(props *HttpServiceDiscoveryIntegrationProps) HttpServiceDiscoveryIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpServiceDiscoveryIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpServiceDiscoveryIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpServiceDiscoveryIntegration_Override(h HttpServiceDiscoveryIntegration, props *HttpServiceDiscoveryIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpServiceDiscoveryIntegration",
		[]interface{}{props},
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
	Method awscdkapigatewayv2alpha.HttpMethod `json:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `json:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awscdkapigatewayv2alpha.IVpcLink `json:"vpcLink"`
	// The discovery service used for the integration.
	// Experimental.
	Service awsservicediscovery.IService `json:"service"`
}

// The Lambda Proxy integration resource for HTTP API.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaProxyIntegration interface {
	awscdkapigatewayv2alpha.IHttpRouteIntegration
	Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
}

// The jsii proxy struct for LambdaProxyIntegration
type jsiiProxy_LambdaProxyIntegration struct {
	internal.Type__awscdkapigatewayv2alphaIHttpRouteIntegration
}

// Experimental.
func NewLambdaProxyIntegration(props *LambdaProxyIntegrationProps) LambdaProxyIntegration {
	_init_.Initialize()

	j := jsiiProxy_LambdaProxyIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.LambdaProxyIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaProxyIntegration_Override(l LambdaProxyIntegration, props *LambdaProxyIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.LambdaProxyIntegration",
		[]interface{}{props},
		l,
	)
}

// (experimental) Bind this integration to the route.
// Experimental.
func (l *jsiiProxy_LambdaProxyIntegration) Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
	var returns *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		l,
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
type LambdaProxyIntegrationProps struct {
	// The handler for this integration.
	// Experimental.
	Handler awslambda.IFunction `json:"handler"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `json:"parameterMapping"`
	// Version of the payload sent to the lambda handler.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Experimental.
	PayloadFormatVersion awscdkapigatewayv2alpha.PayloadFormatVersion `json:"payloadFormatVersion"`
}

// Lambda WebSocket Integration.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaWebSocketIntegration interface {
	awscdkapigatewayv2alpha.IWebSocketRouteIntegration
	Bind(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for LambdaWebSocketIntegration
type jsiiProxy_LambdaWebSocketIntegration struct {
	internal.Type__awscdkapigatewayv2alphaIWebSocketRouteIntegration
}

// Experimental.
func NewLambdaWebSocketIntegration(props *LambdaWebSocketIntegrationProps) LambdaWebSocketIntegration {
	_init_.Initialize()

	j := jsiiProxy_LambdaWebSocketIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.LambdaWebSocketIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaWebSocketIntegration_Override(l LambdaWebSocketIntegration, props *LambdaWebSocketIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.LambdaWebSocketIntegration",
		[]interface{}{props},
		l,
	)
}

// (experimental) Bind this integration to the route.
// Experimental.
func (l *jsiiProxy_LambdaWebSocketIntegration) Bind(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig {
	var returns *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Lambda WebSocket Integration props.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaWebSocketIntegrationProps struct {
	// The handler for this integration.
	// Experimental.
	Handler awslambda.IFunction `json:"handler"`
}

