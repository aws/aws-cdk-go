package awsopensearchserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityLimitsProperty := &CapacityLimitsProperty{
//   	MaxIndexingCapacityInOcu: jsii.Number(123),
//   	MaxSearchCapacityInOcu: jsii.Number(123),
//   	MinIndexingCapacityInOcu: jsii.Number(123),
//   	MinSearchCapacityInOcu: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collectiongroup-capacitylimits.html
//
type CfnCollectionGroup_CapacityLimitsProperty struct {
	// The maximum indexing capacity for collections in the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collectiongroup-capacitylimits.html#cfn-opensearchserverless-collectiongroup-capacitylimits-maxindexingcapacityinocu
	//
	MaxIndexingCapacityInOcu *float64 `field:"optional" json:"maxIndexingCapacityInOcu" yaml:"maxIndexingCapacityInOcu"`
	// The maximum search capacity for collections in the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collectiongroup-capacitylimits.html#cfn-opensearchserverless-collectiongroup-capacitylimits-maxsearchcapacityinocu
	//
	MaxSearchCapacityInOcu *float64 `field:"optional" json:"maxSearchCapacityInOcu" yaml:"maxSearchCapacityInOcu"`
	// The minimum indexing capacity for collections in the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collectiongroup-capacitylimits.html#cfn-opensearchserverless-collectiongroup-capacitylimits-minindexingcapacityinocu
	//
	MinIndexingCapacityInOcu *float64 `field:"optional" json:"minIndexingCapacityInOcu" yaml:"minIndexingCapacityInOcu"`
	// The minimum search capacity for collections in the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchserverless-collectiongroup-capacitylimits.html#cfn-opensearchserverless-collectiongroup-capacitylimits-minsearchcapacityinocu
	//
	MinSearchCapacityInOcu *float64 `field:"optional" json:"minSearchCapacityInOcu" yaml:"minSearchCapacityInOcu"`
}

