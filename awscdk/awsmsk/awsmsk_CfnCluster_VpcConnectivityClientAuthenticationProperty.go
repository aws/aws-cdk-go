package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectivityClientAuthenticationProperty := &vpcConnectivityClientAuthenticationProperty{
//   	sasl: &vpcConnectivitySaslProperty{
//   		iam: &vpcConnectivityIamProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   		scram: &vpcConnectivityScramProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	tls: &vpcConnectivityTlsProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnCluster_VpcConnectivityClientAuthenticationProperty struct {
	// `CfnCluster.VpcConnectivityClientAuthenticationProperty.Sasl`.
	Sasl interface{} `field:"optional" json:"sasl" yaml:"sasl"`
	// `CfnCluster.VpcConnectivityClientAuthenticationProperty.Tls`.
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

