package awsmsk


// Includes all client authentication information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientAuthenticationProperty := &ClientAuthenticationProperty{
//   	Sasl: &SaslProperty{
//   		Iam: &IamProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		Scram: &ScramProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Tls: &TlsProperty{
//   		CertificateAuthorityArnList: []*string{
//   			jsii.String("certificateAuthorityArnList"),
//   		},
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Unauthenticated: &UnauthenticatedProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html
//
type CfnCluster_ClientAuthenticationProperty struct {
	// Details for client authentication using SASL.
	//
	// To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT` . If you choose `TLS_PLAINTEXT` , then you must also set `unauthenticated` to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html#cfn-msk-cluster-clientauthentication-sasl
	//
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// Details for ClientAuthentication using TLS.
	//
	// To turn on TLS access control, you must also turn on `EncryptionInTransit` by setting `inCluster` to true and `clientBroker` to `TLS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html#cfn-msk-cluster-clientauthentication-tls
	//
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Details for ClientAuthentication using no authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html#cfn-msk-cluster-clientauthentication-unauthenticated
	//
	Unauthenticated interface{} `field:"optional" json:"unauthenticated" yaml:"unauthenticated"`
}

