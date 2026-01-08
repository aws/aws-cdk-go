package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicediscovery"
)

// The Service Discovery integration resource for HTTP API.
//
// Example:
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	DefaultIntegration: awscdk.NewHttpServiceDiscoveryIntegration(jsii.String("DefaultIntegration"), service, &HttpServiceDiscoveryIntegrationProps{
//   		VpcLink: *VpcLink,
//   	}),
//   })
//
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
	// Bind this integration to the route.
	Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
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


func NewHttpServiceDiscoveryIntegration(id *string, service interfacesawsservicediscovery.IServiceRef, props *HttpServiceDiscoveryIntegrationProps) HttpServiceDiscoveryIntegration {
	_init_.Initialize()

	if err := validateNewHttpServiceDiscoveryIntegrationParameters(id, service, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpServiceDiscoveryIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpServiceDiscoveryIntegration",
		[]interface{}{id, service, props},
		&j,
	)

	return &j
}

func NewHttpServiceDiscoveryIntegration_Override(h HttpServiceDiscoveryIntegration, id *string, service interfacesawsservicediscovery.IServiceRef, props *HttpServiceDiscoveryIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpServiceDiscoveryIntegration",
		[]interface{}{id, service, props},
		h,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration)SetConnectionType(val awsapigatewayv2.HttpConnectionType) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration)SetHttpMethod(val awsapigatewayv2.HttpMethod) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration)SetIntegrationType(val awsapigatewayv2.HttpIntegrationType) {
	if err := j.validateSetIntegrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_HttpServiceDiscoveryIntegration)SetPayloadFormatVersion(val awsapigatewayv2.PayloadFormatVersion) {
	if err := j.validateSetPayloadFormatVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	if err := h.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpServiceDiscoveryIntegration) CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

