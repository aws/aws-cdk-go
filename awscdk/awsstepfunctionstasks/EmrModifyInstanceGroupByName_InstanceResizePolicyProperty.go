package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Custom policy for requesting termination protection or termination of specific instances when shrinking an instance group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceResizePolicyProperty := &InstanceResizePolicyProperty{
//   	InstancesToProtect: []*string{
//   		jsii.String("instancesToProtect"),
//   	},
//   	InstancesToTerminate: []*string{
//   		jsii.String("instancesToTerminate"),
//   	},
//   	InstanceTerminationTimeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceResizePolicy.html
//
type EmrModifyInstanceGroupByName_InstanceResizePolicyProperty struct {
	// Specific list of instances to be protected when shrinking an instance group.
	// Default: - No instances will be protected when shrinking an instance group.
	//
	InstancesToProtect *[]*string `field:"optional" json:"instancesToProtect" yaml:"instancesToProtect"`
	// Specific list of instances to be terminated when shrinking an instance group.
	// Default: - No instances will be terminated when shrinking an instance group.
	//
	InstancesToTerminate *[]*string `field:"optional" json:"instancesToTerminate" yaml:"instancesToTerminate"`
	// Decommissioning timeout override for the specific list of instances to be terminated.
	// Default: cdk.Duration.seconds
	//
	InstanceTerminationTimeout awscdk.Duration `field:"optional" json:"instanceTerminationTimeout" yaml:"instanceTerminationTimeout"`
}

