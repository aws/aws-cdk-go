package awsdevopsagent


// A group of trigger conditions.
//
// The group matches when all present conditions pass. At least one condition must be present.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   triggerFilterGroupProperty := &TriggerFilterGroupProperty{
//   	Events: []*string{
//   		jsii.String("events"),
//   	},
//   	TargetBranches: &PatternFilterProperty{
//   		Patterns: []*string{
//   			jsii.String("patterns"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-triggerfiltergroup.html
//
type CfnAssociationPropsMixin_TriggerFilterGroupProperty struct {
	// Passes when the webhook event is one of the listed events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-triggerfiltergroup.html#cfn-devopsagent-association-triggerfiltergroup-events
	//
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// A regex-based match condition.
	//
	// Passes when the value matches any pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-triggerfiltergroup.html#cfn-devopsagent-association-triggerfiltergroup-targetbranches
	//
	TargetBranches interface{} `field:"optional" json:"targetBranches" yaml:"targetBranches"`
}

