package mixinsawskendra


// User accounts whose documents should be indexed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   oneDriveUsersProperty := &OneDriveUsersProperty{
//   	OneDriveUserList: []*string{
//   		jsii.String("oneDriveUserList"),
//   	},
//   	OneDriveUserS3Path: &S3PathProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveusers.html
//
type CfnDataSourcePropsMixin_OneDriveUsersProperty struct {
	// A list of users whose documents should be indexed.
	//
	// Specify the user names in email format, for example, `username@tenantdomain` . If you need to index the documents of more than 10 users, use the `OneDriveUserS3Path` field to specify the location of a file containing a list of users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveusers.html#cfn-kendra-datasource-onedriveusers-onedriveuserlist
	//
	OneDriveUserList *[]*string `field:"optional" json:"oneDriveUserList" yaml:"oneDriveUserList"`
	// The S3 bucket location of a file containing a list of users whose documents should be indexed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-onedriveusers.html#cfn-kendra-datasource-onedriveusers-onedriveusers3path
	//
	OneDriveUserS3Path interface{} `field:"optional" json:"oneDriveUserS3Path" yaml:"oneDriveUserS3Path"`
}

