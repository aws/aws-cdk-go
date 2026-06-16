package awsresiliencehubv2


// Configuration for a report output destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportOutputConfigurationProperty := &ReportOutputConfigurationProperty{
//   	S3: &S3ReportOutputConfigurationProperty{
//   		BucketOwner: jsii.String("bucketOwner"),
//   		BucketPath: jsii.String("bucketPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-reportoutputconfiguration.html
//
type CfnService_ReportOutputConfigurationProperty struct {
	// S3 configuration for report output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-reportoutputconfiguration.html#cfn-resiliencehubv2-service-reportoutputconfiguration-s3
	//
	S3 interface{} `field:"required" json:"s3" yaml:"s3"`
}

