package awsodb


// A reference to a CloudVmCluster resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudVmClusterReference := &CloudVmClusterReference{
//   	CloudVmClusterArn: jsii.String("cloudVmClusterArn"),
//   }
//
type CloudVmClusterReference struct {
	// The CloudVmClusterArn of the CloudVmCluster resource.
	CloudVmClusterArn *string `field:"required" json:"cloudVmClusterArn" yaml:"cloudVmClusterArn"`
}

