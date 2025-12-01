package awsec2


// The state of token usage for your instance metadata requests.
//
// Example:
//   var vpc Vpc
//   var instanceType InstanceType
//   var machineImage IMachineImage
//
//
//   // Example 1: Enforce IMDSv2 with comprehensive options
//   // Example 1: Enforce IMDSv2 with comprehensive options
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//   	HttpEndpoint: jsii.Boolean(true),
//   	HttpProtocolIpv6: jsii.Boolean(false),
//   	HttpPutResponseHopLimit: jsii.Number(2),
//   	HttpTokens: ec2.HttpTokens_REQUIRED,
//   	InstanceMetadataTags: jsii.Boolean(true),
//   })
//
//   // Example 2: Enforce IMDSv2 with minimal configuration
//   // Example 2: Enforce IMDSv2 with minimal configuration
//   ec2.NewInstance(this, jsii.String("SecureInstance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: MachineImage,
//   	HttpTokens: ec2.HttpTokens_REQUIRED,
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-metadataoptions.html#cfn-ec2-instance-metadataoptions-httptokens
//
type HttpTokens string

const (
	// If the state is optional, you can choose to retrieve instance metadata with or without a signed token header on your request.
	HttpTokens_OPTIONAL HttpTokens = "OPTIONAL"
	// If the state is required, you must send a signed token header with any instance metadata retrieval requests.
	//
	// In this state,
	// retrieving the IAM role credentials always returns the version 2.0 credentials; the version 1.0 credentials are not available.
	HttpTokens_REQUIRED HttpTokens = "REQUIRED"
)

