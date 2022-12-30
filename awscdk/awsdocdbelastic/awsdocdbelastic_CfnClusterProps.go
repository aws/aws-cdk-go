package awsdocdbelastic

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &cfnClusterProps{
//   	adminUserName: jsii.String("adminUserName"),
//   	authType: jsii.String("authType"),
//   	clusterName: jsii.String("clusterName"),
//   	shardCapacity: jsii.Number(123),
//   	shardCount: jsii.Number(123),
//
//   	// the properties below are optional
//   	adminUserPassword: jsii.String("adminUserPassword"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
type CfnClusterProps struct {
	// `AWS::DocDBElastic::Cluster.AdminUserName`.
	AdminUserName *string `field:"required" json:"adminUserName" yaml:"adminUserName"`
	// `AWS::DocDBElastic::Cluster.AuthType`.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// `AWS::DocDBElastic::Cluster.ClusterName`.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// `AWS::DocDBElastic::Cluster.ShardCapacity`.
	ShardCapacity *float64 `field:"required" json:"shardCapacity" yaml:"shardCapacity"`
	// `AWS::DocDBElastic::Cluster.ShardCount`.
	ShardCount *float64 `field:"required" json:"shardCount" yaml:"shardCount"`
	// `AWS::DocDBElastic::Cluster.AdminUserPassword`.
	AdminUserPassword *string `field:"optional" json:"adminUserPassword" yaml:"adminUserPassword"`
	// `AWS::DocDBElastic::Cluster.KmsKeyId`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// `AWS::DocDBElastic::Cluster.PreferredMaintenanceWindow`.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// `AWS::DocDBElastic::Cluster.SubnetIds`.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// `AWS::DocDBElastic::Cluster.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::DocDBElastic::Cluster.VpcSecurityGroupIds`.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

