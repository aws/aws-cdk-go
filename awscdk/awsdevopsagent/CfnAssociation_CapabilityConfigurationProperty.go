package awsdevopsagent


// Configuration of a single capability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capabilityConfigurationProperty := &CapabilityConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	TriggerFilterGroups: []interface{}{
//   		&TriggerFilterGroupProperty{
//   			Events: []*string{
//   				jsii.String("events"),
//   			},
//   			TargetBranches: &PatternFilterProperty{
//   				Patterns: []*string{
//   					jsii.String("patterns"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-capabilityconfiguration.html
//
type CfnAssociation_CapabilityConfigurationProperty struct {
	// Whether the capability is enabled for this association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-capabilityconfiguration.html#cfn-devopsagent-association-capabilityconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Filter groups controlling when the capability triggers automatically.
	//
	// The capability triggers when any group matches. When omitted, all applicable events trigger the capability. Retained while the capability is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-capabilityconfiguration.html#cfn-devopsagent-association-capabilityconfiguration-triggerfiltergroups
	//
	TriggerFilterGroups interface{} `field:"optional" json:"triggerFilterGroups" yaml:"triggerFilterGroups"`
}

