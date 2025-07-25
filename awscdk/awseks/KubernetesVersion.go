package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Kubernetes cluster version.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv33"
//
//   // or
//   var vpc vpc
//
//
//   eks.NewCluster(this, jsii.String("MyCluster"), &ClusterProps{
//   	KubectlMemory: awscdk.Size_Gibibytes(jsii.Number(4)),
//   	Version: eks.KubernetesVersion_V1_33(),
//   	KubectlLayer: kubectlv33.NewKubectlV33Layer(this, jsii.String("kubectl")),
//   })
//   eks.Cluster_FromClusterAttributes(this, jsii.String("MyCluster"), &ClusterAttributes{
//   	KubectlMemory: awscdk.Size_*Gibibytes(jsii.Number(4)),
//   	Vpc: Vpc,
//   	ClusterName: jsii.String("cluster-name"),
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

func KubernetesVersion_V1_27() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_27",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_28() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_28",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_29() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_29",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_30() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_30",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_31() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_31",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_32() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_32",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_33() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.KubernetesVersion",
		"V1_33",
		&returns,
	)
	return returns
}

