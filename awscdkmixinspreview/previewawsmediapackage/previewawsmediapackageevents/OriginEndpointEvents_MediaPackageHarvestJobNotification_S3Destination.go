package previewawsmediapackageevents


// Type definition for S3_destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3Destination := &S3Destination{
//   	BucketName: []*string{
//   		jsii.String("bucketName"),
//   	},
//   	ManifestKey: []*string{
//   		jsii.String("manifestKey"),
//   	},
//   	RoleArn: []*string{
//   		jsii.String("roleArn"),
//   	},
//   }
//
// Experimental.
type OriginEndpointEvents_MediaPackageHarvestJobNotification_S3Destination struct {
	// bucket_name property.
	//
	// Specify an array of string values to match this event if the actual value of bucket_name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BucketName *[]*string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// manifest_key property.
	//
	// Specify an array of string values to match this event if the actual value of manifest_key is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ManifestKey *[]*string `field:"optional" json:"manifestKey" yaml:"manifestKey"`
	// role_arn property.
	//
	// Specify an array of string values to match this event if the actual value of role_arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RoleArn *[]*string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

