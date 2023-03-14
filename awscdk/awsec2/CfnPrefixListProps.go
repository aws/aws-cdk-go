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
//   	MaxEntries: jsii.Number(123),
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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

