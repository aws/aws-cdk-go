package awsapigateway


// Properties for defining a `CfnMethod`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMethodProps := &CfnMethodProps{
//   	HttpMethod: jsii.String("httpMethod"),
//   	ResourceId: jsii.String("resourceId"),
//   	RestApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	ApiKeyRequired: jsii.Boolean(false),
//   	AuthorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	AuthorizationType: jsii.String("authorizationType"),
//   	AuthorizerId: jsii.String("authorizerId"),
//   	Integration: &IntegrationProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		CacheKeyParameters: []*string{
//   			jsii.String("cacheKeyParameters"),
//   		},
//   		CacheNamespace: jsii.String("cacheNamespace"),
//   		ConnectionId: jsii.String("connectionId"),
//   		ConnectionType: jsii.String("connectionType"),
//   		ContentHandling: jsii.String("contentHandling"),
//   		Credentials: jsii.String("credentials"),
//   		IntegrationHttpMethod: jsii.String("integrationHttpMethod"),
//   		IntegrationResponses: []interface{}{
//   			&IntegrationResponseProperty{
//   				StatusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				ContentHandling: jsii.String("contentHandling"),
//   				ResponseParameters: map[string]*string{
//   					"responseParametersKey": jsii.String("responseParameters"),
//   				},
//   				ResponseTemplates: map[string]*string{
//   					"responseTemplatesKey": jsii.String("responseTemplates"),
//   				},
//   				SelectionPattern: jsii.String("selectionPattern"),
//   			},
//   		},
//   		PassthroughBehavior: jsii.String("passthroughBehavior"),
//   		RequestParameters: map[string]*string{
//   			"requestParametersKey": jsii.String("requestParameters"),
//   		},
//   		RequestTemplates: map[string]*string{
//   			"requestTemplatesKey": jsii.String("requestTemplates"),
//   		},
//   		TimeoutInMillis: jsii.Number(123),
//   		Uri: jsii.String("uri"),
//   	},
//   	MethodResponses: []interface{}{
//   		&MethodResponseProperty{
//   			StatusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			ResponseModels: map[string]*string{
//   				"responseModelsKey": jsii.String("responseModels"),
//   			},
//   			ResponseParameters: map[string]interface{}{
//   				"responseParametersKey": jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	OperationName: jsii.String("operationName"),
//   	RequestModels: map[string]*string{
//   		"requestModelsKey": jsii.String("requestModels"),
//   	},
//   	RequestParameters: map[string]interface{}{
//   		"requestParametersKey": jsii.Boolean(false),
//   	},
//   	RequestValidatorId: jsii.String("requestValidatorId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html
//
type CfnMethodProps struct {
	// The method's HTTP verb.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-httpmethod
	//
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The Resource identifier for the MethodResponse resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-resourceid
	//
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The string identifier of the associated RestApi.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-restapiid
	//
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// A boolean flag specifying whether a valid ApiKey is required to invoke this method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-apikeyrequired
	//
	ApiKeyRequired interface{} `field:"optional" json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// A list of authorization scopes configured on the method.
	//
	// The scopes are used with a `COGNITO_USER_POOLS` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationscopes
	//
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// The method's authorization type.
	//
	// This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html) in the *API Gateway API Reference* .
	//
	// > If you specify the `AuthorizerId` property, specify `CUSTOM` or `COGNITO_USER_POOLS` for this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizationtype
	//
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// The identifier of an authorizer to use on this method.
	//
	// The method's authorization type must be `CUSTOM` or `COGNITO_USER_POOLS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-authorizerid
	//
	AuthorizerId *string `field:"optional" json:"authorizerId" yaml:"authorizerId"`
	// Represents an `HTTP` , `HTTP_PROXY` , `AWS` , `AWS_PROXY` , or Mock integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-integration
	//
	Integration interface{} `field:"optional" json:"integration" yaml:"integration"`
	// Gets a method response associated with a given HTTP status code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-methodresponses
	//
	MethodResponses interface{} `field:"optional" json:"methodResponses" yaml:"methodResponses"`
	// A human-friendly operation identifier for the method.
	//
	// For example, you can assign the `operationName` of `ListPets` for the `GET /pets` method in the `PetStore` example.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-operationname
	//
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// A key-value map specifying data schemas, represented by Model resources, (as the mapped value) of the request payloads of given content types (as the mapping key).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestmodels
	//
	RequestModels interface{} `field:"optional" json:"requestModels" yaml:"requestModels"`
	// A key-value map defining required or optional method request parameters that can be accepted by API Gateway.
	//
	// A key is a method request parameter name matching the pattern of `method.request.{location}.{name}` , where `location` is `querystring` , `path` , or `header` and `name` is a valid and unique parameter name. The value associated with the key is a Boolean flag indicating whether the parameter is required ( `true` ) or optional ( `false` ). The method request parameter names defined here are available in Integration to be mapped to integration request parameters or templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestparameters
	//
	RequestParameters interface{} `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// The identifier of a RequestValidator for request validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html#cfn-apigateway-method-requestvalidatorid
	//
	RequestValidatorId *string `field:"optional" json:"requestValidatorId" yaml:"requestValidatorId"`
}

