package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controller version.
//
// Corresponds to the image tag of 'amazon/aws-load-balancer-controller' image.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv34"
//
//
//   eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_34(),
//   	AlbController: &AlbControllerOptions{
//   		Version: eks.AlbControllerVersion_V2_8_2(),
//   	},
//   	KubectlLayer: kubectlv34.NewKubectlV34Layer(this, jsii.String("kubectl")),
//   })
//
type AlbControllerVersion interface {
	// Whether or not its a custom version.
	Custom() *bool
	// The version of the helm chart to use.
	HelmChartVersion() *string
	// The version string.
	Version() *string
}

// The jsii proxy struct for AlbControllerVersion
type jsiiProxy_AlbControllerVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AlbControllerVersion) Custom() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbControllerVersion) HelmChartVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"helmChartVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbControllerVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Specify a custom version and an associated helm chart version.
//
// Use this if the version you need is not available in one of the predefined versions.
// Note that in this case, you will also need to provide an IAM policy in the controller options.
//
// ALB controller version and helm chart version compatibility information can be found
// here: https://github.com/aws/eks-charts/blob/v0.0.133/stable/aws-load-balancer-controller/Chart.yaml
func AlbControllerVersion_Of(version *string, helmChartVersion *string) AlbControllerVersion {
	_init_.Initialize()

	if err := validateAlbControllerVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns AlbControllerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"of",
		[]interface{}{version, helmChartVersion},
		&returns,
	)

	return returns
}

func AlbControllerVersion_V2_0_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_0_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_0_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_0_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_1_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_1_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_1_2",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_3() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_1_3",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_2_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_2_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_2_2",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_3() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_2_3",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_4() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_2_4",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_3_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_3_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_3_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_3_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_4_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_4_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_4_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_4_2",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_4_3() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_4_3",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_4_4() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_4_4",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_4_5() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_4_5",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_4_6() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_4_6",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_4_7() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_4_7",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_5_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_5_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_5_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_5_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_5_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_5_2",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_5_3() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_5_3",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_5_4() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_5_4",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_6_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_6_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_6_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_6_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_6_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_6_2",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_7_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_7_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_7_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_7_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_7_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_7_2",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_8_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_8_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_8_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_8_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_8_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_8_2",
		&returns,
	)
	return returns
}

