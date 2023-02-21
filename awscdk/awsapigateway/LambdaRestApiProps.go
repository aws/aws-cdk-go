package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Example:
//   var backend function
//
//   api := apigateway.NewLambdaRestApi(this, jsii.String("myapi"), &LambdaRestApiProps{
//   	Handler: backend,
//   	Proxy: jsii.Boolean(false),
//   })
//
//   items := api.Root.AddResource(jsii.String("items"))
//   items.AddMethod(jsii.String("GET")) // GET /items
//   items.AddMethod(jsii.String("POST")) // POST /items
//
//   item := items.AddResource(jsii.String("{item}"))
//   item.AddMethod(jsii.String("GET")) // GET /items/{item}
//
//   // the default integration for methods is "handler", but one can
//   // customize this behavior per method or even a sub path.
//   item.AddMethod(jsii.String("DELETE"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")))
//
type LambdaRestApiProps struct {
	// Adds a CORS preflight OPTIONS method to this resource and all child resources.
	//
	// You can add CORS at the resource-level using `addCorsPreflight`.
	DefaultCorsPreflightOptions *CorsOptions `field:"optional" json:"defaultCorsPreflightOptions" yaml:"defaultCorsPreflightOptions"`
	// An integration to use as a default for all methods created within this API unless an integration is specified.
	DefaultIntegration Integration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// Method options to use as a default for all methods created within this API unless custom options are specified.
	DefaultMethodOptions *MethodOptions `field:"optional" json:"defaultMethodOptions" yaml:"defaultMethodOptions"`
	// Automatically configure an AWS CloudWatch role for API Gateway.
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
	Deploy *bool `field:"optional" json:"deploy" yaml:"deploy"`
	// Options for the API Gateway stage that will always point to the latest deployment when `deploy` is enabled.
	//
	// If `deploy` is disabled,
	// this value cannot be set.
	DeployOptions *StageOptions `field:"optional" json:"deployOptions" yaml:"deployOptions"`
	// A description of the RestApi construct.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke the API using the default execute-api endpoint.
	//
	// To require that clients use a custom domain name to invoke the
	// API, disable the default endpoint.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
	//
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Configure a custom domain name and map it to this API.
	DomainName *DomainNameOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Export name for the CfnOutput containing the API endpoint.
	EndpointExportName *string `field:"optional" json:"endpointExportName" yaml:"endpointExportName"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating
	// an API.
	EndpointTypes *[]EndpointType `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the RestApi resource.
	FailOnWarnings *bool `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Custom header parameters for the request.
	// See: https://docs.aws.amazon.com/cli/latest/reference/apigateway/import-rest-api.html
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for this RestApi.
	Policy awsiam.PolicyDocument `field:"optional" json:"policy" yaml:"policy"`
	// A name for the API Gateway RestApi resource.
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
	// Retains old deployment resources when the API changes.
	//
	// This allows
	// manually reverting stages to point to old deployments via the AWS
	// Console.
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
	// The source of the API key for metering requests according to a usage plan.
	ApiKeySourceType ApiKeySourceType `field:"optional" json:"apiKeySourceType" yaml:"apiKeySourceType"`
	// The list of binary media mime-types that are supported by the RestApi resource, such as "image/png" or "application/octet-stream".
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// The ID of the API Gateway RestApi resource that you want to clone.
	CloneFrom IRestApi `field:"optional" json:"cloneFrom" yaml:"cloneFrom"`
	// The EndpointConfiguration property type specifies the endpoint types of a REST API.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html
	//
	EndpointConfiguration *EndpointConfiguration `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (when undefined) on an API.
	//
	// When compression is enabled, compression or
	// decompression is not applied on the payload if the payload size is
	// smaller than this value. Setting it to zero allows compression for any
	// payload size.
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// The default Lambda function that handles all requests from this API.
	//
	// This handler will be used as a the default integration for all methods in
	// this API, unless specified otherwise in `addMethod`.
	Handler awslambda.IFunction `field:"required" json:"handler" yaml:"handler"`
	// Specific Lambda integration options.
	IntegrationOptions *LambdaIntegrationOptions `field:"optional" json:"integrationOptions" yaml:"integrationOptions"`
	// If true, route all requests to the Lambda Function.
	//
	// If set to false, you will need to explicitly define the API model using
	// `addResource` and `addMethod` (or `addProxy`).
	Proxy *bool `field:"optional" json:"proxy" yaml:"proxy"`
}

