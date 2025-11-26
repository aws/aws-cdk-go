package previewawsfsxmixins


// Use to specify configuration options for a volume’s storage aggregate or aggregates.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aggregateConfigurationProperty := &AggregateConfigurationProperty{
//   	Aggregates: []*string{
//   		jsii.String("aggregates"),
//   	},
//   	ConstituentsPerAggregate: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-aggregateconfiguration.html
//
type CfnVolumePropsMixin_AggregateConfigurationProperty struct {
	// The list of aggregates that this volume resides on.
	//
	// Aggregates are storage pools which make up your primary storage tier. Each high-availability (HA) pair has one aggregate. The names of the aggregates map to the names of the aggregates in the ONTAP CLI and REST API. For FlexVols, there will always be a single entry.
	//
	// Amazon FSx responds with an HTTP status code 400 (Bad Request) for the following conditions:
	//
	// - The strings in the value of `Aggregates` are not are not formatted as `aggrX` , where X is a number between 1 and 12.
	// - The value of `Aggregates` contains aggregates that are not present.
	// - One or more of the aggregates supplied are too close to the volume limit to support adding more volumes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-aggregateconfiguration.html#cfn-fsx-volume-aggregateconfiguration-aggregates
	//
	Aggregates *[]*string `field:"optional" json:"aggregates" yaml:"aggregates"`
	// Used to explicitly set the number of constituents within the FlexGroup per storage aggregate.
	//
	// This field is optional when creating a FlexGroup volume. If unspecified, the default value will be 8. This field cannot be provided when creating a FlexVol volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-volume-aggregateconfiguration.html#cfn-fsx-volume-aggregateconfiguration-constituentsperaggregate
	//
	ConstituentsPerAggregate *float64 `field:"optional" json:"constituentsPerAggregate" yaml:"constituentsPerAggregate"`
}

