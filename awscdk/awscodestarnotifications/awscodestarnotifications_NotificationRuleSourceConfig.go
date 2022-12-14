package awscodestarnotifications


// Information about the Codebuild or CodePipeline associated with a notification source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationRuleSourceConfig := &notificationRuleSourceConfig{
//   	sourceArn: jsii.String("sourceArn"),
//   }
//
type NotificationRuleSourceConfig struct {
	// The Amazon Resource Name (ARN) of the notification source.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
}

