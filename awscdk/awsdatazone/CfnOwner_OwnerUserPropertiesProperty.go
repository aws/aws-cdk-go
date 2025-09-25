package awsdatazone


// The properties of the owner user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ownerUserPropertiesProperty := &OwnerUserPropertiesProperty{
//   	UserIdentifier: jsii.String("userIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-owner-owneruserproperties.html
//
type CfnOwner_OwnerUserPropertiesProperty struct {
	// The ID of the owner user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-owner-owneruserproperties.html#cfn-datazone-owner-owneruserproperties-useridentifier
	//
	UserIdentifier *string `field:"optional" json:"userIdentifier" yaml:"userIdentifier"`
}

