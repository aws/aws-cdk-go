package awsiottwinmaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relationshipValueProperty := &relationshipValueProperty{
//   	targetComponentName: jsii.String("targetComponentName"),
//   	targetEntityId: jsii.String("targetEntityId"),
//   }
//
type CfnComponentType_RelationshipValueProperty struct {
	// `CfnComponentType.RelationshipValueProperty.TargetComponentName`.
	TargetComponentName *string `field:"optional" json:"targetComponentName" yaml:"targetComponentName"`
	// `CfnComponentType.RelationshipValueProperty.TargetEntityId`.
	TargetEntityId *string `field:"optional" json:"targetEntityId" yaml:"targetEntityId"`
}

