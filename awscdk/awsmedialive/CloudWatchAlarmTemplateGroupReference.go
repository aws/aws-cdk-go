package awsmedialive


// A reference to a CloudWatchAlarmTemplateGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchAlarmTemplateGroupReference := &CloudWatchAlarmTemplateGroupReference{
//   	CloudWatchAlarmTemplateGroupArn: jsii.String("cloudWatchAlarmTemplateGroupArn"),
//   	Identifier: jsii.String("identifier"),
//   }
//
type CloudWatchAlarmTemplateGroupReference struct {
	// The ARN of the CloudWatchAlarmTemplateGroup resource.
	CloudWatchAlarmTemplateGroupArn *string `field:"required" json:"cloudWatchAlarmTemplateGroupArn" yaml:"cloudWatchAlarmTemplateGroupArn"`
	// The Identifier of the CloudWatchAlarmTemplateGroup resource.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

