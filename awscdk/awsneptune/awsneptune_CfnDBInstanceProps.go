package awsneptune

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDBInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBInstanceProps := &cfnDBInstanceProps{
//   	dbInstanceClass: jsii.String("dbInstanceClass"),
//
//   	// the properties below are optional
//   	allowMajorVersionUpgrade: jsii.Boolean(false),
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	dbParameterGroupName: jsii.String("dbParameterGroupName"),
//   	dbSnapshotIdentifier: jsii.String("dbSnapshotIdentifier"),
//   	dbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDBInstanceProps struct {
	// Contains the name of the compute and memory capacity class of the DB instance.
	//
	// If you update this property, some interruptions may occur.
	DbInstanceClass *string `field:"required" json:"dbInstanceClass" yaml:"dbInstanceClass"`
	// Indicates that major version upgrades are allowed.
	//
	// Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible. This parameter must be set to true when specifying a value for the EngineVersion parameter that is a different major version than the DB instance's current version.
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Indicates that minor version patches are applied automatically.
	//
	// When updating this property, some interruptions may occur.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Specifies the name of the Availability Zone the DB instance is located in.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of.
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Contains a user-supplied database identifier.
	//
	// This identifier is the unique key that identifies a DB instance.
	DbInstanceIdentifier *string `field:"optional" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The name of an existing DB parameter group or a reference to an AWS::Neptune::DBParameterGroup resource created in the template.
	//
	// If any of the data members of the referenced parameter group are changed during an update, the DB instance might need to be restarted, which causes some interruption. If the parameter group contains static parameters, whether they were changed or not, an update triggers a reboot.
	DbParameterGroupName *string `field:"optional" json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
	// This parameter is not supported.
	//
	// `AWS::Neptune::DBInstance` does not support restoring from snapshots.
	//
	// `AWS::Neptune::DBCluster` does support restoring from snapshots.
	DbSnapshotIdentifier *string `field:"optional" json:"dbSnapshotIdentifier" yaml:"dbSnapshotIdentifier"`
	// A DB subnet group to associate with the DB instance.
	//
	// If you update this value, the new subnet group must be a subnet group in a new virtual private cloud (VPC).
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// An arbitrary set of tags (key-value pairs) for this DB instance.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

