package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations/internal"
)

// AWS WebSocket AWS Type Integration.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiRole role
//   var table table
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AutoDeploy: jsii.Boolean(true),
//   })
//   webSocketApi.AddRoute(jsii.String("$connect"), &WebSocketRouteOptions{
//   	Integration: awscdk.NewWebSocketAwsIntegration(jsii.String("DynamodbPutItem"), &WebSocketAwsIntegrationProps{
//   		IntegrationUri: fmt.Sprintf("arn:aws:apigateway:%v:dynamodb:action/PutItem", this.Region),
//   		IntegrationMethod: apigwv2.HttpMethod_POST,
//   		CredentialsRole: apiRole,
//   		RequestTemplates: map[string]*string{
//   			"application/json": JSON.stringify(map[string]interface{}{
//   				"TableName": table.tableName,
//   				"Item": map[string]map[string]*string{
//   					"id": map[string]*string{
//   						"S": jsii.String("$context.requestId"),
//   					},
//   				},
//   			}),
//   		},
//   	}),
//   })
//
type WebSocketAwsIntegration interface {
	awsapigatewayv2.WebSocketRouteIntegration
	// Bind this integration to the route.
	Bind(_options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketAwsIntegration
type jsiiProxy_WebSocketAwsIntegration struct {
	internal.Type__awsapigatewayv2WebSocketRouteIntegration
}

func NewWebSocketAwsIntegration(id *string, props *WebSocketAwsIntegrationProps) WebSocketAwsIntegration {
	_init_.Initialize()

	if err := validateNewWebSocketAwsIntegrationParameters(id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketAwsIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.WebSocketAwsIntegration",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewWebSocketAwsIntegration_Override(w WebSocketAwsIntegration, id *string, props *WebSocketAwsIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.WebSocketAwsIntegration",
		[]interface{}{id, props},
		w,
	)
}

func (w *jsiiProxy_WebSocketAwsIntegration) Bind(_options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig {
	if err := w.validateBindParameters(_options); err != nil {
		panic(err)
	}
	var returns *awsapigatewayv2.WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

