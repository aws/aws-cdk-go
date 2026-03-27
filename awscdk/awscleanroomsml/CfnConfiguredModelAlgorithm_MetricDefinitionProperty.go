package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDefinitionProperty := &MetricDefinitionProperty{
//   	Name: jsii.String("name"),
//   	Regex: jsii.String("regex"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithm-metricdefinition.html
//
type CfnConfiguredModelAlgorithm_MetricDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithm-metricdefinition.html#cfn-cleanroomsml-configuredmodelalgorithm-metricdefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithm-metricdefinition.html#cfn-cleanroomsml-configuredmodelalgorithm-metricdefinition-regex
	//
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

