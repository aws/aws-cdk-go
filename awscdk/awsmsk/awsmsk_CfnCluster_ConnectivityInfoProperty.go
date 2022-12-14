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
//   connectivityInfoProperty := &connectivityInfoProperty{
//   	publicAccess: &publicAccessProperty{
//   		type: jsii.String("type"),
//   	},
//   	vpcConnectivity: &vpcConnectivityProperty{
//   		clientAuthentication: &vpcConnectivityClientAuthenticationProperty{
//   			sasl: &vpcConnectivitySaslProperty{
//   				iam: &vpcConnectivityIamProperty{
//   					enabled: jsii.Boolean(false),
//   				},
//   				scram: &vpcConnectivityScramProperty{
//   					enabled: jsii.Boolean(false),
//   				},
//   			},
//   			tls: &vpcConnectivityTlsProperty{
//   				enabled: jsii.Boolean(false),
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
	// `CfnCluster.ConnectivityInfoProperty.VpcConnectivity`.
	VpcConnectivity interface{} `field:"optional" json:"vpcConnectivity" yaml:"vpcConnectivity"`
}

