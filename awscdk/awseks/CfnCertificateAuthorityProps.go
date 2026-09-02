package awseks


// Properties for defining a `CfnCertificateAuthority`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateAuthorityProps := &CfnCertificateAuthorityProps{
//   	ClusterName: jsii.String("clusterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-certificateauthority.html
//
type CfnCertificateAuthorityProps struct {
	// The name of the EKS cluster that the certificate authority belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-certificateauthority.html#cfn-eks-certificateauthority-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}

