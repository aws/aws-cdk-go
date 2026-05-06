package awssagemaker


// Specifies the network interface configuration for the instance group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterNetworkInterfaceProperty := &ClusterNetworkInterfaceProperty{
//   	InterfaceType: jsii.String("interfaceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusternetworkinterface.html
//
type CfnCluster_ClusterNetworkInterfaceProperty struct {
	// The type of network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clusternetworkinterface.html#cfn-sagemaker-cluster-clusternetworkinterface-interfacetype
	//
	InterfaceType *string `field:"required" json:"interfaceType" yaml:"interfaceType"`
}

