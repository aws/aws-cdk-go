package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents a Lambda function defined outside of this stack.
//
// Example:
//   fn := lambda.Function_FromFunctionAttributes(this, jsii.String("Function"), &FunctionAttributes{
//   	FunctionArn: jsii.String("arn:aws:lambda:us-east-1:123456789012:function:MyFn"),
//   	// The following are optional properties for specific use cases and should be used with caution:
//
//   	// Use Case: imported function is in the same account as the stack. This tells the CDK that it
//   	// can modify the function's permissions.
//   	SameEnvironment: jsii.Boolean(true),
//
//   	// Use Case: imported function is in a different account and user commits to ensuring that the
//   	// imported function has the correct permissions outside the CDK.
//   	SkipPermissions: jsii.Boolean(true),
//   })
//
type FunctionAttributes struct {
	// The ARN of the Lambda function.
	//
	// Format: arn:<partition>:lambda:<region>:<account-id>:function:<function-name>.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The architecture of this Lambda Function (this is an optional attribute and defaults to X86_64).
	// Default: - Architecture.X86_64
	//
	Architecture Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// The IAM execution role associated with this function.
	//
	// If the role is not specified, any role-related operations will no-op.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Setting this property informs the CDK that the imported function is in the same environment as the stack.
	//
	// This affects certain behaviours such as, whether this function's permission can be modified.
	// When not configured, the CDK attempts to auto-determine this. For environment agnostic stacks, i.e., stacks
	// where the account is not specified with the `env` property, this is determined to be false.
	//
	// Set this to property *ONLY IF* the imported function is in the same account as the stack
	// it's imported in.
	// Default: - depends: true, if the Stack is configured with an explicit `env` (account and region) and the account is the same as this function.
	// For environment-agnostic stacks this will default to `false`.
	//
	SameEnvironment *bool `field:"optional" json:"sameEnvironment" yaml:"sameEnvironment"`
	// The security group of this Lambda, if in a VPC.
	//
	// This needs to be given in order to support allowing connections
	// to this Lambda.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Setting this property informs the CDK that the imported function ALREADY HAS the necessary permissions for what you are trying to do.
	//
	// When not configured, the CDK attempts to auto-determine whether or not
	// additional permissions are necessary on the function when grant APIs are used. If the CDK tried to add
	// permissions on an imported lambda, it will fail.
	//
	// Set this property *ONLY IF* you are committing to manage the imported function's permissions outside of
	// CDK. You are acknowledging that your CDK code alone will have insufficient permissions to access the
	// imported function.
	// Default: false.
	//
	SkipPermissions *bool `field:"optional" json:"skipPermissions" yaml:"skipPermissions"`
}

