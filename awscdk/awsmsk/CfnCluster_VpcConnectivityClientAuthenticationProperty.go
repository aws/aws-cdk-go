package awsmsk


// Not currently supported by AWS CloudFormation .
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
	// `CfnCluster.VpcConnectivityClientAuthenticationProperty.Sasl`.
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// `CfnCluster.VpcConnectivityClientAuthenticationProperty.Tls`.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

