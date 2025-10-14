package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The compute configuration for the fleet.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fleet := codebuild.NewFleet(this, jsii.String("MyFleet"), &FleetProps{
//   	BaseCapacity: jsii.Number(1),
//   	ComputeType: codebuild.FleetComputeType_ATTRIBUTE_BASED,
//   	EnvironmentType: codebuild.EnvironmentType_LINUX_CONTAINER,
//   	ComputeConfiguration: &ComputeConfiguration{
//   		VCpu: jsii.Number(2),
//   		Memory: awscdk.Size_Gibibytes(jsii.Number(4)),
//   		Disk: awscdk.Size_*Gibibytes(jsii.Number(10)),
//   		MachineType: codebuild.MachineType_GENERAL,
//   	},
//   })
//
type ComputeConfiguration struct {
	// When using ATTRIBUTE_BASED, the amount of disk space of the instance type included in your fleet.
	//
	// When using CUSTOM_INSTANCE_TYPE,
	// the additional amount of disk space to provision over the 64GB included by
	// default.
	// Default: - No requirement, the actual value will be based on the other selected configuration properties.
	//
	Disk awscdk.Size `field:"optional" json:"disk" yaml:"disk"`
	// When using CUSTOM_INSTANCE_TYPE, the EC2 instance type to use for fleet instances.
	//
	// Not all instance types are supported by CodeBuild. If you use a disallowed type, the
	// CloudFormation deployment will fail.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html#environment-reserved-capacity.instance-types
	//
	// Default: none.
	//
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// When using ATTRIBUTE_BASED, the machine type of the instance type included in your fleet.
	// Default: - No requirement, the actual value will be based on the other selected configuration properties.
	//
	MachineType MachineType `field:"optional" json:"machineType" yaml:"machineType"`
	// When using ATTRIBUTE_BASED, the amount of memory of the instance type included in your fleet.
	// Default: - No requirement, the actual value will be based on the other selected configuration properties.
	//
	Memory awscdk.Size `field:"optional" json:"memory" yaml:"memory"`
	// When using ATTRIBUTE_BASED, the number of vCPUs of the instance type included in your fleet.
	// Default: - No requirement, the actual value will be based on the other selected configuration properties.
	//
	VCpu *float64 `field:"optional" json:"vCpu" yaml:"vCpu"`
}

