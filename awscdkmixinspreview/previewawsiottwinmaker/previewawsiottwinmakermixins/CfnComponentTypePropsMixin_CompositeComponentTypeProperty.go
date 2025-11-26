package previewawsiottwinmakermixins


// Specifies the ID of the composite component type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   compositeComponentTypeProperty := &CompositeComponentTypeProperty{
//   	ComponentTypeId: jsii.String("componentTypeId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-compositecomponenttype.html
//
type CfnComponentTypePropsMixin_CompositeComponentTypeProperty struct {
	// The ID of the component type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-compositecomponenttype.html#cfn-iottwinmaker-componenttype-compositecomponenttype-componenttypeid
	//
	ComponentTypeId *string `field:"optional" json:"componentTypeId" yaml:"componentTypeId"`
}

