package interfacesawsmedialive


// A reference to a CloudWatchAlarmTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchAlarmTemplateReference := &CloudWatchAlarmTemplateReference{
//   	CloudWatchAlarmTemplateArn: jsii.String("cloudWatchAlarmTemplateArn"),
//   	Identifier: jsii.String("identifier"),
//   }
//
type CloudWatchAlarmTemplateReference struct {
	// The ARN of the CloudWatchAlarmTemplate resource.
	CloudWatchAlarmTemplateArn *string `field:"required" json:"cloudWatchAlarmTemplateArn" yaml:"cloudWatchAlarmTemplateArn"`
	// The Identifier of the CloudWatchAlarmTemplate resource.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
}

