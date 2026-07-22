package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   hyperParameterSpecificationProperty := &HyperParameterSpecificationProperty{
//   	DefaultValue: jsii.String("defaultValue"),
//   	Description: jsii.String("description"),
//   	IsRequired: jsii.Boolean(false),
//   	IsTunable: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	Range: &ParameterRangeProperty{
//   		CategoricalParameterRangeSpecification: &CategoricalParameterRangeSpecificationProperty{
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		ContinuousParameterRangeSpecification: &ContinuousParameterRangeSpecificationProperty{
//   			MaxValue: jsii.String("maxValue"),
//   			MinValue: jsii.String("minValue"),
//   		},
//   		IntegerParameterRangeSpecification: &IntegerParameterRangeSpecificationProperty{
//   			MaxValue: jsii.String("maxValue"),
//   			MinValue: jsii.String("minValue"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparameterspecification.html
//
type CfnAlgorithmPropsMixin_HyperParameterSpecificationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparameterspecification.html#cfn-sagemaker-algorithm-hyperparameterspecification-defaultvalue
	//
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparameterspecification.html#cfn-sagemaker-algorithm-hyperparameterspecification-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparameterspecification.html#cfn-sagemaker-algorithm-hyperparameterspecification-isrequired
	//
	IsRequired interface{} `field:"optional" json:"isRequired" yaml:"isRequired"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparameterspecification.html#cfn-sagemaker-algorithm-hyperparameterspecification-istunable
	//
	IsTunable interface{} `field:"optional" json:"isTunable" yaml:"isTunable"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparameterspecification.html#cfn-sagemaker-algorithm-hyperparameterspecification-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparameterspecification.html#cfn-sagemaker-algorithm-hyperparameterspecification-range
	//
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-hyperparameterspecification.html#cfn-sagemaker-algorithm-hyperparameterspecification-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

