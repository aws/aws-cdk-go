package awsapplicationsignals


// Properties for defining a `CfnGroupingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupingConfigurationProps := &CfnGroupingConfigurationProps{
//   	GroupingAttributeDefinitions: []interface{}{
//   		&GroupingAttributeDefinitionProperty{
//   			GroupingName: jsii.String("groupingName"),
//   			GroupingSourceKeys: []*string{
//   				jsii.String("groupingSourceKeys"),
//   			},
//
//   			// the properties below are optional
//   			DefaultGroupingValue: jsii.String("defaultGroupingValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-groupingconfiguration.html
//
type CfnGroupingConfigurationProps struct {
	// An array of grouping attribute definitions that specify how services should be grouped based on various attributes and source keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-groupingconfiguration.html#cfn-applicationsignals-groupingconfiguration-groupingattributedefinitions
	//
	GroupingAttributeDefinitions interface{} `field:"required" json:"groupingAttributeDefinitions" yaml:"groupingAttributeDefinitions"`
}

