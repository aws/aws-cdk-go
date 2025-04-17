package awsiottwinmaker


// The entity relationship.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationshipValueProperty := &RelationshipValueProperty{
//   	TargetComponentName: jsii.String("targetComponentName"),
//   	TargetEntityId: jsii.String("targetEntityId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-relationshipvalue.html
//
type CfnEntity_RelationshipValueProperty struct {
	// The target component name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-relationshipvalue.html#cfn-iottwinmaker-entity-relationshipvalue-targetcomponentname
	//
	TargetComponentName *string `field:"optional" json:"targetComponentName" yaml:"targetComponentName"`
	// The target entity Id.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-relationshipvalue.html#cfn-iottwinmaker-entity-relationshipvalue-targetentityid
	//
	TargetEntityId *string `field:"optional" json:"targetEntityId" yaml:"targetEntityId"`
}

