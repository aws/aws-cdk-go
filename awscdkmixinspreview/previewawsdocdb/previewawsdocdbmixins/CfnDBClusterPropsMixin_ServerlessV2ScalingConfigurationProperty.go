package previewawsdocdbmixins


// Sets the scaling configuration of an Amazon DocumentDB Serverless cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serverlessV2ScalingConfigurationProperty := &ServerlessV2ScalingConfigurationProperty{
//   	MaxCapacity: jsii.Number(123),
//   	MinCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-docdb-dbcluster-serverlessv2scalingconfiguration.html
//
type CfnDBClusterPropsMixin_ServerlessV2ScalingConfigurationProperty struct {
	// The maximum number of Amazon DocumentDB capacity units (DCUs) for an instance in an Amazon DocumentDB Serverless cluster.
	//
	// You can specify DCU values in half-step increments, such as 32, 32.5, 33, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-docdb-dbcluster-serverlessv2scalingconfiguration.html#cfn-docdb-dbcluster-serverlessv2scalingconfiguration-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum number of Amazon DocumentDB capacity units (DCUs) for an instance in an Amazon DocumentDB Serverless cluster.
	//
	// You can specify DCU values in half-step increments, such as 8, 8.5, 9, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-docdb-dbcluster-serverlessv2scalingconfiguration.html#cfn-docdb-dbcluster-serverlessv2scalingconfiguration-mincapacity
	//
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

