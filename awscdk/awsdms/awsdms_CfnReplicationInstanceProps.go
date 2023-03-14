package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnReplicationInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationInstanceProps := &cfnReplicationInstanceProps{
//   	replicationInstanceClass: jsii.String("replicationInstanceClass"),
//
//   	// the properties below are optional
//   	allocatedStorage: jsii.Number(123),
//   	allowMajorVersionUpgrade: jsii.Boolean(false),
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	engineVersion: jsii.String("engineVersion"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	multiAz: jsii.Boolean(false),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	replicationInstanceIdentifier: jsii.String("replicationInstanceIdentifier"),
//   	replicationSubnetGroupIdentifier: jsii.String("replicationSubnetGroupIdentifier"),
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
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
type CfnReplicationInstanceProps struct {
	// The compute and memory capacity of the replication instance as defined for the specified replication instance class.
	//
	// For example, to specify the instance class dms.c4.large, set this parameter to `"dms.c4.large"` . For more information on the settings and capacities for the available replication instance classes, see [Selecting the right AWS DMS replication instance for your migration](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReplicationInstance.html#CHAP_ReplicationInstance.InDepth) in the *AWS Database Migration Service User Guide* .
	ReplicationInstanceClass *string `field:"required" json:"replicationInstanceClass" yaml:"replicationInstanceClass"`
	// The amount of storage (in gigabytes) to be initially allocated for the replication instance.
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Indicates that major version upgrades are allowed.
	//
	// Changing this parameter does not result in an outage, and the change is asynchronously applied as soon as possible.
	//
	// This parameter must be set to `true` when specifying a value for the `EngineVersion` parameter that is a different major version than the replication instance's current version.
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// A value that indicates whether minor engine upgrades are applied automatically to the replication instance during the maintenance window.
	//
	// This parameter defaults to `true` .
	//
	// Default: `true`.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The Availability Zone that the replication instance will be created in.
	//
	// The default value is a random, system-chosen Availability Zone in the endpoint's AWS Region , for example `us-east-1d` .
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The engine version number of the replication instance.
	//
	// If an engine version number is not specified when a replication instance is created, the default is the latest engine version available.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// An AWS KMS key identifier that is used to encrypt the data on the replication instance.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Region .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies whether the replication instance is a Multi-AZ deployment.
	//
	// You can't set the `AvailabilityZone` parameter if the Multi-AZ parameter is set to `true` .
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// The weekly time range during which system maintenance can occur, in UTC.
	//
	// *Format* : `ddd:hh24:mi-ddd:hh24:mi`
	//
	// *Default* : A 30-minute window selected at random from an 8-hour block of time per AWS Region , occurring on a random day of the week.
	//
	// *Valid days* ( `ddd` ): `Mon` | `Tue` | `Wed` | `Thu` | `Fri` | `Sat` | `Sun`
	//
	// *Constraints* : Minimum 30-minute window.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Specifies the accessibility options for the replication instance.
	//
	// A value of `true` represents an instance with a public IP address. A value of `false` represents an instance with a private IP address. The default value is `true` .
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The replication instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain 1-63 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `myrepinstance`.
	ReplicationInstanceIdentifier *string `field:"optional" json:"replicationInstanceIdentifier" yaml:"replicationInstanceIdentifier"`
	// A subnet group to associate with the replication instance.
	ReplicationSubnetGroupIdentifier *string `field:"optional" json:"replicationSubnetGroupIdentifier" yaml:"replicationSubnetGroupIdentifier"`
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` . For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// One or more tags to be assigned to the replication instance.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the virtual private cloud (VPC) security group to be used with the replication instance.
	//
	// The VPC security group must work with the VPC containing the replication instance.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

