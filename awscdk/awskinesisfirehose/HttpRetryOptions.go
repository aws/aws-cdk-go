package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Describes the retry behavior in case Kinesis Data Firehose is unable to deliver data to the specified Http endpoint destination, or if it doesn't receive a valid acknowledgment of receipt from the specified Http endpoint destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRetryOptions := &HttpRetryOptions{
//   	Duration: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type HttpRetryOptions struct {
	// The total amount of time that Kinesis Data Firehose spends on retries.
	Duration awscdk.Duration `field:"required" json:"duration" yaml:"duration"`
}

