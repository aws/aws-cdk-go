package awscdk


// Options for the combine strategy.
//
// Example:
//   // TableV2 uses a Box internally for replicas.
//   // The mixin defers the merge and appends the new replica at synthesis time.
//   table := dynamodb.NewTableV2(scope, jsii.String("Table"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	// L2 prop: pointInTimeRecovery is a boolean
//   	Replicas: []ReplicaTableProps{
//   		&ReplicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   			PointInTimeRecovery: jsii.Boolean(true),
//   		},
//   	},
//   })
//
//   // Mixins always use L1 (CloudFormation) property names and shapes,
//   // regardless of what the L2 API looks like.
//   table.With(awscdkcfnpropertymixins.NewCfnGlobalTablePropsMixin(&CfnGlobalTableMixinProps{
//   	Replicas: []interface{}{
//   		&ReplicaSpecificationProperty{
//   			Region: jsii.String("eu-west-1"),
//   			// L1 prop: pointInTimeRecoverySpecification is an object
//   			PointInTimeRecoverySpecification: &PointInTimeRecoverySpecificationProperty{
//   				PointInTimeRecoveryEnabled: jsii.Boolean(true),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdk.PropertyMergeStrategy_Combine(&CombineStrategyOptions{
//   		Arrays: awscdk.ArrayMergeStrategy_Append(),
//   	}),
//   }))
//
type CombineStrategyOptions struct {
	// Strategy for merging arrays.
	// Default: ArrayMergeStrategy.replace()
	//
	Arrays IArrayMergeStrategy `field:"optional" json:"arrays" yaml:"arrays"`
}

