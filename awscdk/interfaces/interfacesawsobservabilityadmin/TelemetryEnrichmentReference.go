package interfacesawsobservabilityadmin


// A reference to a TelemetryEnrichment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telemetryEnrichmentReference := &TelemetryEnrichmentReference{
//   	Scope: jsii.String("scope"),
//   }
//
type TelemetryEnrichmentReference struct {
	// The Scope of the TelemetryEnrichment resource.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

