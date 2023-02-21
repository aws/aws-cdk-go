package awscloudfront


// A list of AWS accounts whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trustedSignersProperty := &trustedSignersProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	awsAccountNumbers: []*string{
//   		jsii.String("awsAccountNumbers"),
//   	},
//   }
//
type CfnStreamingDistribution_TrustedSignersProperty struct {
	// This field is `true` if any of the AWS accounts have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies.
	//
	// If not, this field is `false` .
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// An AWS account number that contains active CloudFront key pairs that CloudFront can use to verify the signatures of signed URLs and signed cookies.
	//
	// If the AWS account that owns the key pairs is the same account that owns the CloudFront distribution, the value of this field is `self` .
	AwsAccountNumbers *[]*string `field:"optional" json:"awsAccountNumbers" yaml:"awsAccountNumbers"`
}

