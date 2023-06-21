package customresources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
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
//   	LogRetention: logs.RetentionDays_ONE_DAY,
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
	IsCompleteHandler awslambda.IFunction `field:"optional" json:"isCompleteHandler" yaml:"isCompleteHandler"`
	// The number of days framework log events are kept in CloudWatch Logs.
	//
	// When
	// updating this property, unsetting it doesn't remove the log retention policy.
	// To remove the retention policy, set the value to `INFINITE`.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// Provider Lambda name.
	//
	// The provider lambda function name.
	ProviderFunctionName *string `field:"optional" json:"providerFunctionName" yaml:"providerFunctionName"`
	// Time between calls to the `isComplete` handler which determines if the resource has been stabilized.
	//
	// The first `isComplete` will be called immediately after `handler` and then
	// every `queryInterval` seconds, and until `timeout` has been reached or until
	// `isComplete` returns `true`.
	QueryInterval awscdk.Duration `field:"optional" json:"queryInterval" yaml:"queryInterval"`
	// AWS Lambda execution role.
	//
	// The role that will be assumed by the AWS Lambda.
	// Must be assumable by the 'lambda.amazonaws.com' service principal.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security groups to attach to the provider functions.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Total timeout for the entire operation.
	//
	// The maximum timeout is 2 hours (yes, it can exceed the AWS Lambda 15 minutes).
	TotalTimeout awscdk.Duration `field:"optional" json:"totalTimeout" yaml:"totalTimeout"`
	// The vpc to provision the lambda functions in.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Which subnets from the VPC to place the lambda functions in.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

