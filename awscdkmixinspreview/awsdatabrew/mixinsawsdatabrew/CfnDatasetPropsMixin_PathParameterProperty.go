package mixinsawsdatabrew


// Represents a single entry in the path parameters of a dataset.
//
// Each `PathParameter` consists of a name and a parameter definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pathParameterProperty := &PathParameterProperty{
//   	DatasetParameter: &DatasetParameterProperty{
//   		CreateColumn: jsii.Boolean(false),
//   		DatetimeOptions: &DatetimeOptionsProperty{
//   			Format: jsii.String("format"),
//   			LocaleCode: jsii.String("localeCode"),
//   			TimezoneOffset: jsii.String("timezoneOffset"),
//   		},
//   		Filter: &FilterExpressionProperty{
//   			Expression: jsii.String("expression"),
//   			ValuesMap: []interface{}{
//   				&FilterValueProperty{
//   					Value: jsii.String("value"),
//   					ValueReference: jsii.String("valueReference"),
//   				},
//   			},
//   		},
//   		Name: jsii.String("name"),
//   		Type: jsii.String("type"),
//   	},
//   	PathParameterName: jsii.String("pathParameterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-pathparameter.html
//
type CfnDatasetPropsMixin_PathParameterProperty struct {
	// The path parameter definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-pathparameter.html#cfn-databrew-dataset-pathparameter-datasetparameter
	//
	DatasetParameter interface{} `field:"optional" json:"datasetParameter" yaml:"datasetParameter"`
	// The name of the path parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-pathparameter.html#cfn-databrew-dataset-pathparameter-pathparametername
	//
	PathParameterName *string `field:"optional" json:"pathParameterName" yaml:"pathParameterName"`
}

