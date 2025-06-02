package awsguardduty


// Contains the Amazon Resource Name (ARN) of the resource that receives the published findings, such as an S3 bucket, and the ARN of the KMS key that is used to encrypt these published findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNDestinationPropertiesProperty := &CFNDestinationPropertiesProperty{
//   	DestinationArn: jsii.String("destinationArn"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-publishingdestination-cfndestinationproperties.html
//
type CfnPublishingDestination_CFNDestinationPropertiesProperty struct {
	// The ARN of the resource where the findings are published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-publishingdestination-cfndestinationproperties.html#cfn-guardduty-publishingdestination-cfndestinationproperties-destinationarn
	//
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// The ARN of the KMS key to use for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-publishingdestination-cfndestinationproperties.html#cfn-guardduty-publishingdestination-cfndestinationproperties-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

