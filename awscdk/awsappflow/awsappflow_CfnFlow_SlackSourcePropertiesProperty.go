package awsappflow


// The properties that are applied when Slack is being used as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slackSourcePropertiesProperty := &slackSourcePropertiesProperty{
//   	object: jsii.String("object"),
//   }
//
type CfnFlow_SlackSourcePropertiesProperty struct {
	// The object specified in the Slack flow source.
	Object *string `field:"required" json:"object" yaml:"object"`
}

