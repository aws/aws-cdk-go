package awseksv2


// The configuration for the Kubernetes API server on an Amazon EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kubeApiServerConfigProperty := &KubeApiServerConfigProperty{
//   	EventTtl: jsii.String("eventTtl"),
//   	ServiceNodePortRange: &ServiceNodePortRangeProperty{
//   		MaxPort: jsii.Number(123),
//   		MinPort: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubeapiserverconfig.html
//
type CfnCluster_KubeApiServerConfigProperty struct {
	// The duration that Kubernetes events are retained (e.g., 30m, 1h).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubeapiserverconfig.html#cfn-eks-cluster-kubeapiserverconfig-eventttl
	//
	EventTtl *string `field:"optional" json:"eventTtl" yaml:"eventTtl"`
	// The port range for Kubernetes NodePort services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-kubeapiserverconfig.html#cfn-eks-cluster-kubeapiserverconfig-servicenodeportrange
	//
	ServiceNodePortRange interface{} `field:"optional" json:"serviceNodePortRange" yaml:"serviceNodePortRange"`
}

