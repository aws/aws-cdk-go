package previewawsdatabrewmixins


// Represents a set of options that define how DataBrew selects files for a given Amazon S3 path in a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pathOptionsProperty := &PathOptionsProperty{
//   	FilesLimit: &FilesLimitProperty{
//   		MaxFiles: jsii.Number(123),
//   		Order: jsii.String("order"),
//   		OrderedBy: jsii.String("orderedBy"),
//   	},
//   	LastModifiedDateCondition: &FilterExpressionProperty{
//   		Expression: jsii.String("expression"),
//   		ValuesMap: []interface{}{
//   			&FilterValueProperty{
//   				Value: jsii.String("value"),
//   				ValueReference: jsii.String("valueReference"),
//   			},
//   		},
//   	},
//   	Parameters: []interface{}{
//   		&PathParameterProperty{
//   			DatasetParameter: &DatasetParameterProperty{
//   				CreateColumn: jsii.Boolean(false),
//   				DatetimeOptions: &DatetimeOptionsProperty{
//   					Format: jsii.String("format"),
//   					LocaleCode: jsii.String("localeCode"),
//   					TimezoneOffset: jsii.String("timezoneOffset"),
//   				},
//   				Filter: &FilterExpressionProperty{
//   					Expression: jsii.String("expression"),
//   					ValuesMap: []interface{}{
//   						&FilterValueProperty{
//   							Value: jsii.String("value"),
//   							ValueReference: jsii.String("valueReference"),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   			},
//   			PathParameterName: jsii.String("pathParameterName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-pathoptions.html
//
type CfnDatasetPropsMixin_PathOptionsProperty struct {
	// If provided, this structure imposes a limit on a number of files that should be selected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-pathoptions.html#cfn-databrew-dataset-pathoptions-fileslimit
	//
	FilesLimit interface{} `field:"optional" json:"filesLimit" yaml:"filesLimit"`
	// If provided, this structure defines a date range for matching Amazon S3 objects based on their LastModifiedDate attribute in Amazon S3 .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-pathoptions.html#cfn-databrew-dataset-pathoptions-lastmodifieddatecondition
	//
	LastModifiedDateCondition interface{} `field:"optional" json:"lastModifiedDateCondition" yaml:"lastModifiedDateCondition"`
	// A structure that maps names of parameters used in the Amazon S3 path of a dataset to their definitions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-pathoptions.html#cfn-databrew-dataset-pathoptions-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

