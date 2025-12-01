package awseks


// The control plane scaling tier configuration.
//
// For more information, see EKS Provisioned Control Plane in the Amazon EKS User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   controlPlaneScalingConfigProperty := &ControlPlaneScalingConfigProperty{
//   	Tier: jsii.String("tier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-controlplanescalingconfig.html
//
type CfnCluster_ControlPlaneScalingConfigProperty struct {
	// The control plane scaling tier configuration.
	//
	// Available options are `standard` , `tier-xl` , `tier-2xl` , or `tier-4xl` . For more information, see EKS Provisioned Control Plane in the Amazon EKS User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-controlplanescalingconfig.html#cfn-eks-cluster-controlplanescalingconfig-tier
	//
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

