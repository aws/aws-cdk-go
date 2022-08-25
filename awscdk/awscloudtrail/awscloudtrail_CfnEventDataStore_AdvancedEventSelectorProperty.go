package awscloudtrail


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedEventSelectorProperty := &advancedEventSelectorProperty{
//   	fieldSelectors: []interface{}{
//   		&advancedFieldSelectorProperty{
//   			field: jsii.String("field"),
//
//   			// the properties below are optional
//   			endsWith: []*string{
//   				jsii.String("endsWith"),
//   			},
//   			equalTo: []*string{
//   				jsii.String("equalTo"),
//   			},
//   			notEndsWith: []*string{
//   				jsii.String("notEndsWith"),
//   			},
//   			notEquals: []*string{
//   				jsii.String("notEquals"),
//   			},
//   			notStartsWith: []*string{
//   				jsii.String("notStartsWith"),
//   			},
//   			startsWith: []*string{
//   				jsii.String("startsWith"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnEventDataStore_AdvancedEventSelectorProperty struct {
	// `CfnEventDataStore.AdvancedEventSelectorProperty.FieldSelectors`.
	FieldSelectors interface{} `field:"required" json:"fieldSelectors" yaml:"fieldSelectors"`
	// `CfnEventDataStore.AdvancedEventSelectorProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

