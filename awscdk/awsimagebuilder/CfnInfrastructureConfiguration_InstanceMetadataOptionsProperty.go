package awsimagebuilder


// The instance metadata options that apply to the HTTP requests that pipeline builds use to launch EC2 build and test instances.
//
// For more information about instance metadata options, see [Configure the instance metadata options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-options.html) in the **Amazon EC2 User Guide** for Linux instances, or [Configure the instance metadata options](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/configuring-instance-metadata-options.html) in the **Amazon EC2 Windows Guide** for Windows instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceMetadataOptionsProperty := &InstanceMetadataOptionsProperty{
//   	HttpPutResponseHopLimit: jsii.Number(123),
//   	HttpTokens: jsii.String("httpTokens"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-instancemetadataoptions.html
//
type CfnInfrastructureConfiguration_InstanceMetadataOptionsProperty struct {
	// Limit the number of hops that an instance metadata request can traverse to reach its destination.
	//
	// The default is one hop. However, if HTTP tokens are required, container image builds need a minimum of two hops.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-instancemetadataoptions.html#cfn-imagebuilder-infrastructureconfiguration-instancemetadataoptions-httpputresponsehoplimit
	//
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Indicates whether a signed token header is required for instance metadata retrieval requests.
	//
	// The values affect the response as follows:
	//
	// - *required* – When you retrieve the IAM role credentials, version 2.0 credentials are returned in all cases.
	// - *optional* – You can include a signed token header in your request to retrieve instance metadata, or you can leave it out. If you include it, version 2.0 credentials are returned for the IAM role. Otherwise, version 1.0 credentials are returned.
	//
	// The default setting is *optional* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-instancemetadataoptions.html#cfn-imagebuilder-infrastructureconfiguration-instancemetadataoptions-httptokens
	//
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
}

