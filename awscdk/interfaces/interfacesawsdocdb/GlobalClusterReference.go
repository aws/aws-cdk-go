package interfacesawsdocdb


// A reference to a GlobalCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalClusterReference := &GlobalClusterReference{
//   	GlobalClusterArn: jsii.String("globalClusterArn"),
//   	GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   }
//
type GlobalClusterReference struct {
	// The ARN of the GlobalCluster resource.
	GlobalClusterArn *string `field:"required" json:"globalClusterArn" yaml:"globalClusterArn"`
	// The GlobalClusterIdentifier of the GlobalCluster resource.
	GlobalClusterIdentifier *string `field:"required" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
}

