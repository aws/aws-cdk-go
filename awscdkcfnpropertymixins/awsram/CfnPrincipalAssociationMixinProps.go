package awsram


// Properties for CfnPrincipalAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPrincipalAssociationMixinProps := &CfnPrincipalAssociationMixinProps{
//   	Principal: jsii.String("principal"),
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-principalassociation.html
//
type CfnPrincipalAssociationMixinProps struct {
	// Specifies the principal to associate with the resource share. The possible values are:.
	//
	// - An AWS account ID
	//
	// - An Amazon Resource Name (ARN) of an organization in AWS Organizations
	//
	// - An ARN of an organizational unit (OU) in AWS Organizations
	//
	// - An ARN of an IAM role
	//
	// - An ARN of an IAM user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-principalassociation.html#cfn-ram-principalassociation-principal
	//
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-principalassociation.html#cfn-ram-principalassociation-resourcesharearn
	//
	ResourceShareArn *string `field:"optional" json:"resourceShareArn" yaml:"resourceShareArn"`
}

