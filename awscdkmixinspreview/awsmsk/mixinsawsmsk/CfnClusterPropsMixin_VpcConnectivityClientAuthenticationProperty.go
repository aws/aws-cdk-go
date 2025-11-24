package mixinsawsmsk


// Includes all client authentication information for VpcConnectivity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnClusterPropsMixin_VpcConnectivityClientAuthenticationProperty struct {
	// Details for VpcConnectivity ClientAuthentication using SASL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivityclientauthentication.html#cfn-msk-cluster-vpcconnectivityclientauthentication-sasl
	//
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// Details for VpcConnectivity ClientAuthentication using TLS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-vpcconnectivityclientauthentication.html#cfn-msk-cluster-vpcconnectivityclientauthentication-tls
	//
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

