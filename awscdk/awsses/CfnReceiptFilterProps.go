package awsses


// Properties for defining a `CfnReceiptFilter`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReceiptFilterProps := &CfnReceiptFilterProps{
//   	Filter: &FilterProperty{
//   		IpFilter: &IpFilterProperty{
//   			Cidr: jsii.String("cidr"),
//   			Policy: jsii.String("policy"),
//   		},
//
//   		// the properties below are optional
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptfilter.html
//
type CfnReceiptFilterProps struct {
	// A data structure that describes the IP address filter to create, which consists of a name, an IP address range, and whether to allow or block mail from it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptfilter.html#cfn-ses-receiptfilter-filter
	//
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
}

