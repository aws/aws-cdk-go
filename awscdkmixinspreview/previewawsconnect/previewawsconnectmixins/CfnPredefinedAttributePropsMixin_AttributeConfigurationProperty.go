package previewawsconnectmixins


// Custom metadata associated to a Predefined attribute that controls how the attribute behaves when used by upstream services.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attributeConfigurationProperty := &AttributeConfigurationProperty{
//   	EnableValueValidationOnAssociation: jsii.Boolean(false),
//   	IsReadOnly: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-predefinedattribute-attributeconfiguration.html
//
type CfnPredefinedAttributePropsMixin_AttributeConfigurationProperty struct {
	// Enables customers to enforce strict validation on the specific values that this predefined attribute can hold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-predefinedattribute-attributeconfiguration.html#cfn-connect-predefinedattribute-attributeconfiguration-enablevaluevalidationonassociation
	//
	EnableValueValidationOnAssociation interface{} `field:"optional" json:"enableValueValidationOnAssociation" yaml:"enableValueValidationOnAssociation"`
	// Allows the predefined attribute to show up and be managed in the Amazon Connect UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-predefinedattribute-attributeconfiguration.html#cfn-connect-predefinedattribute-attributeconfiguration-isreadonly
	//
	IsReadOnly interface{} `field:"optional" json:"isReadOnly" yaml:"isReadOnly"`
}

