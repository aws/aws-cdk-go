package awsappsync


// The `Integration` property type specifies the integration data source configuration for the handler.
//
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
	// The unique name of the data source that has been configured on the API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-integration.html#cfn-appsync-channelnamespace-integration-datasourcename
	//
	DataSourceName *string `field:"required" json:"dataSourceName" yaml:"dataSourceName"`
	// The configuration for a Lambda data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-channelnamespace-integration.html#cfn-appsync-channelnamespace-integration-lambdaconfig
	//
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
}

