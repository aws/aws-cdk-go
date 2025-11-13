package interfacesawsses


// A reference to a DedicatedIpPool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dedicatedIpPoolReference := &DedicatedIpPoolReference{
//   	PoolName: jsii.String("poolName"),
//   }
//
type DedicatedIpPoolReference struct {
	// The PoolName of the DedicatedIpPool resource.
	PoolName *string `field:"required" json:"poolName" yaml:"poolName"`
}

