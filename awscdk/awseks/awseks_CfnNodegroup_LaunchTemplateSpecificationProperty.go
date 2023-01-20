package awseks


// An object representing a node group launch template specification.
//
// The launch template cannot include [`SubnetId`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html) , [`IamInstanceProfile`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html) , [`RequestSpotInstances`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RequestSpotInstances.html) , [`HibernationOptions`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_HibernationOptionsRequest.html) , or [`TerminateInstances`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TerminateInstances.html) , or the node group deployment or update will fail. For more information about launch templates, see [`CreateLaunchTemplate`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateLaunchTemplate.html) in the Amazon EC2 API Reference. For more information about using launch templates with Amazon EKS, see [Launch template support](https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the *Amazon EKS User Guide* .
//
// Specify either `name` or `id` , but not both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateSpecificationProperty := &launchTemplateSpecificationProperty{
//   	id: jsii.String("id"),
//   	name: jsii.String("name"),
//   	version: jsii.String("version"),
//   }
//
type CfnNodegroup_LaunchTemplateSpecificationProperty struct {
	// The ID of the launch template.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the launch template.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The version of the launch template to use.
	//
	// If no version is specified, then the template's default version is used.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

