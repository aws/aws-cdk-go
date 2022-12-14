package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectivityProperty := &vpcConnectivityProperty{
//   	clientAuthentication: &vpcConnectivityClientAuthenticationProperty{
//   		sasl: &vpcConnectivitySaslProperty{
//   			iam: &vpcConnectivityIamProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   			scram: &vpcConnectivityScramProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   		},
//   		tls: &vpcConnectivityTlsProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnCluster_VpcConnectivityProperty struct {
	// `CfnCluster.VpcConnectivityProperty.ClientAuthentication`.
	ClientAuthentication interface{} `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
}

