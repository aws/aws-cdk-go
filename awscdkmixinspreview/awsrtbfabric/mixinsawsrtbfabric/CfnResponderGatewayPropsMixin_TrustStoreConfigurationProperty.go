package mixinsawsrtbfabric


// Describes the configuration of a trust store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trustStoreConfigurationProperty := &TrustStoreConfigurationProperty{
//   	CertificateAuthorityCertificates: []*string{
//   		jsii.String("certificateAuthorityCertificates"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-truststoreconfiguration.html
//
type CfnResponderGatewayPropsMixin_TrustStoreConfigurationProperty struct {
	// The certificate authority certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rtbfabric-respondergateway-truststoreconfiguration.html#cfn-rtbfabric-respondergateway-truststoreconfiguration-certificateauthoritycertificates
	//
	CertificateAuthorityCertificates *[]*string `field:"optional" json:"certificateAuthorityCertificates" yaml:"certificateAuthorityCertificates"`
}

