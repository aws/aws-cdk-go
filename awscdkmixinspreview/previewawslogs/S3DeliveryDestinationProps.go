package previewawslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Properties for S3 delivery destination.
//
// Example:
//   import _ "github.com/aws-samples/dummy/awscdkmixinspreview/with"
//   import logDestinations "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import cloudfrontMixins "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   // Create CloudFront distribution
//   var origin IBucket
//
//
//   destinationAccount := "123456789012"
//   sourceAccount := "234567890123"
//   region := "us-east-1"
//
//   app := awscdk.NewApp()
//
//   destStack := awscdk.NewStack(app, jsii.String("destination-stack"), &StackProps{
//   	Env: &Environment{
//   		Account: destinationAccount,
//   		Region: jsii.String(*Region),
//   	},
//   })
//
//   // Create destination bucket
//   destBucket := s3.NewBucket(destStack, jsii.String("DeliveryBucket"))
//   logDestinations.NewS3DeliveryDestination(destStack, jsii.String("Destination"), &S3DeliveryDestinationProps{
//   	Bucket: destBucket,
//   	SourceAccountId: sourceAccount,
//   })
//
//   sourceStack := awscdk.NewStack(app, jsii.String("source-stack"), &StackProps{
//   	Env: &Environment{
//   		Account: sourceAccount,
//   		Region: jsii.String(*Region),
//   	},
//   })
//   distribution := cloudfront.NewDistribution(sourceStack, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.S3BucketOrigin_WithOriginAccessControl(origin),
//   	},
//   })
//
//   destination := logs.CfnDeliveryDestination_FromDeliveryDestinationArn(sourceStack, jsii.String("Destination"), jsii.String("arn of Delivery Destination in destinationAccount"))
//
//   distribution.With(cloudfrontMixins.CfnDistributionLogsMixin_CONNECTION_LOGS().ToDestination(destination))
//
// Experimental.
type S3DeliveryDestinationProps struct {
	// The S3 bucket to deliver logs to.
	// Experimental.
	Bucket interfacesawss3.IBucketRef `field:"required" json:"bucket" yaml:"bucket"`
	// KMS key to use for encrypting logs in the S3 bucket.
	//
	// When provided, grants the logs delivery service permissions to use the key.
	// Default: - No encryption key is configured.
	//
	// Experimental.
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Format of the logs that are sent to this delivery destination.
	// Experimental.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// The permissions version ('V1' or 'V2') to be used for this delivery.
	//
	// Depending on the source of the logs, different permissions are required.
	// Default: "V2".
	//
	// Experimental.
	PermissionsVersion S3LogsDeliveryPermissionsVersion `field:"optional" json:"permissionsVersion" yaml:"permissionsVersion"`
	// Optional acount id for account the delivery source is in for cross account Vended Logs.
	// Experimental.
	SourceAccountId *string `field:"optional" json:"sourceAccountId" yaml:"sourceAccountId"`
}

