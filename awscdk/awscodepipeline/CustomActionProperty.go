package awscodepipeline


// The creation attributes used for defining a configuration property of a custom Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionProperty := &CustomActionProperty{
//   	Name: jsii.String("name"),
//   	Required: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Key: jsii.Boolean(false),
//   	Queryable: jsii.Boolean(false),
//   	Secret: jsii.Boolean(false),
//   	Type: jsii.String("type"),
//   }
//
type CustomActionProperty struct {
	// The name of the property.
	//
	// You use this name in the `configuration` attribute when defining your custom Action class.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether this property is required.
	Required *bool `field:"required" json:"required" yaml:"required"`
	// The description of the property.
	// Default: the description will be empty.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether this property is a key.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-key
	//
	// Default: false.
	//
	Key *bool `field:"optional" json:"key" yaml:"key"`
	// Whether this property is queryable.
	//
	// Note that only a single property of a custom Action can be queryable.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-queryable
	//
	// Default: false.
	//
	Queryable *bool `field:"optional" json:"queryable" yaml:"queryable"`
	// Whether this property is secret, like a password, or access key.
	// Default: false.
	//
	Secret *bool `field:"optional" json:"secret" yaml:"secret"`
	// The type of the property, like 'String', 'Number', or 'Boolean'.
	// Default: 'String'.
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

