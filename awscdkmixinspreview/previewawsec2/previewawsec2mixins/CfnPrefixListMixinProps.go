package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPrefixListPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPrefixListMixinProps := &CfnPrefixListMixinProps{
//   	AddressFamily: jsii.String("addressFamily"),
//   	Entries: []interface{}{
//   		&EntryProperty{
//   			Cidr: jsii.String("cidr"),
//   			Description: jsii.String("description"),
//   		},
//   	},
//   	MaxEntries: jsii.Number(123),
//   	PrefixListName: jsii.String("prefixListName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html
//
type CfnPrefixListMixinProps struct {
	// The IP address type.
	//
	// Valid Values: `IPv4` | `IPv6`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-addressfamily
	//
	AddressFamily *string `field:"optional" json:"addressFamily" yaml:"addressFamily"`
	// The entries for the prefix list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-entries
	//
	Entries interface{} `field:"optional" json:"entries" yaml:"entries"`
	// The maximum number of entries for the prefix list.
	//
	// You can't modify the entries and the size of a prefix list at the same time.
	//
	// This property is required when you create a prefix list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-maxentries
	//
	MaxEntries *float64 `field:"optional" json:"maxEntries" yaml:"maxEntries"`
	// A name for the prefix list.
	//
	// Constraints: Up to 255 characters in length. The name cannot start with `com.amazonaws` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-prefixlistname
	//
	PrefixListName *string `field:"optional" json:"prefixListName" yaml:"prefixListName"`
	// The tags for the prefix list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-prefixlist.html#cfn-ec2-prefixlist-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

