package interfacesawsmediapackage


// A reference to a HarvestJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harvestJobReference := &HarvestJobReference{
//   	HarvestJobArn: jsii.String("harvestJobArn"),
//   }
//
type HarvestJobReference struct {
	// The Arn of the HarvestJob resource.
	HarvestJobArn *string `field:"required" json:"harvestJobArn" yaml:"harvestJobArn"`
}

