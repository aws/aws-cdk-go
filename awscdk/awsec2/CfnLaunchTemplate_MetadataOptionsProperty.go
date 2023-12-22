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
	// Indicates whether IMDSv2 is required.
	//
	// - `optional` - IMDSv2 is optional. You can choose whether to send a session token in your instance metadata retrieval requests. If you retrieve IAM role credentials without a session token, you receive the IMDSv1 role credentials. If you retrieve IAM role credentials using a valid session token, you receive the IMDSv2 role credentials.
	// - `required` - IMDSv2 is required. You must send a session token in your instance metadata retrieval requests. With this option, retrieving the IAM role credentials always returns IMDSv2 credentials; IMDSv1 credentials are not available.
	//
	// Default: If the value of `ImdsSupport` for the Amazon Machine Image (AMI) for your instance is `v2.0` , the default is `required` .
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

