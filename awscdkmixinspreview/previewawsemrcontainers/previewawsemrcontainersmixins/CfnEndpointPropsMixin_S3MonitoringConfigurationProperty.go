package previewawsemrcontainersmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3MonitoringConfigurationProperty := &S3MonitoringConfigurationProperty{
//   	LogUri: jsii.String("logUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-s3monitoringconfiguration.html
//
type CfnEndpointPropsMixin_S3MonitoringConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-s3monitoringconfiguration.html#cfn-emrcontainers-endpoint-s3monitoringconfiguration-loguri
	//
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
}

