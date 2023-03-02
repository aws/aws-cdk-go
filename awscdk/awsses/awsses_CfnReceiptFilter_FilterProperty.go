package awsses


// Specifies an IP address filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &filterProperty{
//   	ipFilter: &ipFilterProperty{
//   		cidr: jsii.String("cidr"),
//   		policy: jsii.String("policy"),
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnReceiptFilter_FilterProperty struct {
	// A structure that provides the IP addresses to block or allow, and whether to block or allow incoming mail from them.
	IpFilter interface{} `field:"required" json:"ipFilter" yaml:"ipFilter"`
	// The name of the IP address filter. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	// - Start and end with a letter or number.
	// - Contain 64 characters or fewer.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

