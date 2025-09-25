package awsaps


// Use this structure to define label sets and the ingestion limits for time series that match label sets, and to specify the retention period of the workspace.
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
	// This is an array of structures, where each structure defines a label set for the workspace, and defines the ingestion limit for active time series for each of those label sets.
	//
	// Each label name in a label set must be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-workspaceconfiguration.html#cfn-aps-workspace-workspaceconfiguration-limitsperlabelsets
	//
	LimitsPerLabelSets interface{} `field:"optional" json:"limitsPerLabelSets" yaml:"limitsPerLabelSets"`
	// Specifies how many days that metrics will be retained in the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-workspaceconfiguration.html#cfn-aps-workspace-workspaceconfiguration-retentionperiodindays
	//
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
}

