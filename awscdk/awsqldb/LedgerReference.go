package awsqldb


// A reference to a Ledger resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ledgerReference := &LedgerReference{
//   	LedgerId: jsii.String("ledgerId"),
//   }
//
type LedgerReference struct {
	// The Id of the Ledger resource.
	LedgerId *string `field:"required" json:"ledgerId" yaml:"ledgerId"`
}

