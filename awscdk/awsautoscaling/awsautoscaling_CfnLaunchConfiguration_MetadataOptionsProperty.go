package awsautoscaling


// `MetadataOptions` is a property of [AWS::AutoScaling::LaunchConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html) that describes metadata options for the instances.
//
// For more information, see [Configure the instance metadata options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html#launch-configurations-imds) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataOptionsProperty := &metadataOptionsProperty{
//   	httpEndpoint: jsii.String("httpEndpoint"),
//   	httpPutResponseHopLimit: jsii.Number(123),
//   	httpTokens: jsii.String("httpTokens"),
//   }
//
type CfnLaunchConfiguration_MetadataOptionsProperty struct {
	// This parameter enables or disables the HTTP metadata endpoint on your instances.
	//
	// If the parameter is not specified, the default state is `enabled` .
	//
	// > If you specify a value of `disabled` , you will not be able to access your instance metadata.
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// The desired HTTP PUT response hop limit for instance metadata requests.
	//
	// The larger the number, the further instance metadata requests can travel.
	//
	// Default: 1.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// The state of token usage for your instance metadata requests.
	//
	// If the parameter is not specified in the request, the default state is `optional` .
	//
	// If the state is `optional` , you can choose to retrieve instance metadata with or without a signed token header on your request. If you retrieve the IAM role credentials without a token, the version 1.0 role credentials are returned. If you retrieve the IAM role credentials using a valid signed token, the version 2.0 role credentials are returned.
	//
	// If the state is `required` , you must send a signed token header with any instance metadata retrieval requests. In this state, retrieving the IAM role credentials always returns the version 2.0 credentials; the version 1.0 credentials are not available.
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
}

