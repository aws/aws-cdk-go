package awscodebuild


// Contains compute attributes.
//
// These attributes only need be specified when your project's or fleet's `computeType` is set to `ATTRIBUTE_BASED_COMPUTE` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeConfigurationProperty := &ComputeConfigurationProperty{
//   	Disk: jsii.Number(123),
//   	MachineType: jsii.String("machineType"),
//   	Memory: jsii.Number(123),
//   	VCpu: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-computeconfiguration.html
//
type CfnFleet_ComputeConfigurationProperty struct {
	// The amount of disk space of the instance type included in your fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-computeconfiguration.html#cfn-codebuild-fleet-computeconfiguration-disk
	//
	Disk *float64 `field:"optional" json:"disk" yaml:"disk"`
	// The machine type of the instance type included in your fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-computeconfiguration.html#cfn-codebuild-fleet-computeconfiguration-machinetype
	//
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The amount of memory of the instance type included in your fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-computeconfiguration.html#cfn-codebuild-fleet-computeconfiguration-memory
	//
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The number of vCPUs of the instance type included in your fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-computeconfiguration.html#cfn-codebuild-fleet-computeconfiguration-vcpu
	//
	VCpu *float64 `field:"optional" json:"vCpu" yaml:"vCpu"`
}

