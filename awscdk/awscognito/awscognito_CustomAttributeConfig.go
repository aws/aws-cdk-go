package awscognito


// Configuration that will be fed into CloudFormation for any custom attribute type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customAttributeConfig := &customAttributeConfig{
//   	dataType: jsii.String("dataType"),
//
//   	// the properties below are optional
//   	mutable: jsii.Boolean(false),
//   	numberConstraints: &numberAttributeConstraints{
//   		max: jsii.Number(123),
//   		min: jsii.Number(123),
//   	},
//   	stringConstraints: &stringAttributeConstraints{
//   		maxLen: jsii.Number(123),
//   		minLen: jsii.Number(123),
//   	},
//   }
//
type CustomAttributeConfig struct {
	// The data type of the custom attribute.
	// See: https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SchemaAttributeType.html#CognitoUserPools-Type-SchemaAttributeType-AttributeDataType
	//
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// Specifies whether the value of the attribute can be changed.
	//
	// For any user pool attribute that's mapped to an identity provider attribute, you must set this parameter to true.
	// Amazon Cognito updates mapped attributes when users sign in to your application through an identity provider.
	// If an attribute is immutable, Amazon Cognito throws an error when it attempts to update the attribute.
	Mutable *bool `field:"optional" json:"mutable" yaml:"mutable"`
	// The constraints for a custom attribute of the 'Number' data type.
	NumberConstraints *NumberAttributeConstraints `field:"optional" json:"numberConstraints" yaml:"numberConstraints"`
	// The constraints for a custom attribute of 'String' data type.
	StringConstraints *StringAttributeConstraints `field:"optional" json:"stringConstraints" yaml:"stringConstraints"`
}

