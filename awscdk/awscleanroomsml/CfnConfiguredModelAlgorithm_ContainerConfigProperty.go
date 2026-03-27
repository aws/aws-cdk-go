package awscleanroomsml


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerConfigProperty := &ContainerConfigProperty{
//   	ImageUri: jsii.String("imageUri"),
//
//   	// the properties below are optional
//   	Arguments: []*string{
//   		jsii.String("arguments"),
//   	},
//   	Entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	MetricDefinitions: []interface{}{
//   		&MetricDefinitionProperty{
//   			Name: jsii.String("name"),
//   			Regex: jsii.String("regex"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithm-containerconfig.html
//
type CfnConfiguredModelAlgorithm_ContainerConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithm-containerconfig.html#cfn-cleanroomsml-configuredmodelalgorithm-containerconfig-imageuri
	//
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithm-containerconfig.html#cfn-cleanroomsml-configuredmodelalgorithm-containerconfig-arguments
	//
	Arguments *[]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithm-containerconfig.html#cfn-cleanroomsml-configuredmodelalgorithm-containerconfig-entrypoint
	//
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanroomsml-configuredmodelalgorithm-containerconfig.html#cfn-cleanroomsml-configuredmodelalgorithm-containerconfig-metricdefinitions
	//
	MetricDefinitions interface{} `field:"optional" json:"metricDefinitions" yaml:"metricDefinitions"`
}

