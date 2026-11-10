package awsdirectconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnResiliencyGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnResiliencyGroupMixinProps := &CfnResiliencyGroupMixinProps{
//   	ConnectionArns: []interface{}{
//   		jsii.String("connectionArns"),
//   	},
//   	IntendedResiliencyModel: jsii.String("intendedResiliencyModel"),
//   	ResiliencyGroupName: jsii.String("resiliencyGroupName"),
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
type CfnResiliencyGroupMixinProps struct {
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
	// The name of the resiliency group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-resiliencygroup.html#cfn-directconnect-resiliencygroup-resiliencygroupname
	//
	ResiliencyGroupName *string `field:"optional" json:"resiliencyGroupName" yaml:"resiliencyGroupName"`
	// The tags associated with the resiliency group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-resiliencygroup.html#cfn-directconnect-resiliencygroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

