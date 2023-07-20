package awsamplifyuibuilder


// The `ComponentBindingPropertiesValueProperties` property specifies the data binding configuration for a specific property using data stored in AWS .
//
// For AWS connected properties, you can bind a property to data stored in an Amazon S3 bucket, an Amplify DataStore model or an authenticated user attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var predicateProperty_ predicateProperty
//
//   componentBindingPropertiesValuePropertiesProperty := &ComponentBindingPropertiesValuePropertiesProperty{
//   	Bucket: jsii.String("bucket"),
//   	DefaultValue: jsii.String("defaultValue"),
//   	Field: jsii.String("field"),
//   	Key: jsii.String("key"),
//   	Model: jsii.String("model"),
//   	Predicates: []interface{}{
//   		&predicateProperty{
//   			And: []interface{}{
//   				predicateProperty_,
//   			},
//   			Field: jsii.String("field"),
//   			Operand: jsii.String("operand"),
//   			Operator: jsii.String("operator"),
//   			Or: []interface{}{
//   				predicateProperty_,
//   			},
//   		},
//   	},
//   	UserAttribute: jsii.String("userAttribute"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalueproperties.html
//
type CfnComponent_ComponentBindingPropertiesValuePropertiesProperty struct {
	// An Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalueproperties.html#cfn-amplifyuibuilder-component-componentbindingpropertiesvalueproperties-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The default value to assign to the property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalueproperties.html#cfn-amplifyuibuilder-component-componentbindingpropertiesvalueproperties-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The field to bind the data to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalueproperties.html#cfn-amplifyuibuilder-component-componentbindingpropertiesvalueproperties-field
	//
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The storage key for an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalueproperties.html#cfn-amplifyuibuilder-component-componentbindingpropertiesvalueproperties-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// An Amplify DataStore model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalueproperties.html#cfn-amplifyuibuilder-component-componentbindingpropertiesvalueproperties-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
	// A list of predicates for binding a component's properties to data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalueproperties.html#cfn-amplifyuibuilder-component-componentbindingpropertiesvalueproperties-predicates
	//
	Predicates interface{} `field:"optional" json:"predicates" yaml:"predicates"`
	// An authenticated user attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalueproperties.html#cfn-amplifyuibuilder-component-componentbindingpropertiesvalueproperties-userattribute
	//
	UserAttribute *string `field:"optional" json:"userAttribute" yaml:"userAttribute"`
}

