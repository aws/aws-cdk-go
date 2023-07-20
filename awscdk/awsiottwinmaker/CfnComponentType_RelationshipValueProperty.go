package awsiottwinmaker


// The component type relationship value.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationshipvalue.html
//
type CfnComponentType_RelationshipValueProperty struct {
	// The target component name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationshipvalue.html#cfn-iottwinmaker-componenttype-relationshipvalue-targetcomponentname
	//
	TargetComponentName *string `field:"optional" json:"targetComponentName" yaml:"targetComponentName"`
	// The target entity Id.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationshipvalue.html#cfn-iottwinmaker-componenttype-relationshipvalue-targetentityid
	//
	TargetEntityId *string `field:"optional" json:"targetEntityId" yaml:"targetEntityId"`
}

