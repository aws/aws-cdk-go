package mixinsawslightsail


// `RelationalDatabaseParameter` is a property of the [AWS::Lightsail::Database](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html) resource. It describes parameters for the database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   relationalDatabaseParameterProperty := &RelationalDatabaseParameterProperty{
//   	AllowedValues: jsii.String("allowedValues"),
//   	ApplyMethod: jsii.String("applyMethod"),
//   	ApplyType: jsii.String("applyType"),
//   	DataType: jsii.String("dataType"),
//   	Description: jsii.String("description"),
//   	IsModifiable: jsii.Boolean(false),
//   	ParameterName: jsii.String("parameterName"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-database-relationaldatabaseparameter.html
//
type CfnDatabasePropsMixin_RelationalDatabaseParameterProperty struct {
	// The valid range of values for the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-database-relationaldatabaseparameter.html#cfn-lightsail-database-relationaldatabaseparameter-allowedvalues
	//
	AllowedValues *string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// Indicates when parameter updates are applied.
	//
	// Can be `immediate` or `pending-reboot` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-database-relationaldatabaseparameter.html#cfn-lightsail-database-relationaldatabaseparameter-applymethod
	//
	ApplyMethod *string `field:"optional" json:"applyMethod" yaml:"applyMethod"`
	// Specifies the engine-specific parameter type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-database-relationaldatabaseparameter.html#cfn-lightsail-database-relationaldatabaseparameter-applytype
	//
	ApplyType *string `field:"optional" json:"applyType" yaml:"applyType"`
	// The valid data type of the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-database-relationaldatabaseparameter.html#cfn-lightsail-database-relationaldatabaseparameter-datatype
	//
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// A description of the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-database-relationaldatabaseparameter.html#cfn-lightsail-database-relationaldatabaseparameter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A Boolean value indicating whether the parameter can be modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-database-relationaldatabaseparameter.html#cfn-lightsail-database-relationaldatabaseparameter-ismodifiable
	//
	IsModifiable interface{} `field:"optional" json:"isModifiable" yaml:"isModifiable"`
	// The name of the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-database-relationaldatabaseparameter.html#cfn-lightsail-database-relationaldatabaseparameter-parametername
	//
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The value for the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-database-relationaldatabaseparameter.html#cfn-lightsail-database-relationaldatabaseparameter-parametervalue
	//
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

