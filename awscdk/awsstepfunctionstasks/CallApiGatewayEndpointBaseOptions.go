package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Base CallApiGatewayEdnpoint Task Props.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var taskInput taskInput
//
//   callApiGatewayEndpointBaseOptions := &CallApiGatewayEndpointBaseOptions{
//   	Method: awscdk.Aws_stepfunctions_tasks.HttpMethod_GET,
//
//   	// the properties below are optional
//   	ApiPath: jsii.String("apiPath"),
//   	AuthType: awscdk.*Aws_stepfunctions_tasks.AuthType_NO_AUTH,
//   	Headers: taskInput,
//   	QueryParameters: taskInput,
//   	RequestBody: taskInput,
//   }
//
type CallApiGatewayEndpointBaseOptions struct {
	// Http method for the API.
	Method HttpMethod `field:"required" json:"method" yaml:"method"`
	// Path parameters appended after API endpoint.
	// Default: - No path.
	//
	ApiPath *string `field:"optional" json:"apiPath" yaml:"apiPath"`
	// Authentication methods.
	// Default: AuthType.NO_AUTH
	//
	AuthType AuthType `field:"optional" json:"authType" yaml:"authType"`
	// HTTP request information that does not relate to contents of the request.
	// Default: - No headers.
	//
	Headers awsstepfunctions.TaskInput `field:"optional" json:"headers" yaml:"headers"`
	// Query strings attatched to end of request.
	// Default: - No query parameters.
	//
	QueryParameters awsstepfunctions.TaskInput `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// HTTP Request body.
	// Default: - No request body.
	//
	RequestBody awsstepfunctions.TaskInput `field:"optional" json:"requestBody" yaml:"requestBody"`
}

