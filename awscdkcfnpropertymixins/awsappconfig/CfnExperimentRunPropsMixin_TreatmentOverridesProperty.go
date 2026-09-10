package awsappconfig


// Treatment overrides for specific entities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   treatmentOverridesProperty := &TreatmentOverridesProperty{
//   	Inline: map[string]*string{
//   		"inlineKey": jsii.String("inline"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-experimentrun-treatmentoverrides.html
//
type CfnExperimentRunPropsMixin_TreatmentOverridesProperty struct {
	// Map of entity ID to treatment key (t1, t2, ..., or c for control).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-experimentrun-treatmentoverrides.html#cfn-appconfig-experimentrun-treatmentoverrides-inline
	//
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
}

