package awsdocdb


// A reference to a DBCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBClusterReference := &DBClusterReference{
//   	DbClusterId: jsii.String("dbClusterId"),
//   }
//
type DBClusterReference struct {
	// The Id of the DBCluster resource.
	DbClusterId *string `field:"required" json:"dbClusterId" yaml:"dbClusterId"`
}

