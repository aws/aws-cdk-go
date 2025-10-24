package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations/internal"
)

// The Sqs integration resource for HTTP API.
//
// Example:
//   import sqs "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue IQueue
//   var httpApi HttpApi
//
//
//   // default integration (send message)
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/default"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: awscdk.NewHttpSqsIntegration(jsii.String("defaultIntegration"), &HttpSqsIntegrationProps{
//   		Queue: *Queue,
//   	}),
//   })
//   // send message integration
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/send-message"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: awscdk.NewHttpSqsIntegration(jsii.String("sendMessageIntegration"), &HttpSqsIntegrationProps{
//   		Queue: *Queue,
//   		Subtype: apigwv2.HttpIntegrationSubtype_SQS_SEND_MESSAGE,
//   	}),
//   })
//   // receive message integration
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/receive-message"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: awscdk.NewHttpSqsIntegration(jsii.String("receiveMessageIntegration"), &HttpSqsIntegrationProps{
//   		Queue: *Queue,
//   		Subtype: apigwv2.HttpIntegrationSubtype_SQS_RECEIVE_MESSAGE,
//   	}),
//   })
//   // delete message integration
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/delete-message"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: awscdk.NewHttpSqsIntegration(jsii.String("deleteMessageIntegration"), &HttpSqsIntegrationProps{
//   		Queue: *Queue,
//   		Subtype: apigwv2.HttpIntegrationSubtype_SQS_DELETE_MESSAGE,
//   	}),
//   })
//   // purge queue integration
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/purge-queue"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: awscdk.NewHttpSqsIntegration(jsii.String("purgeQueueIntegration"), &HttpSqsIntegrationProps{
//   		Queue: *Queue,
//   		Subtype: apigwv2.HttpIntegrationSubtype_SQS_PURGE_QUEUE,
//   	}),
//   })
//
type HttpSqsIntegration interface {
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

// The jsii proxy struct for HttpSqsIntegration
type jsiiProxy_HttpSqsIntegration struct {
	internal.Type__awsapigatewayv2HttpRouteIntegration
}

func NewHttpSqsIntegration(id *string, props *HttpSqsIntegrationProps) HttpSqsIntegration {
	_init_.Initialize()

	if err := validateNewHttpSqsIntegrationParameters(id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpSqsIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpSqsIntegration",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewHttpSqsIntegration_Override(h HttpSqsIntegration, id *string, props *HttpSqsIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.HttpSqsIntegration",
		[]interface{}{id, props},
		h,
	)
}

func (h *jsiiProxy_HttpSqsIntegration) Bind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
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

func (h *jsiiProxy_HttpSqsIntegration) CompleteBind(_options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(_options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{_options},
	)
}

