package awsdevopsagent


// Automated AWS DevOps Agent capabilities for the repository in this association.
//
// Code change risk analysis analyzes pull requests and merge requests for deployment risk, and automated testing extends that analysis by running tests against the change. Capabilities are supported for GitHub and GitLab associations. When this property is set, it fully defines both capabilities below: a capability omitted here is removed from the association. Removing this property from a template that previously set it disables both capabilities. Capabilities configured outside CloudFormation are left unchanged only while this property has never been set. On rollback, both capabilities are restored to the previous template's values, not to settings made outside CloudFormation. Other capabilities on the association are not affected. Updates read and rewrite the association's full capability set, so a capability change made outside CloudFormation while a stack update is in progress may be overwritten.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   capabilitiesProperty := &CapabilitiesProperty{
//   	CodeChangeRiskAnalysis: &CapabilityConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		TriggerFilterGroups: []interface{}{
//   			&TriggerFilterGroupProperty{
//   				Events: []*string{
//   					jsii.String("events"),
//   				},
//   				TargetBranches: &PatternFilterProperty{
//   					Patterns: []*string{
//   						jsii.String("patterns"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	CodeChangeRiskAnalysisAutomatedTesting: &AutomatedTestingCapabilityConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-capabilities.html
//
type CfnAssociationPropsMixin_CapabilitiesProperty struct {
	// Configuration of a single capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-capabilities.html#cfn-devopsagent-association-capabilities-codechangeriskanalysis
	//
	CodeChangeRiskAnalysis interface{} `field:"optional" json:"codeChangeRiskAnalysis" yaml:"codeChangeRiskAnalysis"`
	// Configuration of the code change risk analysis automated testing capability.
	//
	// This capability does not support trigger filter groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-capabilities.html#cfn-devopsagent-association-capabilities-codechangeriskanalysisautomatedtesting
	//
	CodeChangeRiskAnalysisAutomatedTesting interface{} `field:"optional" json:"codeChangeRiskAnalysisAutomatedTesting" yaml:"codeChangeRiskAnalysisAutomatedTesting"`
}

