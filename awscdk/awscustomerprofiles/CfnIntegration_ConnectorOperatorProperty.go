package awscustomerprofiles


// The operation to be performed on the provided source fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorOperatorProperty := &ConnectorOperatorProperty{
//   	Marketo: jsii.String("marketo"),
//   	S3: jsii.String("s3"),
//   	Salesforce: jsii.String("salesforce"),
//   	ServiceNow: jsii.String("serviceNow"),
//   	Zendesk: jsii.String("zendesk"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-connectoroperator.html
//
type CfnIntegration_ConnectorOperatorProperty struct {
	// The operation to be performed on the provided Marketo source fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-connectoroperator.html#cfn-customerprofiles-integration-connectoroperator-marketo
	//
	Marketo *string `field:"optional" json:"marketo" yaml:"marketo"`
	// The operation to be performed on the provided Amazon S3 source fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-connectoroperator.html#cfn-customerprofiles-integration-connectoroperator-s3
	//
	S3 *string `field:"optional" json:"s3" yaml:"s3"`
	// The operation to be performed on the provided Salesforce source fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-connectoroperator.html#cfn-customerprofiles-integration-connectoroperator-salesforce
	//
	Salesforce *string `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The operation to be performed on the provided ServiceNow source fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-connectoroperator.html#cfn-customerprofiles-integration-connectoroperator-servicenow
	//
	ServiceNow *string `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// The operation to be performed on the provided Zendesk source fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-connectoroperator.html#cfn-customerprofiles-integration-connectoroperator-zendesk
	//
	Zendesk *string `field:"optional" json:"zendesk" yaml:"zendesk"`
}

