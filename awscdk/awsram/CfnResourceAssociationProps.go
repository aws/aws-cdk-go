package awsram


// Properties for defining a `CfnResourceAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceAssociationProps := &CfnResourceAssociationProps{
//   	ResourceArn: jsii.String("resourceArn"),
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceassociation.html
//
type CfnResourceAssociationProps struct {
	// Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource to associate with the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceassociation.html#cfn-ram-resourceassociation-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceassociation.html#cfn-ram-resourceassociation-resourcesharearn
	//
	ResourceShareArn *string `field:"required" json:"resourceShareArn" yaml:"resourceShareArn"`
}

