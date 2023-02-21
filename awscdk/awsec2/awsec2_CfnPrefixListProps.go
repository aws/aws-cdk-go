package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPrefixList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPrefixListProps := &cfnPrefixListProps{
//   	addressFamily: jsii.String("addressFamily"),
//   	maxEntries: jsii.Number(123),
//   	prefixListName: jsii.String("prefixListName"),
//
//   	// the properties below are optional
//   	entries: []interface{}{
//   		&entryProperty{
//   			cidr: jsii.String("cidr"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPrefixListProps struct {
	// The IP address type.
	//
	// Valid Values: `IPv4` | `IPv6`.
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// The maximum number of entries for the prefix list.
	MaxEntries *float64 `field:"required" json:"maxEntries" yaml:"maxEntries"`
	// A name for the prefix list.
	//
	// Constraints: Up to 255 characters in length. The name cannot start with `com.amazonaws` .
	PrefixListName *string `field:"required" json:"prefixListName" yaml:"prefixListName"`
	// One or more entries for the prefix list.
	Entries interface{} `field:"optional" json:"entries" yaml:"entries"`
	// The tags for the prefix list.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

