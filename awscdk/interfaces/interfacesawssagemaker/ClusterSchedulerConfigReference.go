package interfacesawssagemaker


// A reference to a ClusterSchedulerConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterSchedulerConfigReference := &ClusterSchedulerConfigReference{
//   	ClusterSchedulerConfigArn: jsii.String("clusterSchedulerConfigArn"),
//   }
//
type ClusterSchedulerConfigReference struct {
	// The ClusterSchedulerConfigArn of the ClusterSchedulerConfig resource.
	ClusterSchedulerConfigArn *string `field:"required" json:"clusterSchedulerConfigArn" yaml:"clusterSchedulerConfigArn"`
}

