package awsiottwinmaker


// The entity relationship.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationshipProperty := &RelationshipProperty{
//   	RelationshipType: jsii.String("relationshipType"),
//   	TargetComponentTypeId: jsii.String("targetComponentTypeId"),
//   }
//
type CfnEntity_RelationshipProperty struct {
	// The relationship type.
	RelationshipType *string `field:"optional" json:"relationshipType" yaml:"relationshipType"`
	// the component type Id target.
	TargetComponentTypeId *string `field:"optional" json:"targetComponentTypeId" yaml:"targetComponentTypeId"`
}

