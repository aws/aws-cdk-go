package previewawsiotmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnThingGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnThingGroupMixinProps := &CfnThingGroupMixinProps{
//   	ParentGroupName: jsii.String("parentGroupName"),
//   	QueryString: jsii.String("queryString"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThingGroupName: jsii.String("thingGroupName"),
//   	ThingGroupProperties: &ThingGroupPropertiesProperty{
//   		AttributePayload: &AttributePayloadProperty{
//   			Attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   		},
//   		ThingGroupDescription: jsii.String("thingGroupDescription"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html
//
type CfnThingGroupMixinProps struct {
	// The parent thing group name.
	//
	// A Dynamic Thing Group does not have `parentGroupName` defined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html#cfn-iot-thinggroup-parentgroupname
	//
	ParentGroupName *string `field:"optional" json:"parentGroupName" yaml:"parentGroupName"`
	// The dynamic thing group search query string.
	//
	// The `queryString` attribute *is* required for `CreateDynamicThingGroup` . The `queryString` attribute *is not* required for `CreateThingGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html#cfn-iot-thinggroup-querystring
	//
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// Metadata which can be used to manage the thing group or dynamic thing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html#cfn-iot-thinggroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The thing group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html#cfn-iot-thinggroup-thinggroupname
	//
	ThingGroupName *string `field:"optional" json:"thingGroupName" yaml:"thingGroupName"`
	// Thing group properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thinggroup.html#cfn-iot-thinggroup-thinggroupproperties
	//
	ThingGroupProperties interface{} `field:"optional" json:"thingGroupProperties" yaml:"thingGroupProperties"`
}

