package awsiot


// Which audit checks are enabled and disabled for this account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auditCheckConfigurationProperty := &auditCheckConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnAccountAuditConfiguration_AuditCheckConfigurationProperty struct {
	// True if this audit check is enabled for this account.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

