package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizer authorizer
//   var duration duration
//   var integration integration
//   var model model
//   var requestValidator requestValidator
//   var resource resource
//
//   resourceProps := &ResourceProps{
//   	Parent: resource,
//   	PathPart: jsii.String("pathPart"),
//
//   	// the properties below are optional
//   	DefaultCorsPreflightOptions: &CorsOptions{
//   		AllowOrigins: []*string{
//   			jsii.String("allowOrigins"),
//   		},
//
//   		// the properties below are optional
//   		AllowCredentials: jsii.Boolean(false),
//   		AllowHeaders: []*string{
//   			jsii.String("allowHeaders"),
//   		},
//   		AllowMethods: []*string{
//   			jsii.String("allowMethods"),
//   		},
//   		DisableCache: jsii.Boolean(false),
//   		ExposeHeaders: []*string{
//   			jsii.String("exposeHeaders"),
//   		},
//   		MaxAge: duration,
//   		StatusCode: jsii.Number(123),
//   	},
//   	DefaultIntegration: integration,
//   	DefaultMethodOptions: &MethodOptions{
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
// Experimental.
type ResourceProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	// Experimental.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	// Experimental.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	// Experimental.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// The parent resource of this resource.
	//
	// You can either pass another
	// `Resource` object or a `RestApi` object here.
	// Experimental.
	Parent IResource `field:"required" json:"parent" yaml:"parent"`
	// A path name for the resource.
	// Experimental.
	PathPart *string `field:"required" json:"pathPart" yaml:"pathPart"`
}

