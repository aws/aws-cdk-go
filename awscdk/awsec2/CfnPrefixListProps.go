package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPrefixList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPrefixListProps := &CfnPrefixListProps{
//   	AddressFamily: jsii.String("addressFamily"),
//   	PrefixListName: jsii.String("prefixListName"),
//
//   	// the properties below are optional
//   	Entries: []interface{}{
//   		&EntryProperty{
//   			Cidr: jsii.String("cidr"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//   	MaxEntries: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html
//
type CfnPrefixListProps struct {
	// The IP address type.
	//
	// Valid Values: `IPv4` | `IPv6`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-addressfamily
	//
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// A name for the prefix list.
	//
	// Constraints: Up to 255 characters in length. The name cannot start with `com.amazonaws` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-prefixlistname
	//
	PrefixListName *string `field:"required" json:"prefixListName" yaml:"prefixListName"`
	// One or more entries for the prefix list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-entries
	//
	Entries interface{} `field:"optional" json:"entries" yaml:"entries"`
	// The maximum number of entries for the prefix list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-maxentries
	//
	MaxEntries *float64 `field:"optional" json:"maxEntries" yaml:"maxEntries"`
	// The tags for the prefix list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

