package awsappsync


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   handlerConfigProperty := &HandlerConfigProperty{
//   	Behavior: jsii.String("behavior"),
//   	Integration: &IntegrationProperty{
//   		DataSourceName: jsii.String("dataSourceName"),
//
//   		// the properties below are optional
//   		LambdaConfig: &LambdaConfigProperty{
//   			InvokeType: jsii.String("invokeType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-handlerconfig.html
//
type CfnChannelNamespace_HandlerConfigProperty struct {
	// Integration behavior for a handler configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-handlerconfig.html#cfn-appsync-channelnamespace-handlerconfig-behavior
	//
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-handlerconfig.html#cfn-appsync-channelnamespace-handlerconfig-integration
	//
	Integration interface{} `field:"required" json:"integration" yaml:"integration"`
}

