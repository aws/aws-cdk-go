package previewawsecsmixins


// Configuration for Service Connect access logging.
//
// Access logs provide detailed information about requests made to your service, including request patterns, response codes, and timing data for debugging and monitoring purposes.
//
// > To enable access logs, you must also specify a `logConfiguration` in the `serviceConnectConfiguration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceConnectAccessLogConfigurationProperty := &ServiceConnectAccessLogConfigurationProperty{
//   	Format: jsii.String("format"),
//   	IncludeQueryParameters: jsii.String("includeQueryParameters"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectaccesslogconfiguration.html
//
type CfnServicePropsMixin_ServiceConnectAccessLogConfigurationProperty struct {
	// The format for Service Connect access log output.
	//
	// Choose TEXT for human-readable logs or JSON for structured data that integrates well with log analysis tools.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectaccesslogconfiguration.html#cfn-ecs-service-serviceconnectaccesslogconfiguration-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Specifies whether to include query parameters in Service Connect access logs.
	//
	// When enabled, query parameters from HTTP requests are included in the access logs. Consider security and privacy implications when enabling this feature, as query parameters may contain sensitive information such as request IDs and tokens. By default, this parameter is `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectaccesslogconfiguration.html#cfn-ecs-service-serviceconnectaccesslogconfiguration-includequeryparameters
	//
	IncludeQueryParameters *string `field:"optional" json:"includeQueryParameters" yaml:"includeQueryParameters"`
}

