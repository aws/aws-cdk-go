package awsiottwinmaker


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
type CfnEntity_RelationshipProperty struct {
	// `CfnEntity.RelationshipProperty.RelationshipType`.
	RelationshipType *string `field:"optional" json:"relationshipType" yaml:"relationshipType"`
	// `CfnEntity.RelationshipProperty.TargetComponentTypeId`.
	TargetComponentTypeId *string `field:"optional" json:"targetComponentTypeId" yaml:"targetComponentTypeId"`
}

