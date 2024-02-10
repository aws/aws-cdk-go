package awsamplifyuibuilder


// The `FieldInputConfig` property specifies the configuration for the default input values to display for a field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var formInputValuePropertyProperty_ formInputValuePropertyProperty
//
//   fieldInputConfigProperty := &FieldInputConfigProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	DefaultChecked: jsii.Boolean(false),
//   	DefaultCountryCode: jsii.String("defaultCountryCode"),
//   	DefaultValue: jsii.String("defaultValue"),
//   	DescriptiveText: jsii.String("descriptiveText"),
//   	FileUploaderConfig: &FileUploaderFieldConfigProperty{
//   		AcceptedFileTypes: []*string{
//   			jsii.String("acceptedFileTypes"),
//   		},
//   		AccessLevel: jsii.String("accessLevel"),
//
//   		// the properties below are optional
//   		IsResumable: jsii.Boolean(false),
//   		MaxFileCount: jsii.Number(123),
//   		MaxSize: jsii.Number(123),
//   		ShowThumbnails: jsii.Boolean(false),
//   	},
//   	IsArray: jsii.Boolean(false),
//   	MaxValue: jsii.Number(123),
//   	MinValue: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Placeholder: jsii.String("placeholder"),
//   	ReadOnly: jsii.Boolean(false),
//   	Required: jsii.Boolean(false),
//   	Step: jsii.Number(123),
//   	Value: jsii.String("value"),
//   	ValueMappings: &ValueMappingsProperty{
//   		Values: []interface{}{
//   			&ValueMappingProperty{
//   				Value: &formInputValuePropertyProperty{
//   					BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   						Property: jsii.String("property"),
//
//   						// the properties below are optional
//   						Field: jsii.String("field"),
//   					},
//   					Concat: []interface{}{
//   						formInputValuePropertyProperty_,
//   					},
//   					Value: jsii.String("value"),
//   				},
//
//   				// the properties below are optional
//   				DisplayValue: &formInputValuePropertyProperty{
//   					BindingProperties: &FormInputValuePropertyBindingPropertiesProperty{
//   						Property: jsii.String("property"),
//
//   						// the properties below are optional
//   						Field: jsii.String("field"),
//   					},
//   					Concat: []interface{}{
//   						formInputValuePropertyProperty_,
//   					},
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		BindingProperties: map[string]interface{}{
//   			"bindingPropertiesKey": &FormInputBindingPropertiesValueProperty{
//   				"bindingProperties": &FormInputBindingPropertiesValuePropertiesProperty{
//   					"model": jsii.String("model"),
//   				},
//   				"type": jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html
//
type CfnForm_FieldInputConfigProperty struct {
	// The input type for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specifies whether a field has a default value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-defaultchecked
	//
	DefaultChecked interface{} `field:"optional" json:"defaultChecked" yaml:"defaultChecked"`
	// The default country code for a phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-defaultcountrycode
	//
	DefaultCountryCode *string `field:"optional" json:"defaultCountryCode" yaml:"defaultCountryCode"`
	// The default value for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The text to display to describe the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-descriptivetext
	//
	DescriptiveText *string `field:"optional" json:"descriptiveText" yaml:"descriptiveText"`
	// The configuration for the file uploader field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-fileuploaderconfig
	//
	FileUploaderConfig interface{} `field:"optional" json:"fileUploaderConfig" yaml:"fileUploaderConfig"`
	// Specifies whether to render the field as an array.
	//
	// This property is ignored if the `dataSourceType` for the form is a Data Store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-isarray
	//
	IsArray interface{} `field:"optional" json:"isArray" yaml:"isArray"`
	// The maximum value to display for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-maxvalue
	//
	MaxValue *float64 `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum value to display for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-minvalue
	//
	MinValue *float64 `field:"optional" json:"minValue" yaml:"minValue"`
	// The name of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The text to display as a placeholder for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-placeholder
	//
	Placeholder *string `field:"optional" json:"placeholder" yaml:"placeholder"`
	// Specifies a read only field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-readonly
	//
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Specifies a field that requires input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The stepping increment for a numeric value in a field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-step
	//
	Step *float64 `field:"optional" json:"step" yaml:"step"`
	// The value for the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
	// The information to use to customize the input fields with data at runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fieldinputconfig.html#cfn-amplifyuibuilder-form-fieldinputconfig-valuemappings
	//
	ValueMappings interface{} `field:"optional" json:"valueMappings" yaml:"valueMappings"`
}

