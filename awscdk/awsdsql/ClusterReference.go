package awsdsql


// A reference to a Cluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterReference := &ClusterReference{
//   	Identifier: jsii.String("identifier"),
//   }
//
type ClusterReference struct {
	// The Identifier of the Cluster resource.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

