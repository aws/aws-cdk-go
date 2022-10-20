package awsdocdb

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
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbInstanceClass: jsii.String("dbInstanceClass"),
//
//   	// the properties below are optional
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	enablePerformanceInsights: jsii.Boolean(false),
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
	// The identifier of the cluster that the instance will belong to.
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The compute and memory capacity of the instance;
	//
	// for example, `db.m4.large` . If you change the class of an instance there can be some interruption in the cluster's service.
	DbInstanceClass *string `field:"required" json:"dbInstanceClass" yaml:"dbInstanceClass"`
	// This parameter does not apply to Amazon DocumentDB.
	//
	// Amazon DocumentDB does not perform minor version upgrades regardless of the value set.
	//
	// Default: `false`.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The Amazon EC2 Availability Zone that the instance is created in.
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS Region .
	//
	// Example: `us-east-1d`.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - The first character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `mydbinstance`.
	DbInstanceIdentifier *string `field:"optional" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// `AWS::DocDB::DBInstance.EnablePerformanceInsights`.
	EnablePerformanceInsights interface{} `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region , occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The tags to be assigned to the instance.
	//
	// You can assign up to 10 tags to an instance.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

