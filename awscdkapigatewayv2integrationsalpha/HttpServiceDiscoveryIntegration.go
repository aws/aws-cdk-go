package awscdkapigatewayv2integrationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/internal"
)

// The Service Discovery integration resource for HTTP API.
//
// Example:
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   vpcLink := apigwv2.NewVpcLink(this, jsii.String("VpcLink"), &VpcLinkProps{
//   	Vpc: Vpc,
//   })
//   namespace := servicediscovery.NewPrivateDnsNamespace(this, jsii.String("Namespace"), &PrivateDnsNamespaceProps{
//   	Name: jsii.String("boobar.com"),
//   	Vpc: Vpc,
//   })
//   service := namespace.CreateService(jsii.String("Service"))
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
//   	DefaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpServiceDiscoveryIntegration(jsii.String("DefaultIntegration"), service, &HttpServiceDiscoveryIntegrationProps{
//   		VpcLink: *VpcLink,
//   	}),
//   })
//
// Experimental.
type HttpServiceDiscoveryIntegration interface {
	awscdkapigatewayv2alpha.HttpRouteIntegration
	// Experimental.
	ConnectionType() awscdkapigatewayv2alpha.HttpConnectionType
	// Experimental.
	SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType)
	// Experimental.
	HttpMethod() awscdkapigatewayv2alpha.HttpMethod
	// Experimental.
	SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod)
	// Experimental.
	IntegrationType() awscdkapigatewayv2alpha.HttpIntegrationType
	// Experimental.
	SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType)
	// Experimental.
	PayloadFormatVersion() awscdkapigatewayv2alpha.PayloadFormatVersion
	// Experimental.
	SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion)
	// Bind this integration to the route.
	// Experimental.
	Bind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Experimental.
	CompleteBind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions)
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

	if err := validateNewHttpServiceDiscoveryIntegrationParameters(id, service, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_HttpServiceDiscoveryIntegration)SetConnectionType(val awscdkapigatewayv2alpha.HttpConnectionType) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration)SetHttpMethod(val awscdkapigatewayv2alpha.HttpMethod) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration)SetIntegrationType(val awscdkapigatewayv2alpha.HttpIntegrationType) {
	if err := j.validateSetIntegrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration)SetPayloadFormatVersion(val awscdkapigatewayv2alpha.PayloadFormatVersion) {
	if err := j.validateSetPayloadFormatVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) Bind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig {
	if err := h.validateBindParameters(_options); err != nil {
		panic(err)
	}
	var returns *awscdkapigatewayv2alpha.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) CompleteBind(_options *awscdkapigatewayv2alpha.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

