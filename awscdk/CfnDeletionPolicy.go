// Version 2 of the AWS Cloud Development Kit library
package awscdk


// With the DeletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted.
//
// You specify a DeletionPolicy attribute for each resource that you want to control. If a resource has no DeletionPolicy
// attribute, AWS CloudFormation deletes the resource by default. Note that this capability also applies to update operations
// that lead to resources being removed.
type CfnDeletionPolicy string

const (
	// AWS CloudFormation deletes the resource and all its content if applicable during stack deletion.
	//
	// You can add this
	// deletion policy to any resource type. By default, if you don't specify a DeletionPolicy, AWS CloudFormation deletes
	// your resources. However, be aware of the following considerations:
	CfnDeletionPolicy_DELETE CfnDeletionPolicy = "DELETE"
	// AWS CloudFormation keeps the resource without deleting the resource or its contents when its stack is deleted.
	//
	// You can add this deletion policy to any resource type. Note that when AWS CloudFormation completes the stack deletion,
	// the stack will be in Delete_Complete state; however, resources that are retained continue to exist and continue to incur
	// applicable charges until you delete those resources.
	CfnDeletionPolicy_RETAIN CfnDeletionPolicy = "RETAIN"
	// For resources that support snapshots (AWS::EC2::Volume, AWS::ElastiCache::CacheCluster, AWS::ElastiCache::ReplicationGroup, AWS::RDS::DBInstance, AWS::RDS::DBCluster, and AWS::Redshift::Cluster), AWS CloudFormation creates a snapshot for the resource before deleting it.
	//
	// Note that when AWS CloudFormation completes the stack deletion, the stack will be in the
	// Delete_Complete state; however, the snapshots that are created with this policy continue to exist and continue to
	// incur applicable charges until you delete those snapshots.
	CfnDeletionPolicy_SNAPSHOT CfnDeletionPolicy = "SNAPSHOT"
)

