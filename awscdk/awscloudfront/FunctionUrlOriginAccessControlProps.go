package awscloudfront


// Properties for creating a Lambda Function URL Origin Access Control resource.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var fn function
//
//
//   fnUrl := fn.AddFunctionUrl(&FunctionUrlOptions{
//   	AuthType: lambda.FunctionUrlAuthType_AWS_IAM,
//   })
//
//   // Define a custom OAC
//   oac := cloudfront.NewFunctionUrlOriginAccessControl(this, jsii.String("MyOAC"), &FunctionUrlOriginAccessControlProps{
//   	OriginAccessControlName: jsii.String("CustomLambdaOAC"),
//   	Signing: cloudfront.Signing_SIGV4_ALWAYS(),
//   })
//
//   // Set up Lambda Function URL with OAC in CloudFront Distribution
//   // Set up Lambda Function URL with OAC in CloudFront Distribution
//   cloudfront.NewDistribution(this, jsii.String("MyDistribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.FunctionUrlOrigin_WithOriginAccessControl(fnUrl, &FunctionUrlOriginWithOACProps{
//   			OriginAccessControl: oac,
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

