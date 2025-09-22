package awss3


// A reference to a MultiRegionAccessPointPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiRegionAccessPointPolicyReference := &MultiRegionAccessPointPolicyReference{
//   	MrapName: jsii.String("mrapName"),
//   }
//
type MultiRegionAccessPointPolicyReference struct {
	// The MrapName of the MultiRegionAccessPointPolicy resource.
	MrapName *string `field:"required" json:"mrapName" yaml:"mrapName"`
}

