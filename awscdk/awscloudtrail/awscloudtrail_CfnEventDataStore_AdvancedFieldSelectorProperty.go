package awscloudtrail


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedFieldSelectorProperty := &advancedFieldSelectorProperty{
//   	field: jsii.String("field"),
//
//   	// the properties below are optional
//   	endsWith: []*string{
//   		jsii.String("endsWith"),
//   	},
//   	equalTo: []*string{
//   		jsii.String("equalTo"),
//   	},
//   	notEndsWith: []*string{
//   		jsii.String("notEndsWith"),
//   	},
//   	notEquals: []*string{
//   		jsii.String("notEquals"),
//   	},
//   	notStartsWith: []*string{
//   		jsii.String("notStartsWith"),
//   	},
//   	startsWith: []*string{
//   		jsii.String("startsWith"),
//   	},
//   }
//
type CfnEventDataStore_AdvancedFieldSelectorProperty struct {
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.Field`.
	Field *string `field:"required" json:"field" yaml:"field"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.EndsWith`.
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.Equals`.
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.NotEndsWith`.
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.NotEquals`.
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.NotStartsWith`.
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// `CfnEventDataStore.AdvancedFieldSelectorProperty.StartsWith`.
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

