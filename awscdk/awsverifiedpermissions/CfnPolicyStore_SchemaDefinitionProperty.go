package awsverifiedpermissions


// Contains a list of principal types, resource types, and actions that can be specified in policies stored in the same .
//
// If the validation mode for the is set to `STRICT` , then policies that can't be validated by this schema are rejected by and can't be stored in the .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-schemadefinition.html
//
type CfnPolicyStore_SchemaDefinitionProperty struct {
	// A JSON string representation of the schema supported by applications that use this .
	//
	// For more information, see [Policy store schema](https://docs.aws.amazon.com/schema.html) in the ** .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policystore-schemadefinition.html#cfn-verifiedpermissions-policystore-schemadefinition-cedarjson
	//
	CedarJson *string `field:"optional" json:"cedarJson" yaml:"cedarJson"`
}

