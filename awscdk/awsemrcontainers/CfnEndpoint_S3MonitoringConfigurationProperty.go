package awsemrcontainers


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3MonitoringConfigurationProperty := &S3MonitoringConfigurationProperty{
//   	LogUri: jsii.String("logUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-s3monitoringconfiguration.html
//
type CfnEndpoint_S3MonitoringConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-s3monitoringconfiguration.html#cfn-emrcontainers-endpoint-s3monitoringconfiguration-loguri
	//
	LogUri *string `field:"required" json:"logUri" yaml:"logUri"`
}

