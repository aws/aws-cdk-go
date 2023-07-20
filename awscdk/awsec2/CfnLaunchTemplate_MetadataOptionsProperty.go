package awsec2


// The metadata options for the instance.
//
// For more information, see [Instance metadata and user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) in the *Amazon EC2 User Guide* .
//
// `MetadataOptions` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataOptionsProperty := &MetadataOptionsProperty{
//   	HttpEndpoint: jsii.String("httpEndpoint"),
//   	HttpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   	HttpPutResponseHopLimit: jsii.Number(123),
//   	HttpTokens: jsii.String("httpTokens"),
//   	InstanceMetadataTags: jsii.String("instanceMetadataTags"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-metadataoptions.html
//
type CfnLaunchTemplate_MetadataOptionsProperty struct {
	// Enables or disables the HTTP metadata endpoint on your instances.
	//
	// If the parameter is not specified, the default state is `enabled` .
	//
	// > If you specify a value of `disabled` , you will not be able to access your instance metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-metadataoptions.html#cfn-ec2-launchtemplate-metadataoptions-httpendpoint
	//
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Enables or disables the IPv6 endpoint for the instance metadata service.
	//
	// Default: `disabled`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-metadataoptions.html#cfn-ec2-launchtemplate-metadataoptions-httpprotocolipv6
	//
	HttpProtocolIpv6 *string `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// The desired HTTP PUT response hop limit for instance metadata requests.
	//
	// The larger the number, the further instance metadata requests can travel.
	//
	// Default: `1`
	//
	// Possible values: Integers from 1 to 64.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-metadataoptions.html#cfn-ec2-launchtemplate-metadataoptions-httpputresponsehoplimit
	//
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// IMDSv2 uses token-backed sessions.
	//
	// Set the use of HTTP tokens to `optional` (in other words, set the use of IMDSv2 to `optional` ) or `required` (in other words, set the use of IMDSv2 to `required` ).
	//
	// - `optional` - When IMDSv2 is optional, you can choose to retrieve instance metadata with or without a session token in your request. If you retrieve the IAM role credentials without a token, the IMDSv1 role credentials are returned. If you retrieve the IAM role credentials using a valid session token, the IMDSv2 role credentials are returned.
	// - `required` - When IMDSv2 is required, you must send a session token with any instance metadata retrieval requests. In this state, retrieving the IAM role credentials always returns IMDSv2 credentials; IMDSv1 credentials are not available.
	//
	// Default: `optional`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-metadataoptions.html#cfn-ec2-launchtemplate-metadataoptions-httptokens
	//
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Set to `enabled` to allow access to instance tags from the instance metadata.
	//
	// Set to `disabled` to turn off access to instance tags from the instance metadata. For more information, see [Work with instance tags using the instance metadata](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#work-with-tags-in-IMDS) .
	//
	// Default: `disabled`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-metadataoptions.html#cfn-ec2-launchtemplate-metadataoptions-instancemetadatatags
	//
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

