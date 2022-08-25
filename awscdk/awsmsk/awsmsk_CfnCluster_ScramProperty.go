package awsmsk


// Details for SASL/SCRAM client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scramProperty := &scramProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnCluster_ScramProperty struct {
	// SASL/SCRAM authentication is enabled or not.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

