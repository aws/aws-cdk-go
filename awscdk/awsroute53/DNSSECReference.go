package awsroute53


// A reference to a DNSSEC resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dNSSECReference := &DNSSECReference{
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   }
//
type DNSSECReference struct {
	// The HostedZoneId of the DNSSEC resource.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
}

