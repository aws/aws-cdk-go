package awseks


// The workload identity configuration of the pod identity association.
//
// When enabled, the association also provisions a SPIFFE identity for the workload. Setting Enabled to false disables workload identity and keeps the ARNs. Removing the property from the template leaves the association unchanged; set Enabled to false to disable workload identity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   workloadIdentityConfigProperty := &WorkloadIdentityConfigProperty{
//   	Enabled: jsii.Boolean(false),
//   	ProviderArn: jsii.String("providerArn"),
//   	TemplateArn: jsii.String("templateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-podidentityassociation-workloadidentityconfig.html
//
type CfnPodIdentityAssociationPropsMixin_WorkloadIdentityConfigProperty struct {
	// Whether workload identity is enabled for the association.
	//
	// Enabling it requires TemplateArn and ProviderArn, and cannot be combined with TargetRoleArn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-podidentityassociation-workloadidentityconfig.html#cfn-eks-podidentityassociation-workloadidentityconfig-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ARN of the customer-owned IAM workload identity provider that produces the trusted context assertion for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-podidentityassociation-workloadidentityconfig.html#cfn-eks-podidentityassociation-workloadidentityconfig-providerarn
	//
	ProviderArn *string `field:"optional" json:"providerArn" yaml:"providerArn"`
	// The ARN of the customer-owned IAM workload identity template that renders the workload's SPIFFE identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-podidentityassociation-workloadidentityconfig.html#cfn-eks-podidentityassociation-workloadidentityconfig-templatearn
	//
	TemplateArn *string `field:"optional" json:"templateArn" yaml:"templateArn"`
}

