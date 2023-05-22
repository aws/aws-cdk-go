package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for calling an REST API Endpoint.
//
// Example:
//   import apigateway "github.com/aws/aws-cdk-go/awscdk"
//   var api restApi
//
//
//   tasks.NewCallApiGatewayRestApiEndpoint(this, jsii.String("Endpoint"), &CallApiGatewayRestApiEndpointProps{
//   	Api: Api,
//   	StageName: jsii.String("Stage"),
//   	Method: tasks.HttpMethod_PUT,
//   	IntegrationPattern: sfn.IntegrationPattern_WAIT_FOR_TASK_TOKEN,
//   	Headers: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"TaskToken": sfn.JsonPath_array(sfn.JsonPath_taskToken()),
//   	}),
//   })
//
// Experimental.
type CallApiGatewayRestApiEndpointProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Http method for the API.
	// Experimental.
	Method HttpMethod `field:"required" json:"method" yaml:"method"`
	// Path parameters appended after API endpoint.
	// Experimental.
	ApiPath *string `field:"optional" json:"apiPath" yaml:"apiPath"`
	// Authentication methods.
	// Experimental.
	AuthType AuthType `field:"optional" json:"authType" yaml:"authType"`
	// HTTP request information that does not relate to contents of the request.
	// Experimental.
	Headers awsstepfunctions.TaskInput `field:"optional" json:"headers" yaml:"headers"`
	// Query strings attatched to end of request.
	// Experimental.
	QueryParameters awsstepfunctions.TaskInput `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// HTTP Request body.
	// Experimental.
	RequestBody awsstepfunctions.TaskInput `field:"optional" json:"requestBody" yaml:"requestBody"`
	// API to call.
	// Experimental.
	Api awsapigateway.IRestApi `field:"required" json:"api" yaml:"api"`
	// Name of the stage where the API is deployed to in API Gateway.
	// Experimental.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}

