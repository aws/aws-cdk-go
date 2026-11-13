package awsdevopsagent


// Configuration of the code change risk analysis automated testing capability.
//
// This capability does not support trigger filter groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   automatedTestingCapabilityConfigurationProperty := &AutomatedTestingCapabilityConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-automatedtestingcapabilityconfiguration.html
//
type CfnAssociationPropsMixin_AutomatedTestingCapabilityConfigurationProperty struct {
	// Whether the capability is enabled for this association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-automatedtestingcapabilityconfiguration.html#cfn-devopsagent-association-automatedtestingcapabilityconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

