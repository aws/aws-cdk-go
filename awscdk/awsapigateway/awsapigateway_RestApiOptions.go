package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Represents the props that all Rest APIs share.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessLogDestination iAccessLogDestination
//   var accessLogFormat accessLogFormat
//   var authorizer authorizer
//   var bucket bucket
//   var certificate certificate
//   var duration duration
//   var integration integration
//   var model model
//   var policyDocument policyDocument
//   var requestValidator requestValidator
//
//   restApiOptions := &RestApiOptions{
//   	CloudWatchRole: jsii.Boolean(false),
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
//   	Deploy: jsii.Boolean(false),
//   	DeployOptions: &StageOptions{
//   		AccessLogDestination: accessLogDestination,
//   		AccessLogFormat: accessLogFormat,
//   		CacheClusterEnabled: jsii.Boolean(false),
//   		CacheClusterSize: jsii.String("cacheClusterSize"),
//   		CacheDataEncrypted: jsii.Boolean(false),
//   		CacheTtl: duration,
//   		CachingEnabled: jsii.Boolean(false),
//   		ClientCertificateId: jsii.String("clientCertificateId"),
//   		DataTraceEnabled: jsii.Boolean(false),
//   		Description: jsii.String("description"),
//   		DocumentationVersion: jsii.String("documentationVersion"),
//   		LoggingLevel: awscdk.*Aws_apigateway.MethodLoggingLevel_OFF,
//   		MethodOptions: map[string]methodDeploymentOptions{
//   			"methodOptionsKey": &methodDeploymentOptions{
//   				"cacheDataEncrypted": jsii.Boolean(false),
//   				"cacheTtl": duration,
//   				"cachingEnabled": jsii.Boolean(false),
//   				"dataTraceEnabled": jsii.Boolean(false),
//   				"loggingLevel": awscdk.*Aws_apigateway.MethodLoggingLevel_OFF,
//   				"metricsEnabled": jsii.Boolean(false),
//   				"throttlingBurstLimit": jsii.Number(123),
//   				"throttlingRateLimit": jsii.Number(123),
//   			},
//   		},
//   		MetricsEnabled: jsii.Boolean(false),
//   		StageName: jsii.String("stageName"),
//   		ThrottlingBurstLimit: jsii.Number(123),
//   		ThrottlingRateLimit: jsii.Number(123),
//   		TracingEnabled: jsii.Boolean(false),
//   		Variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	DisableExecuteApiEndpoint: jsii.Boolean(false),
//   	DomainName: &DomainNameOptions{
//   		Certificate: certificate,
//   		DomainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		BasePath: jsii.String("basePath"),
//   		EndpointType: awscdk.*Aws_apigateway.EndpointType_EDGE,
//   		Mtls: &MTLSConfig{
//   			Bucket: bucket,
//   			Key: jsii.String("key"),
//
//   			// the properties below are optional
//   			Version: jsii.String("version"),
//   		},
//   		SecurityPolicy: awscdk.*Aws_apigateway.SecurityPolicy_TLS_1_0,
//   	},
//   	EndpointExportName: jsii.String("endpointExportName"),
//   	EndpointTypes: []endpointType{
//   		awscdk.*Aws_apigateway.*endpointType_EDGE,
//   	},
//   	FailOnWarnings: jsii.Boolean(false),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Policy: policyDocument,
//   	RestApiName: jsii.String("restApiName"),
//   	RetainDeployments: jsii.Boolean(false),
//   }
//
// Deprecated: - superseded by `RestApiBaseProps`.
type RestApiOptions struct {
	// Automatically configure an AWS CloudWatch role for API Gateway.
	// Deprecated: - superseded by `RestApiBaseProps`.
	CloudWatchRole *bool `field:"optional" json:"cloudWatchRole" yaml:"cloudWatchRole"`
	// Indicates if a Deployment should be automatically created for this API, and recreated when the API model (resources, methods) changes.
	//
	// Since API Gateway deployments are immutable, When this option is enabled
	// (by default), an AWS::ApiGateway::Deployment resource will automatically
	// created with a logical ID that hashes the API model (methods, resources
	// and options). This means that when the model changes, the logical ID of
	// this CloudFormation resource will change, and a new deployment will be
	// created.
	//
	// If this is set, `latestDeployment` will refer to the `Deployment` object
	// and `deploymentStage` will refer to a `Stage` that points to this
	// deployment. To customize the stage options, use the `deployOptions`
	// property.
	//
	// A CloudFormation Output will also be defined with the root URL endpoint
	// of this REST API.
	// Deprecated: - superseded by `RestApiBaseProps`.
	Deploy *bool `field:"optional" json:"deploy" yaml:"deploy"`
	// Options for the API Gateway stage that will always point to the latest deployment when `deploy` is enabled.
	//
	// If `deploy` is disabled,
	// this value cannot be set.
	// Deprecated: - superseded by `RestApiBaseProps`.
	DeployOptions *StageOptions `field:"optional" json:"deployOptions" yaml:"deployOptions"`
	// Specifies whether clients can invoke the API using the default execute-api endpoint.
	//
	// To require that clients use a custom domain name to invoke the
	// API, disable the default endpoint.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
	//
	// Deprecated: - superseded by `RestApiBaseProps`.
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Configure a custom domain name and map it to this API.
	// Deprecated: - superseded by `RestApiBaseProps`.
	DomainName *DomainNameOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Export name for the CfnOutput containing the API endpoint.
	// Deprecated: - superseded by `RestApiBaseProps`.
	EndpointExportName *string `field:"optional" json:"endpointExportName" yaml:"endpointExportName"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating
	// an API.
	// Deprecated: - superseded by `RestApiBaseProps`.
	EndpointTypes *[]EndpointType `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the RestApi resource.
	// Deprecated: - superseded by `RestApiBaseProps`.
	FailOnWarnings *bool `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Custom header parameters for the request.
	// See: https://docs.aws.amazon.com/cli/latest/reference/apigateway/import-rest-api.html
	//
	// Deprecated: - superseded by `RestApiBaseProps`.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for this RestApi.
	// Deprecated: - superseded by `RestApiBaseProps`.
	Policy awsiam.PolicyDocument `field:"optional" json:"policy" yaml:"policy"`
	// A name for the API Gateway RestApi resource.
	// Deprecated: - superseded by `RestApiBaseProps`.
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
	// Retains old deployment resources when the API changes.
	//
	// This allows
	// manually reverting stages to point to old deployments via the AWS
	// Console.
	// Deprecated: - superseded by `RestApiBaseProps`.
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	// Deprecated: - superseded by `RestApiBaseProps`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	// Deprecated: - superseded by `RestApiBaseProps`.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	// Deprecated: - superseded by `RestApiBaseProps`.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
}

