package awsecs


// The settings to use when creating a cluster.
//
// This parameter is used to turn on CloudWatch Container Insights for a cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterSettingsProperty := &clusterSettingsProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnCluster_ClusterSettingsProperty struct {
	// The name of the cluster setting.
	//
	// The only supported value is `containerInsights` .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value to set for the cluster setting.
	//
	// The supported values are `enabled` and `disabled` . If `enabled` is specified, CloudWatch Container Insights will be enabled for the cluster, otherwise it will be disabled unless the `containerInsights` account setting is enabled. If a cluster value is specified, it will override the `containerInsights` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html) .
	Value *string `field:"optional" json:"value" yaml:"value"`
}

