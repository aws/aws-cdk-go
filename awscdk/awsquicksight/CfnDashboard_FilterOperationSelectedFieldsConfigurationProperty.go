package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterOperationSelectedFieldsConfigurationProperty := &FilterOperationSelectedFieldsConfigurationProperty{
//   	SelectedFieldOptions: jsii.String("selectedFieldOptions"),
//   	SelectedFields: []*string{
//   		jsii.String("selectedFields"),
//   	},
//   }
//
type CfnDashboard_FilterOperationSelectedFieldsConfigurationProperty struct {
	// `CfnDashboard.FilterOperationSelectedFieldsConfigurationProperty.SelectedFieldOptions`.
	SelectedFieldOptions *string `field:"optional" json:"selectedFieldOptions" yaml:"selectedFieldOptions"`
	// `CfnDashboard.FilterOperationSelectedFieldsConfigurationProperty.SelectedFields`.
	SelectedFields *[]*string `field:"optional" json:"selectedFields" yaml:"selectedFields"`
}

