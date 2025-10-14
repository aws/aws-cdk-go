package awsobservabilityadmin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   centralizationRuleSourceProperty := &CentralizationRuleSourceProperty{
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//
//   	// the properties below are optional
//   	Scope: jsii.String("scope"),
//   	SourceLogsConfiguration: &SourceLogsConfigurationProperty{
//   		EncryptedLogGroupStrategy: jsii.String("encryptedLogGroupStrategy"),
//   		LogGroupSelectionCriteria: jsii.String("logGroupSelectionCriteria"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrulesource.html
//
type CfnOrganizationCentralizationRule_CentralizationRuleSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrulesource.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationrulesource-regions
	//
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrulesource.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationrulesource-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrulesource.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationrulesource-sourcelogsconfiguration
	//
	SourceLogsConfiguration interface{} `field:"optional" json:"sourceLogsConfiguration" yaml:"sourceLogsConfiguration"`
}

