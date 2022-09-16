package awscloudfront


// Properties for defining a `CfnOriginAccessControl`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginAccessControlProps := &cfnOriginAccessControlProps{
//   	originAccessControlConfig: &originAccessControlConfigProperty{
//   		name: jsii.String("name"),
//   		originAccessControlOriginType: jsii.String("originAccessControlOriginType"),
//   		signingBehavior: jsii.String("signingBehavior"),
//   		signingProtocol: jsii.String("signingProtocol"),
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   	},
//   }
//
type CfnOriginAccessControlProps struct {
	// `AWS::CloudFront::OriginAccessControl.OriginAccessControlConfig`.
	OriginAccessControlConfig interface{} `field:"required" json:"originAccessControlConfig" yaml:"originAccessControlConfig"`
}

