package awsservicediscovery


// A reference to a Instance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceReference := &InstanceReference{
//   	InstanceId: jsii.String("instanceId"),
//   }
//
type InstanceReference struct {
	// The InstanceId of the Instance resource.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
}

