package mixinsawscloudfront


// Properties for CfnOriginAccessControlPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOriginAccessControlMixinProps := &CfnOriginAccessControlMixinProps{
//   	OriginAccessControlConfig: &OriginAccessControlConfigProperty{
//   		Description: jsii.String("description"),
//   		Name: jsii.String("name"),
//   		OriginAccessControlOriginType: jsii.String("originAccessControlOriginType"),
//   		SigningBehavior: jsii.String("signingBehavior"),
//   		SigningProtocol: jsii.String("signingProtocol"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originaccesscontrol.html
//
type CfnOriginAccessControlMixinProps struct {
	// The origin access control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originaccesscontrol.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig
	//
	OriginAccessControlConfig interface{} `field:"optional" json:"originAccessControlConfig" yaml:"originAccessControlConfig"`
}

