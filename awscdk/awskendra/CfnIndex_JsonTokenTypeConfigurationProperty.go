package awskendra


// Provides the configuration information for the JSON token type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonTokenTypeConfigurationProperty := &JsonTokenTypeConfigurationProperty{
//   	GroupAttributeField: jsii.String("groupAttributeField"),
//   	UserNameAttributeField: jsii.String("userNameAttributeField"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jsontokentypeconfiguration.html
//
type CfnIndex_JsonTokenTypeConfigurationProperty struct {
	// The group attribute field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jsontokentypeconfiguration.html#cfn-kendra-index-jsontokentypeconfiguration-groupattributefield
	//
	GroupAttributeField *string `field:"required" json:"groupAttributeField" yaml:"groupAttributeField"`
	// The user name attribute field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-jsontokentypeconfiguration.html#cfn-kendra-index-jsontokentypeconfiguration-usernameattributefield
	//
	UserNameAttributeField *string `field:"required" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

