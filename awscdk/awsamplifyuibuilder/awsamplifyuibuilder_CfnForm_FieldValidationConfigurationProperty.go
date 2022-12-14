package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldValidationConfigurationProperty := &fieldValidationConfigurationProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	numValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   	strValues: []*string{
//   		jsii.String("strValues"),
//   	},
//   	validationMessage: jsii.String("validationMessage"),
//   }
//
type CfnForm_FieldValidationConfigurationProperty struct {
	// `CfnForm.FieldValidationConfigurationProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnForm.FieldValidationConfigurationProperty.NumValues`.
	NumValues interface{} `field:"optional" json:"numValues" yaml:"numValues"`
	// `CfnForm.FieldValidationConfigurationProperty.StrValues`.
	StrValues *[]*string `field:"optional" json:"strValues" yaml:"strValues"`
	// `CfnForm.FieldValidationConfigurationProperty.ValidationMessage`.
	ValidationMessage *string `field:"optional" json:"validationMessage" yaml:"validationMessage"`
}

