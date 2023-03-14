package awskendra


// User accounts whose documents should be indexed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oneDriveUsersProperty := &oneDriveUsersProperty{
//   	oneDriveUserList: []*string{
//   		jsii.String("oneDriveUserList"),
//   	},
//   	oneDriveUserS3Path: &s3PathProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//   }
//
type CfnDataSource_OneDriveUsersProperty struct {
	// A list of users whose documents should be indexed.
	//
	// Specify the user names in email format, for example, `username@tenantdomain` . If you need to index the documents of more than 100 users, use the `OneDriveUserS3Path` field to specify the location of a file containing a list of users.
	OneDriveUserList *[]*string `field:"optional" json:"oneDriveUserList" yaml:"oneDriveUserList"`
	// The S3 bucket location of a file containing a list of users whose documents should be indexed.
	OneDriveUserS3Path interface{} `field:"optional" json:"oneDriveUserS3Path" yaml:"oneDriveUserS3Path"`
}

