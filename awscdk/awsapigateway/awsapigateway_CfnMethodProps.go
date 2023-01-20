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
	// The HTTP method that clients use to call this method.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The ID of an API Gateway [resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html) . For root resource methods, specify the `RestApi` root resource ID, such as `{ "Fn::GetAtt": ["MyRestApi", "RootResourceId"] }` .
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The ID of the [RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource in which API Gateway creates the method.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Indicates whether the method requires clients to submit a valid API key.
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// A list of authorization scopes configured on the method.
	//
	// The scopes are used with a `COGNITO_USER_POOLS` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes match a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// The method's authorization type.
	//
	// This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/api-reference/resource/method/) in the *API Gateway API Reference* .
	//
	// > If you specify the `AuthorizerId` property, specify `CUSTOM` or `COGNITO_USER_POOLS` for this property.
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// The identifier of the [authorizer](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html) to use on this method. If you specify this property, specify `CUSTOM` or `COGNITO_USER_POOLS` for the `AuthorizationType` property.
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// The backend system that the method calls when it receives a request.
	Integration interface{} `field:"optional" json:"integration" yaml:"integration"`
	// The responses that can be sent to the client who calls the method.
	MethodResponses interface{} `field:"optional" json:"methodResponses" yaml:"methodResponses"`
	// A friendly operation name for the method.
	//
	// For example, you can assign the `OperationName` of `ListPets` for the `GET /pets` method.
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// The resources that are used for the request's content type.
	//
	// Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a `Model` resource name as the value. To use the same model regardless of the content type, specify `$default` as the key.
	RequestModels interface{} `field:"optional" json:"requestModels" yaml:"requestModels"`
	// The request parameters that API Gateway accepts.
	//
	// Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value. The Boolean specifies whether a parameter is required. A source must match the format `method.request. *location* . *name*` , where the location is querystring, path, or header, and *name* is a valid, unique parameter name.
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// The ID of the associated request validator.
	RequestValidatorId *string `field:"optional" json:"requestValidatorId" yaml:"requestValidatorId"`
}

