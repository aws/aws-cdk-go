package awseks


// The default capacity type for the cluster.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv34"
//
//
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_34(),
//   	DefaultCapacityType: eks.DefaultCapacityType_EC2,
//   	KubectlLayer: kubectlv34.NewKubectlV34Layer(this, jsii.String("kubectl")),
//   })
//
type DefaultCapacityType string

const (
	// managed node group.
	DefaultCapacityType_NODEGROUP DefaultCapacityType = "NODEGROUP"
	// EC2 autoscaling group.
	DefaultCapacityType_EC2 DefaultCapacityType = "EC2"
)

