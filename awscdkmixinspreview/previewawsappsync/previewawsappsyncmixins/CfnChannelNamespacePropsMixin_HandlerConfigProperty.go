package previewawsappsyncmixins


// The `HandlerConfig` property type specifies the configuration for the handler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   handlerConfigProperty := &HandlerConfigProperty{
//   	Behavior: jsii.String("behavior"),
//   	Integration: &IntegrationProperty{
//   		DataSourceName: jsii.String("dataSourceName"),
//   		LambdaConfig: &LambdaConfigProperty{
//   			InvokeType: jsii.String("invokeType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-handlerconfig.html
//
type CfnChannelNamespacePropsMixin_HandlerConfigProperty struct {
	// The behavior for the handler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-handlerconfig.html#cfn-appsync-channelnamespace-handlerconfig-behavior
	//
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// The integration data source configuration for the handler.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-handlerconfig.html#cfn-appsync-channelnamespace-handlerconfig-integration
	//
	Integration interface{} `field:"optional" json:"integration" yaml:"integration"`
}

