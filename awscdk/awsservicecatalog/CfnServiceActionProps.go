package awsservicecatalog


// Properties for defining a `CfnServiceAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceActionProps := &CfnServiceActionProps{
//   	Definition: []interface{}{
//   		&DefinitionParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	DefinitionType: jsii.String("definitionType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Description: jsii.String("description"),
//   }
//
type CfnServiceActionProps struct {
	// A map that defines the self-service action.
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// The self-service action definition type.
	//
	// For example, `SSM_AUTOMATION` .
	DefinitionType *string `field:"required" json:"definitionType" yaml:"definitionType"`
	// The self-service action name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The self-service action description.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

