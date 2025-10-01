package awsarcregionswitch


// Defines a Kubernetes resource to scale in an Amazon EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kubernetesScalingResourceProperty := &KubernetesScalingResourceProperty{
//   	Name: jsii.String("name"),
//   	Namespace: jsii.String("namespace"),
//
//   	// the properties below are optional
//   	HpaName: jsii.String("hpaName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesscalingresource.html
//
type CfnPlan_KubernetesScalingResourceProperty struct {
	// The name for the Kubernetes resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesscalingresource.html#cfn-arcregionswitch-plan-kubernetesscalingresource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace for the Kubernetes resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesscalingresource.html#cfn-arcregionswitch-plan-kubernetesscalingresource-namespace
	//
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The hpaname for the Kubernetes resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesscalingresource.html#cfn-arcregionswitch-plan-kubernetesscalingresource-hpaname
	//
	HpaName *string `field:"optional" json:"hpaName" yaml:"hpaName"`
}

