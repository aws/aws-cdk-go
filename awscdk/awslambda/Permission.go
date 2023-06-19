package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Represents a permission statement that can be added to a Lambda function's resource policy via the `addPermission()` method.
//
// Example:
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
// Experimental.
type Permission struct {
	// The entity for which you are granting permission to invoke the Lambda function.
	//
	// This entity can be any valid AWS service principal, such as
	// s3.amazonaws.com or sns.amazonaws.com, or, if you are granting
	// cross-account permission, an AWS account ID. For example, you might want
	// to allow a custom application in another AWS account to push events to
	// Lambda by invoking your function.
	//
	// The principal can be either an AccountPrincipal or a ServicePrincipal.
	// Experimental.
	Principal awsiam.IPrincipal `field:"required" json:"principal" yaml:"principal"`
	// The Lambda actions that you want to allow in this statement.
	//
	// For example,
	// you can specify lambda:CreateFunction to specify a certain action, or use
	// a wildcard (``lambda:*``) to grant permission to all Lambda actions. For a
	// list of actions, see Actions and Condition Context Keys for AWS Lambda in
	// the IAM User Guide.
	// Experimental.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// A unique token that must be supplied by the principal invoking the function.
	// Experimental.
	EventSourceToken *string `field:"optional" json:"eventSourceToken" yaml:"eventSourceToken"`
	// The authType for the function URL that you are granting permissions for.
	// Experimental.
	FunctionUrlAuthType FunctionUrlAuthType `field:"optional" json:"functionUrlAuthType" yaml:"functionUrlAuthType"`
	// The scope to which the permission constructs be attached.
	//
	// The default is
	// the Lambda function construct itself, but this would need to be different
	// in cases such as cross-stack references where the Permissions would need
	// to sit closer to the consumer of this permission (i.e., the caller).
	// Experimental.
	Scope awscdk.Construct `field:"optional" json:"scope" yaml:"scope"`
	// The AWS account ID (without hyphens) of the source owner.
	//
	// For example, if
	// you specify an S3 bucket in the SourceArn property, this value is the
	// bucket owner's account ID. You can use this property to ensure that all
	// source principals are owned by a specific account.
	// Experimental.
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// The ARN of a resource that is invoking your function.
	//
	// When granting
	// Amazon Simple Storage Service (Amazon S3) permission to invoke your
	// function, specify this property with the bucket ARN as its value. This
	// ensures that events generated only from the specified bucket, not just
	// any bucket from any AWS account that creates a mapping to your function,
	// can invoke the function.
	// Experimental.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

