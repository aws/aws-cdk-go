package awsneptune

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Construction properties for a DatabaseInstanceNew.
//
// Example:
//   replica1 := neptune.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
//   	Cluster: Cluster,
//   	InstanceType: neptune.InstanceType_R5_LARGE(),
//   })
//
// Experimental.
type DatabaseInstanceProps struct {
	// The Neptune database cluster the instance should launch into.
	// Experimental.
	Cluster IDatabaseCluster `field:"required" json:"cluster" yaml:"cluster"`
	// What type of instance to start for the replicas.
	// Experimental.
	InstanceType InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// The name of the Availability Zone where the DB instance will be located.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	// Experimental.
	DbInstanceName *string `field:"optional" json:"dbInstanceName" yaml:"dbInstanceName"`
	// The DB parameter group to associate with the instance.
	// Experimental.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

