package awsamplifyuibuilder


// The `FieldValidationConfiguration` property specifies the validation configuration for a field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldValidationConfigurationProperty := &FieldValidationConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	NumValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   	StrValues: []*string{
//   		jsii.String("strValues"),
//   	},
//   	ValidationMessage: jsii.String("validationMessage"),
//   }
//
type CfnForm_FieldValidationConfigurationProperty struct {
	// The validation to perform on an object type.
	//
	// ``.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The validation to perform on a number value.
	NumValues interface{} `field:"optional" json:"numValues" yaml:"numValues"`
	// The validation to perform on a string value.
	StrValues *[]*string `field:"optional" json:"strValues" yaml:"strValues"`
	// The validation message to display.
	ValidationMessage *string `field:"optional" json:"validationMessage" yaml:"validationMessage"`
}

