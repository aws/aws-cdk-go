package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizer authorizer
//   var integration integration
//   var model model
//   var requestValidator requestValidator
//   var resource resource
//
//   methodProps := &MethodProps{
//   	HttpMethod: jsii.String("httpMethod"),
//   	Resource: resource,
//
//   	// the properties below are optional
//   	Integration: integration,
//   	Options: &MethodOptions{
//   		ApiKeyRequired: jsii.Boolean(false),
//   		AuthorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		AuthorizationType: awscdk.Aws_apigateway.AuthorizationType_NONE,
//   		Authorizer: authorizer,
//   		MethodResponses: []methodResponse{
//   			&methodResponse{
//   				StatusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				ResponseModels: map[string]iModel{
//   					"responseModelsKey": model,
//   				},
//   				ResponseParameters: map[string]*bool{
//   					"responseParametersKey": jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		OperationName: jsii.String("operationName"),
//   		RequestModels: map[string]*iModel{
//   			"requestModelsKey": model,
//   		},
//   		RequestParameters: map[string]*bool{
//   			"requestParametersKey": jsii.Boolean(false),
//   		},
//   		RequestValidator: requestValidator,
//   		RequestValidatorOptions: &RequestValidatorOptions{
//   			RequestValidatorName: jsii.String("requestValidatorName"),
//   			ValidateRequestBody: jsii.Boolean(false),
//   			ValidateRequestParameters: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type MethodProps struct {
	// The HTTP method ("GET", "POST", "PUT", ...) that clients use to call this method.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// The resource this method is associated with.
	//
	// For root resource methods,
	// specify the `RestApi` object.
	Resource IResource `field:"required" json:"resource" yaml:"resource"`
	// The backend system that the method calls when it receives a request.
	// Default: - a new `MockIntegration`.
	//
	Integration Integration `field:"optional" json:"integration" yaml:"integration"`
	// Method options.
	// Default: - No options.
	//
	Options *MethodOptions `field:"optional" json:"options" yaml:"options"`
}

