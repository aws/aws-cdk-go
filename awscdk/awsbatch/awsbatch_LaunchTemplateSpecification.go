package awsbatch


// Launch template property specification.
//
// Example:
//   var vpc vpc
//   var myLaunchTemplate cfnLaunchTemplate
//
//
//   myComputeEnv := batch.NewComputeEnvironment(this, jsii.String("ComputeEnv"), &computeEnvironmentProps{
//   	computeResources: &computeResources{
//   		launchTemplate: &launchTemplateSpecification{
//   			launchTemplateName: string(myLaunchTemplate.launchTemplateName),
//   		},
//   		vpc: vpc,
//   	},
//   	computeEnvironmentName: jsii.String("MyStorageCapableComputeEnvironment"),
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

