package awsapigatewayv2


// Supported HTTP methods.
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
type HttpMethod string

const (
	// HTTP ANY.
	HttpMethod_ANY HttpMethod = "ANY"
	// HTTP DELETE.
	HttpMethod_DELETE HttpMethod = "DELETE"
	// HTTP GET.
	HttpMethod_GET HttpMethod = "GET"
	// HTTP HEAD.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// HTTP OPTIONS.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	// HTTP PATCH.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// HTTP POST.
	HttpMethod_POST HttpMethod = "POST"
	// HTTP PUT.
	HttpMethod_PUT HttpMethod = "PUT"
)

