package customresources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Initialization properties for the `Provider` construct.
//
// Example:
//   // Create custom resource handler entrypoint
//   handler := lambda.NewFunction(this, jsii.String("my-handler"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_20_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String(`
//   	  exports.handler = async (event, context) => {
//   	    return {
//   	      PhysicalResourceId: '1234',
//   	      NoEcho: true,
//   	      Data: {
//   	        mySecret: 'secret-value',
//   	        hello: 'world',
//   	        ghToken: 'gho_xxxxxxx',
//   	      },
//   	    };
//   	  };`)),
//   })
//
//   // Provision a custom resource provider framework
//   provider := cr.NewProvider(this, jsii.String("my-provider"), &ProviderProps{
//   	OnEventHandler: handler,
//   })
//
//   awscdk.NewCustomResource(this, jsii.String("my-cr"), &CustomResourceProps{
//   	ServiceToken: provider.ServiceToken,
//   })
//
type ProviderProps struct {
	// The AWS Lambda function to invoke for all resource lifecycle operations (CREATE/UPDATE/DELETE).
	//
	// This function is responsible to begin the requested resource operation
	// (CREATE/UPDATE/DELETE) and return any additional properties to add to the
	// event, which will later be passed to `isComplete`. The `PhysicalResourceId`
	// property must be included in the response.
	OnEventHandler awslambda.IFunction `field:"required" json:"onEventHandler" yaml:"onEventHandler"`
	// Whether logging for the waiter state machine is disabled.
	// Default: - false.
	//
	DisableWaiterStateMachineLogging *bool `field:"optional" json:"disableWaiterStateMachineLogging" yaml:"disableWaiterStateMachineLogging"`
	// Lambda execution role for provider framework's isComplete/onTimeout Lambda function.
	//
	// Note that this role
	// must be assumed by the 'lambda.amazonaws.com' service principal. To prevent circular dependency problem
	// in the provider framework, please ensure you specify a different IAM Role for 'frameworkCompleteAndTimeoutRole'
	// from 'frameworkOnEventRole'.
	//
	// This property cannot be used with 'role' property.
	// Default: - A default role will be created.
	//
	FrameworkCompleteAndTimeoutRole awsiam.IRole `field:"optional" json:"frameworkCompleteAndTimeoutRole" yaml:"frameworkCompleteAndTimeoutRole"`
	// Lambda execution role for provider framework's onEvent Lambda function.
	//
	// Note that this role must be assumed
	// by the 'lambda.amazonaws.com' service principal.
	//
	// This property cannot be used with 'role' property.
	// Default: - A default role will be created.
	//
	FrameworkOnEventRole awsiam.IRole `field:"optional" json:"frameworkOnEventRole" yaml:"frameworkOnEventRole"`
	// The AWS Lambda function to invoke in order to determine if the operation is complete.
	//
	// This function will be called immediately after `onEvent` and then
	// periodically based on the configured query interval as long as it returns
	// `false`. If the function still returns `false` and the alloted timeout has
	// passed, the operation will fail.
	// Default: - provider is synchronous. This means that the `onEvent` handler
	// is expected to finish all lifecycle operations within the initial invocation.
	//
	IsCompleteHandler awslambda.IFunction `field:"optional" json:"isCompleteHandler" yaml:"isCompleteHandler"`
	// The Log Group used for logging of events emitted by the custom resource's lambda function.
	//
	// Providing a user-controlled log group was rolled out to commercial regions on 2023-11-16.
	// If you are deploying to another type of region, please check regional availability first.
	// Default: - a default log group created by AWS Lambda.
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The number of days framework log events are kept in CloudWatch Logs.
	//
	// When
	// updating this property, unsetting it doesn't remove the log retention policy.
	// To remove the retention policy, set the value to `INFINITE`.
	//
	// This is a legacy API and we strongly recommend you migrate to `logGroup` if you can.
	// `logGroup` allows you to create a fully customizable log group and instruct the Lambda function to send logs to it.
	// Default: logs.RetentionDays.INFINITE
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// AWS KMS key used to encrypt provider lambda's environment variables.
	// Default: -  AWS Lambda creates and uses an AWS managed customer master key (CMK).
	//
	ProviderFunctionEnvEncryption awskms.IKey `field:"optional" json:"providerFunctionEnvEncryption" yaml:"providerFunctionEnvEncryption"`
	// Provider Lambda name.
	//
	// The provider lambda function name.
	// Default: -  CloudFormation default name from unique physical ID.
	//
	ProviderFunctionName *string `field:"optional" json:"providerFunctionName" yaml:"providerFunctionName"`
	// Time between calls to the `isComplete` handler which determines if the resource has been stabilized.
	//
	// The first `isComplete` will be called immediately after `handler` and then
	// every `queryInterval` seconds, and until `timeout` has been reached or until
	// `isComplete` returns `true`.
	// Default: Duration.seconds(5)
	//
	QueryInterval awscdk.Duration `field:"optional" json:"queryInterval" yaml:"queryInterval"`
	// AWS Lambda execution role.
	//
	// The role is shared by provider framework's onEvent, isComplete lambda, and onTimeout Lambda functions.
	// This role will be assumed by the AWS Lambda, so it must be assumable by the 'lambda.amazonaws.com'
	// service principal.
	// Default: - A default role will be created.
	//
	// Deprecated: - Use frameworkOnEventLambdaRole, frameworkIsCompleteLambdaRole, frameworkOnTimeoutLambdaRole.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security groups to attach to the provider functions.
	//
	// Only used if 'vpc' is supplied.
	// Default: - If `vpc` is not supplied, no security groups are attached. Otherwise, a dedicated security
	// group is created for each function.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Total timeout for the entire operation.
	//
	// The maximum timeout is 1 hour (yes, it can exceed the AWS Lambda 15 minutes).
	// Default: Duration.minutes(30)
	//
	TotalTimeout awscdk.Duration `field:"optional" json:"totalTimeout" yaml:"totalTimeout"`
	// The vpc to provision the lambda functions in.
	// Default: - functions are not provisioned inside a vpc.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Which subnets from the VPC to place the lambda functions in.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Defines what execution history events of the waiter state machine are logged and where they are logged.
	// Default: - A default log group will be created if logging for the waiter state machine is enabled.
	//
	WaiterStateMachineLogOptions *LogOptions `field:"optional" json:"waiterStateMachineLogOptions" yaml:"waiterStateMachineLogOptions"`
}

