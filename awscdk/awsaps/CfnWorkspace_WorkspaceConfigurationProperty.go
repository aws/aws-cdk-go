package awsaps


// Workspace configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workspaceConfigurationProperty := &WorkspaceConfigurationProperty{
//   	LimitsPerLabelSets: []interface{}{
//   		&LimitsPerLabelSetProperty{
//   			LabelSet: []interface{}{
//   				&LabelProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Limits: &LimitsPerLabelSetEntryProperty{
//   				MaxSeries: jsii.Number(123),
//   			},
//   		},
//   	},
//   	RetentionPeriodInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-workspaceconfiguration.html
//
type CfnWorkspace_WorkspaceConfigurationProperty struct {
	// An array of label set and associated limits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-workspaceconfiguration.html#cfn-aps-workspace-workspaceconfiguration-limitsperlabelsets
	//
	LimitsPerLabelSets interface{} `field:"optional" json:"limitsPerLabelSets" yaml:"limitsPerLabelSets"`
	// How many days that metrics are retained in the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-workspaceconfiguration.html#cfn-aps-workspace-workspaceconfiguration-retentionperiodindays
	//
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
}

