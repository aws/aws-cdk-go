package awscognito


// A custom SMS sender AWS Lambda trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customSMSSenderProperty := &customSMSSenderProperty{
//   	lambdaArn: jsii.String("lambdaArn"),
//   	lambdaVersion: jsii.String("lambdaVersion"),
//   }
//
type CfnUserPool_CustomSMSSenderProperty struct {
	// The Amazon Resource Name (ARN) of the AWS Lambda function that Amazon Cognito triggers to send SMS notifications to users.
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
	// The Lambda version represents the signature of the "request" attribute in the "event" information Amazon Cognito passes to your custom SMS sender Lambda function.
	//
	// The only supported value is `V1_0` .
	LambdaVersion *string `field:"optional" json:"lambdaVersion" yaml:"lambdaVersion"`
}

