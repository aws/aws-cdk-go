package awsodb


// A reference to a OdbNetwork resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   odbNetworkReference := &OdbNetworkReference{
//   	OdbNetworkArn: jsii.String("odbNetworkArn"),
//   }
//
type OdbNetworkReference struct {
	// The OdbNetworkArn of the OdbNetwork resource.
	OdbNetworkArn *string `field:"required" json:"odbNetworkArn" yaml:"odbNetworkArn"`
}

