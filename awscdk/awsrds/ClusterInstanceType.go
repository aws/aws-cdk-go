package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The type of Aurora Cluster Instance.
//
// Can be either serverless v2
// or provisioned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterInstanceType := awscdk.Aws_rds.NewClusterInstanceType(jsii.String("instanceType"), awscdk.Aws_rds.InstanceType_PROVISIONED)
//
type ClusterInstanceType interface {
	Type() InstanceType
	// String representation of the instance type that can be used in the CloudFormation resource.
	ToString() *string
}

// The jsii proxy struct for ClusterInstanceType
type jsiiProxy_ClusterInstanceType struct {
	_ byte // padding
}

func (j *jsiiProxy_ClusterInstanceType) Type() InstanceType {
	var returns InstanceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewClusterInstanceType(instanceType *string, type_ InstanceType) ClusterInstanceType {
	_init_.Initialize()

	if err := validateNewClusterInstanceTypeParameters(instanceType, type_); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClusterInstanceType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ClusterInstanceType",
		[]interface{}{instanceType, type_},
		&j,
	)

	return &j
}

func NewClusterInstanceType_Override(c ClusterInstanceType, instanceType *string, type_ InstanceType) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.ClusterInstanceType",
		[]interface{}{instanceType, type_},
		c,
	)
}

// Aurora Provisioned instance type.
func ClusterInstanceType_Provisioned(instanceType awsec2.InstanceType) ClusterInstanceType {
	_init_.Initialize()

	var returns ClusterInstanceType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ClusterInstanceType",
		"provisioned",
		[]interface{}{instanceType},
		&returns,
	)

	return returns
}

// Aurora Serverless V2 instance type.
// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html
//
func ClusterInstanceType_ServerlessV2() ClusterInstanceType {
	_init_.Initialize()

	var returns ClusterInstanceType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.ClusterInstanceType",
		"serverlessV2",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterInstanceType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

