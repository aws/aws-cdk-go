package awsredshift


// A reference to a ClusterSubnetGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterSubnetGroupReference := &ClusterSubnetGroupReference{
//   	ClusterSubnetGroupName: jsii.String("clusterSubnetGroupName"),
//   }
//
type ClusterSubnetGroupReference struct {
	// The ClusterSubnetGroupName of the ClusterSubnetGroup resource.
	ClusterSubnetGroupName *string `field:"required" json:"clusterSubnetGroupName" yaml:"clusterSubnetGroupName"`
}

