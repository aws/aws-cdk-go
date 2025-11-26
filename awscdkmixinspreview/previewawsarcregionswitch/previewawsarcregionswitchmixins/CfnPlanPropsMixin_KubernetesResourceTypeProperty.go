package previewawsarcregionswitchmixins


// Defines the type of Kubernetes resource to scale in an Amazon EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kubernetesResourceTypeProperty := &KubernetesResourceTypeProperty{
//   	ApiVersion: jsii.String("apiVersion"),
//   	Kind: jsii.String("kind"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesresourcetype.html
//
type CfnPlanPropsMixin_KubernetesResourceTypeProperty struct {
	// The API version type for the Kubernetes resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesresourcetype.html#cfn-arcregionswitch-plan-kubernetesresourcetype-apiversion
	//
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The kind for the Kubernetes resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-kubernetesresourcetype.html#cfn-arcregionswitch-plan-kubernetesresourcetype-kind
	//
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

