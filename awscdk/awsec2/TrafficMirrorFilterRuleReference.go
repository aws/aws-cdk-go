package awsec2


// A reference to a TrafficMirrorFilterRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficMirrorFilterRuleReference := &TrafficMirrorFilterRuleReference{
//   	TrafficMirrorFilterRuleId: jsii.String("trafficMirrorFilterRuleId"),
//   }
//
type TrafficMirrorFilterRuleReference struct {
	// The TrafficMirrorFilterRuleId of the TrafficMirrorFilterRule resource.
	TrafficMirrorFilterRuleId *string `field:"required" json:"trafficMirrorFilterRuleId" yaml:"trafficMirrorFilterRuleId"`
}

