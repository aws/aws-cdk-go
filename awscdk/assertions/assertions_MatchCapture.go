package assertions


// Information about a value captured during match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var capture capture
//   var value interface{}
//
//   matchCapture := &matchCapture{
//   	capture: capture,
//   	value: value,
//   }
//
type MatchCapture struct {
	// The instance of Capture class to which this capture is associated with.
	Capture Capture `field:"required" json:"capture" yaml:"capture"`
	// The value that was captured.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

