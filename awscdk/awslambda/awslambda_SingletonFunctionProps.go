package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties for a newly created singleton Lambda.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var architecture architecture
//   var code code
//   var codeSigningConfig codeSigningConfig
//   var destination iDestination
//   var eventSource iEventSource
//   var fileSystem fileSystem
//   var key key
//   var lambdaInsightsVersion lambdaInsightsVersion
//   var layerVersion layerVersion
//   var policyStatement policyStatement
//   var profilingGroup profilingGroup
//   var queue queue
//   var role role
//   var runtime runtime
//   var securityGroup securityGroup
//   var size size
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var topic topic
//   var vpc vpc
//
//   singletonFunctionProps := &singletonFunctionProps{
//   	code: code,
//   	handler: jsii.String("handler"),
//   	runtime: runtime,
//   	uuid: jsii.String("uuid"),
//
//   	// the properties below are optional
//   	allowAllOutbound: jsii.Boolean(false),
//   	allowPublicSubnet: jsii.Boolean(false),
//   	architecture: architecture,
//   	codeSigningConfig: codeSigningConfig,
//   	currentVersionOptions: &versionOptions{
//   		codeSha256: jsii.String("codeSha256"),
//   		description: jsii.String("description"),
//   		maxEventAge: cdk.duration.minutes(jsii.Number(30)),
//   		onFailure: destination,
//   		onSuccess: destination,
//   		provisionedConcurrentExecutions: jsii.Number(123),
//   		removalPolicy: cdk.removalPolicy_DESTROY,
//   		retryAttempts: jsii.Number(123),
//   	},
//   	deadLetterQueue: queue,
//   	deadLetterQueueEnabled: jsii.Boolean(false),
//   	deadLetterTopic: topic,
//   	description: jsii.String("description"),
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	environmentEncryption: key,
//   	ephemeralStorageSize: size,
//   	events: []*iEventSource{
//   		eventSource,
//   	},
//   	filesystem: fileSystem,
//   	functionName: jsii.String("functionName"),
//   	initialPolicy: []*policyStatement{
//   		policyStatement,
//   	},
//   	insightsVersion: lambdaInsightsVersion,
//   	lambdaPurpose: jsii.String("lambdaPurpose"),
//   	layers: []iLayerVersion{
//   		layerVersion,
//   	},
//   	logRetention: awscdk.Aws_logs.retentionDays_ONE_DAY,
//   	logRetentionRetryOptions: &logRetentionRetryOptions{
//   		base: cdk.*duration.minutes(jsii.Number(30)),
//   		maxRetries: jsii.Number(123),
//   	},
//   	logRetentionRole: role,
//   	maxEventAge: cdk.*duration.minutes(jsii.Number(30)),
//   	memorySize: jsii.Number(123),
//   	onFailure: destination,
//   	onSuccess: destination,
//   	profiling: jsii.Boolean(false),
//   	profilingGroup: profilingGroup,
//   	reservedConcurrentExecutions: jsii.Number(123),
//   	retryAttempts: jsii.Number(123),
//   	role: role,
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	timeout: cdk.*duration.minutes(jsii.Number(30)),
//   	tracing: awscdk.Aws_lambda.tracing_ACTIVE,
//   	vpc: vpc,
//   	vpcSubnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type SingletonFunctionProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `field:"optional" json:"onSuccess" yaml:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Whether to allow the Lambda to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Lambda Functions in a public subnet can NOT access the internet.
	//
	// Use this property to acknowledge this limitation and still place the function in a public subnet.
	// See: https://stackoverflow.com/questions/52992085/why-cant-an-aws-lambda-function-inside-a-public-subnet-in-a-vpc-connect-to-the/52994841#52994841
	//
	AllowPublicSubnet *bool `field:"optional" json:"allowPublicSubnet" yaml:"allowPublicSubnet"`
	// The system architectures compatible with this lambda function.
	Architecture Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// Code signing config associated with this function.
	CodeSigningConfig ICodeSigningConfig `field:"optional" json:"codeSigningConfig" yaml:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	CurrentVersionOptions *VersionOptions `field:"optional" json:"currentVersionOptions" yaml:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	//
	// If SNS topic is desired, specify `deadLetterTopic` property instead.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	DeadLetterQueueEnabled *bool `field:"optional" json:"deadLetterQueueEnabled" yaml:"deadLetterQueueEnabled"`
	// The SNS topic to use as a DLQ.
	//
	// Note that if `deadLetterQueueEnabled` is set to `true`, an SQS queue will be created
	// rather than an SNS topic. Using an SNS topic as a DLQ requires this property to be set explicitly.
	DeadLetterTopic awssns.ITopic `field:"optional" json:"deadLetterTopic" yaml:"deadLetterTopic"`
	// A description of the function.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	EnvironmentEncryption awskms.IKey `field:"optional" json:"environmentEncryption" yaml:"environmentEncryption"`
	// The size of the function’s /tmp directory in MiB.
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	Events *[]IEventSource `field:"optional" json:"events" yaml:"events"`
	// The filesystem configuration for the lambda function.
	Filesystem FileSystem `field:"optional" json:"filesystem" yaml:"filesystem"`
	// A name for the function.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	InitialPolicy *[]awsiam.PolicyStatement `field:"optional" json:"initialPolicy" yaml:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-docker.html
	//
	InsightsVersion LambdaInsightsVersion `field:"optional" json:"insightsVersion" yaml:"insightsVersion"`
	// A list of layers to add to the function's execution environment.
	//
	// You can configure your Lambda function to pull in
	// additional code during initialization in the form of layers. Layers are packages of libraries or other dependencies
	// that can be used by multiple functions.
	Layers *[]ILayerVersion `field:"optional" json:"layers" yaml:"layers"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	LogRetentionRetryOptions *LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	LogRetentionRole awsiam.IRole `field:"optional" json:"logRetentionRole" yaml:"logRetentionRole"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// Enable profiling.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	Profiling *bool `field:"optional" json:"profiling" yaml:"profiling"`
	// Profiling Group.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `field:"optional" json:"profilingGroup" yaml:"profilingGroup"`
	// The maximum of concurrent executions you want to reserve for the function.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html
	//
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// Lambda execution role.
	//
	// This is the role that will be assumed by the function upon execution.
	// It controls the permissions that the function will have. The Role must
	// be assumable by the 'lambda.amazonaws.com' service principal.
	//
	// The default Role automatically has permissions granted for Lambda execution. If you
	// provide a Role, you must add the relevant AWS managed policies yourself.
	//
	// The relevant managed policies are "service-role/AWSLambdaBasicExecutionRole" and
	// "service-role/AWSLambdaVPCAccessExecutionRole".
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The list of security groups to associate with the Lambda's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Enable AWS X-Ray Tracing for Lambda Function.
	Tracing Tracing `field:"optional" json:"tracing" yaml:"tracing"`
	// VPC network to place Lambda network interfaces.
	//
	// Specify this if the Lambda function needs to access resources in a VPC.
	// This is required when `vpcSubnets` is specified.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// This requires `vpc` to be specified in order for interfaces to actually be
	// placed in the subnets. If `vpc` is not specify, this will raise an error.
	//
	// Note: Internet access for Lambda Functions requires a NAT Gateway, so picking
	// public subnets is not allowed (unless `allowPublicSubnet` is set to `true`).
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The source code of your Lambda function.
	//
	// You can point to a file in an
	// Amazon Simple Storage Service (Amazon S3) bucket or specify your source
	// code as inline text.
	Code Code `field:"required" json:"code" yaml:"code"`
	// The name of the method within your code that Lambda calls to execute your function.
	//
	// The format includes the file name. It can also include
	// namespaces and other qualifiers, depending on the runtime.
	// For more information, see https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html.
	//
	// Use `Handler.FROM_IMAGE` when defining a function from a Docker image.
	//
	// NOTE: If you specify your source code as inline text by specifying the
	// ZipFile property within the Code property, specify index.function_name as
	// the handler.
	Handler *string `field:"required" json:"handler" yaml:"handler"`
	// The runtime environment for the Lambda function that you are uploading.
	//
	// For valid values, see the Runtime property in the AWS Lambda Developer
	// Guide.
	//
	// Use `Runtime.FROM_IMAGE` when defining a function from a Docker image.
	Runtime Runtime `field:"required" json:"runtime" yaml:"runtime"`
	// A unique identifier to identify this lambda.
	//
	// The identifier should be unique across all custom resource providers.
	// We recommend generating a UUID per provider.
	Uuid *string `field:"required" json:"uuid" yaml:"uuid"`
	// A descriptive name for the purpose of this Lambda.
	//
	// If the Lambda does not have a physical name, this string will be
	// reflected its generated name. The combination of lambdaPurpose
	// and uuid must be unique.
	LambdaPurpose *string `field:"optional" json:"lambdaPurpose" yaml:"lambdaPurpose"`
}

