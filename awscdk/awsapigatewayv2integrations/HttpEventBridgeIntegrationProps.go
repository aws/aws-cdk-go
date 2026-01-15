package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

// Properties to initialize `HttpEventBridgeIntegration`.
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
type HttpEventBridgeIntegrationProps struct {
	// EventBridge event bus that integrates with API Gateway.
	EventBusRef *interfacesawsevents.EventBusReference `field:"required" json:"eventBusRef" yaml:"eventBusRef"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	//
	// When not provided, a default mapping will be used that expects the
	// incoming request body to contain the fields `Detail`, `DetailType`, and
	// `Source`.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html
	//
	// Default: - set `Detail` to `$request.body.Detail`,
	// `DetailType` to `$request.body.DetailType`, and `Source` to `$request.body.Source`.
	//
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// The subtype of the HTTP integration.
	//
	// Only subtypes starting with EVENTBRIDGE_ can be specified.
	// Default: HttpIntegrationSubtype.EVENTBRIDGE_PUT_EVENTS
	//
	Subtype awsapigatewayv2.HttpIntegrationSubtype `field:"optional" json:"subtype" yaml:"subtype"`
}

