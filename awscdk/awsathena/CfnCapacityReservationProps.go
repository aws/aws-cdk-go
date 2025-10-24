package awsathena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCapacityReservation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapacityReservationProps := &CfnCapacityReservationProps{
//   	Name: jsii.String("name"),
//   	TargetDpus: jsii.Number(123),
//
//   	// the properties below are optional
//   	CapacityAssignmentConfiguration: &CapacityAssignmentConfigurationProperty{
//   		CapacityAssignments: []interface{}{
//   			&CapacityAssignmentProperty{
//   				WorkgroupNames: []*string{
//   					jsii.String("workgroupNames"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html
//
type CfnCapacityReservationProps struct {
	// The name of the capacity reservation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html#cfn-athena-capacityreservation-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of data processing units requested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html#cfn-athena-capacityreservation-targetdpus
	//
	TargetDpus *float64 `field:"required" json:"targetDpus" yaml:"targetDpus"`
	// Assigns Athena workgroups (and hence their queries) to capacity reservations.
	//
	// A capacity reservation can have only one capacity assignment configuration, but the capacity assignment configuration can be made up of multiple individual assignments. Each assignment specifies how Athena queries can consume capacity from the capacity reservation that their workgroup is mapped to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html#cfn-athena-capacityreservation-capacityassignmentconfiguration
	//
	CapacityAssignmentConfiguration interface{} `field:"optional" json:"capacityAssignmentConfiguration" yaml:"capacityAssignmentConfiguration"`
	// An array of key-value pairs to apply to the capacity reservation.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html#cfn-athena-capacityreservation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

