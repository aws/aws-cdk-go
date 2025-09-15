package awsstepfunctions


// A reference to a Activity resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activityReference := &ActivityReference{
//   	ActivityArn: jsii.String("activityArn"),
//   }
//
type ActivityReference struct {
	// The Arn of the Activity resource.
	ActivityArn *string `field:"required" json:"activityArn" yaml:"activityArn"`
}

