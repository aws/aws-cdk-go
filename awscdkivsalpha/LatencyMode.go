// The CDK Construct Library for AWS::IVS
package awscdkivsalpha


// Channel latency mode.
// Experimental.
type LatencyMode string

const (
	// Use LOW to minimize broadcaster-to-viewer latency for interactive broadcasts.
	// Experimental.
	LatencyMode_LOW LatencyMode = "LOW"
	// Use NORMAL for broadcasts that do not require viewer interaction.
	// Experimental.
	LatencyMode_NORMAL LatencyMode = "NORMAL"
)

