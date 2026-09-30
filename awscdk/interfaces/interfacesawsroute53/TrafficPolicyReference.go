package interfacesawsroute53


// A reference to a TrafficPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficPolicyReference := &TrafficPolicyReference{
//   	TrafficPolicyArn: jsii.String("trafficPolicyArn"),
//   }
//
type TrafficPolicyReference struct {
	// The Arn of the TrafficPolicy resource.
	TrafficPolicyArn *string `field:"required" json:"trafficPolicyArn" yaml:"trafficPolicyArn"`
}

