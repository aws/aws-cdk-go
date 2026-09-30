package awsemrcontainers


// Amazon S3 configuration for monitoring log publishing.
//
// This data type allows job template parameters to be specified within.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parametricS3MonitoringConfigurationProperty := &ParametricS3MonitoringConfigurationProperty{
//   	LogUri: jsii.String("logUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametrics3monitoringconfiguration.html
//
type CfnJobTemplate_ParametricS3MonitoringConfigurationProperty struct {
	// Amazon S3 destination URI for log publishing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-parametrics3monitoringconfiguration.html#cfn-emrcontainers-jobtemplate-parametrics3monitoringconfiguration-loguri
	//
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
}

