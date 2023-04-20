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
type CfnAnalysis_FilterOperationSelectedFieldsConfigurationProperty struct {
	// `CfnAnalysis.FilterOperationSelectedFieldsConfigurationProperty.SelectedFieldOptions`.
	SelectedFieldOptions *string `field:"optional" json:"selectedFieldOptions" yaml:"selectedFieldOptions"`
	// `CfnAnalysis.FilterOperationSelectedFieldsConfigurationProperty.SelectedFields`.
	SelectedFields *[]*string `field:"optional" json:"selectedFields" yaml:"selectedFields"`
}

