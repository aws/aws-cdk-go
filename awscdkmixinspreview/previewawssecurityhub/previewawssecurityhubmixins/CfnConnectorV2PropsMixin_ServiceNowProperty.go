package previewawssecurityhubmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceNowProperty := &ServiceNowProperty{
//   	AuthStatus: jsii.String("authStatus"),
//   	InstanceName: jsii.String("instanceName"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html
//
type CfnConnectorV2PropsMixin_ServiceNowProperty struct {
	// The auth status of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html#cfn-securityhub-connectorv2-servicenow-authstatus
	//
	AuthStatus *string `field:"optional" json:"authStatus" yaml:"authStatus"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html#cfn-securityhub-connectorv2-servicenow-instancename
	//
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// The ARN of secrets manager containing ClientId and ClientSecret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html#cfn-securityhub-connectorv2-servicenow-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

