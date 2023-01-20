package awsapigateway


// `Integration` is a property of the [AWS::ApiGateway::Method](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html) resource that specifies information about the target backend that a method calls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationProperty := &integrationProperty{
//   	cacheKeyParameters: []*string{
//   		jsii.String("cacheKeyParameters"),
//   	},
//   	cacheNamespace: jsii.String("cacheNamespace"),
//   	connectionId: jsii.String("connectionId"),
//   	connectionType: jsii.String("connectionType"),
//   	contentHandling: jsii.String("contentHandling"),
//   	credentials: jsii.String("credentials"),
//   	integrationHttpMethod: jsii.String("integrationHttpMethod"),
//   	integrationResponses: []interface{}{
//   		&integrationResponseProperty{
//   			statusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			contentHandling: jsii.String("contentHandling"),
//   			responseParameters: map[string]*string{
//   				"responseParametersKey": jsii.String("responseParameters"),
//   			},
//   			responseTemplates: map[string]*string{
//   				"responseTemplatesKey": jsii.String("responseTemplates"),
//   			},
//   			selectionPattern: jsii.String("selectionPattern"),
//   		},
//   	},
//   	passthroughBehavior: jsii.String("passthroughBehavior"),
//   	requestParameters: map[string]*string{
//   		"requestParametersKey": jsii.String("requestParameters"),
//   	},
//   	requestTemplates: map[string]*string{
//   		"requestTemplatesKey": jsii.String("requestTemplates"),
//   	},
//   	timeoutInMillis: jsii.Number(123),
//   	type: jsii.String("type"),
//   	uri: jsii.String("uri"),
//   }
//
type CfnMethod_IntegrationProperty struct {
	// A list of request parameters whose values API Gateway caches.
	//
	// For cases where the integration type allows for RequestParameters to be set, these parameters must also be specified in [RequestParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestparameters) to be supported in `CacheKeyParameters` .
	CacheKeyParameters *[]*string `field:"optional" json:"cacheKeyParameters" yaml:"cacheKeyParameters"`
	// An API-specific tag group of related cached parameters.
	CacheNamespace *string `field:"optional" json:"cacheNamespace" yaml:"cacheNamespace"`
	// The ID of the `VpcLink` used for the integration when `connectionType=VPC_LINK` , otherwise undefined.
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// The type of the network connection to the integration endpoint.
	//
	// The valid value is `INTERNET` for connections through the public routable internet or `VPC_LINK` for private connections between API Gateway and a network load balancer in a VPC. The default value is `INTERNET` .
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Specifies how to handle request payload content type conversions. Valid values are:.
	//
	// - `CONVERT_TO_BINARY` : Converts a request payload from a base64-encoded string to a binary blob.
	// - `CONVERT_TO_TEXT` : Converts a request payload from a binary blob to a base64-encoded string.
	//
	// If this property isn't defined, the request payload is passed through from the method request to the integration request without modification, provided that the `PassthroughBehaviors` property is configured to support payload pass-through.
	ContentHandling *string `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// The credentials that are required for the integration.
	//
	// To specify an AWS Identity and Access Management (IAM) role that API Gateway assumes, specify the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify arn:aws:iam::*:user/*.
	//
	// To use resource-based permissions on the AWS Lambda (Lambda) function, don't specify this property. Use the [AWS::Lambda::Permission](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html) resource to permit API Gateway to call the function. For more information, see [Allow Amazon API Gateway to Invoke a Lambda Function](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html#access-control-resource-based-example-apigateway-invoke-function) in the *AWS Lambda Developer Guide* .
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// The integration's HTTP method type.
	//
	// For the `Type` property, if you specify `MOCK` , this property is optional. For all other types, you must specify this property.
	IntegrationHttpMethod *string `field:"optional" json:"integrationHttpMethod" yaml:"integrationHttpMethod"`
	// The response that API Gateway provides after a method's backend completes processing a request.
	//
	// API Gateway intercepts the response from the backend so that you can control how API Gateway surfaces backend responses. For example, you can map the backend status codes to codes that you define.
	IntegrationResponses interface{} `field:"optional" json:"integrationResponses" yaml:"integrationResponses"`
	// Indicates when API Gateway passes requests to the targeted backend.
	//
	// This behavior depends on the request's `Content-Type` header and whether you defined a mapping template for it.
	//
	// For more information and valid values, see the [passthroughBehavior](https://docs.aws.amazon.com/apigateway/api-reference/link-relation/integration-put/#passthroughBehavior) field in the *API Gateway API Reference* .
	PassthroughBehavior *string `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// The request parameters that API Gateway sends with the backend request.
	//
	// Specify request parameters as key-value pairs (string-to-string mappings), with a destination as the key and a source as the value.
	//
	// Specify the destination by using the following pattern `integration.request. *location* . *name*` , where *location* is query string, path, or header, and *name* is a valid, unique parameter name.
	//
	// The source must be an existing method request parameter or a static value. You must enclose static values in single quotation marks and pre-encode these values based on their destination in the request.
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// A map of Apache Velocity templates that are applied on the request payload.
	//
	// The template that API Gateway uses is based on the value of the `Content-Type` header that's sent by the client. The content type value is the key, and the template is the value (specified as a string), such as the following snippet:
	//
	// `"application/json": "{\n \"statusCode\": 200\n}"`
	//
	// For more information about templates, see [API Gateway Mapping Template and Access Logging Variable Reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html) in the *API Gateway Developer Guide* .
	RequestTemplates interface{} `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// Custom timeout between 50 and 29,000 milliseconds.
	//
	// The default value is 29,000 milliseconds or 29 seconds.
	TimeoutInMillis *float64 `field:"optional" json:"timeoutInMillis" yaml:"timeoutInMillis"`
	// The type of backend that your method is running, such as `HTTP` or `MOCK` .
	//
	// For all of the valid values, see the [type](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/#type) property for the `Integration` resource in the *Amazon API Gateway REST API Reference* .
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The Uniform Resource Identifier (URI) for the integration.
	//
	// If you specify `HTTP` for the `Type` property, specify the API endpoint URL.
	//
	// If you specify `MOCK` for the `Type` property, don't specify this property.
	//
	// If you specify `AWS` for the `Type` property, specify an AWS service that follows this form: arn:aws:apigateway: *region* : *subdomain* . *service|service* : *path|action* / *service_api* . For example, a Lambda function URI follows this form: arn:aws:apigateway: *region* :lambda:path/ *path* . The path is usually in the form /2015-03-31/functions/ *LambdaFunctionARN* /invocations. For more information, see the `uri` property of the [Integration](https://docs.aws.amazon.com/apigateway/api-reference/resource/integration/) resource in the Amazon API Gateway REST API Reference.
	//
	// If you specified `HTTP` or `AWS` for the `Type` property, you must specify this property.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

