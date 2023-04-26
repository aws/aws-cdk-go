package awsbatch


// Launch template property specification.
//
// Example:
//   var vpc vpc
//   var myLaunchTemplate cfnLaunchTemplate
//
//
//   myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &ComputeEnvironmentProps{
//   	ComputeResources: &ComputeResources{
//   		LaunchTemplate: &LaunchTemplateSpecification{
//   			LaunchTemplateName: string(myLaunchTemplate.LaunchTemplateName),
//   		},
//   		Vpc: *Vpc,
//   	},
//   	ComputeEnvironmentName: jsii.String("MyStorageCapableComputeEnvironment"),
//   })
//
// Experimental.
type LaunchTemplateSpecification struct {
	// The Launch template name.
	// Experimental.
	LaunchTemplateName *string `field:"required" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The launch template version to be used (optional).
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

