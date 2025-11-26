package previewawsverifiedpermissionsmixins


// Contains a list of principal types, resource types, and actions that can be specified in policies stored in the same policy store.
//
// If the validation mode for the policy store is set to `STRICT` , then policies that can't be validated by this schema are rejected by Verified Permissions and can't be stored in the policy store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaDefinitionProperty := &SchemaDefinitionProperty{
//   	CedarFormat: jsii.String("cedarFormat"),
//   	CedarJson: jsii.String("cedarJson"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-schemadefinition.html
//
type CfnPolicyStorePropsMixin_SchemaDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-schemadefinition.html#cfn-verifiedpermissions-policystore-schemadefinition-cedarformat
	//
	CedarFormat *string `field:"optional" json:"cedarFormat" yaml:"cedarFormat"`
	// A JSON string representation of the schema supported by applications that use this policy store.
	//
	// For more information, see [Policy store schema](https://docs.aws.amazon.com/verifiedpermissions/latest/userguide/schema.html) in the AVP User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-schemadefinition.html#cfn-verifiedpermissions-policystore-schemadefinition-cedarjson
	//
	CedarJson *string `field:"optional" json:"cedarJson" yaml:"cedarJson"`
}

