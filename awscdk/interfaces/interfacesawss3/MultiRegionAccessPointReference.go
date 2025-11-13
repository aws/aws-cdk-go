package interfacesawss3


// A reference to a MultiRegionAccessPoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiRegionAccessPointReference := &MultiRegionAccessPointReference{
//   	MultiRegionAccessPointName: jsii.String("multiRegionAccessPointName"),
//   }
//
type MultiRegionAccessPointReference struct {
	// The Name of the MultiRegionAccessPoint resource.
	MultiRegionAccessPointName *string `field:"required" json:"multiRegionAccessPointName" yaml:"multiRegionAccessPointName"`
}

