package awsapprunner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The amount of memory reserved for each instance of your App Runner service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memory := awscdk.Aws_apprunner.Memory_FOUR_GB()
//
// Experimental.
type Memory interface {
	// The unit of memory.
	// Experimental.
	Unit() *string
}

// The jsii proxy struct for Memory
type jsiiProxy_Memory struct {
	_ byte // padding
}

func (j *jsiiProxy_Memory) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}


// Custom Memory unit.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-instanceconfiguration.html#cfn-apprunner-service-instanceconfiguration-memory
//
// Experimental.
func Memory_Of(unit *string) Memory {
	_init_.Initialize()

	if err := validateMemory_OfParameters(unit); err != nil {
		panic(err)
	}
	var returns Memory

	_jsii_.StaticInvoke(
		"monocdk.aws_apprunner.Memory",
		"of",
		[]interface{}{unit},
		&returns,
	)

	return returns
}

func Memory_FOUR_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Memory",
		"FOUR_GB",
		&returns,
	)
	return returns
}

func Memory_THREE_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Memory",
		"THREE_GB",
		&returns,
	)
	return returns
}

func Memory_TWO_GB() Memory {
	_init_.Initialize()
	var returns Memory
	_jsii_.StaticGet(
		"monocdk.aws_apprunner.Memory",
		"TWO_GB",
		&returns,
	)
	return returns
}

