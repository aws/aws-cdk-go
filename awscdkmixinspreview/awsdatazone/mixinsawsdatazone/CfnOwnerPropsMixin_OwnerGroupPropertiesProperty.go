package mixinsawsdatazone


// The properties of the domain unit owners group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ownerGroupPropertiesProperty := &OwnerGroupPropertiesProperty{
//   	GroupIdentifier: jsii.String("groupIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-owner-ownergroupproperties.html
//
type CfnOwnerPropsMixin_OwnerGroupPropertiesProperty struct {
	// The ID of the domain unit owners group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-owner-ownergroupproperties.html#cfn-datazone-owner-ownergroupproperties-groupidentifier
	//
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
}

