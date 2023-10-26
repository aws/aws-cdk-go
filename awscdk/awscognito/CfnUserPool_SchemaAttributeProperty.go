package awscognito


// A list of the user attributes and their properties in your user pool.
//
// The attribute schema contains standard attributes, custom attributes with a `custom:` prefix, and developer attributes with a `dev:` prefix. For more information, see [User pool attributes](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html) .
//
// Developer-only attributes are a legacy feature of user pools, are read-only to all app clients. You can create and update developer-only attributes only with IAM-authenticated API operations. Use app client read/write permissions instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaAttributeProperty := &SchemaAttributeProperty{
//   	AttributeDataType: jsii.String("attributeDataType"),
//   	DeveloperOnlyAttribute: jsii.Boolean(false),
//   	Mutable: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	NumberAttributeConstraints: &NumberAttributeConstraintsProperty{
//   		MaxValue: jsii.String("maxValue"),
//   		MinValue: jsii.String("minValue"),
//   	},
//   	Required: jsii.Boolean(false),
//   	StringAttributeConstraints: &StringAttributeConstraintsProperty{
//   		MaxLength: jsii.String("maxLength"),
//   		MinLength: jsii.String("minLength"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html
//
type CfnUserPool_SchemaAttributeProperty struct {
	// The data format of the values for your attribute.
	//
	// When you choose an `AttributeDataType` , Amazon Cognito validates the input against the data type. A custom attribute value in your user's ID token is always a string, for example `"custom:isMember" : "true"` or `"custom:YearsAsMember" : "12"` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-attributedatatype
	//
	AttributeDataType *string `field:"optional" json:"attributeDataType" yaml:"attributeDataType"`
	// > We recommend that you use [WriteAttributes](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolClientType.html#CognitoUserPools-Type-UserPoolClientType-WriteAttributes) in the user pool client to control how attributes can be mutated for new use cases instead of using `DeveloperOnlyAttribute` .
	//
	// Specifies whether the attribute type is developer only. This attribute can only be modified by an administrator. Users will not be able to modify this attribute using their access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-developeronlyattribute
	//
	DeveloperOnlyAttribute interface{} `field:"optional" json:"developerOnlyAttribute" yaml:"developerOnlyAttribute"`
	// Specifies whether the value of the attribute can be changed.
	//
	// Any user pool attribute whose value you map from an IdP attribute must be mutable, with a parameter value of `true` . Amazon Cognito updates mapped attributes when users sign in to your application through an IdP. If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute. For more information, see [Specifying Identity Provider Attribute Mappings for Your User Pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-mutable
	//
	Mutable interface{} `field:"optional" json:"mutable" yaml:"mutable"`
	// The name of your user pool attribute.
	//
	// When you create or update a user pool, adding a schema attribute creates a custom or developer-only attribute. When you add an attribute with a `Name` value of `MyAttribute` , Amazon Cognito creates the custom attribute `custom:MyAttribute` . When `DeveloperOnlyAttribute` is `true` , Amazon Cognito creates your attribute as `dev:MyAttribute` . In an operation that describes a user pool, Amazon Cognito returns this value as `value` for standard attributes, `custom:value` for custom attributes, and `dev:value` for developer-only attributes..
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the constraints for an attribute of the number type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-numberattributeconstraints
	//
	NumberAttributeConstraints interface{} `field:"optional" json:"numberAttributeConstraints" yaml:"numberAttributeConstraints"`
	// Specifies whether a user pool attribute is required.
	//
	// If the attribute is required and the user doesn't provide a value, registration or sign-in will fail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-required
	//
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Specifies the constraints for an attribute of the string type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-stringattributeconstraints
	//
	StringAttributeConstraints interface{} `field:"optional" json:"stringAttributeConstraints" yaml:"stringAttributeConstraints"`
}

