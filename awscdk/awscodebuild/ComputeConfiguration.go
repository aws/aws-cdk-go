package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The compute configuration for the fleet.
//
// Despite what the CloudFormation schema says, the numeric properties (disk, memory, vCpu) are not optional.
//  An `undefined` value will cause the CloudFormation deployment to fail, e.g.
//  > Cannot invoke "java.lang.Integer.intValue()" because the return value of "software.amazon.codebuild.fleet.ComputeConfiguration.getMemory()" is null
// Therefore, these properties default value is set to 0.
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
	// The amount of disk space of the instance type included in your fleet.
	// Default: - No requirement, the actual value will be based on the other selected configuration properties.
	//
	Disk awscdk.Size `field:"optional" json:"disk" yaml:"disk"`
	// The machine type of the instance type included in your fleet.
	// Default: - No requirement, the actual value will be based on the other selected configuration properties.
	//
	MachineType MachineType `field:"optional" json:"machineType" yaml:"machineType"`
	// The amount of memory of the instance type included in your fleet.
	// Default: - No requirement, the actual value will be based on the other selected configuration properties.
	//
	Memory awscdk.Size `field:"optional" json:"memory" yaml:"memory"`
	// The number of vCPUs of the instance type included in your fleet.
	// Default: - No requirement, the actual value will be based on the other selected configuration properties.
	//
	VCpu *float64 `field:"optional" json:"vCpu" yaml:"vCpu"`
}

