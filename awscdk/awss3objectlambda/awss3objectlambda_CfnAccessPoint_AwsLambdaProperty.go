package awss3objectlambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsLambdaProperty := &awsLambdaProperty{
//   	functionArn: jsii.String("functionArn"),
//
//   	// the properties below are optional
//   	functionPayload: jsii.String("functionPayload"),
//   }
//
type CfnAccessPoint_AwsLambdaProperty struct {
	// `CfnAccessPoint.AwsLambdaProperty.FunctionArn`.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// `CfnAccessPoint.AwsLambdaProperty.FunctionPayload`.
	FunctionPayload *string `field:"optional" json:"functionPayload" yaml:"functionPayload"`
}

