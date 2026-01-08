package awsarcregionswitch


// Configuration for report output destinations used in a Region switch plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportOutputConfigurationProperty := &ReportOutputConfigurationProperty{
//   	S3Configuration: &S3ReportOutputConfigurationProperty{
//   		BucketOwner: jsii.String("bucketOwner"),
//   		BucketPath: jsii.String("bucketPath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-reportoutputconfiguration.html
//
type CfnPlan_ReportOutputConfigurationProperty struct {
	// Configuration for delivering reports to an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-reportoutputconfiguration.html#cfn-arcregionswitch-plan-reportoutputconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
}

