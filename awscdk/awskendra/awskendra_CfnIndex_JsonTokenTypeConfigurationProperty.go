package awskendra


// Provides the configuration information for the JSON token type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonTokenTypeConfigurationProperty := &jsonTokenTypeConfigurationProperty{
//   	groupAttributeField: jsii.String("groupAttributeField"),
//   	userNameAttributeField: jsii.String("userNameAttributeField"),
//   }
//
type CfnIndex_JsonTokenTypeConfigurationProperty struct {
	// The group attribute field.
	GroupAttributeField *string `field:"required" json:"groupAttributeField" yaml:"groupAttributeField"`
	// The user name attribute field.
	UserNameAttributeField *string `field:"required" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

