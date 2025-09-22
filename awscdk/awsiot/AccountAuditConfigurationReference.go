package awsiot


// A reference to a AccountAuditConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountAuditConfigurationReference := &AccountAuditConfigurationReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type AccountAuditConfigurationReference struct {
	// The AccountId of the AccountAuditConfiguration resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

