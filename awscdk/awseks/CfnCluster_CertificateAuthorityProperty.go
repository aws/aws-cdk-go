package awseks


// The certificate authority information for the cluster, including the trust bundle and the currently active signing certificate authority.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateAuthorityProperty := &CertificateAuthorityProperty{
//   	Active: &ActiveCertificateAuthorityProperty{
//   		ActivatedBy: jsii.String("activatedBy"),
//   		Id: jsii.String("id"),
//   	},
//   	Data: jsii.String("data"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-certificateauthority.html
//
type CfnCluster_CertificateAuthorityProperty struct {
	// Identifies the certificate authority currently signing certificates for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-certificateauthority.html#cfn-eks-cluster-certificateauthority-active
	//
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// The base64-encoded certificate-authority trust bundle for the cluster (all trusted CAs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-certificateauthority.html#cfn-eks-cluster-certificateauthority-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
}

