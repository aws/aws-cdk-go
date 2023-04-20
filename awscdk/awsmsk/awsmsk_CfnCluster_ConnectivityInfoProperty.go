package awsmsk


// Specifies whether the cluster's brokers are publicly accessible.
//
// By default, they are not.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectivityInfoProperty := &ConnectivityInfoProperty{
//   	PublicAccess: &PublicAccessProperty{
//   		Type: jsii.String("type"),
//   	},
//   	VpcConnectivity: &VpcConnectivityProperty{
//   		ClientAuthentication: &VpcConnectivityClientAuthenticationProperty{
//   			Sasl: &VpcConnectivitySaslProperty{
//   				Iam: &VpcConnectivityIamProperty{
//   					Enabled: jsii.Boolean(false),
//   				},
//   				Scram: &VpcConnectivityScramProperty{
//   					Enabled: jsii.Boolean(false),
//   				},
//   			},
//   			Tls: &VpcConnectivityTlsProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
type CfnCluster_ConnectivityInfoProperty struct {
	// Specifies whether the cluster's brokers are accessible from the internet.
	//
	// Public access is off by default.
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// Not currently supported by AWS CloudFormation .
	VpcConnectivity interface{} `field:"optional" json:"vpcConnectivity" yaml:"vpcConnectivity"`
}

