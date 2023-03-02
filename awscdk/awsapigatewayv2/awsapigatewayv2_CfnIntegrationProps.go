package awsapigatewayv2


// Properties for defining a `CfnIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var requestParameters interface{}
//   var requestTemplates interface{}
//   var responseParameters interface{}
//
//   cfnIntegrationProps := &cfnIntegrationProps{
//   	apiId: jsii.String("apiId"),
//   	integrationType: jsii.String("integrationType"),
//
//   	// the properties below are optional
//   	connectionId: jsii.String("connectionId"),
//   	connectionType: jsii.String("connectionType"),
//   	contentHandlingStrategy: jsii.String("contentHandlingStrategy"),
//   	credentialsArn: jsii.String("credentialsArn"),
//   	description: jsii.String("description"),
//   	integrationMethod: jsii.String("integrationMethod"),
//   	integrationSubtype: jsii.String("integrationSubtype"),
//   	integrationUri: jsii.String("integrationUri"),
//   	passthroughBehavior: jsii.String("passthroughBehavior"),
//   	payloadFormatVersion: jsii.String("payloadFormatVersion"),
//   	requestParameters: requestParameters,
//   	requestTemplates: requestTemplates,
//   	responseParameters: responseParameters,
//   	templateSelectionExpression: jsii.String("templateSelectionExpression"),
//   	timeoutInMillis: jsii.Number(123),
//   	tlsConfig: &tlsConfigProperty{
//   		serverNameToVerify: jsii.String("serverNameToVerify"),
//   	},
//   }
//
type CfnIntegrationProps struct {
	// The API identifier.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The integration type of an integration. One of the following:.
	//
	// `AWS` : for integrating the route or method request with an AWS service action, including the Lambda function-invoking action. With the Lambda function-invoking action, this is referred to as the Lambda custom integration. With any other AWS service action, this is known as AWS integration. Supported only for WebSocket APIs.
	//
	// `AWS_PROXY` : for integrating the route or method request with a Lambda function or other AWS service action. This integration is also referred to as a Lambda proxy integration.
	//
	// `HTTP` : for integrating the route or method request with an HTTP endpoint. This integration is also referred to as the HTTP custom integration. Supported only for WebSocket APIs.
	//
	// `HTTP_PROXY` : for integrating the route or method request with an HTTP endpoint, with the client request passed through as-is. This is also referred to as HTTP proxy integration. For HTTP API private integrations, use an `HTTP_PROXY` integration.
	//
	// `MOCK` : for integrating the route or method request with API Gateway as a "loopback" endpoint without invoking any backend. Supported only for WebSocket APIs.
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// The ID of the VPC link for a private integration.
	//
	// Supported only for HTTP APIs.
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// The type of the network connection to the integration endpoint.
	//
	// Specify `INTERNET` for connections through the public routable internet or `VPC_LINK` for private connections between API Gateway and resources in a VPC. The default value is `INTERNET` .
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Supported only for WebSocket APIs.
	//
	// Specifies how to handle response payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT` , with the following behaviors:
	//
	// `CONVERT_TO_BINARY` : Converts a response payload from a Base64-encoded string to the corresponding binary blob.
	//
	// `CONVERT_TO_TEXT` : Converts a response payload from a binary blob to a Base64-encoded string.
	//
	// If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.
	ContentHandlingStrategy *string `field:"optional" json:"contentHandlingStrategy" yaml:"contentHandlingStrategy"`
	// Specifies the credentials required for the integration, if any.
	//
	// For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::*:user/*` . To use resource-based permissions on supported AWS services, don't specify this parameter.
	CredentialsArn *string `field:"optional" json:"credentialsArn" yaml:"credentialsArn"`
	// The description of the integration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the integration's HTTP method type.
	IntegrationMethod *string `field:"optional" json:"integrationMethod" yaml:"integrationMethod"`
	// Supported only for HTTP API `AWS_PROXY` integrations.
	//
	// Specifies the AWS service action to invoke. To learn more, see [Integration subtype reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html) .
	IntegrationSubtype *string `field:"optional" json:"integrationSubtype" yaml:"integrationSubtype"`
	// For a Lambda integration, specify the URI of a Lambda function.
	//
	// For an HTTP integration, specify a fully-qualified URL.
	//
	// For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service. If you specify the ARN of an AWS Cloud Map service, API Gateway uses `DiscoverInstances` to identify resources. You can use query parameters to target specific resources. To learn more, see [DiscoverInstances](https://docs.aws.amazon.com/cloud-map/latest/api/API_DiscoverInstances.html) . For private integrations, all resources must be owned by the same AWS account .
	IntegrationUri *string `field:"optional" json:"integrationUri" yaml:"integrationUri"`
	// Specifies the pass-through behavior for incoming requests based on the `Content-Type` header in the request, and the available mapping templates specified as the `requestTemplates` property on the `Integration` resource.
	//
	// There are three valid values: `WHEN_NO_MATCH` , `WHEN_NO_TEMPLATES` , and `NEVER` . Supported only for WebSocket APIs.
	//
	// `WHEN_NO_MATCH` passes the request body for unmapped content types through to the integration backend without transformation.
	//
	// `NEVER` rejects unmapped content types with an `HTTP 415 Unsupported Media Type` response.
	//
	// `WHEN_NO_TEMPLATES` allows pass-through when the integration has no content types mapped to templates. However, if there is at least one content type defined, unmapped content types will be rejected with the same `HTTP 415 Unsupported Media Type` response.
	PassthroughBehavior *string `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// Specifies the format of the payload sent to an integration.
	//
	// Required for HTTP APIs. For HTTP APIs, supported values for Lambda proxy integrations are `1.0` and `2.0` . For all other integrations, `1.0` is the only supported value. To learn more, see [Working with AWS Lambda proxy integrations for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html) .
	PayloadFormatVersion *string `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	//
	// The key is an integration request parameter name and the associated value is a method request parameter value or static value that must be enclosed within single quotes and pre-encoded as required by the backend. The method request parameter value must match the pattern of `method.request. {location} . {name}` , where `{location}` is `querystring` , `path` , or `header` ; and `{name}` must be a valid and unique method request parameter name.
	//
	// For HTTP API integrations with a specified `integrationSubtype` , request parameters are a key-value map specifying parameters that are passed to `AWS_PROXY` integrations. You can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Working with AWS service integrations for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services.html) .
	//
	// For HTTP API integrations without a specified `integrationSubtype` request parameters are a key-value map specifying how to transform HTTP requests before sending them to the backend. The key should follow the pattern <action>:<header|querystring|path>.<location> where action can be `append` , `overwrite` or `remove` . For values, you can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Transforming API requests and responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) .
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	//
	// The content type value is the key in this map, and the template (as a String) is the value. Supported only for WebSocket APIs.
	RequestTemplates interface{} `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// Supported only for HTTP APIs.
	//
	// You use response parameters to transform the HTTP response from a backend integration before returning the response to clients. Specify a key-value map from a selection key to response parameters. The selection key must be a valid HTTP status code within the range of 200-599. The value is of type [`ResponseParameterList`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-responseparameterlist.html) . To learn more, see [Transforming API requests and responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html) .
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// The template selection expression for the integration.
	//
	// Supported only for WebSocket APIs.
	TemplateSelectionExpression *string `field:"optional" json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
	// Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs.
	//
	// The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.
	TimeoutInMillis *float64 `field:"optional" json:"timeoutInMillis" yaml:"timeoutInMillis"`
	// The TLS configuration for a private integration.
	//
	// If you specify a TLS configuration, private integration traffic uses the HTTPS protocol. Supported only for HTTP APIs.
	TlsConfig interface{} `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
}

