package previewawseksmixins


// Configuration for integrating Argo CD with IAM Identity Center.
//
// This allows you to use your organization's identity provider for authentication to Argo CD.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   awsIdcProperty := &AwsIdcProperty{
//   	IdcInstanceArn: jsii.String("idcInstanceArn"),
//   	IdcManagedApplicationArn: jsii.String("idcManagedApplicationArn"),
//   	IdcRegion: jsii.String("idcRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-awsidc.html
//
type CfnCapabilityPropsMixin_AwsIdcProperty struct {
	// The ARN of the IAM Identity Center instance to use for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-awsidc.html#cfn-eks-capability-awsidc-idcinstancearn
	//
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// The ARN of the managed application created in IAM Identity Center for this Argo CD capability.
	//
	// This application is automatically created and managed by EKS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-awsidc.html#cfn-eks-capability-awsidc-idcmanagedapplicationarn
	//
	IdcManagedApplicationArn *string `field:"optional" json:"idcManagedApplicationArn" yaml:"idcManagedApplicationArn"`
	// The Region where your IAM Identity Center instance is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-capability-awsidc.html#cfn-eks-capability-awsidc-idcregion
	//
	IdcRegion *string `field:"optional" json:"idcRegion" yaml:"idcRegion"`
}

