package awsconnect


// Properties for defining a `CfnDataLakeAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataLakeAssociationProps := &CfnDataLakeAssociationProps{
//   	DataSetId: jsii.String("dataSetId"),
//   	InstanceId: jsii.String("instanceId"),
//
//   	// the properties below are optional
//   	TargetAccountId: jsii.String("targetAccountId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datalakeassociation.html
//
type CfnDataLakeAssociationProps struct {
	// The identifier of the analytics data set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datalakeassociation.html#cfn-connect-datalakeassociation-datasetid
	//
	DataSetId *string `field:"required" json:"dataSetId" yaml:"dataSetId"`
	// The identifier of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datalakeassociation.html#cfn-connect-datalakeassociation-instanceid
	//
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The identifier of the target account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-datalakeassociation.html#cfn-connect-datalakeassociation-targetaccountid
	//
	TargetAccountId *string `field:"optional" json:"targetAccountId" yaml:"targetAccountId"`
}

