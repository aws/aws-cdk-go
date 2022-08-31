package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2integrations/internal"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

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

