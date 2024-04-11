package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Instance type for EC2 instances.
//
// This class takes a literal string, good if you already
// know the identifier of the type you want.
//
// Example:
//   var vpc vpc
//
//
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	MasterUser: &Login{
//   		Username: jsii.String("myuser"),
//   	},
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   	Vpc: Vpc,
//   	RemovalPolicy: awscdk.RemovalPolicy_SNAPSHOT,
//   	InstanceRemovalPolicy: awscdk.RemovalPolicy_RETAIN,
//   })
//
type InstanceType interface {
	// The instance's CPU architecture.
	Architecture() InstanceArchitecture
	// Return whether this instance type is a burstable instance type.
	IsBurstable() *bool
	SameInstanceClassAs(other InstanceType) *bool
	// Return the instance type as a dotted string.
	ToString() *string
}

// The jsii proxy struct for InstanceType
type jsiiProxy_InstanceType struct {
	_ byte // padding
}

func (j *jsiiProxy_InstanceType) Architecture() InstanceArchitecture {
	var returns InstanceArchitecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}


func NewInstanceType(instanceTypeIdentifier *string) InstanceType {
	_init_.Initialize()

	if err := validateNewInstanceTypeParameters(instanceTypeIdentifier); err != nil {
		panic(err)
	}
	j := jsiiProxy_InstanceType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InstanceType",
		[]interface{}{instanceTypeIdentifier},
		&j,
	)

	return &j
}

func NewInstanceType_Override(i InstanceType, instanceTypeIdentifier *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InstanceType",
		[]interface{}{instanceTypeIdentifier},
		i,
	)
}

// Instance type for EC2 instances.
//
// This class takes a combination of a class and size.
//
// Be aware that not all combinations of class and size are available, and not all
// classes are available in all regions.
func InstanceType_Of(instanceClass InstanceClass, instanceSize InstanceSize) InstanceType {
	_init_.Initialize()

	if err := validateInstanceType_OfParameters(instanceClass, instanceSize); err != nil {
		panic(err)
	}
	var returns InstanceType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InstanceType",
		"of",
		[]interface{}{instanceClass, instanceSize},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceType) IsBurstable() *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"isBurstable",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceType) SameInstanceClassAs(other InstanceType) *bool {
	if err := i.validateSameInstanceClassAsParameters(other); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"sameInstanceClassAs",
		[]interface{}{other},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

