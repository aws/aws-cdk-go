package awscdkapprunneralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The number of CPU units reserved for each instance of your App Runner service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apprunner_alpha "github.com/aws/aws-cdk-go/awscdkapprunneralpha"
//
//   cpu := apprunner_alpha.Cpu_FOUR_VCPU()
//
// Experimental.
type Cpu interface {
	// The unit of CPU.
	// Experimental.
	Unit() *string
}

// The jsii proxy struct for Cpu
type jsiiProxy_Cpu struct {
	_ byte // padding
}

func (j *jsiiProxy_Cpu) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}


// Custom CPU unit.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-instanceconfiguration.html#cfn-apprunner-service-instanceconfiguration-cpu
//
// Experimental.
func Cpu_Of(unit *string) Cpu {
	_init_.Initialize()

	if err := validateCpu_OfParameters(unit); err != nil {
		panic(err)
	}
	var returns Cpu

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apprunner-alpha.Cpu",
		"of",
		[]interface{}{unit},
		&returns,
	)

	return returns
}

func Cpu_FOUR_VCPU() Cpu {
	_init_.Initialize()
	var returns Cpu
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Cpu",
		"FOUR_VCPU",
		&returns,
	)
	return returns
}

func Cpu_HALF_VCPU() Cpu {
	_init_.Initialize()
	var returns Cpu
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Cpu",
		"HALF_VCPU",
		&returns,
	)
	return returns
}

func Cpu_ONE_VCPU() Cpu {
	_init_.Initialize()
	var returns Cpu
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Cpu",
		"ONE_VCPU",
		&returns,
	)
	return returns
}

func Cpu_QUARTER_VCPU() Cpu {
	_init_.Initialize()
	var returns Cpu
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Cpu",
		"QUARTER_VCPU",
		&returns,
	)
	return returns
}

func Cpu_TWO_VCPU() Cpu {
	_init_.Initialize()
	var returns Cpu
	_jsii_.StaticGet(
		"@aws-cdk/aws-apprunner-alpha.Cpu",
		"TWO_VCPU",
		&returns,
	)
	return returns
}

