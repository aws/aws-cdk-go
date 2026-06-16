package awsobservabilityadmin


// Properties for defining a `CfnTelemetryEnrichment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTelemetryEnrichmentProps := &CfnTelemetryEnrichmentProps{
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetryenrichment.html
//
type CfnTelemetryEnrichmentProps struct {
	// Scope of the Telemetry Enrichment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetryenrichment.html#cfn-observabilityadmin-telemetryenrichment-scope
	//
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

