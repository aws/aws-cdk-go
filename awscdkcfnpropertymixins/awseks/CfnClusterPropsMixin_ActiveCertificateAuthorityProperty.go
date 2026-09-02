package awseks


// Identifies the certificate authority currently signing certificates for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   activeCertificateAuthorityProperty := &ActiveCertificateAuthorityProperty{
//   	ActivatedBy: jsii.String("activatedBy"),
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-activecertificateauthority.html
//
type CfnClusterPropsMixin_ActiveCertificateAuthorityProperty struct {
	// Indicates whether the active certificate authority was activated by EKS or by the customer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-activecertificateauthority.html#cfn-eks-cluster-activecertificateauthority-activatedby
	//
	ActivatedBy *string `field:"optional" json:"activatedBy" yaml:"activatedBy"`
	// The ID of the active (signing) certificate authority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-activecertificateauthority.html#cfn-eks-cluster-activecertificateauthority-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

