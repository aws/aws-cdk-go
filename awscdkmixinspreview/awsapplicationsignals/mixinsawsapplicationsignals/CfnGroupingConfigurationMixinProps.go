package mixinsawsapplicationsignals


// Properties for CfnGroupingConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGroupingConfigurationMixinProps := &CfnGroupingConfigurationMixinProps{
//   	GroupingAttributeDefinitions: []interface{}{
//   		&GroupingAttributeDefinitionProperty{
//   			DefaultGroupingValue: jsii.String("defaultGroupingValue"),
//   			GroupingName: jsii.String("groupingName"),
//   			GroupingSourceKeys: []*string{
//   				jsii.String("groupingSourceKeys"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-groupingconfiguration.html
//
type CfnGroupingConfigurationMixinProps struct {
	// An array of grouping attribute definitions that specify how services should be grouped based on various attributes and source keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-groupingconfiguration.html#cfn-applicationsignals-groupingconfiguration-groupingattributedefinitions
	//
	GroupingAttributeDefinitions interface{} `field:"optional" json:"groupingAttributeDefinitions" yaml:"groupingAttributeDefinitions"`
}

