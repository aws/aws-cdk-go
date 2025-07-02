package awsbedrock


// A Lambda function that processes documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformationFunctionProperty := &TransformationFunctionProperty{
//   	TransformationLambdaConfiguration: &TransformationLambdaConfigurationProperty{
//   		LambdaArn: jsii.String("lambdaArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-transformationfunction.html
//
type CfnDataSource_TransformationFunctionProperty struct {
	// The Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-transformationfunction.html#cfn-bedrock-datasource-transformationfunction-transformationlambdaconfiguration
	//
	TransformationLambdaConfiguration interface{} `field:"required" json:"transformationLambdaConfiguration" yaml:"transformationLambdaConfiguration"`
}

