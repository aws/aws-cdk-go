package mixinsawscloudfront


// Properties for CfnKeyGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKeyGroupMixinProps := &CfnKeyGroupMixinProps{
//   	KeyGroupConfig: &KeyGroupConfigProperty{
//   		Comment: jsii.String("comment"),
//   		Items: []*string{
//   			jsii.String("items"),
//   		},
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keygroup.html
//
type CfnKeyGroupMixinProps struct {
	// The key group configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keygroup.html#cfn-cloudfront-keygroup-keygroupconfig
	//
	KeyGroupConfig interface{} `field:"optional" json:"keyGroupConfig" yaml:"keyGroupConfig"`
}

