package awsfms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProtocolsList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProtocolsListProps := &CfnProtocolsListProps{
//   	ListName: jsii.String("listName"),
//   	ProtocolsList: []*string{
//   		jsii.String("protocolsList"),
//   	},
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-protocolslist.html
//
type CfnProtocolsListProps struct {
	// The name of the Firewall Manager protocols list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-protocolslist.html#cfn-fms-protocolslist-listname
	//
	ListName *string `field:"required" json:"listName" yaml:"listName"`
	// An array of protocols in the Firewall Manager protocols list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-protocolslist.html#cfn-fms-protocolslist-protocolslist
	//
	ProtocolsList *[]*string `field:"required" json:"protocolsList" yaml:"protocolsList"`
	// An array of key-value pairs to apply to the protocols list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fms-protocolslist.html#cfn-fms-protocolslist-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

