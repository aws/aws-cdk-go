package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMailManagerAddressList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMailManagerAddressListProps := &CfnMailManagerAddressListProps{
//   	AddressListName: jsii.String("addressListName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageraddresslist.html
//
type CfnMailManagerAddressListProps struct {
	// A user-friendly name for the address list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageraddresslist.html#cfn-ses-mailmanageraddresslist-addresslistname
	//
	AddressListName *string `field:"optional" json:"addressListName" yaml:"addressListName"`
	// The tags used to organize, track, or control access for the resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageraddresslist.html#cfn-ses-mailmanageraddresslist-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

