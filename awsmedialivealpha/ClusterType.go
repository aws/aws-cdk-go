package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The hardware type for the cluster.
//
// Example:
//   var stack Stack
//   var instanceRole IRole
//
//
//   cluster := medialive.NewCluster(stack, jsii.String("Cluster"), &ClusterProps{
//   	ClusterName: jsii.String("on-prem-cluster"),
//   	ClusterType: medialive.ClusterType_ON_PREMISES(),
//   	InstanceRole: InstanceRole,
//   })
//
// Experimental.
type ClusterType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ClusterType
type jsiiProxy_ClusterType struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func ClusterType_Of(value *string) ClusterType {
	_init_.Initialize()

	if err := validateClusterType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ClusterType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ClusterType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ClusterType_EC2() ClusterType {
	_init_.Initialize()
	var returns ClusterType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ClusterType",
		"EC2",
		&returns,
	)
	return returns
}

func ClusterType_ON_PREMISES() ClusterType {
	_init_.Initialize()
	var returns ClusterType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ClusterType",
		"ON_PREMISES",
		&returns,
	)
	return returns
}

func ClusterType_OUTPOSTS_RACK() ClusterType {
	_init_.Initialize()
	var returns ClusterType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ClusterType",
		"OUTPOSTS_RACK",
		&returns,
	)
	return returns
}

func ClusterType_OUTPOSTS_SERVER() ClusterType {
	_init_.Initialize()
	var returns ClusterType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ClusterType",
		"OUTPOSTS_SERVER",
		&returns,
	)
	return returns
}

