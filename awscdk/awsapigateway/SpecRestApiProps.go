package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props to instantiate a new SpecRestApi.
//
// Example:
//   var integration integration
//
//
//   api := apigateway.NewSpecRestApi(this, jsii.String("books-api"), &SpecRestApiProps{
//   	ApiDefinition: apigateway.ApiDefinition_FromAsset(jsii.String("path-to-file.json")),
//   })
//
//   booksResource := api.Root.AddResource(jsii.String("books"))
//   booksResource.AddMethod(jsii.String("GET"), integration)
//
type SpecRestApiProps struct {
	// Automatically configure an AWS CloudWatch role for API Gateway.
	// Default: - false if `@aws-cdk/aws-apigateway:disableCloudWatchRole` is enabled, true otherwise.
	//
	CloudWatchRole *bool `field:"optional" json:"cloudWatchRole" yaml:"cloudWatchRole"`
	// The removal policy applied to the AWS CloudWatch role when this resource is removed from the application.
	//
	// Requires `cloudWatchRole` to be enabled.
	// Default: - RemovalPolicy.RETAIN
	//
	CloudWatchRoleRemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"cloudWatchRoleRemovalPolicy" yaml:"cloudWatchRoleRemovalPolicy"`
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
	// Default: true.
	//
	Deploy *bool `field:"optional" json:"deploy" yaml:"deploy"`
	// Options for the API Gateway stage that will always point to the latest deployment when `deploy` is enabled.
	//
	// If `deploy` is disabled,
	// this value cannot be set.
	// Default: - Based on defaults of `StageOptions`.
	//
	DeployOptions *StageOptions `field:"optional" json:"deployOptions" yaml:"deployOptions"`
	// A description of the RestApi construct.
	// Default: - 'Automatically created by the RestApi construct'.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke the API using the default execute-api endpoint.
	//
	// To require that clients use a custom domain name to invoke the
	// API, disable the default endpoint.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
	//
	// Default: false.
	//
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// Configure a custom domain name and map it to this API.
	// Default: - no domain name is defined, use `addDomainName` or directly define a `DomainName`.
	//
	DomainName *DomainNameOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Export name for the CfnOutput containing the API endpoint.
	// Default: - when no export name is given, output will be created without export.
	//
	EndpointExportName *string `field:"optional" json:"endpointExportName" yaml:"endpointExportName"`
	// A list of the endpoint types of the API.
	//
	// Use this property when creating
	// an API.
	// Default: EndpointType.EDGE
	//
	EndpointTypes *[]EndpointType `field:"optional" json:"endpointTypes" yaml:"endpointTypes"`
	// Indicates whether to roll back the resource if a warning occurs while API Gateway is creating the RestApi resource.
	// Default: false.
	//
	FailOnWarnings *bool `field:"optional" json:"failOnWarnings" yaml:"failOnWarnings"`
	// Custom header parameters for the request.
	// See: https://docs.aws.amazon.com/cli/latest/reference/apigateway/import-rest-api.html
	//
	// Default: - No parameters.
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// A policy document that contains the permissions for this RestApi.
	// Default: - No policy.
	//
	Policy awsiam.PolicyDocument `field:"optional" json:"policy" yaml:"policy"`
	// A name for the API Gateway RestApi resource.
	// Default: - ID of the RestApi construct.
	//
	RestApiName *string `field:"optional" json:"restApiName" yaml:"restApiName"`
	// Retains old deployment resources when the API changes.
	//
	// This allows
	// manually reverting stages to point to old deployments via the AWS
	// Console.
	// Default: false.
	//
	RetainDeployments *bool `field:"optional" json:"retainDeployments" yaml:"retainDeployments"`
	// An OpenAPI definition compatible with API Gateway.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api.html
	//
	ApiDefinition ApiDefinition `field:"required" json:"apiDefinition" yaml:"apiDefinition"`
	// A Size(in bytes, kibibytes, mebibytes etc) that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (when undefined) on an API.
	//
	// When compression is enabled, compression or
	// decompression is not applied on the payload if the payload size is
	// smaller than this value. Setting it to zero allows compression for any
	// payload size.
	// Default: - Compression is disabled.
	//
	MinCompressionSize awscdk.Size `field:"optional" json:"minCompressionSize" yaml:"minCompressionSize"`
}

