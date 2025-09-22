package awsappsync


// The `LambdaConfig` property type specifies the integration configuration for a Lambda data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConfigProperty := &LambdaConfigProperty{
//   	InvokeType: jsii.String("invokeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-lambdaconfig.html
//
type CfnChannelNamespace_LambdaConfigProperty struct {
	// The invocation type for a Lambda data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-lambdaconfig.html#cfn-appsync-channelnamespace-lambdaconfig-invoketype
	//
	InvokeType *string `field:"required" json:"invokeType" yaml:"invokeType"`
}

