package previewawssesmixins


// Specifies an IP address filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   filterProperty := &FilterProperty{
//   	IpFilter: &IpFilterProperty{
//   		Cidr: jsii.String("cidr"),
//   		Policy: jsii.String("policy"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptfilter-filter.html
//
type CfnReceiptFilterPropsMixin_FilterProperty struct {
	// A structure that provides the IP addresses to block or allow, and whether to block or allow incoming mail from them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptfilter-filter.html#cfn-ses-receiptfilter-filter-ipfilter
	//
	IpFilter interface{} `field:"optional" json:"ipFilter" yaml:"ipFilter"`
	// The name of the IP address filter. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-).
	// - Start and end with a letter or number.
	// - Contain 64 characters or fewer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptfilter-filter.html#cfn-ses-receiptfilter-filter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

