package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Properties for defining a `WebFunction`.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var codeBucket IBucket
//
//
//   lambda.NewWebFunction(this, jsii.String("PerRegionApi"), &WebFunctionProps{
//   	Code: lambda.WebFunctionCode_FromBucket(codeBucket, jsii.String("app.zip"), jsii.String("version-id")),
//   	Runtime: lambda.WebFunctionRuntime_NODEJS_24_X(),
//   	DefaultEndpointAuth: lambda.WebFunctionAuthType_IAM_AUTH,
//   	DefaultEndpoint: &WebFunctionDefaultEndpointOptions{
//   		EndpointLocation: lambda.WebFunctionEndpointLocation_PerRegion([]*string{
//   			jsii.String("us-east-1"),
//   			jsii.String("us-west-2"),
//   		}),
//   	},
//   })
//
type WebFunctionProps struct {
	// The deployment package for the WebFunction.
	Code WebFunctionCode `field:"required" json:"code" yaml:"code"`
	// The authentication type of the endpoint the WebFunction creates, `endpoint`.
	//
	// There is no default: you must choose explicitly so an endpoint is never accidentally public or
	// accidentally locked down. Specify `WebFunctionAuthType.IAM_AUTH` (Lambda authorizes requests with SigV4/IAM)
	// or `WebFunctionAuthType.APPLICATION_MANAGED` (the endpoint is public and your function handles authorization).
	//
	// With `IAM_AUTH`, requests must be signed with SigV4 and carry the `X-Amz-Content-SHA256`
	// header, also for a request without a body; without it the endpoint returns 403 even for a
	// valid signature. A `MULTI_REGION` endpoint is signed with SigV4a, for the Region set `*`.
	DefaultEndpointAuth WebFunctionAuthType `field:"required" json:"defaultEndpointAuth" yaml:"defaultEndpointAuth"`
	// The runtime environment for the WebFunction.
	Runtime WebFunctionRuntime `field:"required" json:"runtime" yaml:"runtime"`
	// The minimum log level for application logs (logs emitted by your code, e.g. `console.log()`, `console.warn()`). Only logs at this level or above will be captured in CloudWatch.
	//
	// For example, setting this to `INFO` will capture `console.info()` and
	// `console.warn()` but filter out `console.debug()`.
	// Default: - No application log level filtering.
	//
	ApplicationLogLevel ApplicationLogLevel `field:"optional" json:"applicationLogLevel" yaml:"applicationLogLevel"`
	// The other settings of the endpoint the WebFunction creates, `endpoint`.
	//
	// Endpoints added with `addEndpoint()` take their own settings, auth type included.
	// Default: - a home-Region endpoint named `prod`, with no description and no scaling or rate
	// limit beyond your account's quotas.
	//
	DefaultEndpoint *WebFunctionDefaultEndpointOptions `field:"optional" json:"defaultEndpoint" yaml:"defaultEndpoint"`
	// The network connector that gives the function egress into a VPC.
	//
	// The revision pins to the connector's latest version at the time it is created. For a
	// connector defined in the same stack, every update to the connector therefore replaces the
	// revision as part of that deployment, and traffic moves to the new version with it. To hold
	// the function on a fixed version instead — or to use a connector from another stack, whose
	// changing version cannot be exported — import the version with
	// `NetworkConnector.fromNetworkConnectorVersionArn()`.
	//
	// The connector must have been created for `NetworkConnectorComputeType.WEB_FUNCTION`, and
	// the endpoint must be `WebFunctionEndpointType.HOME_REGION`.
	//
	// The full `INetworkConnector` is required (rather than a ref) because the revision pins to
	// `latestVersionArn`, which `INetworkConnectorRef` does not carry.
	//
	// [disable-awslint:prefer-ref-interface].
	// Default: - no VPC egress.
	//
	EgressNetworkConnector INetworkConnector `field:"optional" json:"egressNetworkConnector" yaml:"egressNetworkConnector"`
	// The entry file the runtime executes, relative to the root of the deployment package.
	//
	// Only supported for Node.js runtimes. When set, it is passed to the runtime via the reserved
	// `AWS_LAMBDA_NODEJS_ENTRYPOINT` environment variable. Do not set that
	// variable through `environment` as well.
	// Default: - the runtime's default entry point, `index.js`
	//
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Key-value pairs of environment variables to pass to the WebFunction.
	// Default: - No environment variables.
	//
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The KMS key used to encrypt environment variables.
	//
	// The Lambda service principal is granted the key usage it needs to encrypt
	// and decrypt the revision's environment variables.
	// Default: - AWS managed key.
	//
	KmsKey interfacesawskms.IKeyRef `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The CloudWatch log group to send logs to.
	// Default: - A log group is created, named by CloudFormation and retained when the stack is
	// deleted. Read its name from the `logGroup` attribute.
	//
	LogGroup interfacesawslogs.ILogGroupRef `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The maximum number of concurrent requests handled by a single execution environment.
	//
	// Must be an integer between 1 and 128. When omitted, nothing is sent and the service
	// applies its default.
	// Default: - the service default of 64 is applied.
	//
	MaxConcurrencyPerEnvironment *float64 `field:"optional" json:"maxConcurrencyPerEnvironment" yaml:"maxConcurrencyPerEnvironment"`
	// A description of the WebFunction revision.
	//
	// Must be 256 characters or fewer.
	// Default: - No description.
	//
	RevisionDescription *string `field:"optional" json:"revisionDescription" yaml:"revisionDescription"`
	// The IAM execution role for the WebFunction.
	// Default: - A new role with `AWSLambdaBasicExecutionRole` is created.
	//
	Role interface{ interfacesawsiam.IRoleRef;awsiam.IGrantable } `field:"optional" json:"role" yaml:"role"`
	// The minimum log level for system logs (logs emitted by the Lambda runtime, e.g. start/end/report lines and runtime errors). Only logs at this level or above will be captured in CloudWatch.
	// Default: - No system log level filtering.
	//
	SystemLogLevel SystemLogLevel `field:"optional" json:"systemLogLevel" yaml:"systemLogLevel"`
	// Tags to apply to the WebFunction.
	//
	// Only the WebFunction resource carries tags. Its revision and endpoint are not
	// taggable, so nothing is applied to them.
	// Default: - No tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The function execution time (in seconds) after which the WebFunction is terminated.
	//
	// Must be an integer between 3 and 900 seconds.
	// Default: Duration.seconds(30)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The name of the WebFunction.
	//
	// Must be 2 to 64 characters of letters, numbers, hyphens or underscores, starting and
	// ending with a letter or number.
	// Default: - A unique name is generated by CDK.
	//
	WebFunctionName *string `field:"optional" json:"webFunctionName" yaml:"webFunctionName"`
}

