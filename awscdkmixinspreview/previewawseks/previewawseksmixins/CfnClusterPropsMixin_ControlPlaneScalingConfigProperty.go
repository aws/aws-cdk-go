package previewawseksmixins


// Configuration for provisioned control plane scaling.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   controlPlaneScalingConfigProperty := &ControlPlaneScalingConfigProperty{
//   	Tier: jsii.String("tier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-controlplanescalingconfig.html
//
type CfnClusterPropsMixin_ControlPlaneScalingConfigProperty struct {
	// The scaling tier for the provisioned control plane.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-controlplanescalingconfig.html#cfn-eks-cluster-controlplanescalingconfig-tier
	//
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

