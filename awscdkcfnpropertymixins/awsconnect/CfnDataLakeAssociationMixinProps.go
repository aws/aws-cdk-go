package awsconnect


// Properties for CfnDataLakeAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDataLakeAssociationMixinProps := &CfnDataLakeAssociationMixinProps{
//   	DataSetId: jsii.String("dataSetId"),
//   	InstanceId: jsii.String("instanceId"),
//   	TargetAccountId: jsii.String("targetAccountId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datalakeassociation.html
//
type CfnDataLakeAssociationMixinProps struct {
	// The identifier of the analytics data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datalakeassociation.html#cfn-connect-datalakeassociation-datasetid
	//
	DataSetId *string `field:"optional" json:"dataSetId" yaml:"dataSetId"`
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datalakeassociation.html#cfn-connect-datalakeassociation-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The identifier of the target account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datalakeassociation.html#cfn-connect-datalakeassociation-targetaccountid
	//
	TargetAccountId *string `field:"optional" json:"targetAccountId" yaml:"targetAccountId"`
}

