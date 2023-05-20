package awsmsk


// Includes all client authentication information for VpcConnectivity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectivityClientAuthenticationProperty := &VpcConnectivityClientAuthenticationProperty{
//   	Sasl: &VpcConnectivitySaslProperty{
//   		Iam: &VpcConnectivityIamProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		Scram: &VpcConnectivityScramProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Tls: &VpcConnectivityTlsProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnCluster_VpcConnectivityClientAuthenticationProperty struct {
	// Details for VpcConnectivity ClientAuthentication using SASL.
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// Details for VpcConnectivity ClientAuthentication using TLS.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

