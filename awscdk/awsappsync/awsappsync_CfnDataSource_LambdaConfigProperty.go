package awsappsync


// The `LambdaConfig` property type specifies the Lambda function ARN for an AWS AppSync data source.
//
// `LambdaConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConfigProperty := &lambdaConfigProperty{
//   	lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   }
//
type CfnDataSource_LambdaConfigProperty struct {
	// The ARN for the Lambda function.
	LambdaFunctionArn *string `field:"required" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
}

