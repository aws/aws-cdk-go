package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   metricDefinitionProperty := &MetricDefinitionProperty{
//   	Name: jsii.String("name"),
//   	Regex: jsii.String("regex"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-metricdefinition.html
//
type CfnAlgorithmPropsMixin_MetricDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-metricdefinition.html#cfn-sagemaker-algorithm-metricdefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-metricdefinition.html#cfn-sagemaker-algorithm-metricdefinition-regex
	//
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

