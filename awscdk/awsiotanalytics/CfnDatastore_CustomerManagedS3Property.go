package awsiotanalytics


// S3-customer-managed;
//
// When you choose customer-managed storage, the `retentionPeriod` parameter is ignored. You can't change the choice of Amazon S3 storage after your data store is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerManagedS3Property := &CustomerManagedS3Property{
//   	Bucket: jsii.String("bucket"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	KeyPrefix: jsii.String("keyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-customermanageds3.html
//
type CfnDatastore_CustomerManagedS3Property struct {
	// The name of the Amazon S3 bucket where your data is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-customermanageds3.html#cfn-iotanalytics-datastore-customermanageds3-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The ARN of the role that grants ITA permission to interact with your Amazon S3 resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-customermanageds3.html#cfn-iotanalytics-datastore-customermanageds3-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// (Optional) The prefix used to create the keys of the data store data objects.
	//
	// Each object in an Amazon S3 bucket has a key that is its unique identifier in the bucket. Each object in a bucket has exactly one key. The prefix must end with a forward slash (/).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-datastore-customermanageds3.html#cfn-iotanalytics-datastore-customermanageds3-keyprefix
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

