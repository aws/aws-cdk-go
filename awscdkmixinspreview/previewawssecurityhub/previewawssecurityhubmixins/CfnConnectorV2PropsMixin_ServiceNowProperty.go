package previewawssecurityhubmixins


// Information about a ServiceNow ITSM integration.
//
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
	// The status of the authorization between ServiceNow and the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html#cfn-securityhub-connectorv2-servicenow-authstatus
	//
	AuthStatus *string `field:"optional" json:"authStatus" yaml:"authStatus"`
	// The instanceName of ServiceNow ITSM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html#cfn-securityhub-connectorv2-servicenow-instancename
	//
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the ServiceNow credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html#cfn-securityhub-connectorv2-servicenow-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

