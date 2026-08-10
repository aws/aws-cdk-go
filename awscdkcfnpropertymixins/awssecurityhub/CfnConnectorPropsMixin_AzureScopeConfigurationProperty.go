package awssecurityhub


// The scope configuration for an Azure connector, defining the tenant or subscription scope.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   azureScopeConfigurationProperty := &AzureScopeConfigurationProperty{
//   	ScopeType: jsii.String("scopeType"),
//   	ScopeValues: []*string{
//   		jsii.String("scopeValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-azurescopeconfiguration.html
//
type CfnConnectorPropsMixin_AzureScopeConfigurationProperty struct {
	// The type of scope.
	//
	// Valid values are ``tenant`` and ``subscription``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-azurescopeconfiguration.html#cfn-securityhub-connector-azurescopeconfiguration-scopetype
	//
	ScopeType *string `field:"optional" json:"scopeType" yaml:"scopeType"`
	// The list of scope values, such as subscription IDs, when the scope type is ``subscription``.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-connector-azurescopeconfiguration.html#cfn-securityhub-connector-azurescopeconfiguration-scopevalues
	//
	ScopeValues *[]*string `field:"optional" json:"scopeValues" yaml:"scopeValues"`
}

