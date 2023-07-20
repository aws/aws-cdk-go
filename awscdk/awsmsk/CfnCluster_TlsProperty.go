package awsmsk


// Details for client authentication using TLS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsProperty := &TlsProperty{
//   	CertificateAuthorityArnList: []*string{
//   		jsii.String("certificateAuthorityArnList"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-tls.html
//
type CfnCluster_TlsProperty struct {
	// List of AWS Private CA ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-tls.html#cfn-msk-cluster-tls-certificateauthorityarnlist
	//
	CertificateAuthorityArnList *[]*string `field:"optional" json:"certificateAuthorityArnList" yaml:"certificateAuthorityArnList"`
	// TLS authentication is enabled or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-tls.html#cfn-msk-cluster-tls-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

