package awsbedrock


// A Lambda function that processes documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformationLambdaConfigurationProperty := &TransformationLambdaConfigurationProperty{
//   	LambdaArn: jsii.String("lambdaArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-transformationlambdaconfiguration.html
//
type CfnDataSource_TransformationLambdaConfigurationProperty struct {
	// The function's ARN identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-transformationlambdaconfiguration.html#cfn-bedrock-datasource-transformationlambdaconfiguration-lambdaarn
	//
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
}

