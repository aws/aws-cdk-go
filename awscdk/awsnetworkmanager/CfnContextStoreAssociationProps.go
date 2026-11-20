package awsnetworkmanager


// Properties for defining a `CfnContextStoreAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContextStoreAssociationProps := &CfnContextStoreAssociationProps{
//   	ContextStoreId: jsii.String("contextStoreId"),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-contextstoreassociation.html
//
type CfnContextStoreAssociationProps struct {
	// The ID of the context store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-contextstoreassociation.html#cfn-networkmanager-contextstoreassociation-contextstoreid
	//
	ContextStoreId *string `field:"required" json:"contextStoreId" yaml:"contextStoreId"`
	// The ARN of the Transit Gateway to associate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-contextstoreassociation.html#cfn-networkmanager-contextstoreassociation-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

