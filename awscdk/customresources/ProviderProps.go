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
//   var onEvent function
//   var isComplete function
//   var myRole role
//
//   myProvider := cr.NewProvider(this, jsii.String("MyProvider"), &ProviderProps{
//   	OnEventHandler: onEvent,
//   	IsCompleteHandler: isComplete,
//   	LogGroup: logs.NewLogGroup(this, jsii.String("MyProviderLogs"), &LogGroupProps{
//   		Retention: logs.RetentionDays_ONE_DAY,
//   	}),
//   	Role: myRole,
//   	ProviderFunctionName: jsii.String("the-lambda-name"),
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
	// Default: - a default log group created by AWS Lambda.
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The number of days framework log events are kept in CloudWatch Logs.
	//
	// When
	// updating this property, unsetting it doesn't remove the log retention policy.
	// To remove the retention policy, set the value to `INFINITE`.
	// Default: logs.RetentionDays.INFINITE
	//
	// Deprecated: Use logGroup for full control over the custom resource log group.
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
	// The role that will be assumed by the AWS Lambda.
	// Must be assumable by the 'lambda.amazonaws.com' service principal.
	// Default: - A default role will be created.
	//
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
}

