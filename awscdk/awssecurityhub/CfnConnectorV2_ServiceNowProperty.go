package awssecurityhub


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowProperty := &ServiceNowProperty{
//   	InstanceName: jsii.String("instanceName"),
//   	SecretArn: jsii.String("secretArn"),
//
//   	// the properties below are optional
//   	AuthStatus: jsii.String("authStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html
//
type CfnConnectorV2_ServiceNowProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html#cfn-securityhub-connectorv2-servicenow-instancename
	//
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// The ARN of secrets manager containing ClientId and ClientSecret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html#cfn-securityhub-connectorv2-servicenow-secretarn
	//
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The auth status of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-servicenow.html#cfn-securityhub-connectorv2-servicenow-authstatus
	//
	AuthStatus *string `field:"optional" json:"authStatus" yaml:"authStatus"`
}

