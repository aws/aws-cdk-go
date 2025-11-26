package previewawsneptunemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDBInstancePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBInstanceMixinProps := &CfnDBInstanceMixinProps{
//   	AllowMajorVersionUpgrade: jsii.Boolean(false),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbInstanceClass: jsii.String("dbInstanceClass"),
//   	DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	DbParameterGroupName: jsii.String("dbParameterGroupName"),
//   	DbSnapshotIdentifier: jsii.String("dbSnapshotIdentifier"),
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html
//
type CfnDBInstanceMixinProps struct {
	// Indicates that major version upgrades are allowed.
	//
	// Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible. This parameter must be set to true when specifying a value for the EngineVersion parameter that is a different major version than the DB instance's current version.
	//
	// When you change this parameter for an existing DB cluster, CloudFormation will replace your existing DB cluster with a new, empty one that uses the engine version you specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-allowmajorversionupgrade
	//
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Indicates that minor version patches are applied automatically.
	//
	// When updating this property, some interruptions may occur.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-autominorversionupgrade
	//
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Specifies the name of the Availability Zone the DB instance is located in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-dbclusteridentifier
	//
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Contains the name of the compute and memory capacity class of the DB instance.
	//
	// If you update this property, some interruptions may occur.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-dbinstanceclass
	//
	DbInstanceClass *string `field:"optional" json:"dbInstanceClass" yaml:"dbInstanceClass"`
	// Contains a user-supplied database identifier.
	//
	// This identifier is the unique key that identifies a DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-dbinstanceidentifier
	//
	DbInstanceIdentifier *string `field:"optional" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The name of an existing DB parameter group or a reference to an AWS::Neptune::DBParameterGroup resource created in the template.
	//
	// If any of the data members of the referenced parameter group are changed during an update, the DB instance might need to be restarted, which causes some interruption. If the parameter group contains static parameters, whether they were changed or not, an update triggers a reboot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-dbparametergroupname
	//
	DbParameterGroupName *string `field:"optional" json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
	// This parameter is not supported.
	//
	// `AWS::Neptune::DBInstance` does not support restoring from snapshots.
	//
	// `AWS::Neptune::DBCluster` does support restoring from snapshots.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-dbsnapshotidentifier
	//
	// Deprecated: this property has been deprecated.
	DbSnapshotIdentifier *string `field:"optional" json:"dbSnapshotIdentifier" yaml:"dbSnapshotIdentifier"`
	// A DB subnet group to associate with the DB instance.
	//
	// If you update this value, the new subnet group must be a subnet group in a new virtual private cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-dbsubnetgroupname
	//
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Indicates whether the DB instance is publicly accessible.
	//
	// When the DB instance is publicly accessible and you connect from outside of the DB instance's virtual private cloud (VPC), its Domain Name System (DNS) endpoint resolves to the public IP address. When you connect from within the same VPC as the DB instance, the endpoint resolves to the private IP address. Access to the DB instance is ultimately controlled by the security group it uses. That public access isn't permitted if the security group assigned to the DB cluster doesn't permit it.
	//
	// When the DB instance isn't publicly accessible, it is an internal DB instance with a DNS name that resolves to a private IP address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-publiclyaccessible
	//
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An arbitrary set of tags (key-value pairs) for this DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html#cfn-neptune-dbinstance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

