package awsopensearchservice


// Specifies options for node-to-node encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeToNodeEncryptionOptionsProperty := &nodeToNodeEncryptionOptionsProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnDomain_NodeToNodeEncryptionOptionsProperty struct {
	// Specifies to enable or disable node-to-node encryption on the domain.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

