package previewawsomicsmixins


// If you are using the ECR pull through cache feature, the registry mapping maps between the ECR repository and the upstream registry where container images are pulled and synchronized.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   registryMappingProperty := &RegistryMappingProperty{
//   	EcrAccountId: jsii.String("ecrAccountId"),
//   	EcrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   	UpstreamRegistryUrl: jsii.String("upstreamRegistryUrl"),
//   	UpstreamRepositoryPrefix: jsii.String("upstreamRepositoryPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-registrymapping.html
//
type CfnWorkflowVersionPropsMixin_RegistryMappingProperty struct {
	// Account ID of the account that owns the upstream container image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-registrymapping.html#cfn-omics-workflowversion-registrymapping-ecraccountid
	//
	EcrAccountId *string `field:"optional" json:"ecrAccountId" yaml:"ecrAccountId"`
	// The repository prefix to use in the ECR private repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-registrymapping.html#cfn-omics-workflowversion-registrymapping-ecrrepositoryprefix
	//
	EcrRepositoryPrefix *string `field:"optional" json:"ecrRepositoryPrefix" yaml:"ecrRepositoryPrefix"`
	// The URI of the upstream registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-registrymapping.html#cfn-omics-workflowversion-registrymapping-upstreamregistryurl
	//
	UpstreamRegistryUrl *string `field:"optional" json:"upstreamRegistryUrl" yaml:"upstreamRegistryUrl"`
	// The repository prefix of the corresponding repository in the upstream registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflowversion-registrymapping.html#cfn-omics-workflowversion-registrymapping-upstreamrepositoryprefix
	//
	UpstreamRepositoryPrefix *string `field:"optional" json:"upstreamRepositoryPrefix" yaml:"upstreamRepositoryPrefix"`
}

