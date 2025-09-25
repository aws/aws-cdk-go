package awsroute53


// A reference to a HostedZone resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostedZoneReference := &HostedZoneReference{
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   }
//
type HostedZoneReference struct {
	// The Id of the HostedZone resource.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
}

