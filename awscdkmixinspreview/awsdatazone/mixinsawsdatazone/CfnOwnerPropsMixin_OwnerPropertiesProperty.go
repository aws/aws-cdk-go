package mixinsawsdatazone


// The properties of a domain unit's owner.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ownerPropertiesProperty := &OwnerPropertiesProperty{
//   	Group: &OwnerGroupPropertiesProperty{
//   		GroupIdentifier: jsii.String("groupIdentifier"),
//   	},
//   	User: &OwnerUserPropertiesProperty{
//   		UserIdentifier: jsii.String("userIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-owner-ownerproperties.html
//
type CfnOwnerPropsMixin_OwnerPropertiesProperty struct {
	// Specifies that the domain unit owner is a group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-owner-ownerproperties.html#cfn-datazone-owner-ownerproperties-group
	//
	Group interface{} `field:"optional" json:"group" yaml:"group"`
	// Specifies that the domain unit owner is a user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-owner-ownerproperties.html#cfn-datazone-owner-ownerproperties-user
	//
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

