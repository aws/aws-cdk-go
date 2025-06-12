package awscodepipelineactions


// The CodePipeline variables emitted by the InspectorScan Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inspectorScanVariables := &InspectorScanVariables{
//   	HighestScannedSeverity: jsii.String("highestScannedSeverity"),
//   }
//
type InspectorScanVariables struct {
	// The highest severity output from the scan.
	//
	// Valid values are medium | high | critical.
	HighestScannedSeverity *string `field:"required" json:"highestScannedSeverity" yaml:"highestScannedSeverity"`
}

