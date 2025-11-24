package mixinsawsresourceexplorer2


// Properties for CfnDefaultViewAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDefaultViewAssociationMixinProps := &CfnDefaultViewAssociationMixinProps{
//   	ViewArn: jsii.String("viewArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-defaultviewassociation.html
//
type CfnDefaultViewAssociationMixinProps struct {
	// The ARN of the view to set as the default for the AWS Region and AWS account in which you call this operation.
	//
	// The specified view must already exist in the specified Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-defaultviewassociation.html#cfn-resourceexplorer2-defaultviewassociation-viewarn
	//
	ViewArn *string `field:"optional" json:"viewArn" yaml:"viewArn"`
}

