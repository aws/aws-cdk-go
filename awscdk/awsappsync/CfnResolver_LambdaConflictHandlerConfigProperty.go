package awsappsync


// The `LambdaConflictHandlerConfig` when configuring LAMBDA as the Conflict Handler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConflictHandlerConfigProperty := &LambdaConflictHandlerConfigProperty{
//   	LambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-resolver-lambdaconflicthandlerconfig.html
//
type CfnResolver_LambdaConflictHandlerConfigProperty struct {
	// The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-resolver-lambdaconflicthandlerconfig.html#cfn-appsync-resolver-lambdaconflicthandlerconfig-lambdaconflicthandlerarn
	//
	LambdaConflictHandlerArn *string `field:"optional" json:"lambdaConflictHandlerArn" yaml:"lambdaConflictHandlerArn"`
}

