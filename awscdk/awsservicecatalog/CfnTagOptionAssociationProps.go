package awsservicecatalog


// Properties for defining a `CfnTagOptionAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTagOptionAssociationProps := &CfnTagOptionAssociationProps{
//   	ResourceId: jsii.String("resourceId"),
//   	TagOptionId: jsii.String("tagOptionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html
//
type CfnTagOptionAssociationProps struct {
	// The resource identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html#cfn-servicecatalog-tagoptionassociation-resourceid
	//
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The TagOption identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html#cfn-servicecatalog-tagoptionassociation-tagoptionid
	//
	TagOptionId *string `field:"required" json:"tagOptionId" yaml:"tagOptionId"`
}

