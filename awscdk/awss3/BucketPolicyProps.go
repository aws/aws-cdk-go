package awss3

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Example:
//   import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"
//   import cloudfrontMixins "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   // Create CloudFront distribution
//   var bucket Bucket
//
//   distribution := cloudfront.NewDistribution(scope, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.S3BucketOrigin_WithOriginAccessControl(bucket),
//   	},
//   })
//
//   // Create destination bucket
//   destBucket := s3.NewBucket(scope, jsii.String("DeliveryBucket"))
//   // Add permissions to bucket to facilitate log delivery
//   bucketPolicy := s3.NewBucketPolicy(scope, jsii.String("DeliveryBucketPolicy"), &BucketPolicyProps{
//   	Bucket: destBucket,
//   	Document: iam.NewPolicyDocument(),
//   })
//   // Create S3 delivery destination for logs
//   destination := logs.NewCfnDeliveryDestination(scope, jsii.String("Destination"), &CfnDeliveryDestinationProps{
//   	DestinationResourceArn: destBucket.bucketArn,
//   	Name: jsii.String("unique-destination-name"),
//   	DeliveryDestinationType: jsii.String("S3"),
//   })
//
//   distribution.With(cloudfrontMixins.CfnDistributionLogsMixin_CONNECTION_LOGS().ToDestination(destination))
//
type BucketPolicyProps struct {
	// The Amazon S3 bucket that the policy applies to.
	Bucket IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Policy document to apply to the bucket.
	// Default: - A new empty PolicyDocument will be created.
	//
	Document awsiam.PolicyDocument `field:"optional" json:"document" yaml:"document"`
	// Policy to apply when the policy is removed from this stack.
	// Default: - RemovalPolicy.DESTROY.
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

