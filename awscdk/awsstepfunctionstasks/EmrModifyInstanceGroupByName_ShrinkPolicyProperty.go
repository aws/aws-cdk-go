package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Policy for customizing shrink operations.
//
// Allows configuration of decommissioning timeout and targeted instance shrinking.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shrinkPolicyProperty := &ShrinkPolicyProperty{
//   	DecommissionTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	InstanceResizePolicy: &InstanceResizePolicyProperty{
//   		InstancesToProtect: []*string{
//   			jsii.String("instancesToProtect"),
//   		},
//   		InstancesToTerminate: []*string{
//   			jsii.String("instancesToTerminate"),
//   		},
//   		InstanceTerminationTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ShrinkPolicy.html
//
type EmrModifyInstanceGroupByName_ShrinkPolicyProperty struct {
	// The desired timeout for decommissioning an instance.
	//
	// Overrides the default YARN decommissioning timeout.
	// Default: - EMR selected default.
	//
	DecommissionTimeout awscdk.Duration `field:"optional" json:"decommissionTimeout" yaml:"decommissionTimeout"`
	// Custom policy for requesting termination protection or termination of specific instances when shrinking an instance group.
	// Default: - None.
	//
	InstanceResizePolicy *EmrModifyInstanceGroupByName_InstanceResizePolicyProperty `field:"optional" json:"instanceResizePolicy" yaml:"instanceResizePolicy"`
}

