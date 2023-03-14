package awsevidently


// A structure that contains information about where Evidently is to store evaluation events for longer term storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataDeliveryObjectProperty := &dataDeliveryObjectProperty{
//   	logGroup: jsii.String("logGroup"),
//   	s3: &s3DestinationProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
type CfnProject_DataDeliveryObjectProperty struct {
	// If the project stores evaluation events in CloudWatch Logs , this structure stores the log group name.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// If the project stores evaluation events in an Amazon S3 bucket, this structure stores the bucket name and bucket prefix.
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

