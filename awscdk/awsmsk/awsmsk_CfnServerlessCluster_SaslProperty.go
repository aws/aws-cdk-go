package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   saslProperty := &saslProperty{
//   	iam: &iamProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnServerlessCluster_SaslProperty struct {
	// `CfnServerlessCluster.SaslProperty.Iam`.
	Iam interface{} `field:"required" json:"iam" yaml:"iam"`
}

