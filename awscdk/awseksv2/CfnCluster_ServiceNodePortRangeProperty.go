package awseksv2


// The port range for Kubernetes NodePort services.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNodePortRangeProperty := &ServiceNodePortRangeProperty{
//   	MaxPort: jsii.Number(123),
//   	MinPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-servicenodeportrange.html
//
type CfnCluster_ServiceNodePortRangeProperty struct {
	// The maximum port number in the range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-servicenodeportrange.html#cfn-eks-cluster-servicenodeportrange-maxport
	//
	MaxPort *float64 `field:"optional" json:"maxPort" yaml:"maxPort"`
	// The minimum port number in the range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-servicenodeportrange.html#cfn-eks-cluster-servicenodeportrange-minport
	//
	MinPort *float64 `field:"optional" json:"minPort" yaml:"minPort"`
}

