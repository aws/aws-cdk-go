package awsec2


// A reference to a TrafficMirrorTarget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficMirrorTargetReference := &TrafficMirrorTargetReference{
//   	TrafficMirrorTargetId: jsii.String("trafficMirrorTargetId"),
//   }
//
type TrafficMirrorTargetReference struct {
	// The Id of the TrafficMirrorTarget resource.
	TrafficMirrorTargetId *string `field:"required" json:"trafficMirrorTargetId" yaml:"trafficMirrorTargetId"`
}

