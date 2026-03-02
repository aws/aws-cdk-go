package previewawslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Properties for Cloudwatch delivery destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroupRef ILogGroupRef
//
//   cloudwatchDeliveryDestinationProps := &CloudwatchDeliveryDestinationProps{
//   	LogGroup: logGroupRef,
//
//   	// the properties below are optional
//   	OutputFormat: jsii.String("outputFormat"),
//   }
//
// Experimental.
type CloudwatchDeliveryDestinationProps struct {
	// Log group to deliver logs to.
	// Experimental.
	LogGroup interfacesawslogs.ILogGroupRef `field:"required" json:"logGroup" yaml:"logGroup"`
	// Format of the logs that are sent to this delivery destination.
	// Experimental.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

