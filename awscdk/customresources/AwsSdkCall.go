package customresources


// An AWS SDK call.
//
// Example:
//   cr.NewAwsCustomResource(this, jsii.String("GetParameterCustomResource"), &AwsCustomResourceProps{
//   	OnUpdate: &AwsSdkCall{
//   		 // will also be called for a CREATE event
//   		Service: jsii.String("SSM"),
//   		Action: jsii.String("getParameter"),
//   		Parameters: map[string]interface{}{
//   			"Name": jsii.String("my-parameter"),
//   			"WithDecryption": jsii.Boolean(true),
//   		},
//   		PhysicalResourceId: cr.PhysicalResourceId_FromResponse(jsii.String("Parameter.ARN")),
//   	},
//   	Policy: cr.AwsCustomResourcePolicy_FromSdkCalls(&SdkCallsPolicyOptions{
//   		Resources: cr.AwsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
type AwsSdkCall struct {
	// The service action to call.
	//
	// This is the name of an AWS API call, in one of the following forms:
	//
	// - An API call name as found in the API Reference documentation (`GetObject`)
	// - The API call name starting with a lowercase letter (`getObject`)
	// - The AWS SDK for JavaScript v3 command class name (`GetObjectCommand`).
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The service to call.
	//
	// This is the name of an AWS service, in one of the following forms:
	//
	// - An AWS SDK for JavaScript v3 package name (`@aws-sdk/client-api-gateway`)
	// - An AWS SDK for JavaScript v3 client name (`api-gateway`)
	// - An AWS SDK for JavaScript v2 constructor name (`APIGateway`)
	// - A lowercase AWS SDK for JavaScript v2 constructor name (`apigateway`).
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// API version to use for the service.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/locking-api-versions.html
	//
	// Default: - use latest available API version.
	//
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Used for running the SDK calls in underlying lambda with a different role.
	//
	// Can be used primarily for cross-account requests to for example connect
	// hostedzone with a shared vpc.
	// Region controls where assumeRole call is made.
	//
	// Example for Route53 / associateVPCWithHostedZone.
	// Default: - run without assuming role.
	//
	AssumedRoleArn *string `field:"optional" json:"assumedRoleArn" yaml:"assumedRoleArn"`
	// The regex pattern to use to catch API errors.
	//
	// The `code` property of the
	// `Error` object will be tested against this pattern. If there is a match an
	// error will not be thrown.
	// Default: - do not catch errors.
	//
	IgnoreErrorCodesMatching *string `field:"optional" json:"ignoreErrorCodesMatching" yaml:"ignoreErrorCodesMatching"`
	// A property used to configure logging during lambda function execution.
	//
	// Note: The default Logging configuration is all. This configuration will enable logging on all logged data
	// in the lambda handler. This includes:
	//  - The event object that is received by the lambda handler
	//  - The response received after making a API call
	//  - The response object that the lambda handler will return
	//  - SDK versioning information
	// - Caught and uncaught errors.
	// Default: Logging.all()
	//
	Logging Logging `field:"optional" json:"logging" yaml:"logging"`
	// Restrict the data returned by the custom resource to specific paths in the API response.
	//
	// Use this to limit the data returned by the custom
	// resource if working with API calls that could potentially result in custom
	// response objects exceeding the hard limit of 4096 bytes.
	//
	// Example for ECS / updateService: ['service.deploymentConfiguration.maximumPercent']
	// Default: - return all data.
	//
	OutputPaths *[]*string `field:"optional" json:"outputPaths" yaml:"outputPaths"`
	// The parameters for the service action.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Default: - no parameters.
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The physical resource id of the custom resource for this call.
	//
	// Mandatory for onCreate call.
	// In onUpdate, you can omit this to passthrough it from request.
	// Default: - no physical resource id.
	//
	PhysicalResourceId PhysicalResourceId `field:"optional" json:"physicalResourceId" yaml:"physicalResourceId"`
	// The region to send service requests to.
	//
	// **Note: Cross-region operations are generally considered an anti-pattern.**
	// **Consider first deploying a stack in that region.**
	// Default: - the region where this custom resource is deployed.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

