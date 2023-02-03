package awss3deployment


// Storage class used for storing the object.
//
// Example:
//   websiteBucket := s3.NewBucket(this, jsii.String("WebsiteBucket"), &bucketProps{
//   	websiteIndexDocument: jsii.String("index.html"),
//   	publicReadAccess: jsii.Boolean(true),
//   })
//
//   s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &bucketDeploymentProps{
//   	sources: []iSource{
//   		s3deploy.source.asset(jsii.String("./website-dist")),
//   	},
//   	destinationBucket: websiteBucket,
//   	destinationKeyPrefix: jsii.String("web/static"),
//   	 // optional prefix in destination bucket
//   	metadata: &userDefinedObjectMetadata{
//   		a: jsii.String("1"),
//   		b: jsii.String("2"),
//   	},
//   	 // user-defined metadata
//
//   	// system-defined metadata
//   	contentType: jsii.String("text/html"),
//   	contentLanguage: jsii.String("en"),
//   	storageClass: s3deploy.storageClass_INTELLIGENT_TIERING,
//   	serverSideEncryption: s3deploy.serverSideEncryption_AES_256,
//   	cacheControl: []cacheControl{
//   		s3deploy.*cacheControl.setPublic(),
//   		s3deploy.*cacheControl.maxAge(awscdk.Duration.hours(jsii.Number(1))),
//   	},
//   	accessControl: s3.bucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
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

