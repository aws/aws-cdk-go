package awskinesisanalyticsv2


// The information required to deploy a Kinesis Data Analytics Studio notebook as an application with durable state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deployAsApplicationConfigurationProperty := &DeployAsApplicationConfigurationProperty{
//   	S3ContentLocation: &S3ContentBaseLocationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//
//   		// the properties below are optional
//   		BasePath: jsii.String("basePath"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-deployasapplicationconfiguration.html
//
type CfnApplication_DeployAsApplicationConfigurationProperty struct {
	// The description of an Amazon S3 object that contains the Amazon Data Analytics application, including the Amazon Resource Name (ARN) of the S3 bucket, the name of the Amazon S3 object that contains the data, and the version number of the Amazon S3 object that contains the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-deployasapplicationconfiguration.html#cfn-kinesisanalyticsv2-application-deployasapplicationconfiguration-s3contentlocation
	//
	S3ContentLocation interface{} `field:"required" json:"s3ContentLocation" yaml:"s3ContentLocation"`
}

