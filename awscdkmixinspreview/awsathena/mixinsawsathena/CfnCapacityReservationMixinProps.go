package mixinsawsathena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCapacityReservationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapacityReservationMixinProps := &CfnCapacityReservationMixinProps{
//   	CapacityAssignmentConfiguration: &CapacityAssignmentConfigurationProperty{
//   		CapacityAssignments: []interface{}{
//   			&CapacityAssignmentProperty{
//   				WorkgroupNames: []*string{
//   					jsii.String("workgroupNames"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetDpus: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html
//
type CfnCapacityReservationMixinProps struct {
	// Assigns Athena workgroups (and hence their queries) to capacity reservations.
	//
	// A capacity reservation can have only one capacity assignment configuration, but the capacity assignment configuration can be made up of multiple individual assignments. Each assignment specifies how Athena queries can consume capacity from the capacity reservation that their workgroup is mapped to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html#cfn-athena-capacityreservation-capacityassignmentconfiguration
	//
	CapacityAssignmentConfiguration interface{} `field:"optional" json:"capacityAssignmentConfiguration" yaml:"capacityAssignmentConfiguration"`
	// The name of the capacity reservation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html#cfn-athena-capacityreservation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to the capacity reservation.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html#cfn-athena-capacityreservation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The number of data processing units requested.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html#cfn-athena-capacityreservation-targetdpus
	//
	TargetDpus *float64 `field:"optional" json:"targetDpus" yaml:"targetDpus"`
}

