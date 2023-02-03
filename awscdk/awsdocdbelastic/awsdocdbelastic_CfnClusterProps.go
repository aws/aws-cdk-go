package awsdocdbelastic

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
	// The name of the Amazon DocumentDB elastic clusters administrator.
	//
	// *Constraints* :
	//
	// - Must be from 1 to 63 letters or numbers.
	// - The first character must be a letter.
	// - Cannot be a reserved word.
	AdminUserName *string `field:"required" json:"adminUserName" yaml:"adminUserName"`
	// The authentication type used to determine where to fetch the password used for accessing the elastic cluster.
	//
	// Valid types are `PLAIN_TEXT` or `SECRET_ARN` .
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// The name of the new elastic cluster. This parameter is stored as a lowercase string.
	//
	// *Constraints* :
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - The first character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// *Example* : `my-cluster`.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The number of vCPUs assigned to each elastic cluster shard.
	//
	// Maximum is 64. Allowed values are 2, 4, 8, 16, 32, 64.
	ShardCapacity *float64 `field:"required" json:"shardCapacity" yaml:"shardCapacity"`
	// The number of shards assigned to the elastic cluster.
	//
	// Maximum is 32.
	ShardCount *float64 `field:"required" json:"shardCount" yaml:"shardCount"`
	// The password for the Elastic DocumentDB cluster administrator and can contain any printable ASCII characters.
	//
	// *Constraints* :
	//
	// - Must contain from 8 to 100 characters.
	// - Cannot contain a forward slash (/), double quote ("), or the "at" symbol (@).
	// - A valid `AdminUserName` entry is also required.
	AdminUserPassword *string `field:"optional" json:"adminUserPassword" yaml:"adminUserPassword"`
	// The KMS key identifier to use to encrypt the new elastic cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption key. If you are creating a cluster using the same Amazon account that owns this KMS encryption key, you can use the KMS key alias instead of the ARN as the KMS encryption key.
	//
	// If an encryption key is not specified, Amazon DocumentDB uses the default encryption key that KMS creates for your account. Your account has a different default encryption key for each Amazon Region.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// *Format* : `ddd:hh24:mi-ddd:hh24:mi`
	//
	// *Default* : a 30-minute window selected at random from an 8-hour block of time for each AWS Region , occurring on a random day of the week.
	//
	// *Valid days* : Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// *Constraints* : Minimum 30-minute window.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The Amazon EC2 subnet IDs for the new elastic cluster.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The tags to be assigned to the new elastic cluster.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A list of EC2 VPC security groups to associate with the new elastic cluster.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

