package awsiot


// Parameters used when defining a mitigation action that move a set of things to a thing group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addThingsToThingGroupParamsProperty := &AddThingsToThingGroupParamsProperty{
//   	ThingGroupNames: []*string{
//   		jsii.String("thingGroupNames"),
//   	},
//
//   	// the properties below are optional
//   	OverrideDynamicGroups: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-addthingstothinggroupparams.html
//
type CfnMitigationAction_AddThingsToThingGroupParamsProperty struct {
	// The list of groups to which you want to add the things that triggered the mitigation action.
	//
	// You can add a thing to a maximum of 10 groups, but you can't add a thing to more than one group in the same hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-addthingstothinggroupparams.html#cfn-iot-mitigationaction-addthingstothinggroupparams-thinggroupnames
	//
	ThingGroupNames *[]*string `field:"required" json:"thingGroupNames" yaml:"thingGroupNames"`
	// Specifies if this mitigation action can move the things that triggered the mitigation action even if they are part of one or more dynamic thing groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-addthingstothinggroupparams.html#cfn-iot-mitigationaction-addthingstothinggroupparams-overridedynamicgroups
	//
	OverrideDynamicGroups interface{} `field:"optional" json:"overrideDynamicGroups" yaml:"overrideDynamicGroups"`
}

