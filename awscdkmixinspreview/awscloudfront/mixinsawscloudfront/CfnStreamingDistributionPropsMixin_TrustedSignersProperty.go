package mixinsawscloudfront


// A list of AWS accounts whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trustedSignersProperty := &TrustedSignersProperty{
//   	AwsAccountNumbers: []*string{
//   		jsii.String("awsAccountNumbers"),
//   	},
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-trustedsigners.html
//
type CfnStreamingDistributionPropsMixin_TrustedSignersProperty struct {
	// An AWS account number that contains active CloudFront key pairs that CloudFront can use to verify the signatures of signed URLs and signed cookies.
	//
	// If the AWS account that owns the key pairs is the same account that owns the CloudFront distribution, the value of this field is `self` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-trustedsigners.html#cfn-cloudfront-streamingdistribution-trustedsigners-awsaccountnumbers
	//
	AwsAccountNumbers *[]*string `field:"optional" json:"awsAccountNumbers" yaml:"awsAccountNumbers"`
	// This field is `true` if any of the AWS accounts in the list are configured as trusted signers.
	//
	// If not, this field is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-trustedsigners.html#cfn-cloudfront-streamingdistribution-trustedsigners-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

