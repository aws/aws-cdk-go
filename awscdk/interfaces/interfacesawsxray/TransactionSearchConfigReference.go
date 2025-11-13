package interfacesawsxray


// A reference to a TransactionSearchConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transactionSearchConfigReference := &TransactionSearchConfigReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type TransactionSearchConfigReference struct {
	// The AccountId of the TransactionSearchConfig resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

