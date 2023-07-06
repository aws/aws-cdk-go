package awsverifiedpermissions


// Contains a list of principal types, resource types, and actions that can be specified in policies stored in the same policy store.
//
// If the validation mode for the policy store is set to `STRICT` , then policies that can't be validated by this schema are rejected by Verified Permissions and can't be stored in the policy store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaDefinitionProperty := &SchemaDefinitionProperty{
//   	CedarJson: jsii.String("cedarJson"),
//   }
//
type CfnPolicyStore_SchemaDefinitionProperty struct {
	// A JSON string representation of the schema supported by applications that use this policy store.
	//
	// For more information, see [Policy store schema](https://docs.aws.amazon.com/verifiedpermissions/latest/userguide/schema.html) in the *Amazon Verified Permissions User Guide* .
	CedarJson *string `field:"optional" json:"cedarJson" yaml:"cedarJson"`
}

