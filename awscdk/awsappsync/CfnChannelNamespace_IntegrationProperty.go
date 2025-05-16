package awsappsync


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationProperty := &IntegrationProperty{
//   	DataSourceName: jsii.String("dataSourceName"),
//
//   	// the properties below are optional
//   	LambdaConfig: &LambdaConfigProperty{
//   		InvokeType: jsii.String("invokeType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-integration.html
//
type CfnChannelNamespace_IntegrationProperty struct {
	// Data source to invoke for this integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-integration.html#cfn-appsync-channelnamespace-integration-datasourcename
	//
	DataSourceName *string `field:"required" json:"dataSourceName" yaml:"dataSourceName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-integration.html#cfn-appsync-channelnamespace-integration-lambdaconfig
	//
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
}

