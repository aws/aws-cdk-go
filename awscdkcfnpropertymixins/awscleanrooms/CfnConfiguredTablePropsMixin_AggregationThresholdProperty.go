package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   aggregationThresholdProperty := &AggregationThresholdProperty{
//   	AllowedAggregateExpressionType: jsii.String("allowedAggregateExpressionType"),
//   	IdentityColumns: []*string{
//   		jsii.String("identityColumns"),
//   	},
//   	MinimumIdentityCount: jsii.Number(123),
//   	OutputColumnThresholds: []interface{}{
//   		&OutputColumnThresholdProperty{
//   			MinimumIdentityCount: jsii.Number(123),
//   			OutputColumnName: jsii.String("outputColumnName"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregationthreshold.html
//
type CfnConfiguredTablePropsMixin_AggregationThresholdProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregationthreshold.html#cfn-cleanrooms-configuredtable-aggregationthreshold-allowedaggregateexpressiontype
	//
	AllowedAggregateExpressionType *string `field:"optional" json:"allowedAggregateExpressionType" yaml:"allowedAggregateExpressionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregationthreshold.html#cfn-cleanrooms-configuredtable-aggregationthreshold-identitycolumns
	//
	IdentityColumns *[]*string `field:"optional" json:"identityColumns" yaml:"identityColumns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregationthreshold.html#cfn-cleanrooms-configuredtable-aggregationthreshold-minimumidentitycount
	//
	MinimumIdentityCount *float64 `field:"optional" json:"minimumIdentityCount" yaml:"minimumIdentityCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregationthreshold.html#cfn-cleanrooms-configuredtable-aggregationthreshold-outputcolumnthresholds
	//
	OutputColumnThresholds interface{} `field:"optional" json:"outputColumnThresholds" yaml:"outputColumnThresholds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregationthreshold.html#cfn-cleanrooms-configuredtable-aggregationthreshold-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

