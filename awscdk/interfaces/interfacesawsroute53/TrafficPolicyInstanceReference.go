package interfacesawsroute53


// A reference to a TrafficPolicyInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficPolicyInstanceReference := &TrafficPolicyInstanceReference{
//   	TrafficPolicyInstanceArn: jsii.String("trafficPolicyInstanceArn"),
//   }
//
type TrafficPolicyInstanceReference struct {
	// The Arn of the TrafficPolicyInstance resource.
	TrafficPolicyInstanceArn *string `field:"required" json:"trafficPolicyInstanceArn" yaml:"trafficPolicyInstanceArn"`
}

