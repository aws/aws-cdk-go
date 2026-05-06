package interfacesawscloudwatch


// A reference to a OTelEnrichment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oTelEnrichmentReference := &OTelEnrichmentReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type OTelEnrichmentReference struct {
	// The AccountId of the OTelEnrichment resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

