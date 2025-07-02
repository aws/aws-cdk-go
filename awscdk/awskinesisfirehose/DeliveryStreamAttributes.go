package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// A full specification of a delivery stream that can be used to import it fluently into the CDK application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   deliveryStreamAttributes := &DeliveryStreamAttributes{
//   	DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   	DeliveryStreamName: jsii.String("deliveryStreamName"),
//   	Role: role,
//   }
//
type DeliveryStreamAttributes struct {
	// The ARN of the delivery stream.
	//
	// At least one of deliveryStreamArn and deliveryStreamName must be provided.
	// Default: - derived from `deliveryStreamName`.
	//
	DeliveryStreamArn *string `field:"optional" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// The name of the delivery stream.
	//
	// At least one of deliveryStreamName and deliveryStreamArn  must be provided.
	// Default: - derived from `deliveryStreamArn`.
	//
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Amazon Data Firehose to read from sources and encrypt data server-side.
	// Default: - the imported stream cannot be granted access to other resources as an `iam.IGrantable`.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

