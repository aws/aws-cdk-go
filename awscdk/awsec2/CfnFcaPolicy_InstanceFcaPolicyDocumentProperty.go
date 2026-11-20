package awsec2


// The FCA policy document for EC2 instance cost allocation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFcaPolicyDocumentProperty := &InstanceFcaPolicyDocumentProperty{
//   	Dimensions: []*string{
//   		jsii.String("dimensions"),
//   	},
//   	ReportingGranularity: []*string{
//   		jsii.String("reportingGranularity"),
//   	},
//
//   	// the properties below are optional
//   	CidrRangeSpecifications: &CidrRangeSpecificationsProperty{
//   		Prefixes: []interface{}{
//   			&CidrMappingProperty{
//   				CidrLabel: jsii.String("cidrLabel"),
//   				CidrRange: jsii.String("cidrRange"),
//   			},
//   		},
//   		PrefixesType: jsii.String("prefixesType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fcapolicy-instancefcapolicydocument.html
//
type CfnFcaPolicy_InstanceFcaPolicyDocumentProperty struct {
	// The cost dimensions to allocate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fcapolicy-instancefcapolicydocument.html#cfn-ec2-fcapolicy-instancefcapolicydocument-dimensions
	//
	Dimensions *[]*string `field:"required" json:"dimensions" yaml:"dimensions"`
	// The granularity at which costs are reported.
	//
	// Contains exactly one value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fcapolicy-instancefcapolicydocument.html#cfn-ec2-fcapolicy-instancefcapolicydocument-reportinggranularity
	//
	ReportingGranularity *[]*string `field:"required" json:"reportingGranularity" yaml:"reportingGranularity"`
	// CIDR range specifications, used when the reporting granularity is 'cidr'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-fcapolicy-instancefcapolicydocument.html#cfn-ec2-fcapolicy-instancefcapolicydocument-cidrrangespecifications
	//
	CidrRangeSpecifications interface{} `field:"optional" json:"cidrRangeSpecifications" yaml:"cidrRangeSpecifications"`
}

