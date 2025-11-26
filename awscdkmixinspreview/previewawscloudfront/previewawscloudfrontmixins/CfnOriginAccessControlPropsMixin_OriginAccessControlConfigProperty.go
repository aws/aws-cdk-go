package previewawscloudfrontmixins


// Creates a new origin access control in CloudFront.
//
// After you create an origin access control, you can add it to an origin in a CloudFront distribution so that CloudFront sends authenticated (signed) requests to the origin.
//
// This makes it possible to block public access to the origin, allowing viewers (users) to access the origin's content only through CloudFront.
//
// For more information about using a CloudFront origin access control, see [Restricting access to an AWS origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-origin.html) in the *Amazon CloudFront Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   originAccessControlConfigProperty := &OriginAccessControlConfigProperty{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	OriginAccessControlOriginType: jsii.String("originAccessControlOriginType"),
//   	SigningBehavior: jsii.String("signingBehavior"),
//   	SigningProtocol: jsii.String("signingProtocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html
//
type CfnOriginAccessControlPropsMixin_OriginAccessControlConfigProperty struct {
	// A description of the origin access control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name to identify the origin access control.
	//
	// You can specify up to 64 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of origin that this origin access control is for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-originaccesscontrolorigintype
	//
	OriginAccessControlOriginType *string `field:"optional" json:"originAccessControlOriginType" yaml:"originAccessControlOriginType"`
	// Specifies which requests CloudFront signs (adds authentication information to).
	//
	// Specify `always` for the most common use case. For more information, see [origin access control advanced settings](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html#oac-advanced-settings) in the *Amazon CloudFront Developer Guide* .
	//
	// This field can have one of the following values:
	//
	// - `always` – CloudFront signs all origin requests, overwriting the `Authorization` header from the viewer request if one exists.
	// - `never` – CloudFront doesn't sign any origin requests. This value turns off origin access control for all origins in all distributions that use this origin access control.
	// - `no-override` – If the viewer request doesn't contain the `Authorization` header, then CloudFront signs the origin request. If the viewer request contains the `Authorization` header, then CloudFront doesn't sign the origin request and instead passes along the `Authorization` header from the viewer request. *WARNING: To pass along the `Authorization` header from the viewer request, you *must* add the `Authorization` header to a [cache policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html) for all cache behaviors that use origins associated with this origin access control.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-signingbehavior
	//
	SigningBehavior *string `field:"optional" json:"signingBehavior" yaml:"signingBehavior"`
	// The signing protocol of the origin access control, which determines how CloudFront signs (authenticates) requests.
	//
	// The only valid value is `sigv4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-signingprotocol
	//
	SigningProtocol *string `field:"optional" json:"signingProtocol" yaml:"signingProtocol"`
}

