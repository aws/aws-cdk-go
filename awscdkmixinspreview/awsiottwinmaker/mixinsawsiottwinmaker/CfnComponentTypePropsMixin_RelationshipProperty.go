package mixinsawsiottwinmaker


// An object that specifies a relationship with another component type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   relationshipProperty := &RelationshipProperty{
//   	RelationshipType: jsii.String("relationshipType"),
//   	TargetComponentTypeId: jsii.String("targetComponentTypeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationship.html
//
type CfnComponentTypePropsMixin_RelationshipProperty struct {
	// The type of the relationship.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationship.html#cfn-iottwinmaker-componenttype-relationship-relationshiptype
	//
	RelationshipType *string `field:"optional" json:"relationshipType" yaml:"relationshipType"`
	// The ID of the target component type associated with this relationship.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-relationship.html#cfn-iottwinmaker-componenttype-relationship-targetcomponenttypeid
	//
	TargetComponentTypeId *string `field:"optional" json:"targetComponentTypeId" yaml:"targetComponentTypeId"`
}

