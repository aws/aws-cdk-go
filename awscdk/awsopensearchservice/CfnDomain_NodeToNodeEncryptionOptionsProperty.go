package awsopensearchservice


// Specifies options for node-to-node encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeToNodeEncryptionOptionsProperty := &NodeToNodeEncryptionOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodetonodeencryptionoptions.html
//
type CfnDomain_NodeToNodeEncryptionOptionsProperty struct {
	// Specifies to enable or disable node-to-node encryption on the domain.
	//
	// Required if you enable fine-grained access control in [AdvancedSecurityOptionsInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodetonodeencryptionoptions.html#cfn-opensearchservice-domain-nodetonodeencryptionoptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

