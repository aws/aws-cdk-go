package awsappsync


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   handlerConfigsProperty := &HandlerConfigsProperty{
//   	OnPublish: &HandlerConfigProperty{
//   		Behavior: jsii.String("behavior"),
//   		Integration: &IntegrationProperty{
//   			DataSourceName: jsii.String("dataSourceName"),
//
//   			// the properties below are optional
//   			LambdaConfig: &LambdaConfigProperty{
//   				InvokeType: jsii.String("invokeType"),
//   			},
//   		},
//   	},
//   	OnSubscribe: &HandlerConfigProperty{
//   		Behavior: jsii.String("behavior"),
//   		Integration: &IntegrationProperty{
//   			DataSourceName: jsii.String("dataSourceName"),
//
//   			// the properties below are optional
//   			LambdaConfig: &LambdaConfigProperty{
//   				InvokeType: jsii.String("invokeType"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-handlerconfigs.html
//
type CfnChannelNamespace_HandlerConfigsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-handlerconfigs.html#cfn-appsync-channelnamespace-handlerconfigs-onpublish
	//
	OnPublish interface{} `field:"optional" json:"onPublish" yaml:"onPublish"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-handlerconfigs.html#cfn-appsync-channelnamespace-handlerconfigs-onsubscribe
	//
	OnSubscribe interface{} `field:"optional" json:"onSubscribe" yaml:"onSubscribe"`
}

