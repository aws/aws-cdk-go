package awsiot


// Parameters used when defining a mitigation action that move a set of things to a thing group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addThingsToThingGroupParamsProperty := &addThingsToThingGroupParamsProperty{
//   	thingGroupNames: []*string{
//   		jsii.String("thingGroupNames"),
//   	},
//
//   	// the properties below are optional
//   	overrideDynamicGroups: jsii.Boolean(false),
//   }
//
type CfnMitigationAction_AddThingsToThingGroupParamsProperty struct {
	// The list of groups to which you want to add the things that triggered the mitigation action.
	//
	// You can add a thing to a maximum of 10 groups, but you can't add a thing to more than one group in the same hierarchy.
	ThingGroupNames *[]*string `field:"required" json:"thingGroupNames" yaml:"thingGroupNames"`
	// Specifies if this mitigation action can move the things that triggered the mitigation action even if they are part of one or more dynamic thing groups.
	OverrideDynamicGroups interface{} `field:"optional" json:"overrideDynamicGroups" yaml:"overrideDynamicGroups"`
}

