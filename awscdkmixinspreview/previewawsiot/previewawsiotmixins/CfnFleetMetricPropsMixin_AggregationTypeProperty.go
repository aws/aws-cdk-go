package previewawsiotmixins


// The type of aggregation queries.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aggregationTypeProperty := &AggregationTypeProperty{
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-fleetmetric-aggregationtype.html
//
type CfnFleetMetricPropsMixin_AggregationTypeProperty struct {
	// The name of the aggregation type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-fleetmetric-aggregationtype.html#cfn-iot-fleetmetric-aggregationtype-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the values of aggregation types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-fleetmetric-aggregationtype.html#cfn-iot-fleetmetric-aggregationtype-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

