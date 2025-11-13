package interfacesawsec2


// A reference to a TrafficMirrorSession resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficMirrorSessionReference := &TrafficMirrorSessionReference{
//   	TrafficMirrorSessionId: jsii.String("trafficMirrorSessionId"),
//   }
//
type TrafficMirrorSessionReference struct {
	// The Id of the TrafficMirrorSession resource.
	TrafficMirrorSessionId *string `field:"required" json:"trafficMirrorSessionId" yaml:"trafficMirrorSessionId"`
}

