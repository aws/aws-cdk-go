package awsec2


// The state of token usage for your instance metadata requests.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   ec2.NewLaunchTemplate(this, jsii.String("LaunchTemplate"), &launchTemplateProps{
//   	httpEndpoint: jsii.Boolean(true),
//   	httpProtocolIpv6: jsii.Boolean(true),
//   	httpPutResponseHopLimit: jsii.Number(1),
//   	httpTokens: ec2.launchTemplateHttpTokens_REQUIRED,
//   	instanceMetadataTags: jsii.Boolean(true),
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-metadataoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-metadataoptions-httptokens
//
type LaunchTemplateHttpTokens string

const (
	// If the state is optional, you can choose to retrieve instance metadata with or without a signed token header on your request.
	LaunchTemplateHttpTokens_OPTIONAL LaunchTemplateHttpTokens = "OPTIONAL"
	// If the state is required, you must send a signed token header with any instance metadata retrieval requests.
	//
	// In this state,
	// retrieving the IAM role credentials always returns the version 2.0 credentials; the version 1.0 credentials are not available.
	LaunchTemplateHttpTokens_REQUIRED LaunchTemplateHttpTokens = "REQUIRED"
)

