package mixinsawsiottwinmaker


// The property group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   propertyGroupProperty := &PropertyGroupProperty{
//   	GroupType: jsii.String("groupType"),
//   	PropertyNames: []*string{
//   		jsii.String("propertyNames"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-propertygroup.html
//
type CfnComponentTypePropsMixin_PropertyGroupProperty struct {
	// The group type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-propertygroup.html#cfn-iottwinmaker-componenttype-propertygroup-grouptype
	//
	GroupType *string `field:"optional" json:"groupType" yaml:"groupType"`
	// The property names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-propertygroup.html#cfn-iottwinmaker-componenttype-propertygroup-propertynames
	//
	PropertyNames *[]*string `field:"optional" json:"propertyNames" yaml:"propertyNames"`
}

