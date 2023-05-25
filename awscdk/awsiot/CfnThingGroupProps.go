package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnThingGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnThingGroupProps := &CfnThingGroupProps{
//   	ParentGroupName: jsii.String("parentGroupName"),
//   	QueryString: jsii.String("queryString"),
//   	Tags: []cfnTag{
//   		&cfnTag{
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
type CfnThingGroupProps struct {
	// The parent thing group name.
	//
	// A Dynamic Thing Group does not have `parentGroupName` defined.
	ParentGroupName *string `field:"optional" json:"parentGroupName" yaml:"parentGroupName"`
	// The dynamic thing group search query string.
	//
	// The `queryString` attribute *is* required for `CreateDynamicThingGroup` . The `queryString` attribute *is not* required for `CreateThingGroup` .
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// Metadata which can be used to manage the thing group or dynamic thing group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The thing group name.
	ThingGroupName *string `field:"optional" json:"thingGroupName" yaml:"thingGroupName"`
	// Thing group properties.
	ThingGroupProperties interface{} `field:"optional" json:"thingGroupProperties" yaml:"thingGroupProperties"`
}

