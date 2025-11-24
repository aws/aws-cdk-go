package mixinsawsobservabilityadmin


// Configuration specifying the source of telemetry data to be centralized.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   centralizationRuleSourceProperty := &CentralizationRuleSourceProperty{
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	Scope: jsii.String("scope"),
//   	SourceLogsConfiguration: &SourceLogsConfigurationProperty{
//   		EncryptedLogGroupStrategy: jsii.String("encryptedLogGroupStrategy"),
//   		LogGroupSelectionCriteria: jsii.String("logGroupSelectionCriteria"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrulesource.html
//
type CfnOrganizationCentralizationRulePropsMixin_CentralizationRuleSourceProperty struct {
	// The list of source regions from which telemetry data should be centralized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrulesource.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationrulesource-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// The organizational scope from which telemetry data should be centralized, specified using organization id, accounts or organizational unit ids.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrulesource.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationrulesource-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Log specific configuration for centralization source log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrulesource.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationrulesource-sourcelogsconfiguration
	//
	SourceLogsConfiguration interface{} `field:"optional" json:"sourceLogsConfiguration" yaml:"sourceLogsConfiguration"`
}

