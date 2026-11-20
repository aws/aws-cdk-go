package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFcaPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnFcaPolicyMixinProps := &CfnFcaPolicyMixinProps{
//   	Description: jsii.String("description"),
//   	InstanceFcaPolicyDocument: &InstanceFcaPolicyDocumentProperty{
//   		CidrRangeSpecifications: &CidrRangeSpecificationsProperty{
//   			Prefixes: []interface{}{
//   				&CidrMappingProperty{
//   					CidrLabel: jsii.String("cidrLabel"),
//   					CidrRange: jsii.String("cidrRange"),
//   				},
//   			},
//   			PrefixesType: jsii.String("prefixesType"),
//   		},
//   		Dimensions: []*string{
//   			jsii.String("dimensions"),
//   		},
//   		ReportingGranularity: []*string{
//   			jsii.String("reportingGranularity"),
//   		},
//   	},
//   	NatFcaPolicyDocument: &NatFcaPolicyDocumentProperty{
//   		CidrRangeSpecifications: &CidrRangeSpecificationsProperty{
//   			Prefixes: []interface{}{
//   				&CidrMappingProperty{
//   					CidrLabel: jsii.String("cidrLabel"),
//   					CidrRange: jsii.String("cidrRange"),
//   				},
//   			},
//   			PrefixesType: jsii.String("prefixesType"),
//   		},
//   		Dimensions: []*string{
//   			jsii.String("dimensions"),
//   		},
//   		ReportingGranularity: []*string{
//   			jsii.String("reportingGranularity"),
//   		},
//   	},
//   	NetworkResourceType: jsii.String("networkResourceType"),
//   	PcxFcaPolicyDocument: &PcxFcaPolicyDocumentProperty{
//   		CidrRangeSpecifications: &CidrRangeSpecificationsProperty{
//   			Prefixes: []interface{}{
//   				&CidrMappingProperty{
//   					CidrLabel: jsii.String("cidrLabel"),
//   					CidrRange: jsii.String("cidrRange"),
//   				},
//   			},
//   			PrefixesType: jsii.String("prefixesType"),
//   		},
//   		Dimensions: []*string{
//   			jsii.String("dimensions"),
//   		},
//   		ReportingGranularity: []*string{
//   			jsii.String("reportingGranularity"),
//   		},
//   	},
//   	ShowBackEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicy.html
//
type CfnFcaPolicyMixinProps struct {
	// The description of the FCA policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicy.html#cfn-ec2-fcapolicy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The FCA policy document for EC2 instance cost allocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicy.html#cfn-ec2-fcapolicy-instancefcapolicydocument
	//
	InstanceFcaPolicyDocument interface{} `field:"optional" json:"instanceFcaPolicyDocument" yaml:"instanceFcaPolicyDocument"`
	// The FCA policy document for NAT gateway cost allocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicy.html#cfn-ec2-fcapolicy-natfcapolicydocument
	//
	NatFcaPolicyDocument interface{} `field:"optional" json:"natFcaPolicyDocument" yaml:"natFcaPolicyDocument"`
	// The type of network resource the FCA policy applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicy.html#cfn-ec2-fcapolicy-networkresourcetype
	//
	NetworkResourceType *string `field:"optional" json:"networkResourceType" yaml:"networkResourceType"`
	// The FCA policy document for VPC peering connection cost allocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicy.html#cfn-ec2-fcapolicy-pcxfcapolicydocument
	//
	PcxFcaPolicyDocument interface{} `field:"optional" json:"pcxFcaPolicyDocument" yaml:"pcxFcaPolicyDocument"`
	// Indicates whether show-back reporting is enabled for the FCA policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicy.html#cfn-ec2-fcapolicy-showbackenabled
	//
	ShowBackEnabled interface{} `field:"optional" json:"showBackEnabled" yaml:"showBackEnabled"`
	// The tags for the FCA policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicy.html#cfn-ec2-fcapolicy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

