package awsdirectconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnResiliencyGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResiliencyGroupProps := &CfnResiliencyGroupProps{
//   	ResiliencyGroupName: jsii.String("resiliencyGroupName"),
//
//   	// the properties below are optional
//   	ConnectionArns: []interface{}{
//   		jsii.String("connectionArns"),
//   	},
//   	IntendedResiliencyModel: jsii.String("intendedResiliencyModel"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-resiliencygroup.html
//
type CfnResiliencyGroupProps struct {
	// The name of the resiliency group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-resiliencygroup.html#cfn-directconnect-resiliencygroup-resiliencygroupname
	//
	ResiliencyGroupName *string `field:"required" json:"resiliencyGroupName" yaml:"resiliencyGroupName"`
	// The ARNs of the connections that are members of the resiliency group.
	//
	// When specified, this is the complete, authoritative membership list: on update, connections not present are disassociated, and an explicit empty list removes all members. Omit the property to leave the existing membership unmanaged (unchanged on update). On create, an omitted or empty list means no connections are associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-resiliencygroup.html#cfn-directconnect-resiliencygroup-connectionarns
	//
	ConnectionArns *[]interface{} `field:"optional" json:"connectionArns" yaml:"connectionArns"`
	// The resiliency model that the resiliency group is intended to meet.
	//
	// Required when creating a resiliency group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-resiliencygroup.html#cfn-directconnect-resiliencygroup-intendedresiliencymodel
	//
	IntendedResiliencyModel *string `field:"optional" json:"intendedResiliencyModel" yaml:"intendedResiliencyModel"`
	// The tags associated with the resiliency group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-resiliencygroup.html#cfn-directconnect-resiliencygroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

