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
//   methodProps := &methodProps{
//   	httpMethod: jsii.String("httpMethod"),
//   	resource: resource,
//
//   	// the properties below are optional
//   	integration: integration,
//   	options: &methodOptions{
//   		apiKeyRequired: jsii.Boolean(false),
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizationType: awscdk.Aws_apigateway.authorizationType_NONE,
//   		authorizer: authorizer,
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("statusCode"),
//
//   				// the properties below are optional
//   				responseModels: map[string]iModel{
//   					"responseModelsKey": model,
//   				},
//   				responseParameters: map[string]*bool{
//   					"responseParametersKey": jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		operationName: jsii.String("operationName"),
//   		requestModels: map[string]*iModel{
//   			"requestModelsKey": model,
//   		},
//   		requestParameters: map[string]*bool{
//   			"requestParametersKey": jsii.Boolean(false),
//   		},
//   		requestValidator: requestValidator,
//   		requestValidatorOptions: &requestValidatorOptions{
//   			requestValidatorName: jsii.String("requestValidatorName"),
//   			validateRequestBody: jsii.Boolean(false),
//   			validateRequestParameters: jsii.Boolean(false),
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
	Integration Integration `field:"optional" json:"integration" yaml:"integration"`
	// Method options.
	Options *MethodOptions `field:"optional" json:"options" yaml:"options"`
}

