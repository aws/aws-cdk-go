package previewawsiotanalyticsevents


// Type definition for CustomerManagedS3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customerManagedS3 := &CustomerManagedS3{
//   	Bucket: []*string{
//   		jsii.String("bucket"),
//   	},
//   	KeyPrefix: []*string{
//   		jsii.String("keyPrefix"),
//   	},
//   	RoleArn: []*string{
//   		jsii.String("roleArn"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_CustomerManagedS3 struct {
	// bucket property.
	//
	// Specify an array of string values to match this event if the actual value of bucket is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Bucket *[]*string `field:"optional" json:"bucket" yaml:"bucket"`
	// keyPrefix property.
	//
	// Specify an array of string values to match this event if the actual value of keyPrefix is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KeyPrefix *[]*string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
	// roleArn property.
	//
	// Specify an array of string values to match this event if the actual value of roleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RoleArn *[]*string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

