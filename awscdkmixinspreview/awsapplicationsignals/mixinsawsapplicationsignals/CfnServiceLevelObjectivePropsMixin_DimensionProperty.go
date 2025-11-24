package mixinsawsapplicationsignals


// A dimension is a name/value pair that is part of the identity of a metric.
//
// Because dimensions are part of the unique identifier for a metric, whenever you add a unique name/value pair to one of your metrics, you are creating a new variation of that metric. For example, many Amazon EC2 metrics publish `InstanceId` as a dimension name, and the actual instance ID as the value for that dimension.
//
// You can assign up to 30 dimensions to a metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dimensionProperty := &DimensionProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-dimension.html
//
type CfnServiceLevelObjectivePropsMixin_DimensionProperty struct {
	// The name of the dimension.
	//
	// Dimension names must contain only ASCII characters, must include at least one non-whitespace character, and cannot start with a colon ( `:` ). ASCII control characters are not supported as part of dimension names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-dimension.html#cfn-applicationsignals-servicelevelobjective-dimension-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the dimension.
	//
	// Dimension values must contain only ASCII characters and must include at least one non-whitespace character. ASCII control characters are not supported as part of dimension values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-dimension.html#cfn-applicationsignals-servicelevelobjective-dimension-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

