package awsram


// Properties for CfnSourceAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSourceAssociationMixinProps := &CfnSourceAssociationMixinProps{
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//   	SourceId: jsii.String("sourceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-sourceassociation.html
//
type CfnSourceAssociationMixinProps struct {
	// Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-sourceassociation.html#cfn-ram-sourceassociation-resourcesharearn
	//
	ResourceShareArn *string `field:"optional" json:"resourceShareArn" yaml:"resourceShareArn"`
	// Specifies the ID of the source account to associate with the resource share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-sourceassociation.html#cfn-ram-sourceassociation-sourceid
	//
	SourceId *string `field:"optional" json:"sourceId" yaml:"sourceId"`
}

