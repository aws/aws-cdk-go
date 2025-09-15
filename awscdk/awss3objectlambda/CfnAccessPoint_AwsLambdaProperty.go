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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-awslambda.html
//
type CfnAccessPoint_AwsLambdaProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-awslambda.html#cfn-s3objectlambda-accesspoint-awslambda-functionarn
	//
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3objectlambda-accesspoint-awslambda.html#cfn-s3objectlambda-accesspoint-awslambda-functionpayload
	//
	FunctionPayload *string `field:"optional" json:"functionPayload" yaml:"functionPayload"`
}

