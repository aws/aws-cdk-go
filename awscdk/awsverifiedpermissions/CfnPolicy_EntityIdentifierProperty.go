package awsverifiedpermissions


// Contains the identifier of an entity in a policy, including its ID and type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityIdentifierProperty := &EntityIdentifierProperty{
//   	EntityId: jsii.String("entityId"),
//   	EntityType: jsii.String("entityType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-entityidentifier.html
//
type CfnPolicy_EntityIdentifierProperty struct {
	// The identifier of an entity.
	//
	// `"entityId":" *identifier* "`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-entityidentifier.html#cfn-verifiedpermissions-policy-entityidentifier-entityid
	//
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// The type of an entity.
	//
	// Example: `"entityType":" *typeName* "`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-policy-entityidentifier.html#cfn-verifiedpermissions-policy-entityidentifier-entitytype
	//
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
}

