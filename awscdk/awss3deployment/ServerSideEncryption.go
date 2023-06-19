package awss3deployment


// Indicates whether server-side encryption is enabled for the object, and whether that encryption is from the AWS Key Management Service (AWS KMS) or from Amazon S3 managed encryption (SSE-S3).
//
// Example:
//   websiteBucket := s3.NewBucket(this, jsii.String("WebsiteBucket"), &BucketProps{
//   	WebsiteIndexDocument: jsii.String("index.html"),
//   	PublicReadAccess: jsii.Boolean(true),
//   })
//
//   s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &BucketDeploymentProps{
//   	Sources: []iSource{
//   		s3deploy.Source_Asset(jsii.String("./website-dist")),
//   	},
//   	DestinationBucket: websiteBucket,
//   	DestinationKeyPrefix: jsii.String("web/static"),
//   	 // optional prefix in destination bucket
//   	Metadata: &UserDefinedObjectMetadata{
//   		A: jsii.String("1"),
//   		B: jsii.String("2"),
//   	},
//   	 // user-defined metadata
//
//   	// system-defined metadata
//   	ContentType: jsii.String("text/html"),
//   	ContentLanguage: jsii.String("en"),
//   	StorageClass: s3deploy.StorageClass_INTELLIGENT_TIERING,
//   	ServerSideEncryption: s3deploy.ServerSideEncryption_AES_256,
//   	CacheControl: []cacheControl{
//   		s3deploy.*cacheControl_SetPublic(),
//   		s3deploy.*cacheControl_MaxAge(awscdk.Duration_Hours(jsii.Number(1))),
//   	},
//   	AccessControl: s3.BucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Experimental.
type ServerSideEncryption string

const (
	// 'AES256'.
	// Experimental.
	ServerSideEncryption_AES_256 ServerSideEncryption = "AES_256"
	// 'aws:kms'.
	// Experimental.
	ServerSideEncryption_AWS_KMS ServerSideEncryption = "AWS_KMS"
)

