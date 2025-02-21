package awsmsk


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivityclientauthentication.html
//
type CfnCluster_VpcConnectivityClientAuthenticationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivityclientauthentication.html#cfn-msk-cluster-vpcconnectivityclientauthentication-sasl
	//
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivityclientauthentication.html#cfn-msk-cluster-vpcconnectivityclientauthentication-tls
	//
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

