package awsec2


// Specifies the metadata options for the instance.
//
// `MetadataOptions` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataOptionsProperty := &metadataOptionsProperty{
//   	httpEndpoint: jsii.String("httpEndpoint"),
//   	httpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   	httpPutResponseHopLimit: jsii.Number(123),
//   	httpTokens: jsii.String("httpTokens"),
//   	instanceMetadataTags: jsii.String("instanceMetadataTags"),
//   }
//
type CfnLaunchTemplate_MetadataOptionsProperty struct {
	// Enables or disables the HTTP metadata endpoint on your instances.
	//
	// If the parameter is not specified, the default state is `enabled` .
	//
	// > If you specify a value of `disabled` , you will not be able to access your instance metadata.
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Enables or disables the IPv6 endpoint for the instance metadata service.
	//
	// Default: `disabled`.
	HttpProtocolIpv6 *string `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// The desired HTTP PUT response hop limit for instance metadata requests.
	//
	// The larger the number, the further instance metadata requests can travel.
	//
	// Default: 1
	//
	// Possible values: Integers from 1 to 64.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// The state of token usage for your instance metadata requests.
	//
	// If the parameter is not specified in the request, the default state is `optional` .
	//
	// If the state is `optional` , you can choose to retrieve instance metadata with or without a signed token header on your request. If you retrieve the IAM role credentials without a token, the version 1.0 role credentials are returned. If you retrieve the IAM role credentials using a valid signed token, the version 2.0 role credentials are returned.
	//
	// If the state is `required` , you must send a signed token header with any instance metadata retrieval requests. In this state, retrieving the IAM role credentials always returns the version 2.0 credentials; the version 1.0 credentials are not available.
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Set to `enabled` to allow access to instance tags from the instance metadata.
	//
	// Set to `disabled` to turn off access to instance tags from the instance metadata. For more information, see [Work with instance tags using the instance metadata](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#work-with-tags-in-IMDS) .
	//
	// Default: `disabled`.
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

