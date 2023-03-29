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
type CfnTemplate_FilterOperationSelectedFieldsConfigurationProperty struct {
	// `CfnTemplate.FilterOperationSelectedFieldsConfigurationProperty.SelectedFieldOptions`.
	SelectedFieldOptions *string `field:"optional" json:"selectedFieldOptions" yaml:"selectedFieldOptions"`
	// `CfnTemplate.FilterOperationSelectedFieldsConfigurationProperty.SelectedFields`.
	SelectedFields *[]*string `field:"optional" json:"selectedFields" yaml:"selectedFields"`
}

