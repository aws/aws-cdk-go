package awsamplifyuibuilder


// The `FormInputValueProperty` property specifies the configuration for an input field on a form.
//
// Use `FormInputValueProperty` to specify the values to render or bind by default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formInputValuePropertyProperty := &FormInputValuePropertyProperty{
//   	Value: jsii.String("value"),
//   }
//
type CfnForm_FormInputValuePropertyProperty struct {
	// The value to assign to the input field.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

