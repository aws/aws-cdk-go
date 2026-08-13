package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The buffering options that can be used before data is delivered to the specified destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var size Size
//
//   httpBufferingHints := &HttpBufferingHints{
//   	Interval: cdk.Duration_Minutes(jsii.Number(30)),
//   	Size: size,
//   }
//
type HttpBufferingHints struct {
	// The higher interval allows more time to collect data and the size of data may be bigger.
	//
	// The lower interval sends the data more frequently and may be more advantageous when looking at shorter cycles of data activity.
	// Default: 300 seconds.
	//
	Interval awscdk.Duration `field:"optional" json:"interval" yaml:"interval"`
	// The higher buffer size may be lower in cost with higher latency.
	//
	// The lower buffer size will be faster in delivery with higher cost and less latency.
	// Default: 5 MiB.
	//
	Size awscdk.Size `field:"optional" json:"size" yaml:"size"`
}

