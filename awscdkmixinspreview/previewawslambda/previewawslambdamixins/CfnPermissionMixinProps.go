package previewawslambdamixins


// Properties for CfnPermissionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPermissionMixinProps := &CfnPermissionMixinProps{
//   	Action: jsii.String("action"),
//   	EventSourceToken: jsii.String("eventSourceToken"),
//   	FunctionName: jsii.String("functionName"),
//   	FunctionUrlAuthType: jsii.String("functionUrlAuthType"),
//   	InvokedViaFunctionUrl: jsii.Boolean(false),
//   	Principal: jsii.String("principal"),
//   	PrincipalOrgId: jsii.String("principalOrgId"),
//   	SourceAccount: jsii.String("sourceAccount"),
//   	SourceArn: jsii.String("sourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html
//
type CfnPermissionMixinProps struct {
	// The action that the principal can use on the function.
	//
	// For example, `lambda:InvokeFunction` or `lambda:GetFunction` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// For Alexa Smart Home functions, a token that the invoker must supply.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-eventsourcetoken
	//
	EventSourceToken *string `field:"optional" json:"eventSourceToken" yaml:"eventSourceToken"`
	// The name or ARN of the Lambda function, version, or alias.
	//
	// **Name formats** - *Function name* – `my-function` (name-only), `my-function:v1` (with alias).
	// - *Function ARN* – `arn:aws:lambda:us-west-2:123456789012:function:my-function` .
	// - *Partial ARN* – `123456789012:function:my-function` .
	//
	// You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-functionname
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The type of authentication that your function URL uses.
	//
	// Set to `AWS_IAM` if you want to restrict access to authenticated users only. Set to `NONE` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Control access to Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-functionurlauthtype
	//
	FunctionUrlAuthType *string `field:"optional" json:"functionUrlAuthType" yaml:"functionUrlAuthType"`
	// Restricts the `lambda:InvokeFunction` action to function URL calls.
	//
	// When specified, this option prevents the principal from invoking the function by any means other than the function URL. For more information, see [Control access to Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-invokedviafunctionurl
	//
	InvokedViaFunctionUrl interface{} `field:"optional" json:"invokedViaFunctionUrl" yaml:"invokedViaFunctionUrl"`
	// The AWS service , AWS account , IAM user, or IAM role that invokes the function.
	//
	// If you specify a service, use `SourceArn` or `SourceAccount` to limit who can invoke the function through that service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// The identifier for your organization in AWS Organizations .
	//
	// Use this to grant permissions to all the AWS accounts under this organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-principalorgid
	//
	PrincipalOrgId *string `field:"optional" json:"principalOrgId" yaml:"principalOrgId"`
	// For AWS service , the ID of the AWS account that owns the resource.
	//
	// Use this together with `SourceArn` to ensure that the specified account owns the resource. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-sourceaccount
	//
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// For AWS services , the ARN of the AWS resource that invokes the function.
	//
	// For example, an Amazon S3 bucket or Amazon SNS topic.
	//
	// Note that Lambda configures the comparison using the `StringLike` operator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html#cfn-lambda-permission-sourcearn
	//
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

