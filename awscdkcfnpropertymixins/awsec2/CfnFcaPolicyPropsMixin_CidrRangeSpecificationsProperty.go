package awsec2


// CIDR range specifications, used when the reporting granularity is 'cidr'.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cidrRangeSpecificationsProperty := &CidrRangeSpecificationsProperty{
//   	Prefixes: []interface{}{
//   		&CidrMappingProperty{
//   			CidrLabel: jsii.String("cidrLabel"),
//   			CidrRange: jsii.String("cidrRange"),
//   		},
//   	},
//   	PrefixesType: jsii.String("prefixesType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fcapolicy-cidrrangespecifications.html
//
type CfnFcaPolicyPropsMixin_CidrRangeSpecificationsProperty struct {
	// The CIDR range to label mappings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fcapolicy-cidrrangespecifications.html#cfn-ec2-fcapolicy-cidrrangespecifications-prefixes
	//
	Prefixes interface{} `field:"optional" json:"prefixes" yaml:"prefixes"`
	// The type of prefixes the mappings apply to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fcapolicy-cidrrangespecifications.html#cfn-ec2-fcapolicy-cidrrangespecifications-prefixestype
	//
	PrefixesType *string `field:"optional" json:"prefixesType" yaml:"prefixesType"`
}

