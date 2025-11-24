package mixinsawsathena


// Assigns Athena workgroups (and hence their queries) to capacity reservations.
//
// A capacity reservation can have only one capacity assignment configuration, but the capacity assignment configuration can be made up of multiple individual assignments. Each assignment specifies how Athena queries can consume capacity from the capacity reservation that their workgroup is mapped to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   capacityAssignmentConfigurationProperty := &CapacityAssignmentConfigurationProperty{
//   	CapacityAssignments: []interface{}{
//   		&CapacityAssignmentProperty{
//   			WorkgroupNames: []*string{
//   				jsii.String("workgroupNames"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-capacityreservation-capacityassignmentconfiguration.html
//
type CfnCapacityReservationPropsMixin_CapacityAssignmentConfigurationProperty struct {
	// The list of assignments that make up the capacity assignment configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-capacityreservation-capacityassignmentconfiguration.html#cfn-athena-capacityreservation-capacityassignmentconfiguration-capacityassignments
	//
	CapacityAssignments interface{} `field:"optional" json:"capacityAssignments" yaml:"capacityAssignments"`
}

