package awsstepfunctions


// Properties for defining a new Step Functions Activity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activityProps := &activityProps{
//   	activityName: jsii.String("activityName"),
//   }
//
type ActivityProps struct {
	// The name for this activity.
	ActivityName *string `field:"optional" json:"activityName" yaml:"activityName"`
}

