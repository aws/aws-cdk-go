package awscloudfront


// Properties for creating a S3 Origin Access Control resource.
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   oac := cloudfront.NewS3OriginAccessControl(this, jsii.String("MyOAC"), &S3OriginAccessControlProps{
//   	Signing: cloudfront.Signing_SIGV4_NO_OVERRIDE(),
//   })
//   s3Origin := origins.S3BucketOrigin_WithOriginAccessControl(myBucket, &S3BucketOriginWithOACProps{
//   	OriginAccessControl: oac,
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: s3Origin,
//   	},
//   })
//
type S3OriginAccessControlProps struct {
	// A description of the origin access control.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name to identify the origin access control, with a maximum length of 64 characters.
	// Default: - a generated name.
	//
	OriginAccessControlName *string `field:"optional" json:"originAccessControlName" yaml:"originAccessControlName"`
	// Specifies which requests CloudFront signs and the signing protocol.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html#cfn-cloudfront-originaccesscontrol-originaccesscontrolconfig-signingbehavior
	//
	// Default: SIGV4_ALWAYS.
	//
	Signing Signing `field:"optional" json:"signing" yaml:"signing"`
}

