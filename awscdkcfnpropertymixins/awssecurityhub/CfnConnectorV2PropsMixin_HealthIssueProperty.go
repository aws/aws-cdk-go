package awssecurityhub


// A health issue associated with the connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   healthIssueProperty := &HealthIssueProperty{
//   	Code: jsii.String("code"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-healthissue.html
//
type CfnConnectorV2PropsMixin_HealthIssueProperty struct {
	// The code identifying the type of health issue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-healthissue.html#cfn-securityhub-connectorv2-healthissue-code
	//
	Code *string `field:"optional" json:"code" yaml:"code"`
	// The message describing the health issue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-healthissue.html#cfn-securityhub-connectorv2-healthissue-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

