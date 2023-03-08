package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Props to create a new instance of RestApi.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   	definition: stepfunctions.chain.start(stepfunctions.NewPass(this, jsii.String("Pass"))),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("Api"), &restApiProps{
//   	restApiName: jsii.String("MyApi"),
//   })
//   api.root.addMethod(jsii.String("GET"), apigateway.stepFunctionsIntegration.startExecution(stateMachine))
//
// Experimental.
type RestApiProps struct {
	// Automatically configure an AWS CloudWatch role for API Gateway.
	// Experimental.
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
	// Experimental.
	Deploy *bool `field:"optional" json:"deploy" yaml:"deploy"`
	// Options for the API Gateway stage that will always point to the latest deployment when `deploy` is enabled.
	//
	// If `deploy` is disabled,
	// this value cannot be set.
	// Experimental.
	DeployOptions *StageOptions `field:"optional" json:"deployOptions" yaml:"deployOptions"`
	// Specifies whether clients can invoke the API using the default execute-api endpoint.
	//
	// To require that clients use a custom domain name to invoke the
	// API, disable the default endpoint.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
	//
	// Experimental.
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Configure a custom domain name and map it to this API.
	// Experimental.
	DomainName *DomainNameOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Export name for the CfnOutput containing the API endpoint.
	// Experimental.
	EndpointExportName *string `field:"optional" json:"endpointExportName" yaml:"endpointExportName"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating
	// an API.
	// Experimental.
	EndpointTypes *[]EndpointType `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the RestApi resource.
	// Experimental.
	FailOnWarnings *bool `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Custom header parameters for the request.
	// See: https://docs.aws.amazon.com/cli/latest/reference/apigateway/import-rest-api.html
	//
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for this RestApi.
	// Experimental.
	Policy awsiam.PolicyDocument `field:"optional" json:"policy" yaml:"policy"`
	// A name for the API Gateway RestApi resource.
	// Experimental.
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
	// Retains old deployment resources when the API changes.
	//
	// This allows
	// manually reverting stages to point to old deployments via the AWS
	// Console.
	// Experimental.
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
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
	// The source of the API key for metering requests according to a usage plan.
	// Experimental.
	ApiKeySourceType ApiKeySourceType `field:"optional" json:"apiKeySourceType" yaml:"apiKeySourceType"`
	// The list of binary media mime-types that are supported by the RestApi resource, such as "image/png" or "application/octet-stream".
	// Experimental.
	BinaryMediaTypes *[]*string `field:"optional" json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// The ID of the API Gateway RestApi resource that you want to clone.
	// Experimental.
	CloneFrom IRestApi `field:"optional" json:"cloneFrom" yaml:"cloneFrom"`
	// A description of the purpose of this API Gateway RestApi resource.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The EndpointConfiguration property type specifies the endpoint types of a REST API.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html
	//
	// Experimental.
	EndpointConfiguration *EndpointConfiguration `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (when undefined) on an API.
	//
	// When compression is enabled, compression or
	// decompression is not applied on the payload if the payload size is
	// smaller than this value. Setting it to zero allows compression for any
	// payload size.
	// Experimental.
	MinimumCompressionSize *float64 `field:"optional" json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
}

