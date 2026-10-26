package awscloudfront


// Properties for creating a Lambda Function URL Origin Access Control resource.
//
// Example:
//   var regionalEndpointDomainName string
//
//
//   oac := cloudfront.NewFunctionUrlOriginAccessControl(this, jsii.String("WebFunctionOAC"), &FunctionUrlOriginAccessControlProps{
//   	Signing: cloudfront.Signing_SIGV4_ALWAYS_AMZ_AUTH(),
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("MyDistribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewHttpOrigin(regionalEndpointDomainName, &HttpOriginProps{
//   			OriginAccessControlId: oac.OriginAccessControlId,
//   		}),
//   	},
//   })
//
type FunctionUrlOriginAccessControlProps struct {
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

