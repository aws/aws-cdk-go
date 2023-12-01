package awscdkapigatewayv2integrationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/internal"
)

// The Network Load Balancer integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   })
//   listener := lb.AddListener(jsii.String("listener"), &BaseNetworkListenerProps{
//   	Port: jsii.Number(80),
//   })
//   listener.AddTargets(jsii.String("target"), &AddNetworkTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
//   	DefaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
// Deprecated.
type HttpNlbIntegration interface {
	awscdkapigatewayv2alpha.HttpRouteIntegration
	// Deprecated.
	ConnectionType() awscdkapigatewayv2alpha.HttpConnectionType
	// Deprecated.
	SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType)
	// Deprecated.
	HttpMethod() awscdkapigatewayv2alpha.HttpMethod
	// Deprecated.
	SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod)
	// Deprecated.
	IntegrationType() awscdkapigatewayv2alpha.HttpIntegrationType
	// Deprecated.
	SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType)
	// Deprecated.
	PayloadFormatVersion() awscdkapigatewayv2alpha.PayloadFormatVersion
	// Deprecated.
	SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion)
	// Bind this integration to the route.
	// Deprecated.
	Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Deprecated.
	CompleteBind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions)
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


// Deprecated.
func NewHttpNlbIntegration(id *string, listener awselasticloadbalancingv2.INetworkListener, props *HttpNlbIntegrationProps) HttpNlbIntegration {
	_init_.Initialize()

	if err := validateNewHttpNlbIntegrationParameters(id, listener, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpNlbIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpNlbIntegration",
		[]interface{}{id, listener, props},
		&j,
	)

	return &j
}

// Deprecated.
func NewHttpNlbIntegration_Override(h HttpNlbIntegration, id *string, listener awselasticloadbalancingv2.INetworkListener, props *HttpNlbIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpNlbIntegration",
		[]interface{}{id, listener, props},
		h,
	)
}

func (j *jsiiProxy_HttpNlbIntegration)SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_HttpNlbIntegration)SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_HttpNlbIntegration)SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType) {
	if err := j.validateSetIntegrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_HttpNlbIntegration)SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion) {
	if err := j.validateSetPayloadFormatVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (h *jsiiProxy_HttpNlbIntegration) Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
	if err := h.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpNlbIntegration) CompleteBind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

