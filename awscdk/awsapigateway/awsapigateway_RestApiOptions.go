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
//   restApiOptions := &restApiOptions{
//   	cloudWatchRole: jsii.Boolean(false),
//   	defaultCorsPreflightOptions: &corsOptions{
//   		allowOrigins: []*string{
//   			jsii.String("allowOrigins"),
//   		},
//
//   		// the properties below are optional
//   		allowCredentials: jsii.Boolean(false),
//   		allowHeaders: []*string{
//   			jsii.String("allowHeaders"),
//   		},
//   		allowMethods: []*string{
//   			jsii.String("allowMethods"),
//   		},
//   		disableCache: jsii.Boolean(false),
//   		exposeHeaders: []*string{
//   			jsii.String("exposeHeaders"),
//   		},
//   		maxAge: duration,
//   		statusCode: jsii.Number(123),
//   	},
//   	defaultIntegration: integration,
//   	defaultMethodOptions: &methodOptions{
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
//   	deploy: jsii.Boolean(false),
//   	deployOptions: &stageOptions{
//   		accessLogDestination: accessLogDestination,
//   		accessLogFormat: accessLogFormat,
//   		cacheClusterEnabled: jsii.Boolean(false),
//   		cacheClusterSize: jsii.String("cacheClusterSize"),
//   		cacheDataEncrypted: jsii.Boolean(false),
//   		cacheTtl: duration,
//   		cachingEnabled: jsii.Boolean(false),
//   		clientCertificateId: jsii.String("clientCertificateId"),
//   		dataTraceEnabled: jsii.Boolean(false),
//   		description: jsii.String("description"),
//   		documentationVersion: jsii.String("documentationVersion"),
//   		loggingLevel: awscdk.*Aws_apigateway.methodLoggingLevel_OFF,
//   		methodOptions: map[string]methodDeploymentOptions{
//   			"methodOptionsKey": &methodDeploymentOptions{
//   				"cacheDataEncrypted": jsii.Boolean(false),
//   				"cacheTtl": duration,
//   				"cachingEnabled": jsii.Boolean(false),
//   				"dataTraceEnabled": jsii.Boolean(false),
//   				"loggingLevel": awscdk.*Aws_apigateway.*methodLoggingLevel_OFF,
//   				"metricsEnabled": jsii.Boolean(false),
//   				"throttlingBurstLimit": jsii.Number(123),
//   				"throttlingRateLimit": jsii.Number(123),
//   			},
//   		},
//   		metricsEnabled: jsii.Boolean(false),
//   		stageName: jsii.String("stageName"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   		tracingEnabled: jsii.Boolean(false),
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	disableExecuteApiEndpoint: jsii.Boolean(false),
//   	domainName: &domainNameOptions{
//   		certificate: certificate,
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		basePath: jsii.String("basePath"),
//   		endpointType: awscdk.*Aws_apigateway.endpointType_EDGE,
//   		mtls: &mTLSConfig{
//   			bucket: bucket,
//   			key: jsii.String("key"),
//
//   			// the properties below are optional
//   			version: jsii.String("version"),
//   		},
//   		securityPolicy: awscdk.*Aws_apigateway.securityPolicy_TLS_1_0,
//   	},
//   	endpointExportName: jsii.String("endpointExportName"),
//   	endpointTypes: []*endpointType{
//   		awscdk.*Aws_apigateway.*endpointType_EDGE,
//   	},
//   	failOnWarnings: jsii.Boolean(false),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	policy: policyDocument,
//   	restApiName: jsii.String("restApiName"),
//   	retainDeployments: jsii.Boolean(false),
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

