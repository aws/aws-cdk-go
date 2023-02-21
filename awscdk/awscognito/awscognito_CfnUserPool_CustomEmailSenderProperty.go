package awscognito


// A custom email sender AWS Lambda trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customEmailSenderProperty := &customEmailSenderProperty{
//   	lambdaArn: jsii.String("lambdaArn"),
//   	lambdaVersion: jsii.String("lambdaVersion"),
//   }
//
type CfnUserPool_CustomEmailSenderProperty struct {
	// The Amazon Resource Name (ARN) of the AWS Lambda function that Amazon Cognito triggers to send email notifications to users.
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
	// The Lambda version represents the signature of the "request" attribute in the "event" information that Amazon Cognito passes to your custom email sender AWS Lambda function.
	//
	// The only supported value is `V1_0` .
	LambdaVersion *string `field:"optional" json:"lambdaVersion" yaml:"lambdaVersion"`
}

