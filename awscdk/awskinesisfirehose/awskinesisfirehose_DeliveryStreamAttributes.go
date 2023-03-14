package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
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
//   deliveryStreamAttributes := &deliveryStreamAttributes{
//   	deliveryStreamArn: jsii.String("deliveryStreamArn"),
//   	deliveryStreamName: jsii.String("deliveryStreamName"),
//   	role: role,
//   }
//
// Experimental.
type DeliveryStreamAttributes struct {
	// The ARN of the delivery stream.
	//
	// At least one of deliveryStreamArn and deliveryStreamName must be provided.
	// Experimental.
	DeliveryStreamArn *string `field:"optional" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// The name of the delivery stream.
	//
	// At least one of deliveryStreamName and deliveryStreamArn  must be provided.
	// Experimental.
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// The IAM role associated with this delivery stream.
	//
	// Assumed by Kinesis Data Firehose to read from sources and encrypt data server-side.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

