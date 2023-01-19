package awsapigateway


// Properties for defining a `CfnMethod`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMethodProps := &cfnMethodProps{
//   	httpMethod: jsii.String("httpMethod"),
//   	resourceId: jsii.String("resourceId"),
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	apiKeyRequired: jsii.Boolean(false),
//   	authorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	authorizationType: jsii.String("authorizationType"),
//   	authorizerId: jsii.String("authorizerId"),
//   	integration: &integrationProperty{
//   		cacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		cacheNamespace: jsii.String("cacheNamespace"),
//   		connectionId: jsii.String("connectionId"),
//   		connectionType: jsii.String("connectionType"),
//   		contentHandling: jsii.String("contentHandling"),
//   		credentials: jsii.String("credentials"),
//   		integrationHttpMethod: jsii.String("integrationHttpMethod"),
//   		integrationResponses: []interface{}{
//   			&integrationResponseProperty{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				contentHandling: jsii.String("contentHandling"),
//   				responseParameters: map[string]*string{
//   					"responseParametersKey": jsii.String("responseParameters"),
//   				},
//   				responseTemplates: map[string]*string{
//   					"responseTemplatesKey": jsii.String("responseTemplates"),
//   				},
//   				selectionPattern: jsii.String("selectionPattern"),
//   			},
//   		},
//   		passthroughBehavior: jsii.String("passthroughBehavior"),
//   		requestParameters: map[string]*string{
//   			"requestParametersKey": jsii.String("requestParameters"),
//   		},
//   		requestTemplates: map[string]*string{
//   			"requestTemplatesKey": jsii.String("requestTemplates"),
//   		},
//   		timeoutInMillis: jsii.Number(123),
//   		type: jsii.String("type"),
//   		uri: jsii.String("uri"),
//   	},
//   	methodResponses: []interface{}{
//   		&methodResponseProperty{
//   			statusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			responseModels: map[string]*string{
//   				"responseModelsKey": jsii.String("responseModels"),
//   			},
//   			responseParameters: map[string]interface{}{
//   				"responseParametersKey": jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	operationName: jsii.String("operationName"),
//   	requestModels: map[string]*string{
//   		"requestModelsKey": jsii.String("requestModels"),
//   	},
//   	requestParameters: map[string]interface{}{
//   		"requestParametersKey": jsii.Boolean(false),
//   	},
//   	requestValidatorId: jsii.String("requestValidatorId"),
//   }
//
type CfnMethodProps struct {
	// The method's HTTP verb.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The Resource identifier for the MethodResponse resource.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The string identifier of the associated RestApi.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// A boolean flag specifying whether a valid ApiKey is required to invoke this method.
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// A list of authorization scopes configured on the method.
	//
	// The scopes are used with a `COGNITO_USER_POOLS` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// The method's authorization type.
	//
	// This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html) in the *API Gateway API Reference* .
	//
	// > If you specify the `AuthorizerId` property, specify `CUSTOM` or `COGNITO_USER_POOLS` for this property.
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// The identifier of an Authorizer to use on this method.
	//
	// The `authorizationType` must be `CUSTOM` .
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// Represents an `HTTP` , `HTTP_PROXY` , `AWS` , `AWS_PROXY` , or Mock integration.
	Integration interface{} `field:"optional" json:"integration" yaml:"integration"`
	// Gets a method response associated with a given HTTP status code.
	MethodResponses interface{} `field:"optional" json:"methodResponses" yaml:"methodResponses"`
	// A human-friendly operation identifier for the method.
	//
	// For example, you can assign the `operationName` of `ListPets` for the `GET /pets` method in the `PetStore` example.
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// A key-value map specifying data schemas, represented by Model resources, (as the mapped value) of the request payloads of given content types (as the mapping key).
	RequestModels interface{} `field:"optional" json:"requestModels" yaml:"requestModels"`
	// A key-value map defining required or optional method request parameters that can be accepted by API Gateway.
	//
	// A key is a method request parameter name matching the pattern of `method.request.{location}.{name}` , where `location` is `querystring` , `path` , or `header` and `name` is a valid and unique parameter name. The value associated with the key is a Boolean flag indicating whether the parameter is required ( `true` ) or optional ( `false` ). The method request parameter names defined here are available in Integration to be mapped to integration request parameters or templates.
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// The identifier of a RequestValidator for request validation.
	RequestValidatorId *string `field:"optional" json:"requestValidatorId" yaml:"requestValidatorId"`
}

