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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html
//
type CfnForm_FieldValidationConfigurationProperty struct {
	// The validation to perform on an object type.
	//
	// ``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html#cfn-amplifyuibuilder-form-fieldvalidationconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The validation to perform on a number value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html#cfn-amplifyuibuilder-form-fieldvalidationconfiguration-numvalues
	//
	NumValues interface{} `field:"optional" json:"numValues" yaml:"numValues"`
	// The validation to perform on a string value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html#cfn-amplifyuibuilder-form-fieldvalidationconfiguration-strvalues
	//
	StrValues *[]*string `field:"optional" json:"strValues" yaml:"strValues"`
	// The validation message to display.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html#cfn-amplifyuibuilder-form-fieldvalidationconfiguration-validationmessage
	//
	ValidationMessage *string `field:"optional" json:"validationMessage" yaml:"validationMessage"`
}

