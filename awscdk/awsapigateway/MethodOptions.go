package awsapigateway


// Example:
//   var api restApi
//   var userLambda function
//
//
//   userModel := api.AddModel(jsii.String("UserModel"), &ModelOptions{
//   	Schema: &JsonSchema{
//   		Type: apigateway.JsonSchemaType_OBJECT,
//   		Properties: map[string]jsonSchema{
//   			"userId": &jsonSchema{
//   				"type": apigateway.JsonSchemaType_STRING,
//   			},
//   			"name": &jsonSchema{
//   				"type": apigateway.JsonSchemaType_STRING,
//   			},
//   		},
//   		Required: []*string{
//   			jsii.String("userId"),
//   		},
//   	},
//   })
//   api.Root.AddResource(jsii.String("user")).AddMethod(jsii.String("POST"),
//   apigateway.NewLambdaIntegration(userLambda), &MethodOptions{
//   	RequestModels: map[string]iModel{
//   		"application/json": userModel,
//   	},
//   })
//
type MethodOptions struct {
	// Indicates whether the method requires clients to submit a valid API key.
	// Default: false.
	//
	ApiKeyRequired *bool `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// A list of authorization scopes configured on the method.
	//
	// The scopes are used with
	// a COGNITO_USER_POOLS authorizer to authorize the method invocation.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationscopes
	//
	// Default: - no authorization scopes.
	//
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// Method authorization. If the value is set of `Custom`, an `authorizer` must also be specified.
	//
	// If you're using one of the authorizers that are available via the `Authorizer` class, such as `Authorizer#token()`,
	// it is recommended that this option not be specified. The authorizer will take care of setting the correct authorization type.
	// However, specifying an authorization type using this property that conflicts with what is expected by the `Authorizer`
	// will result in an error.
	// Default: - open access unless `authorizer` is specified.
	//
	AuthorizationType AuthorizationType `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// If `authorizationType` is `Custom`, this specifies the ID of the method authorizer resource.
	//
	// If specified, the value of `authorizationType` must be set to `Custom`.
	Authorizer IAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
	// The responses that can be sent to the client who calls the method.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-settings-method-response.html
	//
	// Default: None
	//
	// This property is not required, but if these are not supplied for a Lambda
	// proxy integration, the Lambda function must return a value of the correct format,
	// for the integration response to be correctly mapped to a response to the client.
	//
	MethodResponses *[]*MethodResponse `field:"optional" json:"methodResponses" yaml:"methodResponses"`
	// A friendly operation name for the method.
	//
	// For example, you can assign the
	// OperationName of ListPets for the GET /pets method.
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// The models which describe data structure of request payload.
	//
	// When
	// combined with `requestValidator` or `requestValidatorOptions`, the service
	// will validate the API request payload before it reaches the API's Integration (including proxies).
	// Specify `requestModels` as key-value pairs, with a content type
	// (e.g. `'application/json'`) as the key and an API Gateway Model as the value.
	//
	// Example:
	//   var api restApi
	//   var userLambda function
	//
	//
	//   userModel := api.AddModel(jsii.String("UserModel"), &ModelOptions{
	//   	Schema: &JsonSchema{
	//   		Type: apigateway.JsonSchemaType_OBJECT,
	//   		Properties: map[string]jsonSchema{
	//   			"userId": &jsonSchema{
	//   				"type": apigateway.JsonSchemaType_STRING,
	//   			},
	//   			"name": &jsonSchema{
	//   				"type": apigateway.JsonSchemaType_STRING,
	//   			},
	//   		},
	//   		Required: []*string{
	//   			jsii.String("userId"),
	//   		},
	//   	},
	//   })
	//   api.Root.AddResource(jsii.String("user")).AddMethod(jsii.String("POST"),
	//   apigateway.NewLambdaIntegration(userLambda), &MethodOptions{
	//   	RequestModels: map[string]iModel{
	//   		"application/json": userModel,
	//   	},
	//   })
	//
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-settings-method-request.html#setup-method-request-model
	//
	RequestModels *map[string]IModel `field:"optional" json:"requestModels" yaml:"requestModels"`
	// The request parameters that API Gateway accepts.
	//
	// Specify request parameters
	// as key-value pairs (string-to-Boolean mapping), with a source as the key and
	// a Boolean as the value. The Boolean specifies whether a parameter is required.
	// A source must match the format method.request.location.name, where the location
	// is querystring, path, or header, and name is a valid, unique parameter name.
	// Default: None.
	//
	RequestParameters *map[string]*bool `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// The ID of the associated request validator.
	//
	// Only one of `requestValidator` or `requestValidatorOptions` must be specified.
	// Works together with `requestModels` or `requestParameters` to validate
	// the request before it reaches integration like Lambda Proxy Integration.
	// Default: - No default validator.
	//
	RequestValidator IRequestValidator `field:"optional" json:"requestValidator" yaml:"requestValidator"`
	// Request validator options to create new validator Only one of `requestValidator` or `requestValidatorOptions` must be specified.
	//
	// Works together with `requestModels` or `requestParameters` to validate
	// the request before it reaches integration like Lambda Proxy Integration.
	// Default: - No default validator.
	//
	RequestValidatorOptions *RequestValidatorOptions `field:"optional" json:"requestValidatorOptions" yaml:"requestValidatorOptions"`
}

