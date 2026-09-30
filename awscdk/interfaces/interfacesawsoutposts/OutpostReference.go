package interfacesawsoutposts


// A reference to a Outpost resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outpostReference := &OutpostReference{
//   	OutpostArn: jsii.String("outpostArn"),
//   }
//
type OutpostReference struct {
	// The OutpostArn of the Outpost resource.
	OutpostArn *string `field:"required" json:"outpostArn" yaml:"outpostArn"`
}

