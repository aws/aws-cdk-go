package awsec2


// AWS Verified Access provides server side encryption by default to data at rest using AWS -owned KMS keys.
//
// You also have the option of using customer managed KMS keys, which can be specified using the options below.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sseSpecificationProperty := &SseSpecificationProperty{
//   	CustomerManagedKeyEnabled: jsii.Boolean(false),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-ssespecification.html
//
type CfnVerifiedAccessEndpoint_SseSpecificationProperty struct {
	// Enable or disable the use of customer managed KMS keys for server side encryption.
	//
	// Valid values: `True` | `False`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-ssespecification.html#cfn-ec2-verifiedaccessendpoint-ssespecification-customermanagedkeyenabled
	//
	CustomerManagedKeyEnabled interface{} `field:"optional" json:"customerManagedKeyEnabled" yaml:"customerManagedKeyEnabled"`
	// The ARN of the KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccessendpoint-ssespecification.html#cfn-ec2-verifiedaccessendpoint-ssespecification-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

