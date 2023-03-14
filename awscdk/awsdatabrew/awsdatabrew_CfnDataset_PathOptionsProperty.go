package awsdatabrew


// Represents a set of options that define how DataBrew selects files for a given Amazon S3 path in a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathOptionsProperty := &pathOptionsProperty{
//   	filesLimit: &filesLimitProperty{
//   		maxFiles: jsii.Number(123),
//
//   		// the properties below are optional
//   		order: jsii.String("order"),
//   		orderedBy: jsii.String("orderedBy"),
//   	},
//   	lastModifiedDateCondition: &filterExpressionProperty{
//   		expression: jsii.String("expression"),
//   		valuesMap: []interface{}{
//   			&filterValueProperty{
//   				value: jsii.String("value"),
//   				valueReference: jsii.String("valueReference"),
//   			},
//   		},
//   	},
//   	parameters: []interface{}{
//   		&pathParameterProperty{
//   			datasetParameter: &datasetParameterProperty{
//   				name: jsii.String("name"),
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				createColumn: jsii.Boolean(false),
//   				datetimeOptions: &datetimeOptionsProperty{
//   					format: jsii.String("format"),
//
//   					// the properties below are optional
//   					localeCode: jsii.String("localeCode"),
//   					timezoneOffset: jsii.String("timezoneOffset"),
//   				},
//   				filter: &filterExpressionProperty{
//   					expression: jsii.String("expression"),
//   					valuesMap: []interface{}{
//   						&filterValueProperty{
//   							value: jsii.String("value"),
//   							valueReference: jsii.String("valueReference"),
//   						},
//   					},
//   				},
//   			},
//   			pathParameterName: jsii.String("pathParameterName"),
//   		},
//   	},
//   }
//
type CfnDataset_PathOptionsProperty struct {
	// If provided, this structure imposes a limit on a number of files that should be selected.
	FilesLimit interface{} `field:"optional" json:"filesLimit" yaml:"filesLimit"`
	// If provided, this structure defines a date range for matching Amazon S3 objects based on their LastModifiedDate attribute in Amazon S3 .
	LastModifiedDateCondition interface{} `field:"optional" json:"lastModifiedDateCondition" yaml:"lastModifiedDateCondition"`
	// A structure that maps names of parameters used in the Amazon S3 path of a dataset to their definitions.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

