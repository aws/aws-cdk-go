package interfacesawsec2


// A reference to a TrafficMirrorFilter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficMirrorFilterReference := &TrafficMirrorFilterReference{
//   	TrafficMirrorFilterId: jsii.String("trafficMirrorFilterId"),
//   }
//
type TrafficMirrorFilterReference struct {
	// The Id of the TrafficMirrorFilter resource.
	TrafficMirrorFilterId *string `field:"required" json:"trafficMirrorFilterId" yaml:"trafficMirrorFilterId"`
}

