package awsses


// A receipt IP address filter enables you to specify whether to accept or reject mail originating from an IP address or range of IP addresses.
//
// For information about setting up IP address filters, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/receiving-email-ip-filtering-console-walkthrough.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipFilterProperty := &ipFilterProperty{
//   	cidr: jsii.String("cidr"),
//   	policy: jsii.String("policy"),
//   }
//
type CfnReceiptFilter_IpFilterProperty struct {
	// A single IP address or a range of IP addresses to block or allow, specified in Classless Inter-Domain Routing (CIDR) notation.
	//
	// An example of a single email address is 10.0.0.1. An example of a range of IP addresses is 10.0.0.1/24. For more information about CIDR notation, see [RFC 2317](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc2317) .
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Indicates whether to block or allow incoming mail from the specified IP addresses.
	Policy *string `field:"required" json:"policy" yaml:"policy"`
}

