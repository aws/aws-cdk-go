package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsConfigurationPolicyProperty := &MetricsConfigurationPolicyProperty{
//   	NoiseLevel: jsii.String("noiseLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-metricsconfigurationpolicy.html
//
type CfnConfiguredModelAlgorithmAssociation_MetricsConfigurationPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-metricsconfigurationpolicy.html#cfn-cleanroomsml-configuredmodelalgorithmassociation-metricsconfigurationpolicy-noiselevel
	//
	NoiseLevel *string `field:"required" json:"noiseLevel" yaml:"noiseLevel"`
}

