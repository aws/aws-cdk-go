package interfacesawsemr


// A reference to a InstanceGroupConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceGroupConfigReference := &InstanceGroupConfigReference{
//   	InstanceGroupId: jsii.String("instanceGroupId"),
//   	JobFlowId: jsii.String("jobFlowId"),
//   }
//
type InstanceGroupConfigReference struct {
	// The InstanceGroupId of the InstanceGroupConfig resource.
	InstanceGroupId *string `field:"required" json:"instanceGroupId" yaml:"instanceGroupId"`
	// The JobFlowId of the InstanceGroupConfig resource.
	JobFlowId *string `field:"required" json:"jobFlowId" yaml:"jobFlowId"`
}

