package awsmsk


// Details for SASL/IAM client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamProperty := &IamProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
type CfnServerlessCluster_IamProperty struct {
	// SASL/IAM authentication is enabled or not.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

