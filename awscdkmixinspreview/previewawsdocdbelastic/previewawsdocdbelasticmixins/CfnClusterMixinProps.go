package previewawsdocdbelasticmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnClusterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterMixinProps := &CfnClusterMixinProps{
//   	AdminUserName: jsii.String("adminUserName"),
//   	AdminUserPassword: jsii.String("adminUserPassword"),
//   	AuthType: jsii.String("authType"),
//   	BackupRetentionPeriod: jsii.Number(123),
//   	ClusterName: jsii.String("clusterName"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	PreferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	ShardCapacity: jsii.Number(123),
//   	ShardCount: jsii.Number(123),
//   	ShardInstanceCount: jsii.Number(123),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html
//
type CfnClusterMixinProps struct {
	// The name of the Amazon DocumentDB elastic clusters administrator.
	//
	// *Constraints* :
	//
	// - Must be from 1 to 63 letters or numbers.
	// - The first character must be a letter.
	// - Cannot be a reserved word.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-adminusername
	//
	AdminUserName *string `field:"optional" json:"adminUserName" yaml:"adminUserName"`
	// The password for the Elastic DocumentDB cluster administrator and can contain any printable ASCII characters.
	//
	// *Constraints* :
	//
	// - Must contain from 8 to 100 characters.
	// - Cannot contain a forward slash (/), double quote ("), or the "at" symbol (@).
	// - A valid `AdminUserName` entry is also required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-adminuserpassword
	//
	AdminUserPassword *string `field:"optional" json:"adminUserPassword" yaml:"adminUserPassword"`
	// The authentication type used to determine where to fetch the password used for accessing the elastic cluster.
	//
	// Valid types are `PLAIN_TEXT` or `SECRET_ARN` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-authtype
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The number of days for which automatic snapshots are retained.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-backupretentionperiod
	//
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// The name of the new elastic cluster. This parameter is stored as a lowercase string.
	//
	// *Constraints* :
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - The first character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// *Example* : `my-cluster`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The KMS key identifier to use to encrypt the new elastic cluster.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption key. If you are creating a cluster using the same Amazon account that owns this KMS encryption key, you can use the KMS key alias instead of the ARN as the KMS encryption key.
	//
	// If an encryption key is not specified, Amazon DocumentDB uses the default encryption key that KMS creates for your account. Your account has a different default encryption key for each Amazon Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The daily time range during which automated backups are created if automated backups are enabled, as determined by `backupRetentionPeriod` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-preferredbackupwindow
	//
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// *Format* : `ddd:hh24:mi-ddd:hh24:mi`
	//
	// *Default* : a 30-minute window selected at random from an 8-hour block of time for each AWS Region , occurring on a random day of the week.
	//
	// *Valid days* : Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// *Constraints* : Minimum 30-minute window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The number of vCPUs assigned to each elastic cluster shard.
	//
	// Maximum is 64. Allowed values are 2, 4, 8, 16, 32, 64.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-shardcapacity
	//
	ShardCapacity *float64 `field:"optional" json:"shardCapacity" yaml:"shardCapacity"`
	// The number of shards assigned to the elastic cluster.
	//
	// Maximum is 32.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-shardcount
	//
	ShardCount *float64 `field:"optional" json:"shardCount" yaml:"shardCount"`
	// The number of replica instances applying to all shards in the cluster.
	//
	// A `shardInstanceCount` value of 1 means there is one writer instance, and any additional instances are replicas that can be used for reads and to improve availability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-shardinstancecount
	//
	ShardInstanceCount *float64 `field:"optional" json:"shardInstanceCount" yaml:"shardInstanceCount"`
	// The Amazon EC2 subnet IDs for the new elastic cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The tags to be assigned to the new elastic cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A list of EC2 VPC security groups to associate with the new elastic cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdbelastic-cluster.html#cfn-docdbelastic-cluster-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

