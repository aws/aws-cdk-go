package interfacesawsevents


// A reference to a EventBus resource.
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
type EventBusReference struct {
	// The ARN of the EventBus resource.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
	// The Name of the EventBus resource.
	EventBusName *string `field:"required" json:"eventBusName" yaml:"eventBusName"`
}

