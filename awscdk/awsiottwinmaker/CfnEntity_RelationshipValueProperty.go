package awsiottwinmaker


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
type CfnEntity_RelationshipValueProperty struct {
	// `CfnEntity.RelationshipValueProperty.TargetComponentName`.
	TargetComponentName *string `field:"optional" json:"targetComponentName" yaml:"targetComponentName"`
	// `CfnEntity.RelationshipValueProperty.TargetEntityId`.
	TargetEntityId *string `field:"optional" json:"targetEntityId" yaml:"targetEntityId"`
}

