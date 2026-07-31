package awseks


// The default capacity type for the cluster.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv36"
//
//
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_36(),
//   	DefaultCapacityType: eks.DefaultCapacityType_EC2,
//   	KubectlLayer: kubectlv36.NewKubectlV36Layer(this, jsii.String("kubectl")),
//   })
//
type DefaultCapacityType string

const (
	// managed node group.
	DefaultCapacityType_NODEGROUP DefaultCapacityType = "NODEGROUP"
	// EC2 autoscaling group.
	DefaultCapacityType_EC2 DefaultCapacityType = "EC2"
)

