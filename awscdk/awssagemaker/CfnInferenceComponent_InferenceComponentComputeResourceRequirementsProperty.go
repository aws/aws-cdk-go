package awssagemaker


// Defines the compute resources to allocate to run a model that you assign to an inference component.
//
// These resources include CPU cores, accelerators, and memory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceComponentComputeResourceRequirementsProperty := &InferenceComponentComputeResourceRequirementsProperty{
//   	MaxMemoryRequiredInMb: jsii.Number(123),
//   	MinMemoryRequiredInMb: jsii.Number(123),
//   	NumberOfAcceleratorDevicesRequired: jsii.Number(123),
//   	NumberOfCpuCoresRequired: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.html
//
type CfnInferenceComponent_InferenceComponentComputeResourceRequirementsProperty struct {
	// The maximum MB of memory to allocate to run a model that you assign to an inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.html#cfn-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements-maxmemoryrequiredinmb
	//
	MaxMemoryRequiredInMb *float64 `field:"optional" json:"maxMemoryRequiredInMb" yaml:"maxMemoryRequiredInMb"`
	// The minimum MB of memory to allocate to run a model that you assign to an inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.html#cfn-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements-minmemoryrequiredinmb
	//
	MinMemoryRequiredInMb *float64 `field:"optional" json:"minMemoryRequiredInMb" yaml:"minMemoryRequiredInMb"`
	// The number of accelerators to allocate to run a model that you assign to an inference component.
	//
	// Accelerators include GPUs and AWS Inferentia.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.html#cfn-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements-numberofacceleratordevicesrequired
	//
	NumberOfAcceleratorDevicesRequired *float64 `field:"optional" json:"numberOfAcceleratorDevicesRequired" yaml:"numberOfAcceleratorDevicesRequired"`
	// The number of CPU cores to allocate to run a model that you assign to an inference component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.html#cfn-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements-numberofcpucoresrequired
	//
	NumberOfCpuCoresRequired *float64 `field:"optional" json:"numberOfCpuCoresRequired" yaml:"numberOfCpuCoresRequired"`
}

