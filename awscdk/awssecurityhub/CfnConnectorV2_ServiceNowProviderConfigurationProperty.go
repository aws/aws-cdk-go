package awssecurityhub


// The initial configuration settings required to establish an integration between Security Hub CSPM and ServiceNow ITSM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowProviderConfigurationProperty := &ServiceNowProviderConfigurationProperty{
//   	InstanceName: jsii.String("instanceName"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenowproviderconfiguration.html
//
type CfnConnectorV2_ServiceNowProviderConfigurationProperty struct {
	// The instance name of ServiceNow ITSM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenowproviderconfiguration.html#cfn-securityhub-connectorv2-servicenowproviderconfiguration-instancename
	//
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the ServiceNow credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenowproviderconfiguration.html#cfn-securityhub-connectorv2-servicenowproviderconfiguration-secretarn
	//
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

