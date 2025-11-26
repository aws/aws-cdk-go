package previewawsguarddutyevents


// Type definition for AffectedResources_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   affectedResources1 := &AffectedResources1{
//   	AwsCloudTrailTrail: []*string{
//   		jsii.String("awsCloudTrailTrail"),
//   	},
//   	AwsEc2Instance: []*string{
//   		jsii.String("awsEc2Instance"),
//   	},
//   	AwsS3Bucket: []*string{
//   		jsii.String("awsS3Bucket"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_AffectedResources1 struct {
	// AWS-CloudTrail-Trail property.
	//
	// Specify an array of string values to match this event if the actual value of AWS-CloudTrail-Trail is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsCloudTrailTrail *[]*string `field:"optional" json:"awsCloudTrailTrail" yaml:"awsCloudTrailTrail"`
	// AWS-EC2-Instance property.
	//
	// Specify an array of string values to match this event if the actual value of AWS-EC2-Instance is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsEc2Instance *[]*string `field:"optional" json:"awsEc2Instance" yaml:"awsEc2Instance"`
	// AWS-S3-Bucket property.
	//
	// Specify an array of string values to match this event if the actual value of AWS-S3-Bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsS3Bucket *[]*string `field:"optional" json:"awsS3Bucket" yaml:"awsS3Bucket"`
}

