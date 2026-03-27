package interfacesawsroute53globalresolver


// A reference to a DnsView resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dnsViewReference := &DnsViewReference{
//   	DnsViewArn: jsii.String("dnsViewArn"),
//   	DnsViewId: jsii.String("dnsViewId"),
//   }
//
type DnsViewReference struct {
	// The ARN of the DnsView resource.
	DnsViewArn *string `field:"required" json:"dnsViewArn" yaml:"dnsViewArn"`
	// The DnsViewId of the DnsView resource.
	DnsViewId *string `field:"required" json:"dnsViewId" yaml:"dnsViewId"`
}

