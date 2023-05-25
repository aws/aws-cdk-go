package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Kubernetes cluster version.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_26(),
//   	DefaultCapacityType: eks.DefaultCapacityType_EC2,
//   })
//
// See: https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions.html#kubernetes-release-calendar
//
type KubernetesVersion interface {
	// cluster version number.
	Version() *string
}

// The jsii proxy struct for KubernetesVersion
type jsiiProxy_KubernetesVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_KubernetesVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom cluster version.
func KubernetesVersion_Of(version *string) KubernetesVersion {
	_init_.Initialize()

	if err := validateKubernetesVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns KubernetesVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func KubernetesVersion_V1_14() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_14",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_15() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_15",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_16() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_16",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_17() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_17",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_18() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_18",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_19() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_19",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_20() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_20",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_21() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_21",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_22() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_22",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_23() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_23",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_24() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_24",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_25() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_25",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_26() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_26",
		&returns,
	)
	return returns
}

