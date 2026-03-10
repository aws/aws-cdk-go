package previewawsiotanalyticsevents


// Type definition for ChannelStorage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   channelStorage := &ChannelStorage{
//   	CustomerManagedS3: &CustomerManagedS3{
//   		Bucket: []*string{
//   			jsii.String("bucket"),
//   		},
//   		KeyPrefix: []*string{
//   			jsii.String("keyPrefix"),
//   		},
//   		RoleArn: []*string{
//   			jsii.String("roleArn"),
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ChannelStorage struct {
	// customerManagedS3 property.
	//
	// Specify an array of string values to match this event if the actual value of customerManagedS3 is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CustomerManagedS3 *AWSAPICallViaCloudTrail_CustomerManagedS3 `field:"optional" json:"customerManagedS3" yaml:"customerManagedS3"`
}

