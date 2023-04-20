package awsec2


// The state of token usage for your instance metadata requests.
//
// Example:
//   ec2.NewLaunchTemplate(this, jsii.String("LaunchTemplate"), &LaunchTemplateProps{
//   	HttpEndpoint: jsii.Boolean(true),
//   	HttpProtocolIpv6: jsii.Boolean(true),
//   	HttpPutResponseHopLimit: jsii.Number(1),
//   	HttpTokens: ec2.LaunchTemplateHttpTokens_REQUIRED,
//   	InstanceMetadataTags: jsii.Boolean(true),
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

