package awscloudfront


// Common properties for creating a Origin Access Control resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var signing signing
//
//   originAccessControlBaseProps := &OriginAccessControlBaseProps{
//   	Description: jsii.String("description"),
//   	OriginAccessControlName: jsii.String("originAccessControlName"),
//   	Signing: signing,
//   }
//
type OriginAccessControlBaseProps struct {
	// A description of the origin access control.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name to identify the origin access control, with a maximum length of 64 characters.
	// Default: - a generated name.
	//
	OriginAccessControlName *string `field:"optional" json:"originAccessControlName" yaml:"originAccessControlName"`
	// Specifies which requests CloudFront signs and the signing protocol.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-signingbehavior
	//
	// Default: SIGV4_ALWAYS.
	//
	Signing Signing `field:"optional" json:"signing" yaml:"signing"`
}

