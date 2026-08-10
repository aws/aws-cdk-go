package interfacesawscloudformation


// A reference to a ResourceScan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceScanReference := &ResourceScanReference{
//   	ResourceScanId: jsii.String("resourceScanId"),
//   }
//
type ResourceScanReference struct {
	// The ResourceScanId of the ResourceScan resource.
	ResourceScanId *string `field:"required" json:"resourceScanId" yaml:"resourceScanId"`
}

