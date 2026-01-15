package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations/internal"
)

// The EventBridge PutEvents integration resource for HTTP API.
//
// Example:
//   import events "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bus IEventBus
//   var httpApi HttpApi
//
//
//   // default integration (PutEvents)
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/default"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: awscdk.NewHttpEventBridgeIntegration(jsii.String("DefaultEventBridgeIntegration"), &HttpEventBridgeIntegrationProps{
//   		EventBusRef: bus.EventBusRef,
//   	}),
//   })
//
//   // explicit subtype
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/put-events"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: awscdk.NewHttpEventBridgeIntegration(jsii.String("ExplicitSubtypeIntegration"), &HttpEventBridgeIntegrationProps{
//   		EventBusRef: bus.*EventBusRef,
//   		Subtype: apigwv2.HttpIntegrationSubtype_EVENTBRIDGE_PUT_EVENTS,
//   	}),
//   })
//
type HttpEventBridgeIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
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

// The jsii proxy struct for HttpEventBridgeIntegration
type jsiiProxy_HttpEventBridgeIntegration struct {
	internal.Type__awsapigatewayv2HttpRouteIntegration
}

func NewHttpEventBridgeIntegration(id *string, props *HttpEventBridgeIntegrationProps) HttpEventBridgeIntegration {
	_init_.Initialize()

	if err := validateNewHttpEventBridgeIntegrationParameters(id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpEventBridgeIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpEventBridgeIntegration",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewHttpEventBridgeIntegration_Override(h HttpEventBridgeIntegration, id *string, props *HttpEventBridgeIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpEventBridgeIntegration",
		[]interface{}{id, props},
		h,
	)
}

func (h *jsiiProxy_HttpEventBridgeIntegration) Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
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

func (h *jsiiProxy_HttpEventBridgeIntegration) CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

