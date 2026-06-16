package awsobservabilityadmin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sourceMetricsConfigurationProperty := &SourceMetricsConfigurationProperty{
//   	MetricsSelectionCriteria: jsii.String("metricsSelectionCriteria"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-sourcemetricsconfiguration.html
//
type CfnOrganizationCentralizationRulePropsMixin_SourceMetricsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-sourcemetricsconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-sourcemetricsconfiguration-metricsselectioncriteria
	//
	MetricsSelectionCriteria *string `field:"optional" json:"metricsSelectionCriteria" yaml:"metricsSelectionCriteria"`
}

