package awsmsk


// Includes information related to client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientAuthenticationProperty := &clientAuthenticationProperty{
//   	sasl: &saslProperty{
//   		iam: &iamProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		scram: &scramProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	tls: &tlsProperty{
//   		certificateAuthorityArnList: []*string{
//   			jsii.String("certificateAuthorityArnList"),
//   		},
//   		enabled: jsii.Boolean(false),
//   	},
//   	unauthenticated: &unauthenticatedProperty{
//   		enabled: jsii.Boolean(false),
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

