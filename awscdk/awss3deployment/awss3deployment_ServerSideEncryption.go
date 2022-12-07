package awss3deployment


// Indicates whether server-side encryption is enabled for the object, and whether that encryption is from the AWS Key Management Service (AWS KMS) or from Amazon S3 managed encryption (SSE-S3).
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
type ServerSideEncryption string

const (
	// 'AES256'.
	ServerSideEncryption_AES_256 ServerSideEncryption = "AES_256"
	// 'aws:kms'.
	ServerSideEncryption_AWS_KMS ServerSideEncryption = "AWS_KMS"
)

