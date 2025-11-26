package previewawsapigatewayv2mixins


// Properties for CfnApiPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var body interface{}
//
//   cfnApiMixinProps := &CfnApiMixinProps{
//   	ApiKeySelectionExpression: jsii.String("apiKeySelectionExpression"),
//   	BasePath: jsii.String("basePath"),
//   	Body: body,
//   	BodyS3Location: &BodyS3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Etag: jsii.String("etag"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   	CorsConfiguration: &CorsProperty{
//   		AllowCredentials: jsii.Boolean(false),
//   		AllowHeaders: []*string{
//   			jsii.String("allowHeaders"),
//   		},
//   		AllowMethods: []*string{
//   			jsii.String("allowMethods"),
//   		},
//   		AllowOrigins: []*string{
//   			jsii.String("allowOrigins"),
//   		},
//   		ExposeHeaders: []*string{
//   			jsii.String("exposeHeaders"),
//   		},
//   		MaxAge: jsii.Number(123),
//   	},
//   	CredentialsArn: jsii.String("credentialsArn"),
//   	Description: jsii.String("description"),
//   	DisableExecuteApiEndpoint: jsii.Boolean(false),
//   	DisableSchemaValidation: jsii.Boolean(false),
//   	FailOnWarnings: jsii.Boolean(false),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Name: jsii.String("name"),
//   	ProtocolType: jsii.String("protocolType"),
//   	RouteKey: jsii.String("routeKey"),
//   	RouteSelectionExpression: jsii.String("routeSelectionExpression"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Target: jsii.String("target"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html
//
type CfnApiMixinProps struct {
	// An API key selection expression.
	//
	// Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-apikeyselectionexpression
	//
	ApiKeySelectionExpression *string `field:"optional" json:"apiKeySelectionExpression" yaml:"apiKeySelectionExpression"`
	// Specifies how to interpret the base path of the API during import.
	//
	// Valid values are `ignore` , `prepend` , and `split` . The default value is `ignore` . To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html) . Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-basepath
	//
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The OpenAPI definition.
	//
	// Supported only for HTTP APIs. To import an HTTP API, you must specify a `Body` or `BodyS3Location` . If you specify a `Body` or `BodyS3Location` , don't specify CloudFormation resources such as `AWS::ApiGatewayV2::Authorizer` or `AWS::ApiGatewayV2::Route` . API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-body
	//
	Body interface{} `field:"optional" json:"body" yaml:"body"`
	// The S3 location of an OpenAPI definition.
	//
	// Supported only for HTTP APIs. To import an HTTP API, you must specify a `Body` or `BodyS3Location` . If you specify a `Body` or `BodyS3Location` , don't specify CloudFormation resources such as `AWS::ApiGatewayV2::Authorizer` or `AWS::ApiGatewayV2::Route` . API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-bodys3location
	//
	BodyS3Location interface{} `field:"optional" json:"bodyS3Location" yaml:"bodyS3Location"`
	// A CORS configuration.
	//
	// Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-corsconfiguration
	//
	CorsConfiguration interface{} `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// This property is part of quick create.
	//
	// It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify `arn:aws:iam::*:user/*` . To use resource-based permissions on supported AWS services, specify `null` . Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-credentialsarn
	//
	CredentialsArn *string `field:"optional" json:"credentialsArn" yaml:"credentialsArn"`
	// The description of the API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke your API by using the default `execute-api` endpoint.
	//
	// By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-disableexecuteapiendpoint
	//
	DisableExecuteApiEndpoint interface{} `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Avoid validating models when creating a deployment.
	//
	// Supported only for WebSocket APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-disableschemavalidation
	//
	DisableSchemaValidation interface{} `field:"optional" json:"disableSchemaValidation" yaml:"disableSchemaValidation"`
	// Specifies whether to rollback the API creation when a warning is encountered.
	//
	// By default, API creation continues if a warning is encountered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-failonwarnings
	//
	FailOnWarnings interface{} `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// The IP address types that can invoke the API.
	//
	// Use `ipv4` to allow only IPv4 addresses to invoke your API, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke your API.
	//
	// Don’t use IP address type for an HTTP API based on an OpenAPI specification. Instead, specify the IP address type in the OpenAPI specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The name of the API.
	//
	// Required unless you specify an OpenAPI definition for `Body` or `S3BodyLocation` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The API protocol.
	//
	// Valid values are `WEBSOCKET` or `HTTP` . Required unless you specify an OpenAPI definition for `Body` or `S3BodyLocation` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-protocoltype
	//
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// This property is part of quick create.
	//
	// If you don't specify a `routeKey` , a default route of `$default` is created. The `$default` route acts as a catch-all for any request made to your API, for a particular stage. The `$default` route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-routekey
	//
	RouteKey *string `field:"optional" json:"routeKey" yaml:"routeKey"`
	// The route selection expression for the API.
	//
	// For HTTP APIs, the `routeSelectionExpression` must be `${request.method} ${request.path}` . If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-routeselectionexpression
	//
	RouteSelectionExpression *string `field:"optional" json:"routeSelectionExpression" yaml:"routeSelectionExpression"`
	// The collection of tags.
	//
	// Each tag element is associated with a given resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// This property is part of quick create.
	//
	// Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
	// A version identifier for the API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html#cfn-apigatewayv2-api-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

