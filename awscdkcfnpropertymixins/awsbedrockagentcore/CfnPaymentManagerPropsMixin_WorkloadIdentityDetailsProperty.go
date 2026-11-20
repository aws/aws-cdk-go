package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   workloadIdentityDetailsProperty := &WorkloadIdentityDetailsProperty{
//   	WorkloadIdentityArn: jsii.String("workloadIdentityArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-workloadidentitydetails.html
//
type CfnPaymentManagerPropsMixin_WorkloadIdentityDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-workloadidentitydetails.html#cfn-bedrockagentcore-paymentmanager-workloadidentitydetails-workloadidentityarn
	//
	WorkloadIdentityArn *string `field:"optional" json:"workloadIdentityArn" yaml:"workloadIdentityArn"`
}

