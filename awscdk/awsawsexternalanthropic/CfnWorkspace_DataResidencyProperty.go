package awsawsexternalanthropic


// Data residency configuration for the workspace.
//
// WorkspaceGeo is immutable after creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataResidencyProperty := &DataResidencyProperty{
//   	AllowedInferenceGeos: []*string{
//   		jsii.String("allowedInferenceGeos"),
//   	},
//   	DefaultInferenceGeo: jsii.String("defaultInferenceGeo"),
//   	WorkspaceGeo: jsii.String("workspaceGeo"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-awsexternalanthropic-workspace-dataresidency.html
//
type CfnWorkspace_DataResidencyProperty struct {
	// Permitted inference geo values.
	//
	// Omit to allow all geos (the service default of 'unrestricted'); otherwise list specific geos.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-awsexternalanthropic-workspace-dataresidency.html#cfn-awsexternalanthropic-workspace-dataresidency-allowedinferencegeos
	//
	AllowedInferenceGeos *[]*string `field:"optional" json:"allowedInferenceGeos" yaml:"allowedInferenceGeos"`
	// Default inference geo applied when requests omit the parameter.
	//
	// Defaults to 'global' if omitted. Must be a member of AllowedInferenceGeos unless AllowedInferenceGeos is omitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-awsexternalanthropic-workspace-dataresidency.html#cfn-awsexternalanthropic-workspace-dataresidency-defaultinferencegeo
	//
	DefaultInferenceGeo *string `field:"optional" json:"defaultInferenceGeo" yaml:"defaultInferenceGeo"`
	// Geographic region for workspace data storage.
	//
	// Immutable after creation. Defaults to 'us' if omitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-awsexternalanthropic-workspace-dataresidency.html#cfn-awsexternalanthropic-workspace-dataresidency-workspacegeo
	//
	WorkspaceGeo *string `field:"optional" json:"workspaceGeo" yaml:"workspaceGeo"`
}

