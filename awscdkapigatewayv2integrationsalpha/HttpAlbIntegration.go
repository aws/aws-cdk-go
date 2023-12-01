package awscdkapigatewayv2integrationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/internal"
)

// The Application Load Balancer integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("lb"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   })
//   listener := lb.AddListener(jsii.String("listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//   })
//   listener.AddTargets(jsii.String("target"), &AddApplicationTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
//   	DefaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
// Deprecated.
type HttpAlbIntegration interface {
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


// Deprecated.
func NewHttpAlbIntegration(id *string, listener awselasticloadbalancingv2.IApplicationListener, props *HttpAlbIntegrationProps) HttpAlbIntegration {
	_init_.Initialize()

	if err := validateNewHttpAlbIntegrationParameters(id, listener, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpAlbIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpAlbIntegration",
		[]interface{}{id, listener, props},
		&j,
	)

	return &j
}

// Deprecated.
func NewHttpAlbIntegration_Override(h HttpAlbIntegration, id *string, listener awselasticloadbalancingv2.IApplicationListener, props *HttpAlbIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.HttpAlbIntegration",
		[]interface{}{id, listener, props},
		h,
	)
}

func (j *jsiiProxy_HttpAlbIntegration)SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_HttpAlbIntegration)SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_HttpAlbIntegration)SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType) {
	if err := j.validateSetIntegrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_HttpAlbIntegration)SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion) {
	if err := j.validateSetPayloadFormatVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (h *jsiiProxy_HttpAlbIntegration) Bind(options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
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

func (h *jsiiProxy_HttpAlbIntegration) CompleteBind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

