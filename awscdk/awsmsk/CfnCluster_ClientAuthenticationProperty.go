package awsmsk


// Includes information related to client authentication.
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
type CfnCluster_ClientAuthenticationProperty struct {
	// Details for ClientAuthentication using SASL.
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// Details for client authentication using TLS.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Details for ClientAuthentication using no authentication.
	Unauthenticated interface{} `field:"optional" json:"unauthenticated" yaml:"unauthenticated"`
}

