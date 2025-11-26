package previewawsmediapackageevents


// Type definition for Harvest_job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   harvestJob := &HarvestJob{
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	CreatedAt: []*string{
//   		jsii.String("createdAt"),
//   	},
//   	EndTime: []*string{
//   		jsii.String("endTime"),
//   	},
//   	Id: []*string{
//   		jsii.String("id"),
//   	},
//   	OriginEndpointId: []*string{
//   		jsii.String("originEndpointId"),
//   	},
//   	S3Destination: &S3Destination{
//   		BucketName: []*string{
//   			jsii.String("bucketName"),
//   		},
//   		ManifestKey: []*string{
//   			jsii.String("manifestKey"),
//   		},
//   		RoleArn: []*string{
//   			jsii.String("roleArn"),
//   		},
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type OriginEndpointEvents_MediaPackageHarvestJobNotification_HarvestJob struct {
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// created_at property.
	//
	// Specify an array of string values to match this event if the actual value of created_at is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreatedAt *[]*string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// end_time property.
	//
	// Specify an array of string values to match this event if the actual value of end_time is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTime *[]*string `field:"optional" json:"endTime" yaml:"endTime"`
	// id property.
	//
	// Specify an array of string values to match this event if the actual value of id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// origin_endpoint_id property.
	//
	// Specify an array of string values to match this event if the actual value of origin_endpoint_id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the OriginEndpoint reference.
	//
	// Experimental.
	OriginEndpointId *[]*string `field:"optional" json:"originEndpointId" yaml:"originEndpointId"`
	// s3_destination property.
	//
	// Specify an array of string values to match this event if the actual value of s3_destination is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	S3Destination *OriginEndpointEvents_MediaPackageHarvestJobNotification_S3Destination `field:"optional" json:"s3Destination" yaml:"s3Destination"`
	// start_time property.
	//
	// Specify an array of string values to match this event if the actual value of start_time is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}

