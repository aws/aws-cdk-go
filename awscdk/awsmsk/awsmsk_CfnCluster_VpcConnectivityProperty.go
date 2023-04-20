package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectivityProperty := &VpcConnectivityProperty{
//   	ClientAuthentication: &VpcConnectivityClientAuthenticationProperty{
//   		Sasl: &VpcConnectivitySaslProperty{
//   			Iam: &VpcConnectivityIamProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			Scram: &VpcConnectivityScramProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   		Tls: &VpcConnectivityTlsProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnCluster_VpcConnectivityProperty struct {
	// Not currently supported by AWS CloudFormation .
	ClientAuthentication interface{} `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
}

