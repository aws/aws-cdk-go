package mixinsawscognito


// The configuration of a custom email sender Lambda trigger.
//
// This trigger routes all email notifications from a user pool to a Lambda function that delivers the message using custom logic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customEmailSenderProperty := &CustomEmailSenderProperty{
//   	LambdaArn: jsii.String("lambdaArn"),
//   	LambdaVersion: jsii.String("lambdaVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-customemailsender.html
//
type CfnUserPoolPropsMixin_CustomEmailSenderProperty struct {
	// The Amazon Resource Name (ARN) of the function that you want to assign to your Lambda trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-customemailsender.html#cfn-cognito-userpool-customemailsender-lambdaarn
	//
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
	// The user pool trigger version of the request that Amazon Cognito sends to your Lambda function.
	//
	// Higher-numbered versions add fields that support new features.
	//
	// You must use a `LambdaVersion` of `V1_0` with a custom sender function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-customemailsender.html#cfn-cognito-userpool-customemailsender-lambdaversion
	//
	LambdaVersion *string `field:"optional" json:"lambdaVersion" yaml:"lambdaVersion"`
}

