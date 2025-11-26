package previewawsarcregionswitchmixins


// Defines a Kubernetes resource to scale in an Amazon EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kubernetesScalingResourceProperty := &KubernetesScalingResourceProperty{
//   	HpaName: jsii.String("hpaName"),
//   	Name: jsii.String("name"),
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesscalingresource.html
//
type CfnPlanPropsMixin_KubernetesScalingResourceProperty struct {
	// The hpaname for the Kubernetes resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesscalingresource.html#cfn-arcregionswitch-plan-kubernetesscalingresource-hpaname
	//
	HpaName *string `field:"optional" json:"hpaName" yaml:"hpaName"`
	// The name for the Kubernetes resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesscalingresource.html#cfn-arcregionswitch-plan-kubernetesscalingresource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace for the Kubernetes resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesscalingresource.html#cfn-arcregionswitch-plan-kubernetesscalingresource-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

