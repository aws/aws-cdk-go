package mixinsawsdatazone


// Properties for CfnOwnerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOwnerMixinProps := &CfnOwnerMixinProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EntityIdentifier: jsii.String("entityIdentifier"),
//   	EntityType: jsii.String("entityType"),
//   	Owner: &OwnerPropertiesProperty{
//   		Group: &OwnerGroupPropertiesProperty{
//   			GroupIdentifier: jsii.String("groupIdentifier"),
//   		},
//   		User: &OwnerUserPropertiesProperty{
//   			UserIdentifier: jsii.String("userIdentifier"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-owner.html
//
type CfnOwnerMixinProps struct {
	// The ID of the domain in which you want to add the entity owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-owner.html#cfn-datazone-owner-domainidentifier
	//
	DomainIdentifier *string `field:"optional" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the entity to which you want to add an owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-owner.html#cfn-datazone-owner-entityidentifier
	//
	EntityIdentifier *string `field:"optional" json:"entityIdentifier" yaml:"entityIdentifier"`
	// The type of an entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-owner.html#cfn-datazone-owner-entitytype
	//
	EntityType *string `field:"optional" json:"entityType" yaml:"entityType"`
	// The owner that you want to add to the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-owner.html#cfn-datazone-owner-owner
	//
	Owner interface{} `field:"optional" json:"owner" yaml:"owner"`
}

