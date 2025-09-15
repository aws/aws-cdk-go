package awspinpointemail


// A reference to a DedicatedIpPool resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dedicatedIpPoolReference := &DedicatedIpPoolReference{
//   	DedicatedIpPoolId: jsii.String("dedicatedIpPoolId"),
//   }
//
type DedicatedIpPoolReference struct {
	// The Id of the DedicatedIpPool resource.
	DedicatedIpPoolId *string `field:"required" json:"dedicatedIpPoolId" yaml:"dedicatedIpPoolId"`
}

