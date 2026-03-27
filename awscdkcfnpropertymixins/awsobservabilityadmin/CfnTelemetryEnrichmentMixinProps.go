package awsobservabilityadmin


// Properties for CfnTelemetryEnrichmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTelemetryEnrichmentMixinProps := &CfnTelemetryEnrichmentMixinProps{
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetryenrichment.html
//
type CfnTelemetryEnrichmentMixinProps struct {
	// Scope of the Telemetry Enrichment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-telemetryenrichment.html#cfn-observabilityadmin-telemetryenrichment-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

