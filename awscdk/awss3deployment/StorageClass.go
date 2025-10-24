package awss3deployment


// Storage class used for storing the object.
//
// Example:
//   websiteBucket := s3.NewBucket(this, jsii.String("WebsiteBucket"), &BucketProps{
//   	WebsiteIndexDocument: jsii.String("index.html"),
//   	PublicReadAccess: jsii.Boolean(true),
//   })
//
//   s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &BucketDeploymentProps{
//   	Sources: []ISource{
//   		s3deploy.Source_Asset(jsii.String("./website-dist")),
//   	},
//   	DestinationBucket: websiteBucket,
//   	DestinationKeyPrefix: jsii.String("web/static"),
//   	 // optional prefix in destination bucket
//   	Metadata: map[string]*string{
//   		"A": jsii.String("1"),
//   		"b": jsii.String("2"),
//   	},
//   	 // user-defined metadata
//
//   	// system-defined metadata
//   	ContentType: jsii.String("text/html"),
//   	ContentLanguage: jsii.String("en"),
//   	StorageClass: s3deploy.StorageClass_INTELLIGENT_TIERING,
//   	ServerSideEncryption: s3deploy.ServerSideEncryption_AES_256,
//   	CacheControl: []CacheControl{
//   		s3deploy.CacheControl_SetPublic(),
//   		s3deploy.CacheControl_MaxAge(awscdk.Duration_Hours(jsii.Number(1))),
//   	},
//   	AccessControl: s3.BucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
type StorageClass string

const (
	// 'STANDARD'.
	StorageClass_STANDARD StorageClass = "STANDARD"
	// 'REDUCED_REDUNDANCY'.
	StorageClass_REDUCED_REDUNDANCY StorageClass = "REDUCED_REDUNDANCY"
	// 'STANDARD_IA'.
	StorageClass_STANDARD_IA StorageClass = "STANDARD_IA"
	// 'ONEZONE_IA'.
	StorageClass_ONEZONE_IA StorageClass = "ONEZONE_IA"
	// 'INTELLIGENT_TIERING'.
	StorageClass_INTELLIGENT_TIERING StorageClass = "INTELLIGENT_TIERING"
	// 'GLACIER'.
	StorageClass_GLACIER StorageClass = "GLACIER"
	// 'DEEP_ARCHIVE'.
	StorageClass_DEEP_ARCHIVE StorageClass = "DEEP_ARCHIVE"
)

