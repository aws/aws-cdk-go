package awseksv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The scaling tier for the EKS cluster provisioned control plane.
//
// Amazon EKS Provisioned Control Plane lets you select a scaling tier to ensure high and
// predictable control plane performance for demanding workloads such as AI training/inference,
// high-performance computing, or large-scale data processing.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HighPerformanceCluster"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_36(),
//   	ControlPlaneScalingTier: eks.ControlPlaneScalingTier_TIER_XL(),
//   })
//
// See: https://docs.aws.amazon.com/eks/latest/userguide/eks-provisioned-control-plane.html
//
type ControlPlaneScalingTier interface {
	// The string value of the scaling tier as expected by the EKS API.
	Value() *string
}

// The jsii proxy struct for ControlPlaneScalingTier
type jsiiProxy_ControlPlaneScalingTier struct {
	_ byte // padding
}

func (j *jsiiProxy_ControlPlaneScalingTier) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A custom scaling tier, for values not yet available as a static member.
func ControlPlaneScalingTier_Of(tier *string) ControlPlaneScalingTier {
	_init_.Initialize()

	if err := validateControlPlaneScalingTier_OfParameters(tier); err != nil {
		panic(err)
	}
	var returns ControlPlaneScalingTier

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks_v2.ControlPlaneScalingTier",
		"of",
		[]interface{}{tier},
		&returns,
	)

	return returns
}

func ControlPlaneScalingTier_STANDARD() ControlPlaneScalingTier {
	_init_.Initialize()
	var returns ControlPlaneScalingTier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks_v2.ControlPlaneScalingTier",
		"STANDARD",
		&returns,
	)
	return returns
}

func ControlPlaneScalingTier_TIER_2XL() ControlPlaneScalingTier {
	_init_.Initialize()
	var returns ControlPlaneScalingTier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks_v2.ControlPlaneScalingTier",
		"TIER_2XL",
		&returns,
	)
	return returns
}

func ControlPlaneScalingTier_TIER_4XL() ControlPlaneScalingTier {
	_init_.Initialize()
	var returns ControlPlaneScalingTier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks_v2.ControlPlaneScalingTier",
		"TIER_4XL",
		&returns,
	)
	return returns
}

func ControlPlaneScalingTier_TIER_8XL() ControlPlaneScalingTier {
	_init_.Initialize()
	var returns ControlPlaneScalingTier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks_v2.ControlPlaneScalingTier",
		"TIER_8XL",
		&returns,
	)
	return returns
}

func ControlPlaneScalingTier_TIER_XL() ControlPlaneScalingTier {
	_init_.Initialize()
	var returns ControlPlaneScalingTier
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks_v2.ControlPlaneScalingTier",
		"TIER_XL",
		&returns,
	)
	return returns
}

