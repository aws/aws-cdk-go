package awseks


// A pod identity to associate with an add-on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   podIdentityAssociationProperty := &PodIdentityAssociationProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	ServiceAccount: jsii.String("serviceAccount"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-addon-podidentityassociation.html
//
type CfnAddon_PodIdentityAssociationProperty struct {
	// The IAM role ARN that the pod identity association is created for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-addon-podidentityassociation.html#cfn-eks-addon-podidentityassociation-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Kubernetes service account that the pod identity association is created for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-addon-podidentityassociation.html#cfn-eks-addon-podidentityassociation-serviceaccount
	//
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

