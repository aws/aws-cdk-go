package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a permission statement that can be added to a Lambda function's resource policy via the `addPermission()` method.
//
// Example:
//   // Grant permissions to a service
//   var fn function
//
//   principal := iam.NewServicePrincipal(jsii.String("my-service"))
//
//   fn.GrantInvoke(principal)
//
//   // Equivalent to:
//   fn.AddPermission(jsii.String("my-service Invocation"), &Permission{
//   	Principal: principal,
//   })
//
type Permission struct {
	// The entity for which you are granting permission to invoke the Lambda function.
	//
	// This entity can be any of the following:
	//
	// - a valid AWS service principal, such as `s3.amazonaws.com` or `sns.amazonaws.com`
	// - an AWS account ID for cross-account permissions. For example, you might want
	//   to allow a custom application in another AWS account to push events to
	//   Lambda by invoking your function.
	// - an AWS organization principal to grant permissions to an entire organization.
	//
	// The principal can be an AccountPrincipal, an ArnPrincipal, a ServicePrincipal,
	// or an OrganizationPrincipal.
	Principal awsiam.IPrincipal `field:"required" json:"principal" yaml:"principal"`
	// The Lambda actions that you want to allow in this statement.
	//
	// For example,
	// you can specify lambda:CreateFunction to specify a certain action, or use
	// a wildcard (``lambda:*``) to grant permission to all Lambda actions. For a
	// list of actions, see Actions and Condition Context Keys for AWS Lambda in
	// the IAM User Guide.
	// Default: 'lambda:InvokeFunction'.
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// A unique token that must be supplied by the principal invoking the function.
	// Default: - The caller would not need to present a token.
	//
	EventSourceToken *string `field:"optional" json:"eventSourceToken" yaml:"eventSourceToken"`
	// The authType for the function URL that you are granting permissions for.
	// Default: - No functionUrlAuthType.
	//
	FunctionUrlAuthType FunctionUrlAuthType `field:"optional" json:"functionUrlAuthType" yaml:"functionUrlAuthType"`
	// The organization you want to grant permissions to.
	//
	// Use this ONLY if you
	// need to grant permissions to a subset of the organization. If you want to
	// grant permissions to the entire organization, sending the organization principal
	// through the `principal` property will suffice.
	//
	// You can use this property to ensure that all source principals are owned by
	// a specific organization.
	// Default: - No organizationId.
	//
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// The scope to which the permission constructs be attached.
	//
	// The default is
	// the Lambda function construct itself, but this would need to be different
	// in cases such as cross-stack references where the Permissions would need
	// to sit closer to the consumer of this permission (i.e., the caller).
	// Default: - The instance of lambda.IFunction
	//
	Scope constructs.Construct `field:"optional" json:"scope" yaml:"scope"`
	// The AWS account ID (without hyphens) of the source owner.
	//
	// For example, if
	// you specify an S3 bucket in the SourceArn property, this value is the
	// bucket owner's account ID. You can use this property to ensure that all
	// source principals are owned by a specific account.
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// The ARN of a resource that is invoking your function.
	//
	// When granting
	// Amazon Simple Storage Service (Amazon S3) permission to invoke your
	// function, specify this property with the bucket ARN as its value. This
	// ensures that events generated only from the specified bucket, not just
	// any bucket from any AWS account that creates a mapping to your function,
	// can invoke the function.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

