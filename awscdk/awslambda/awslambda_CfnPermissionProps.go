package awslambda


// Properties for defining a `CfnPermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermissionProps := &cfnPermissionProps{
//   	action: jsii.String("action"),
//   	functionName: jsii.String("functionName"),
//   	principal: jsii.String("principal"),
//
//   	// the properties below are optional
//   	eventSourceToken: jsii.String("eventSourceToken"),
//   	functionUrlAuthType: jsii.String("functionUrlAuthType"),
//   	principalOrgId: jsii.String("principalOrgId"),
//   	sourceAccount: jsii.String("sourceAccount"),
//   	sourceArn: jsii.String("sourceArn"),
//   }
//
type CfnPermissionProps struct {
	// The action that the principal can use on the function.
	//
	// For example, `lambda:InvokeFunction` or `lambda:GetFunction` .
	Action *string `field:"required" json:"action" yaml:"action"`
	// The name of the Lambda function, version, or alias.
	//
	// **Name formats** - *Function name* - `my-function` (name-only), `my-function:v1` (with alias).
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:my-function` .
	// - *Partial ARN* - `123456789012:function:my-function` .
	//
	// You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The AWS service or account that invokes the function.
	//
	// If you specify a service, use `SourceArn` or `SourceAccount` to limit who can invoke the function through that service.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// For Alexa Smart Home functions, a token that must be supplied by the invoker.
	EventSourceToken *string `field:"optional" json:"eventSourceToken" yaml:"eventSourceToken"`
	// The type of authentication that your function URL uses.
	//
	// Set to `AWS_IAM` if you want to restrict access to authenticated `IAM` users only. Set to `NONE` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html) .
	FunctionUrlAuthType *string `field:"optional" json:"functionUrlAuthType" yaml:"functionUrlAuthType"`
	// The identifier for your organization in AWS Organizations .
	//
	// Use this to grant permissions to all the AWS accounts under this organization.
	PrincipalOrgId *string `field:"optional" json:"principalOrgId" yaml:"principalOrgId"`
	// For Amazon S3, the ID of the account that owns the resource.
	//
	// Use this together with `SourceArn` to ensure that the resource is owned by the specified account. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// For AWS services, the ARN of the AWS resource that invokes the function.
	//
	// For example, an Amazon S3 bucket or Amazon SNS topic.
	//
	// Note that Lambda configures the comparison using the `StringLike` operator.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

