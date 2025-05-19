package awsdatazone


// Properties for defining a `CfnOwner`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOwnerProps := &CfnOwnerProps{
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
type CfnOwnerProps struct {
	// The ID of the domain in which you want to add the entity owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-owner.html#cfn-datazone-owner-domainidentifier
	//
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The ID of the entity to which you want to add an owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-owner.html#cfn-datazone-owner-entityidentifier
	//
	EntityIdentifier *string `field:"required" json:"entityIdentifier" yaml:"entityIdentifier"`
	// The type of an entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-owner.html#cfn-datazone-owner-entitytype
	//
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// The owner that you want to add to the entity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-owner.html#cfn-datazone-owner-owner
	//
	Owner interface{} `field:"required" json:"owner" yaml:"owner"`
}

