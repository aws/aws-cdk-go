package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties to initialize `HttpSqsIntegration`.
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
type HttpSqsIntegrationProps struct {
	// SQS queue that Integrates with API Gateway.
	Queue awssqs.IQueue `field:"required" json:"queue" yaml:"queue"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html
	//
	// Default: - specify `QueueUrl`. Additionally, set `MessageBody` to `$request.body.MessageBody` for `SQS_SEND_MESSAGE` subtype
	// and set `ReceiptHandle` to `$request.body.ReceiptHandle` for `SQS_DELETE_MESSAGE` subtype.
	//
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// The subtype of the HTTP integration.
	//
	// Only subtypes starting with SQS_ can be specified.
	// Default: HttpIntegrationSubtype.SQS_SEND_MESSAGE
	//
	Subtype awsapigatewayv2.HttpIntegrationSubtype `field:"optional" json:"subtype" yaml:"subtype"`
}

