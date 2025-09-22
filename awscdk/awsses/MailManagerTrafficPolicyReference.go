package awsses


// A reference to a MailManagerTrafficPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mailManagerTrafficPolicyReference := &MailManagerTrafficPolicyReference{
//   	TrafficPolicyId: jsii.String("trafficPolicyId"),
//   }
//
type MailManagerTrafficPolicyReference struct {
	// The TrafficPolicyId of the MailManagerTrafficPolicy resource.
	TrafficPolicyId *string `field:"required" json:"trafficPolicyId" yaml:"trafficPolicyId"`
}

