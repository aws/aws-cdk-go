package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Policy for customizing shrink operations.
//
// Allows configuration of decommissioning timeout and targeted instance shrinking.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   shrinkPolicyProperty := &ShrinkPolicyProperty{
//   	DecommissionTimeout: duration,
//   	InstanceResizePolicy: &InstanceResizePolicyProperty{
//   		InstancesToProtect: []*string{
//   			jsii.String("instancesToProtect"),
//   		},
//   		InstancesToTerminate: []*string{
//   			jsii.String("instancesToTerminate"),
//   		},
//   		InstanceTerminationTimeout: duration,
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ShrinkPolicy.html
//
// Experimental.
type EmrModifyInstanceGroupByName_ShrinkPolicyProperty struct {
	// The desired timeout for decommissioning an instance.
	//
	// Overrides the default YARN decommissioning timeout.
	// Experimental.
	DecommissionTimeout awscdk.Duration `field:"optional" json:"decommissionTimeout" yaml:"decommissionTimeout"`
	// Custom policy for requesting termination protection or termination of specific instances when shrinking an instance group.
	// Experimental.
	InstanceResizePolicy *EmrModifyInstanceGroupByName_InstanceResizePolicyProperty `field:"optional" json:"instanceResizePolicy" yaml:"instanceResizePolicy"`
}

