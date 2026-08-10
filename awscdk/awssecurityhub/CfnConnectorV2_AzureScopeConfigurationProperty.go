package awssecurityhub


// The scope configuration for an Azure connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   azureScopeConfigurationProperty := &AzureScopeConfigurationProperty{
//   	ScopeType: jsii.String("scopeType"),
//
//   	// the properties below are optional
//   	ScopeValues: []*string{
//   		jsii.String("scopeValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-azurescopeconfiguration.html
//
type CfnConnectorV2_AzureScopeConfigurationProperty struct {
	// The scope type for the Azure connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-azurescopeconfiguration.html#cfn-securityhub-connectorv2-azurescopeconfiguration-scopetype
	//
	ScopeType *string `field:"required" json:"scopeType" yaml:"scopeType"`
	// The list of scope values for the Azure connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connectorv2-azurescopeconfiguration.html#cfn-securityhub-connectorv2-azurescopeconfiguration-scopevalues
	//
	ScopeValues *[]*string `field:"optional" json:"scopeValues" yaml:"scopeValues"`
}

