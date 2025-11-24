package mixinsawsamplifyuibuilder


// Represents the data binding configuration for a specific property using data stored in AWS .
//
// For AWS connected properties, you can bind a property to data stored in an Amplify DataStore model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   formInputBindingPropertiesValuePropertiesProperty := &FormInputBindingPropertiesValuePropertiesProperty{
//   	Model: jsii.String("model"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-forminputbindingpropertiesvalueproperties.html
//
type CfnFormPropsMixin_FormInputBindingPropertiesValuePropertiesProperty struct {
	// An Amplify DataStore model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-forminputbindingpropertiesvalueproperties.html#cfn-amplifyuibuilder-form-forminputbindingpropertiesvalueproperties-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
}

