package awssesactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Construction properties for a stop action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   stopProps := &stopProps{
//   	topic: topic,
//   }
//
// Experimental.
type StopProps struct {
	// The SNS topic to notify when the stop action is taken.
	// Experimental.
	Topic awssns.ITopic `field:"optional" json:"topic" yaml:"topic"`
}

