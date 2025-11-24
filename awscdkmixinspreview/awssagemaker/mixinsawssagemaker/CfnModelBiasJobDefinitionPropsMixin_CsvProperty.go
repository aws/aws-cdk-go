package mixinsawssagemaker


// The CSV format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   csvProperty := &CsvProperty{
//   	Header: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-csv.html
//
type CfnModelBiasJobDefinitionPropsMixin_CsvProperty struct {
	// A boolean flag indicating if given CSV has header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelbiasjobdefinition-csv.html#cfn-sagemaker-modelbiasjobdefinition-csv-header
	//
	Header interface{} `field:"optional" json:"header" yaml:"header"`
}

