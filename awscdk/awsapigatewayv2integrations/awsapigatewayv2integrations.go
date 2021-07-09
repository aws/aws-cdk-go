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
// Experimental.
type HttpAlbIntegration interface {
	awsapigatewayv2.IHttpRouteIntegration
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
	internal.Type__awsapigatewayv2IHttpRouteIntegration
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
func NewHttpAlbIntegration(props *HttpAlbIntegrationProps) HttpAlbIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpAlbIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpAlbIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpAlbIntegration_Override(h HttpAlbIntegration, props *HttpAlbIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpAlbIntegration",
		[]interface{}{props},
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
// Experimental.
type HttpAlbIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `json:"method"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `json:"vpcLink"`
	// The listener to the application load balancer used for the integration.
	// Experimental.
	Listener awselasticloadbalancingv2.IApplicationListener `json:"listener"`
}

// The Network Load Balancer integration resource for HTTP API.
// Experimental.
type HttpNlbIntegration interface {
	awsapigatewayv2.IHttpRouteIntegration
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
	internal.Type__awsapigatewayv2IHttpRouteIntegration
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
func NewHttpNlbIntegration(props *HttpNlbIntegrationProps) HttpNlbIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpNlbIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpNlbIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpNlbIntegration_Override(h HttpNlbIntegration, props *HttpNlbIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpNlbIntegration",
		[]interface{}{props},
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
// Experimental.
type HttpNlbIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `json:"method"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `json:"vpcLink"`
	// The listener to the network load balancer used for the integration.
	// Experimental.
	Listener awselasticloadbalancingv2.INetworkListener `json:"listener"`
}

// Base options for private integration.
// Experimental.
type HttpPrivateIntegrationOptions struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `json:"method"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `json:"vpcLink"`
}

// The HTTP Proxy integration resource for HTTP API.
// Experimental.
type HttpProxyIntegration interface {
	awsapigatewayv2.IHttpRouteIntegration
	Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
}

// The jsii proxy struct for HttpProxyIntegration
type jsiiProxy_HttpProxyIntegration struct {
	internal.Type__awsapigatewayv2IHttpRouteIntegration
}

// Experimental.
func NewHttpProxyIntegration(props *HttpProxyIntegrationProps) HttpProxyIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpProxyIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpProxyIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpProxyIntegration_Override(h HttpProxyIntegration, props *HttpProxyIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpProxyIntegration",
		[]interface{}{props},
		h,
	)
}

// Bind this integration to the route.
// Experimental.
func (h *jsiiProxy_HttpProxyIntegration) Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
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
// Experimental.
type HttpProxyIntegrationProps struct {
	// The full-qualified HTTP URL for the HTTP integration.
	// Experimental.
	Url *string `json:"url"`
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `json:"method"`
}

// The Service Discovery integration resource for HTTP API.
// Experimental.
type HttpServiceDiscoveryIntegration interface {
	awsapigatewayv2.IHttpRouteIntegration
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
	internal.Type__awsapigatewayv2IHttpRouteIntegration
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
func NewHttpServiceDiscoveryIntegration(props *HttpServiceDiscoveryIntegrationProps) HttpServiceDiscoveryIntegration {
	_init_.Initialize()

	j := jsiiProxy_HttpServiceDiscoveryIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpServiceDiscoveryIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpServiceDiscoveryIntegration_Override(h HttpServiceDiscoveryIntegration, props *HttpServiceDiscoveryIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpServiceDiscoveryIntegration",
		[]interface{}{props},
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
// Experimental.
type HttpServiceDiscoveryIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `json:"method"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `json:"vpcLink"`
	// The discovery service used for the integration.
	// Experimental.
	Service awsservicediscovery.IService `json:"service"`
}

// The Lambda Proxy integration resource for HTTP API.
// Experimental.
type LambdaProxyIntegration interface {
	awsapigatewayv2.IHttpRouteIntegration
	Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
}

// The jsii proxy struct for LambdaProxyIntegration
type jsiiProxy_LambdaProxyIntegration struct {
	internal.Type__awsapigatewayv2IHttpRouteIntegration
}

// Experimental.
func NewLambdaProxyIntegration(props *LambdaProxyIntegrationProps) LambdaProxyIntegration {
	_init_.Initialize()

	j := jsiiProxy_LambdaProxyIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.LambdaProxyIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaProxyIntegration_Override(l LambdaProxyIntegration, props *LambdaProxyIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.LambdaProxyIntegration",
		[]interface{}{props},
		l,
	)
}

// Bind this integration to the route.
// Experimental.
func (l *jsiiProxy_LambdaProxyIntegration) Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Lambda Proxy integration properties.
// Experimental.
type LambdaProxyIntegrationProps struct {
	// The handler for this integration.
	// Experimental.
	Handler awslambda.IFunction `json:"handler"`
	// Version of the payload sent to the lambda handler.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Experimental.
	PayloadFormatVersion awsapigatewayv2.PayloadFormatVersion `json:"payloadFormatVersion"`
}

// Lambda WebSocket Integration.
// Experimental.
type LambdaWebSocketIntegration interface {
	awsapigatewayv2.IWebSocketRouteIntegration
	Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for LambdaWebSocketIntegration
type jsiiProxy_LambdaWebSocketIntegration struct {
	internal.Type__awsapigatewayv2IWebSocketRouteIntegration
}

// Experimental.
func NewLambdaWebSocketIntegration(props *LambdaWebSocketIntegrationProps) LambdaWebSocketIntegration {
	_init_.Initialize()

	j := jsiiProxy_LambdaWebSocketIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.LambdaWebSocketIntegration",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaWebSocketIntegration_Override(l LambdaWebSocketIntegration, props *LambdaWebSocketIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.LambdaWebSocketIntegration",
		[]interface{}{props},
		l,
	)
}

// Bind this integration to the route.
// Experimental.
func (l *jsiiProxy_LambdaWebSocketIntegration) Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig {
	var returns *awsapigatewayv2.WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Lambda WebSocket Integration props.
// Experimental.
type LambdaWebSocketIntegrationProps struct {
	// The handler for this integration.
	// Experimental.
	Handler awslambda.IFunction `json:"handler"`
}

