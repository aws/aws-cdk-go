package awsiottwinmaker


// An object that specifies a relationship with another component type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationshipProperty := &relationshipProperty{
//   	relationshipType: jsii.String("relationshipType"),
//   	targetComponentTypeId: jsii.String("targetComponentTypeId"),
//   }
//
type CfnComponentType_RelationshipProperty struct {
	// The type of the relationship.
	RelationshipType *string `field:"optional" json:"relationshipType" yaml:"relationshipType"`
	// The ID of the target component type associated with this relationship.
	TargetComponentTypeId *string `field:"optional" json:"targetComponentTypeId" yaml:"targetComponentTypeId"`
}

