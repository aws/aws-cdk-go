package awscustomerprofiles


// The operation to be performed on the provided source fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorOperatorProperty := &connectorOperatorProperty{
//   	marketo: jsii.String("marketo"),
//   	s3: jsii.String("s3"),
//   	salesforce: jsii.String("salesforce"),
//   	serviceNow: jsii.String("serviceNow"),
//   	zendesk: jsii.String("zendesk"),
//   }
//
type CfnIntegration_ConnectorOperatorProperty struct {
	// The operation to be performed on the provided Marketo source fields.
	Marketo *string `field:"optional" json:"marketo" yaml:"marketo"`
	// The operation to be performed on the provided Amazon S3 source fields.
	S3 *string `field:"optional" json:"s3" yaml:"s3"`
	// The operation to be performed on the provided Salesforce source fields.
	Salesforce *string `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The operation to be performed on the provided ServiceNow source fields.
	ServiceNow *string `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// The operation to be performed on the provided Zendesk source fields.
	Zendesk *string `field:"optional" json:"zendesk" yaml:"zendesk"`
}

