package previewawsworkspacesinstancesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceMetadataOptionsRequestProperty := &InstanceMetadataOptionsRequestProperty{
//   	HttpEndpoint: jsii.String("httpEndpoint"),
//   	HttpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   	HttpPutResponseHopLimit: jsii.Number(123),
//   	HttpTokens: jsii.String("httpTokens"),
//   	InstanceMetadataTags: jsii.String("instanceMetadataTags"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest.html
//
type CfnWorkspaceInstancePropsMixin_InstanceMetadataOptionsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest.html#cfn-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest-httpendpoint
	//
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest.html#cfn-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest-httpprotocolipv6
	//
	HttpProtocolIpv6 *string `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest.html#cfn-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest-httpputresponsehoplimit
	//
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest.html#cfn-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest-httptokens
	//
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest.html#cfn-workspacesinstances-workspaceinstance-instancemetadataoptionsrequest-instancemetadatatags
	//
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

