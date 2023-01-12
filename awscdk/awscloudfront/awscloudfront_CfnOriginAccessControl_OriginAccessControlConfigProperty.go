package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originAccessControlConfigProperty := &originAccessControlConfigProperty{
//   	name: jsii.String("name"),
//   	originAccessControlOriginType: jsii.String("originAccessControlOriginType"),
//   	signingBehavior: jsii.String("signingBehavior"),
//   	signingProtocol: jsii.String("signingProtocol"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnOriginAccessControl_OriginAccessControlConfigProperty struct {
	// `CfnOriginAccessControl.OriginAccessControlConfigProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnOriginAccessControl.OriginAccessControlConfigProperty.OriginAccessControlOriginType`.
	OriginAccessControlOriginType *string `field:"required" json:"originAccessControlOriginType" yaml:"originAccessControlOriginType"`
	// `CfnOriginAccessControl.OriginAccessControlConfigProperty.SigningBehavior`.
	SigningBehavior *string `field:"required" json:"signingBehavior" yaml:"signingBehavior"`
	// `CfnOriginAccessControl.OriginAccessControlConfigProperty.SigningProtocol`.
	SigningProtocol *string `field:"required" json:"signingProtocol" yaml:"signingProtocol"`
	// `CfnOriginAccessControl.OriginAccessControlConfigProperty.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

