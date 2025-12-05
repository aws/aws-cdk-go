package previewawsconnectmixins


// Properties for CfnDataTableAttributePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataTableAttributeMixinProps := &CfnDataTableAttributeMixinProps{
//   	DataTableArn: jsii.String("dataTableArn"),
//   	Description: jsii.String("description"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	Primary: jsii.Boolean(false),
//   	Validation: &ValidationProperty{
//   		Enum: &EnumProperty{
//   			Strict: jsii.Boolean(false),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		ExclusiveMaximum: jsii.Number(123),
//   		ExclusiveMinimum: jsii.Number(123),
//   		Maximum: jsii.Number(123),
//   		MaxLength: jsii.Number(123),
//   		MaxValues: jsii.Number(123),
//   		Minimum: jsii.Number(123),
//   		MinLength: jsii.Number(123),
//   		MinValues: jsii.Number(123),
//   		MultipleOf: jsii.Number(123),
//   	},
//   	ValueType: jsii.String("valueType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatableattribute.html
//
type CfnDataTableAttributeMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatableattribute.html#cfn-connect-datatableattribute-datatablearn
	//
	DataTableArn *string `field:"optional" json:"dataTableArn" yaml:"dataTableArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatableattribute.html#cfn-connect-datatableattribute-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatableattribute.html#cfn-connect-datatableattribute-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatableattribute.html#cfn-connect-datatableattribute-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatableattribute.html#cfn-connect-datatableattribute-primary
	//
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatableattribute.html#cfn-connect-datatableattribute-validation
	//
	Validation interface{} `field:"optional" json:"validation" yaml:"validation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datatableattribute.html#cfn-connect-datatableattribute-valuetype
	//
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

