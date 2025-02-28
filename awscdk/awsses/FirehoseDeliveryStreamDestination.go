package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
)

// An object that defines an Amazon Kinesis Data Firehose destination for email events.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//
//   var myConfigurationSet configurationSet
//   var firehoseDeliveryStream iDeliveryStream
//   var iamRole iRole
//
//
//   // Create IAM Role automatically
//   myConfigurationSet.AddEventDestination(jsii.String("ToFirehose"), &ConfigurationSetEventDestinationOptions{
//   	Destination: ses.EventDestination_FirehoseDeliveryStream(&FirehoseDeliveryStreamDestination{
//   		DeliveryStream: firehoseDeliveryStream,
//   	}),
//   })
//
//   // Specify your IAM Role
//   myConfigurationSet.AddEventDestination(jsii.String("ToFirehose"), &ConfigurationSetEventDestinationOptions{
//   	Destination: ses.EventDestination_*FirehoseDeliveryStream(&FirehoseDeliveryStreamDestination{
//   		DeliveryStream: firehoseDeliveryStream,
//   		Role: iamRole,
//   	}),
//   })
//
type FirehoseDeliveryStreamDestination struct {
	// The Amazon Kinesis Data Firehose stream that the Amazon SES API v2 sends email events to.
	DeliveryStream awskinesisfirehose.IDeliveryStream `field:"required" json:"deliveryStream" yaml:"deliveryStream"`
	// The IAM role that the Amazon SES API v2 uses to send email events to the Amazon Kinesis Data Firehose stream.
	// Default: - Create IAM Role for Kinesis Data Firehose Delivery stream.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

