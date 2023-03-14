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
type CfnOriginAccessControlProps struct {
	// The origin access control.
	OriginAccessControlConfig interface{} `field:"required" json:"originAccessControlConfig" yaml:"originAccessControlConfig"`
}

