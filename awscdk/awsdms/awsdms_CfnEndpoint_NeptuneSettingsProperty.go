package awsdms


// Provides information that defines an Amazon Neptune endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For more information about the available settings, see [Specifying endpoint settings for Amazon Neptune as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   neptuneSettingsProperty := &neptuneSettingsProperty{
//   	errorRetryDuration: jsii.Number(123),
//   	iamAuthEnabled: jsii.Boolean(false),
//   	maxFileSize: jsii.Number(123),
//   	maxRetryCount: jsii.Number(123),
//   	s3BucketFolder: jsii.String("s3BucketFolder"),
//   	s3BucketName: jsii.String("s3BucketName"),
//   	serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   }
//
type CfnEndpoint_NeptuneSettingsProperty struct {
	// The number of milliseconds for AWS DMS to wait to retry a bulk-load of migrated graph data to the Neptune target database before raising an error.
	//
	// The default is 250.
	ErrorRetryDuration *float64 `field:"optional" json:"errorRetryDuration" yaml:"errorRetryDuration"`
	// If you want IAM authorization enabled for this endpoint, set this parameter to `true` .
	//
	// Then attach the appropriate IAM policy document to your service role specified by `ServiceAccessRoleArn` . The default is `false` .
	IamAuthEnabled interface{} `field:"optional" json:"iamAuthEnabled" yaml:"iamAuthEnabled"`
	// The maximum size in kilobytes of migrated graph data stored in a .csv file before AWS DMS bulk-loads the data to the Neptune target database. The default is 1,048,576 KB. If the bulk load is successful, AWS DMS clears the bucket, ready to store the next batch of migrated graph data.
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// The number of times for AWS DMS to retry a bulk load of migrated graph data to the Neptune target database before raising an error.
	//
	// The default is 5.
	MaxRetryCount *float64 `field:"optional" json:"maxRetryCount" yaml:"maxRetryCount"`
	// A folder path where you want AWS DMS to store migrated graph data in the S3 bucket specified by `S3BucketName`.
	S3BucketFolder *string `field:"optional" json:"s3BucketFolder" yaml:"s3BucketFolder"`
	// The name of the Amazon S3 bucket where AWS DMS can temporarily store migrated graph data in .csv files before bulk-loading it to the Neptune target database. AWS DMS maps the SQL source data to graph data before storing it in these .csv files.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// The Amazon Resource Name (ARN) of the service role that you created for the Neptune target endpoint.
	//
	// The role must allow the `iam:PassRole` action.
	//
	// For more information, see [Creating an IAM Service Role for Accessing Amazon Neptune as a Target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.ServiceRole) in the *AWS Database Migration Service User Guide* .
	ServiceAccessRoleArn *string `field:"optional" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
}

