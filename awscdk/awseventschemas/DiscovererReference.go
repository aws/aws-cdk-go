package awseventschemas


// A reference to a Discoverer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   discovererReference := &DiscovererReference{
//   	DiscovererArn: jsii.String("discovererArn"),
//   }
//
type DiscovererReference struct {
	// The DiscovererArn of the Discoverer resource.
	DiscovererArn *string `field:"required" json:"discovererArn" yaml:"discovererArn"`
}

