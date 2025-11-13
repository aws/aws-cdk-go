package interfacesawsneptune


// A reference to a DBCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBClusterReference := &DBClusterReference{
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   }
//
type DBClusterReference struct {
	// The DBClusterIdentifier of the DBCluster resource.
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
}

