package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentProps := &CfnDeploymentProps{
//   	BlueGreenDeploymentName: jsii.String("blueGreenDeploymentName"),
//   	Source: jsii.String("source"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetAllocatedStorage: jsii.Number(123),
//   	TargetDbClusterParameterGroupName: jsii.String("targetDbClusterParameterGroupName"),
//   	TargetDbInstanceClass: jsii.String("targetDbInstanceClass"),
//   	TargetDbParameterGroupName: jsii.String("targetDbParameterGroupName"),
//   	TargetEngineVersion: jsii.String("targetEngineVersion"),
//   	TargetIops: jsii.Number(123),
//   	TargetStorageThroughput: jsii.Number(123),
//   	TargetStorageType: jsii.String("targetStorageType"),
//   	UpgradeTargetStorageConfig: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html
//
type CfnDeploymentProps struct {
	// The name of the Blue/Green Deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-bluegreendeploymentname
	//
	BlueGreenDeploymentName *string `field:"required" json:"blueGreenDeploymentName" yaml:"blueGreenDeploymentName"`
	// The ARN of the source production database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The tags assigned to the Blue/Green Deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The allocated storage size in GiB for the target DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-targetallocatedstorage
	//
	TargetAllocatedStorage *float64 `field:"optional" json:"targetAllocatedStorage" yaml:"targetAllocatedStorage"`
	// The DB cluster parameter group for the target Aurora DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-targetdbclusterparametergroupname
	//
	TargetDbClusterParameterGroupName *string `field:"optional" json:"targetDbClusterParameterGroupName" yaml:"targetDbClusterParameterGroupName"`
	// The DB instance class for target RDS DB instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-targetdbinstanceclass
	//
	TargetDbInstanceClass *string `field:"optional" json:"targetDbInstanceClass" yaml:"targetDbInstanceClass"`
	// The DB parameter group for the target DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-targetdbparametergroupname
	//
	TargetDbParameterGroupName *string `field:"optional" json:"targetDbParameterGroupName" yaml:"targetDbParameterGroupName"`
	// The DB engine version for the target database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-targetengineversion
	//
	TargetEngineVersion *string `field:"optional" json:"targetEngineVersion" yaml:"targetEngineVersion"`
	// The Provisioned IOPS value for the target DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-targetiops
	//
	TargetIops *float64 `field:"optional" json:"targetIops" yaml:"targetIops"`
	// The storage throughput value for the target DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-targetstoragethroughput
	//
	TargetStorageThroughput *float64 `field:"optional" json:"targetStorageThroughput" yaml:"targetStorageThroughput"`
	// The storage type for the target DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-targetstoragetype
	//
	TargetStorageType *string `field:"optional" json:"targetStorageType" yaml:"targetStorageType"`
	// Whether to upgrade the storage file system configuration on the target database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-deployment.html#cfn-rds-deployment-upgradetargetstorageconfig
	//
	UpgradeTargetStorageConfig interface{} `field:"optional" json:"upgradeTargetStorageConfig" yaml:"upgradeTargetStorageConfig"`
}

