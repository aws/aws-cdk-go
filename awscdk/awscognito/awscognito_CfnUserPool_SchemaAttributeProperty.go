package awscognito


// Contains information about the schema attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaAttributeProperty := &schemaAttributeProperty{
//   	attributeDataType: jsii.String("attributeDataType"),
//   	developerOnlyAttribute: jsii.Boolean(false),
//   	mutable: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	numberAttributeConstraints: &numberAttributeConstraintsProperty{
//   		maxValue: jsii.String("maxValue"),
//   		minValue: jsii.String("minValue"),
//   	},
//   	required: jsii.Boolean(false),
//   	stringAttributeConstraints: &stringAttributeConstraintsProperty{
//   		maxLength: jsii.String("maxLength"),
//   		minLength: jsii.String("minLength"),
//   	},
//   }
//
type CfnUserPool_SchemaAttributeProperty struct {
	// The attribute data type.
	AttributeDataType *string `field:"optional" json:"attributeDataType" yaml:"attributeDataType"`
	// > We recommend that you use [WriteAttributes](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolClientType.html#CognitoUserPools-Type-UserPoolClientType-WriteAttributes) in the user pool client to control how attributes can be mutated for new use cases instead of using `DeveloperOnlyAttribute` .
	//
	// Specifies whether the attribute type is developer only. This attribute can only be modified by an administrator. Users will not be able to modify this attribute using their access token.
	DeveloperOnlyAttribute interface{} `field:"optional" json:"developerOnlyAttribute" yaml:"developerOnlyAttribute"`
	// Specifies whether the value of the attribute can be changed.
	//
	// For any user pool attribute that is mapped to an IdP attribute, you must set this parameter to `true` . Amazon Cognito updates mapped attributes when users sign in to your application through an IdP. If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute. For more information, see [Specifying Identity Provider Attribute Mappings for Your User Pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html) .
	Mutable interface{} `field:"optional" json:"mutable" yaml:"mutable"`
	// A schema attribute of the name type.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the constraints for an attribute of the number type.
	NumberAttributeConstraints interface{} `field:"optional" json:"numberAttributeConstraints" yaml:"numberAttributeConstraints"`
	// Specifies whether a user pool attribute is required.
	//
	// If the attribute is required and the user doesn't provide a value, registration or sign-in will fail.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Specifies the constraints for an attribute of the string type.
	StringAttributeConstraints interface{} `field:"optional" json:"stringAttributeConstraints" yaml:"stringAttributeConstraints"`
}

