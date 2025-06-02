package awsdax


// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnClusterProps := &CfnClusterProps{
//   	IamRoleArn: jsii.String("iamRoleArn"),
//   	NodeType: jsii.String("nodeType"),
//   	ReplicationFactor: jsii.Number(123),
//
//   	// the properties below are optional
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	ClusterEndpointEncryptionType: jsii.String("clusterEndpointEncryptionType"),
//   	ClusterName: jsii.String("clusterName"),
//   	Description: jsii.String("description"),
//   	NotificationTopicArn: jsii.String("notificationTopicArn"),
//   	ParameterGroupName: jsii.String("parameterGroupName"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SseSpecification: &SSESpecificationProperty{
//   		SseEnabled: jsii.Boolean(false),
//   	},
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html
//
type CfnClusterProps struct {
	// A valid Amazon Resource Name (ARN) that identifies an IAM role.
	//
	// At runtime, DAX will assume this role and use the role's permissions to access DynamoDB on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-iamrolearn
	//
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The node type for the nodes in the cluster.
	//
	// (All nodes in a DAX cluster are of the same type.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-nodetype
	//
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// The number of nodes in the DAX cluster.
	//
	// A replication factor of 1 will create a single-node cluster, without any read replicas. For additional fault tolerance, you can create a multiple node cluster with one or more read replicas. To do this, set `ReplicationFactor` to a number between 3 (one primary and two read replicas) and 10 (one primary and nine read replicas). `If the AvailabilityZones` parameter is provided, its length must equal the `ReplicationFactor` .
	//
	// > AWS recommends that you have at least two read replicas per cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-replicationfactor
	//
	ReplicationFactor *float64 `field:"required" json:"replicationFactor" yaml:"replicationFactor"`
	// The Availability Zones (AZs) in which the cluster nodes will reside after the cluster has been created or updated.
	//
	// If provided, the length of this list must equal the `ReplicationFactor` parameter. If you omit this parameter, DAX will spread the nodes across Availability Zones for the highest availability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-availabilityzones
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The encryption type of the cluster's endpoint. Available values are:.
	//
	// - `NONE` - The cluster's endpoint will be unencrypted.
	// - `TLS` - The cluster's endpoint will be encrypted with Transport Layer Security, and will provide an x509 certificate for authentication.
	//
	// The default value is `NONE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-clusterendpointencryptiontype
	//
	ClusterEndpointEncryptionType *string `field:"optional" json:"clusterEndpointEncryptionType" yaml:"clusterEndpointEncryptionType"`
	// The name of the DAX cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The description of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications will be sent.
	//
	// > The Amazon SNS topic owner must be same as the DAX cluster owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-notificationtopicarn
	//
	NotificationTopicArn *string `field:"optional" json:"notificationTopicArn" yaml:"notificationTopicArn"`
	// The parameter group to be associated with the DAX cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-parametergroupname
	//
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// A range of time when maintenance of DAX cluster software will be performed.
	//
	// For example: `sun:01:00-sun:09:00` . Cluster maintenance normally takes less than 30 minutes, and is performed automatically within the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// A list of security group IDs to be assigned to each node in the DAX cluster.
	//
	// (Each of the security group ID is system-generated.)
	//
	// If this parameter is not specified, DAX assigns the default VPC security group to each node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Represents the settings used to enable server-side encryption on the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-ssespecification
	//
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// The name of the subnet group to be used for the replication group.
	//
	// > DAX clusters can only run in an Amazon VPC environment. All of the subnets that you specify in a subnet group must exist in the same VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-subnetgroupname
	//
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// A set of tags to associate with the DAX cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html#cfn-dax-cluster-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

