package awsmsk


// Broker access controls.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-connectivityinfo.html
//
type CfnCluster_ConnectivityInfoProperty struct {
	// Access control settings for the cluster's brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-connectivityinfo.html#cfn-msk-cluster-connectivityinfo-publicaccess
	//
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// VPC connection control settings for brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-connectivityinfo.html#cfn-msk-cluster-connectivityinfo-vpcconnectivity
	//
	VpcConnectivity interface{} `field:"optional" json:"vpcConnectivity" yaml:"vpcConnectivity"`
}

