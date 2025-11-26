package previewawsservicecatalogmixins


// Properties for CfnServiceActionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceActionMixinProps := &CfnServiceActionMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Definition: []interface{}{
//   		&DefinitionParameterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	DefinitionType: jsii.String("definitionType"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceaction.html
//
type CfnServiceActionMixinProps struct {
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceaction.html#cfn-servicecatalog-serviceaction-acceptlanguage
	//
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// A map that defines the self-service action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceaction.html#cfn-servicecatalog-serviceaction-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// The self-service action definition type.
	//
	// For example, `SSM_AUTOMATION` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceaction.html#cfn-servicecatalog-serviceaction-definitiontype
	//
	DefinitionType *string `field:"optional" json:"definitionType" yaml:"definitionType"`
	// The self-service action description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceaction.html#cfn-servicecatalog-serviceaction-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The self-service action name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-serviceaction.html#cfn-servicecatalog-serviceaction-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

