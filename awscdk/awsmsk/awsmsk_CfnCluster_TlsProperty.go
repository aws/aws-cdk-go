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
type CfnCluster_TlsProperty struct {
	// List of ACM Certificate Authority ARNs.
	CertificateAuthorityArnList *[]*string `field:"optional" json:"certificateAuthorityArnList" yaml:"certificateAuthorityArnList"`
	// TLS authentication is enabled or not.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

