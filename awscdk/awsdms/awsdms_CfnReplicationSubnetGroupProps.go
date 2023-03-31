package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnReplicationSubnetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationSubnetGroupProps := &cfnReplicationSubnetGroupProps{
//   	replicationSubnetGroupDescription: jsii.String("replicationSubnetGroupDescription"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	replicationSubnetGroupIdentifier: jsii.String("replicationSubnetGroupIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnReplicationSubnetGroupProps struct {
	// The description for the subnet group.
	ReplicationSubnetGroupDescription *string `field:"required" json:"replicationSubnetGroupDescription" yaml:"replicationSubnetGroupDescription"`
	// One or more subnet IDs to be assigned to the subnet group.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The identifier for the replication subnet group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the identifier.
	ReplicationSubnetGroupIdentifier *string `field:"optional" json:"replicationSubnetGroupIdentifier" yaml:"replicationSubnetGroupIdentifier"`
	// One or more tags to be assigned to the subnet group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

