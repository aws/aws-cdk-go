// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The interface that various route integration classes will inherit.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
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
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener, &httpAlbIntegrationProps{
//   		parameterMapping: apigwv2.NewParameterMapping().custom(jsii.String("myKey"), jsii.String("myValue")),
//   	}),
//   })
//
// Experimental.
type HttpRouteIntegration interface {
	// Bind this integration to the route.
	// Experimental.
	Bind(options *HttpRouteIntegrationBindOptions) *HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Experimental.
	CompleteBind(_options *HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpRouteIntegration
type jsiiProxy_HttpRouteIntegration struct {
	_ byte // padding
}

// Initialize an integration for a route on http api.
// Experimental.
func NewHttpRouteIntegration_Override(h HttpRouteIntegration, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpRouteIntegration",
		[]interface{}{id},
		h,
	)
}

func (h *jsiiProxy_HttpRouteIntegration) Bind(options *HttpRouteIntegrationBindOptions) *HttpRouteIntegrationConfig {
	if err := h.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpRouteIntegration) CompleteBind(_options *HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

