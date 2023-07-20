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
//   clusterSettingsProperty := &ClusterSettingsProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-clustersettings.html
//
type CfnCluster_ClusterSettingsProperty struct {
	// The name of the cluster setting.
	//
	// The value is `containerInsights` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-clustersettings.html#cfn-ecs-cluster-clustersettings-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value to set for the cluster setting. The supported values are `enabled` and `disabled` .
	//
	// If you set `name` to `containerInsights` and `value` to `enabled` , CloudWatch Container Insights will be on for the cluster, otherwise it will be off unless the `containerInsights` account setting is turned on. If a cluster value is specified, it will override the `containerInsights` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-clustersettings.html#cfn-ecs-cluster-clustersettings-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

