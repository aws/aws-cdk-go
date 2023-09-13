package awsconfig


// Specifies whether the configuration recorder excludes certain resource types from being recorded.
//
// Use the `resourceTypes` field to enter a comma-separated list of resource types you want to exclude from recording.
//
// By default, when AWS Config adds support for a new resource type in the Region where you set up the configuration recorder, including global resource types, AWS Config starts recording resources of that type automatically.
//
// > *How to use*
// >
// > To use this option, you must set the `useOnly` field of [AWS::Config::ConfigurationRecorder RecordingStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordingstrategy.html) to `EXCLUSION_BY_RESOURCE_TYPES` .
// >
// > AWS Config will then record configuration changes for all supported resource types, except the resource types that you specify to exclude from being recorded.
// >
// > *Globally recorded resources*
// >
// > Unless specifically listed as exclusions, `AWS::RDS::GlobalCluster` will be recorded automatically in all supported AWS Config Regions were the configuration recorder is enabled. IAM users, groups, roles, and customer managed policies will be recorded automatically in all enabled AWS Config Regions where AWS Config was available before February 2022. This list does not include the following Regions:
// >
// > - Asia Pacific (Hyderabad)
// > - Asia Pacific (Melbourne)
// > - Europe (Spain)
// > - Europe (Zurich)
// > - Israel (Tel Aviv)
// > - Middle East (UAE).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exclusionByResourceTypesProperty := &ExclusionByResourceTypesProperty{
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-exclusionbyresourcetypes.html
//
type CfnConfigurationRecorder_ExclusionByResourceTypesProperty struct {
	// A comma-separated list of resource types to exclude from recording by the configuration recorder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-exclusionbyresourcetypes.html#cfn-config-configurationrecorder-exclusionbyresourcetypes-resourcetypes
	//
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
}

