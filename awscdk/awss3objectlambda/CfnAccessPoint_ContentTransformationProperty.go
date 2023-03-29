package awss3objectlambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentTransformationProperty := &ContentTransformationProperty{
//   	AwsLambda: &AwsLambdaProperty{
//   		FunctionArn: jsii.String("functionArn"),
//
//   		// the properties below are optional
//   		FunctionPayload: jsii.String("functionPayload"),
//   	},
//   }
//
type CfnAccessPoint_ContentTransformationProperty struct {
	// `CfnAccessPoint.ContentTransformationProperty.AwsLambda`.
	AwsLambda interface{} `field:"required" json:"awsLambda" yaml:"awsLambda"`
}

