package awsappflow


// Specifies elements that Amazon AppFlow includes in the file and folder names in the flow destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixConfigProperty := &PrefixConfigProperty{
//   	PathPrefixHierarchy: []*string{
//   		jsii.String("pathPrefixHierarchy"),
//   	},
//   	PrefixFormat: jsii.String("prefixFormat"),
//   	PrefixType: jsii.String("prefixType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-prefixconfig.html
//
type CfnFlow_PrefixConfigProperty struct {
	// Specifies whether the destination file path includes either or both of the following elements:.
	//
	// - **EXECUTION_ID** - The ID that Amazon AppFlow assigns to the flow run.
	// - **SCHEMA_VERSION** - The version number of your data schema. Amazon AppFlow assigns this version number. The version number increases by one when you change any of the following settings in your flow configuration:
	//
	// - Source-to-destination field mappings
	// - Field data types
	// - Partition keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-prefixconfig.html#cfn-appflow-flow-prefixconfig-pathprefixhierarchy
	//
	PathPrefixHierarchy *[]*string `field:"optional" json:"pathPrefixHierarchy" yaml:"pathPrefixHierarchy"`
	// Determines the level of granularity for the date and time that's included in the prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-prefixconfig.html#cfn-appflow-flow-prefixconfig-prefixformat
	//
	PrefixFormat *string `field:"optional" json:"prefixFormat" yaml:"prefixFormat"`
	// Determines the format of the prefix, and whether it applies to the file name, file path, or both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-prefixconfig.html#cfn-appflow-flow-prefixconfig-prefixtype
	//
	PrefixType *string `field:"optional" json:"prefixType" yaml:"prefixType"`
}

