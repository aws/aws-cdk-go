package awsworkspacesweb


// A reference to a SessionLogger resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionLoggerReference := &SessionLoggerReference{
//   	SessionLoggerArn: jsii.String("sessionLoggerArn"),
//   }
//
type SessionLoggerReference struct {
	// The SessionLoggerArn of the SessionLogger resource.
	SessionLoggerArn *string `field:"required" json:"sessionLoggerArn" yaml:"sessionLoggerArn"`
}

