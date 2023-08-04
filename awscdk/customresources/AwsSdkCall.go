package customresources


// An AWS SDK call.
//
// Example:
//   awsCustom := cr.NewAwsCustomResource(this, jsii.String("aws-custom"), &AwsCustomResourceProps{
//   	OnCreate: &AwsSdkCall{
//   		Service: jsii.String("..."),
//   		Action: jsii.String("..."),
//   		Parameters: map[string]*string{
//   			"text": jsii.String("..."),
//   		},
//   		PhysicalResourceId: cr.PhysicalResourceId_Of(jsii.String("...")),
//   	},
//   	OnUpdate: &AwsSdkCall{
//   		Service: jsii.String("..."),
//   		Action: jsii.String("..."),
//   		Parameters: map[string]interface{}{
//   			"text": jsii.String("..."),
//   			"resourceId": cr.NewPhysicalResourceIdReference(),
//   		},
//   	},
//   	Policy: cr.AwsCustomResourcePolicy_FromSdkCalls(&SdkCallsPolicyOptions{
//   		Resources: cr.AwsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
type AwsSdkCall struct {
	// The service action to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The service to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	Service *string `field:"required" json:"service" yaml:"service"`
	// API version to use for the service.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/locking-api-versions.html
	//
	// Default: - use latest available API version.
	//
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Used for running the SDK calls in underlying lambda with a different role Can be used primarily for cross-account requests to for example connect hostedzone with a shared vpc.
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

