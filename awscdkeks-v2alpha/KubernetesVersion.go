package awscdkeks-v2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeks-v2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Kubernetes cluster version.
//
// Example:
//   cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_31(),
//   	EndpointAccess: eks.EndpointAccess_PRIVATE(),
//   })
//
// See: https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions.html#kubernetes-release-calendar
//
// Experimental.
type KubernetesVersion interface {
	// cluster version number.
	// Experimental.
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
// Experimental.
func KubernetesVersion_Of(version *string) KubernetesVersion {
	_init_.Initialize()

	if err := validateKubernetesVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns KubernetesVersion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func KubernetesVersion_V1_25() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesVersion",
		"V1_25",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_26() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesVersion",
		"V1_26",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_27() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesVersion",
		"V1_27",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_28() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesVersion",
		"V1_28",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_29() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesVersion",
		"V1_29",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_30() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesVersion",
		"V1_30",
		&returns,
	)
	return returns
}

func KubernetesVersion_V1_31() KubernetesVersion {
	_init_.Initialize()
	var returns KubernetesVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesVersion",
		"V1_31",
		&returns,
	)
	return returns
}

