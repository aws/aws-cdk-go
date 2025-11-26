package previewawsecrmixins


// Properties for CfnPullTimeUpdateExclusionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPullTimeUpdateExclusionMixinProps := &CfnPullTimeUpdateExclusionMixinProps{
//   	PrincipalArn: jsii.String("principalArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pulltimeupdateexclusion.html
//
type CfnPullTimeUpdateExclusionMixinProps struct {
	// Principal arn that should not update image pull times.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pulltimeupdateexclusion.html#cfn-ecr-pulltimeupdateexclusion-principalarn
	//
	PrincipalArn *string `field:"optional" json:"principalArn" yaml:"principalArn"`
}

