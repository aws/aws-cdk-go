package awsiotsitewise


// Contains an asset transform property.
//
// A transform is a one-to-one mapping of a property's data points from one form to another. For example, you can use a transform to convert a Celsius data stream to Fahrenheit by applying the transformation expression to each data point of the Celsius stream. Transforms can only input properties that are `INTEGER` , `DOUBLE` , or `BOOLEAN` type. Booleans convert to `0` ( `FALSE` ) and `1` ( `TRUE` )..
//
// For more information, see [Defining data properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html#transforms) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformProperty := &transformProperty{
//   	expression: jsii.String("expression"),
//   	variables: []interface{}{
//   		&expressionVariableProperty{
//   			name: jsii.String("name"),
//   			value: &variableValueProperty{
//   				propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   				// the properties below are optional
//   				hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   			},
//   		},
//   	},
//   }
//
type CfnAssetModel_TransformProperty struct {
	// The mathematical expression that defines the transformation function.
	//
	// You can specify up to 10 variables per expression. You can specify up to 10 functions per expression.
	//
	// For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The list of variables used in the expression.
	Variables interface{} `field:"required" json:"variables" yaml:"variables"`
}

