package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectAccessLogConfigurationProperty := &ServiceConnectAccessLogConfigurationProperty{
//   	Format: jsii.String("format"),
//
//   	// the properties below are optional
//   	IncludeQueryParameters: jsii.String("includeQueryParameters"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectaccesslogconfiguration.html
//
type CfnService_ServiceConnectAccessLogConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectaccesslogconfiguration.html#cfn-ecs-service-serviceconnectaccesslogconfiguration-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-serviceconnectaccesslogconfiguration.html#cfn-ecs-service-serviceconnectaccesslogconfiguration-includequeryparameters
	//
	IncludeQueryParameters *string `field:"optional" json:"includeQueryParameters" yaml:"includeQueryParameters"`
}

