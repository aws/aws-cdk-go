package awsfis


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsConfigurationProperty := &cloudWatchLogsConfigurationProperty{
//   	logGroupArn: jsii.String("logGroupArn"),
//   }
//
type CfnExperimentTemplate_CloudWatchLogsConfigurationProperty struct {
	// `CfnExperimentTemplate.CloudWatchLogsConfigurationProperty.LogGroupArn`.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

