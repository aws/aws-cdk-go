package awscloudfront


// Properties of CloudFront OriginAccessIdentity.
//
// Example:
//   myBucket := s3.NewBucket(this, jsii.String("myBucket"))
//   myOai := cloudfront.NewOriginAccessIdentity(this, jsii.String("myOAI"), &OriginAccessIdentityProps{
//   	Comment: jsii.String("My custom OAI"),
//   })
//   s3Origin := origins.S3BucketOrigin_WithOriginAccessIdentity(myBucket, &S3BucketOriginWithOAIProps{
//   	OriginAccessIdentity: myOai,
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: s3Origin,
//   	},
//   })
//
type OriginAccessIdentityProps struct {
	// Any comments you want to include about the origin access identity.
	// Default: "Allows CloudFront to reach the bucket".
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

