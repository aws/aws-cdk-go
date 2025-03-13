package awsappintegrations


// The event integration filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventFilterProperty := &EventFilterProperty{
//   	Source: jsii.String("source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-eventintegration-eventfilter.html
//
type CfnEventIntegration_EventFilterProperty struct {
	// The source of the events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-eventintegration-eventfilter.html#cfn-appintegrations-eventintegration-eventfilter-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
}

