package awsec2


// Specifies the metadata options for the instance.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-metadataoptions.html
//
type CfnInstance_MetadataOptionsProperty struct {
	// Enables or disables the HTTP metadata endpoint on your instances.
	//
	// If you specify a value of `disabled` , you cannot access your instance metadata.
	//
	// Default: `enabled`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-metadataoptions.html#cfn-ec2-instance-metadataoptions-httpendpoint
	//
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Enables or disables the IPv6 endpoint for the instance metadata service.
	//
	// Default: `disabled`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-metadataoptions.html#cfn-ec2-instance-metadataoptions-httpprotocolipv6
	//
	HttpProtocolIpv6 *string `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// The maximum number of hops that the metadata token can travel.
	//
	// Possible values: Integers from 1 to 64.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-metadataoptions.html#cfn-ec2-instance-metadataoptions-httpputresponsehoplimit
	//
	// Default: - 1.
	//
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Indicates whether IMDSv2 is required.
	//
	// - `optional` - IMDSv2 is optional, which means that you can use either IMDSv2 or IMDSv1.
	// - `required` - IMDSv2 is required, which means that IMDSv1 is disabled, and you must use IMDSv2.
	//
	// Default:
	//
	// - If the value of `ImdsSupport` for the Amazon Machine Image (AMI) for your instance is `v2.0` and the account level default is set to `no-preference` , the default is `required` .
	// - If the value of `ImdsSupport` for the Amazon Machine Image (AMI) for your instance is `v2.0` , but the account level default is set to `V1 or V2` , the default is `optional` .
	//
	// The default value can also be affected by other combinations of parameters. For more information, see [Order of precedence for instance metadata options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-options.html#instance-metadata-options-order-of-precedence) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-metadataoptions.html#cfn-ec2-instance-metadataoptions-httptokens
	//
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Set to `enabled` to allow access to instance tags from the instance metadata.
	//
	// Set to `disabled` to turn off access to instance tags from the instance metadata. For more information, see [Work with instance tags using the instance metadata](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#work-with-tags-in-IMDS) .
	//
	// Default: `disabled`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-metadataoptions.html#cfn-ec2-instance-metadataoptions-instancemetadatatags
	//
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

