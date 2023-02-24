// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


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
	// The Launch template ID.
	//
	// Mutually exclusive with `launchTemplateName`.
	// Experimental.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The Launch template name.
	//
	// Mutually exclusive with `launchTemplateId`.
	// Experimental.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// Use security groups defined in the launch template network interfaces.
	//
	// In some cases, such as specifying Elastic Fabric Adapters,
	// network interfaces must be used to specify security groups.  This
	// parameter tells the Compute Environment construct that this is your
	// intention, and stops it from creating its own security groups.  This
	// parameter is mutually exclusive with securityGroups in the Compute
	// Environment.
	// Experimental.
	UseNetworkInterfaceSecurityGroups *bool `field:"optional" json:"useNetworkInterfaceSecurityGroups" yaml:"useNetworkInterfaceSecurityGroups"`
	// The launch template version to be used (optional).
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

