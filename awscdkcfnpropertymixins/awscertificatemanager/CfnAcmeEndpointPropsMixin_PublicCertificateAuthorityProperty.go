package awscertificatemanager


// Configuration for the public certificate authority.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   publicCertificateAuthorityProperty := &PublicCertificateAuthorityProperty{
//   	AllowedKeyAlgorithms: []*string{
//   		jsii.String("allowedKeyAlgorithms"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmeendpoint-publiccertificateauthority.html
//
type CfnAcmeEndpointPropsMixin_PublicCertificateAuthorityProperty struct {
	// The allowed key algorithms for certificates issued via this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmeendpoint-publiccertificateauthority.html#cfn-certificatemanager-acmeendpoint-publiccertificateauthority-allowedkeyalgorithms
	//
	AllowedKeyAlgorithms *[]*string `field:"optional" json:"allowedKeyAlgorithms" yaml:"allowedKeyAlgorithms"`
}

