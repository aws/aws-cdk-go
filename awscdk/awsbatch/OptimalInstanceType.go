package awsbatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Not a real instance type!
//
// Indicates that Batch will choose one it determines to be optimal
// for the workload.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optimalInstanceType := awscdk.Aws_batch.NewOptimalInstanceType()
//
type OptimalInstanceType interface {
	awsec2.InstanceType
	// The instance's CPU architecture.
	Architecture() awsec2.InstanceArchitecture
	// Return whether this instance type is a burstable instance type.
	IsBurstable() *bool
	SameInstanceClassAs(other awsec2.InstanceType) *bool
	// Return the instance type as a dotted string.
	ToString() *string
}

// The jsii proxy struct for OptimalInstanceType
type jsiiProxy_OptimalInstanceType struct {
	internal.Type__awsec2InstanceType
}

func (j *jsiiProxy_OptimalInstanceType) Architecture() awsec2.InstanceArchitecture {
	var returns awsec2.InstanceArchitecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}


func NewOptimalInstanceType() OptimalInstanceType {
	_init_.Initialize()

	j := jsiiProxy_OptimalInstanceType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.OptimalInstanceType",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewOptimalInstanceType_Override(o OptimalInstanceType) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_batch.OptimalInstanceType",
		nil, // no parameters
		o,
	)
}

// Instance type for EC2 instances.
//
// This class takes a combination of a class and size.
//
// Be aware that not all combinations of class and size are available, and not all
// classes are available in all regions.
func OptimalInstanceType_Of(instanceClass awsec2.InstanceClass, instanceSize awsec2.InstanceSize) awsec2.InstanceType {
	_init_.Initialize()

	if err := validateOptimalInstanceType_OfParameters(instanceClass, instanceSize); err != nil {
		panic(err)
	}
	var returns awsec2.InstanceType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_batch.OptimalInstanceType",
		"of",
		[]interface{}{instanceClass, instanceSize},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OptimalInstanceType) IsBurstable() *bool {
	var returns *bool

	_jsii_.Invoke(
		o,
		"isBurstable",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OptimalInstanceType) SameInstanceClassAs(other awsec2.InstanceType) *bool {
	if err := o.validateSameInstanceClassAsParameters(other); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		o,
		"sameInstanceClassAs",
		[]interface{}{other},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OptimalInstanceType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

