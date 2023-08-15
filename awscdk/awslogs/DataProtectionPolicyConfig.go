package awslogs


// Interface representing a data protection policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var statement interface{}
//
//   dataProtectionPolicyConfig := &DataProtectionPolicyConfig{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Statement: statement,
//   	Version: jsii.String("version"),
//   }
//
type DataProtectionPolicyConfig struct {
	// Description of the data protection policy.
	// Default: - 'cdk generated data protection policy'.
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name of the data protection policy.
	// Default: - 'data-protection-policy-cdk'.
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Statements within the data protection policy.
	//
	// Must contain one Audit and one Redact statement.
	Statement interface{} `field:"required" json:"statement" yaml:"statement"`
	// Version of the data protection policy.
	Version *string `field:"required" json:"version" yaml:"version"`
}

