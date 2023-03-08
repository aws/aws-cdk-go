package awss3objectlambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsLambdaProperty := &AwsLambdaProperty{
//   	FunctionArn: jsii.String("functionArn"),
//
//   	// the properties below are optional
//   	FunctionPayload: jsii.String("functionPayload"),
//   }
//
type CfnAccessPoint_AwsLambdaProperty struct {
	// `CfnAccessPoint.AwsLambdaProperty.FunctionArn`.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// `CfnAccessPoint.AwsLambdaProperty.FunctionPayload`.
	FunctionPayload *string `field:"optional" json:"functionPayload" yaml:"functionPayload"`
}

