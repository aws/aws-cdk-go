package awscloudfront


// Properties for defining a `CfnOriginAccessControl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginAccessControlProps := &CfnOriginAccessControlProps{
//   	OriginAccessControlConfig: &OriginAccessControlConfigProperty{
//   		Name: jsii.String("name"),
//   		OriginAccessControlOriginType: jsii.String("originAccessControlOriginType"),
//   		SigningBehavior: jsii.String("signingBehavior"),
//   		SigningProtocol: jsii.String("signingProtocol"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originaccesscontrol.html
//
type CfnOriginAccessControlProps struct {
	// The origin access control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originaccesscontrol.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig
	//
	OriginAccessControlConfig interface{} `field:"required" json:"originAccessControlConfig" yaml:"originAccessControlConfig"`
}

