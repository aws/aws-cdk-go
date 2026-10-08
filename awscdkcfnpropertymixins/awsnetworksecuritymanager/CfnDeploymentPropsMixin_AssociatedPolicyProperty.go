package awsnetworksecuritymanager


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   associatedPolicyProperty := &AssociatedPolicyProperty{
//   	PolicyArn: jsii.String("policyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-deployment-associatedpolicy.html
//
type CfnDeploymentPropsMixin_AssociatedPolicyProperty struct {
	// ARN of the associated policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-deployment-associatedpolicy.html#cfn-networksecuritymanager-deployment-associatedpolicy-policyarn
	//
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
}

