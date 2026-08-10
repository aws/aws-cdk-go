package awsmediaconnectalpha


// Properties for the flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   var flow Flow
//   var sourceConfiguration SourceConfiguration
//
//   flowSourceProps := &FlowSourceProps{
//   	Flow: flow,
//   	Source: sourceConfiguration,
//   }
//
// Experimental.
type FlowSourceProps struct {
	// The flow this source is connected to.
	//
	// The flow must have Failover enabled to add an additional source.
	// Experimental.
	Flow IFlow `field:"required" json:"flow" yaml:"flow"`
	// Flow Source Configuration.
	// Experimental.
	Source SourceConfiguration `field:"required" json:"source" yaml:"source"`
}

