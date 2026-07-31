package awseks


// Control plane scaling tier for EKS Provisioned Control Plane.
//
// Provisioned Control Plane allows cluster administrators to select from a set
// of scaling tiers to ensure high and predictable performance for demanding workloads
// such as AI training/inference, high-performance computing, or large-scale data processing.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv35"
//
//
//   eks.NewCluster(this, jsii.String("HighPerformanceCluster"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_35(),
//   	KubectlLayer: kubectlv35.NewKubectlV35Layer(this, jsii.String("kubectl")),
//   	ControlPlaneScalingTier: eks.ControlPlaneScalingTier_TIER_XL,
//   })
//
// See: https://docs.aws.amazon.com/eks/latest/userguide/eks-provisioned-control-plane.html
//
type ControlPlaneScalingTier string

const (
	// Standard control plane (default, no additional cost).
	ControlPlaneScalingTier_STANDARD ControlPlaneScalingTier = "STANDARD"
	// Extra-large provisioned tier.
	ControlPlaneScalingTier_TIER_XL ControlPlaneScalingTier = "TIER_XL"
	// 2x extra-large provisioned tier.
	ControlPlaneScalingTier_TIER_2XL ControlPlaneScalingTier = "TIER_2XL"
	// 4x extra-large provisioned tier.
	ControlPlaneScalingTier_TIER_4XL ControlPlaneScalingTier = "TIER_4XL"
	// 8x extra-large provisioned tier.
	ControlPlaneScalingTier_TIER_8XL ControlPlaneScalingTier = "TIER_8XL"
)

