package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientAuthenticationProperty := &clientAuthenticationProperty{
//   	sasl: &saslProperty{
//   		iam: &iamProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnServerlessCluster_ClientAuthenticationProperty struct {
	// `CfnServerlessCluster.ClientAuthenticationProperty.Sasl`.
	Sasl interface{} `field:"required" json:"sasl" yaml:"sasl"`
}

