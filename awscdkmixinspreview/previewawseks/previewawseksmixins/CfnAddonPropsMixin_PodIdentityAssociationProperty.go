package previewawseksmixins


// Amazon EKS Pod Identity associations provide the ability to manage credentials for your applications, similar to the way that Amazon EC2 instance profiles provide credentials to Amazon EC2 instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   podIdentityAssociationProperty := &PodIdentityAssociationProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	ServiceAccount: jsii.String("serviceAccount"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-addon-podidentityassociation.html
//
type CfnAddonPropsMixin_PodIdentityAssociationProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role to associate with the service account.
	//
	// The EKS Pod Identity agent manages credentials to assume this role for applications in the containers in the Pods that use this service account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-addon-podidentityassociation.html#cfn-eks-addon-podidentityassociation-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The name of the Kubernetes service account inside the cluster to associate the IAM credentials with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-addon-podidentityassociation.html#cfn-eks-addon-podidentityassociation-serviceaccount
	//
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

