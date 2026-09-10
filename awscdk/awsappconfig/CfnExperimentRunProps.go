package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnExperimentRun`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExperimentRunProps := &CfnExperimentRunProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	ExperimentDefinitionIdentifier: jsii.String("experimentDefinitionIdentifier"),
//   	ExposurePercentage: jsii.Number(123),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TreatmentOverrides: &TreatmentOverridesProperty{
//   		Inline: map[string]*string{
//   			"inlineKey": jsii.String("inline"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-experimentrun.html
//
type CfnExperimentRunProps struct {
	// The application name or ID used to create the experiment run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-experimentrun.html#cfn-appconfig-experimentrun-applicationidentifier
	//
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The experiment definition name or ID used to create the experiment run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-experimentrun.html#cfn-appconfig-experimentrun-experimentdefinitionidentifier
	//
	ExperimentDefinitionIdentifier *string `field:"required" json:"experimentDefinitionIdentifier" yaml:"experimentDefinitionIdentifier"`
	// Percentage of traffic exposed to the experiment (0-100).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-experimentrun.html#cfn-appconfig-experimentrun-exposurepercentage
	//
	ExposurePercentage *float64 `field:"required" json:"exposurePercentage" yaml:"exposurePercentage"`
	// Description of the experiment run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-experimentrun.html#cfn-appconfig-experimentrun-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags to associate with the experiment run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-experimentrun.html#cfn-appconfig-experimentrun-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Treatment overrides for specific entities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-experimentrun.html#cfn-appconfig-experimentrun-treatmentoverrides
	//
	TreatmentOverrides interface{} `field:"optional" json:"treatmentOverrides" yaml:"treatmentOverrides"`
}

