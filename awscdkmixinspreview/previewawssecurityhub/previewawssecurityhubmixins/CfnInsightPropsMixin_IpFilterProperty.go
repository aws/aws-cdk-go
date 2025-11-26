package previewawssecurityhubmixins


// The IP filter for querying findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ipFilterProperty := &IpFilterProperty{
//   	Cidr: jsii.String("cidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-ipfilter.html
//
type CfnInsightPropsMixin_IpFilterProperty struct {
	// A finding's CIDR value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-insight-ipfilter.html#cfn-securityhub-insight-ipfilter-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

