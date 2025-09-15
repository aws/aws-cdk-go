package awsemr


// A reference to a InstanceGroupConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceGroupConfigReference := &InstanceGroupConfigReference{
//   	InstanceGroupConfigId: jsii.String("instanceGroupConfigId"),
//   }
//
type InstanceGroupConfigReference struct {
	// The Id of the InstanceGroupConfig resource.
	InstanceGroupConfigId *string `field:"required" json:"instanceGroupConfigId" yaml:"instanceGroupConfigId"`
}

