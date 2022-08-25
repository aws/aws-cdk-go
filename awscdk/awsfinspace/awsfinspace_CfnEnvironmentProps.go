package awsfinspace


// Properties for defining a `CfnEnvironment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributeMap interface{}
//
//   cfnEnvironmentProps := &cfnEnvironmentProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	dataBundles: []*string{
//   		jsii.String("dataBundles"),
//   	},
//   	description: jsii.String("description"),
//   	federationMode: jsii.String("federationMode"),
//   	federationParameters: &federationParametersProperty{
//   		applicationCallBackUrl: jsii.String("applicationCallBackUrl"),
//   		attributeMap: attributeMap,
//   		federationProviderName: jsii.String("federationProviderName"),
//   		federationUrn: jsii.String("federationUrn"),
//   		samlMetadataDocument: jsii.String("samlMetadataDocument"),
//   		samlMetadataUrl: jsii.String("samlMetadataUrl"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	superuserParameters: &superuserParametersProperty{
//   		emailAddress: jsii.String("emailAddress"),
//   		firstName: jsii.String("firstName"),
//   		lastName: jsii.String("lastName"),
//   	},
//   }
//
type CfnEnvironmentProps struct {
	// The name of the FinSpace environment.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of Amazon Resource Names (ARN) of the data bundles to install. Currently supported data bundle ARNs:.
	//
	// - `arn:aws:finspace:${Region}::data-bundle/capital-markets-sample` - Contains sample Capital Markets datasets, categories and controlled vocabularies.
	// - `arn:aws:finspace:${Region}::data-bundle/taq` (default) - Contains trades and quotes data in addition to sample Capital Markets data.
	DataBundles *[]*string `field:"optional" json:"dataBundles" yaml:"dataBundles"`
	// The description of the FinSpace environment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The authentication mode for the environment.
	FederationMode *string `field:"optional" json:"federationMode" yaml:"federationMode"`
	// Configuration information when authentication mode is FEDERATED.
	FederationParameters interface{} `field:"optional" json:"federationParameters" yaml:"federationParameters"`
	// The KMS key id used to encrypt in the FinSpace environment.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Configuration information for the superuser.
	SuperuserParameters interface{} `field:"optional" json:"superuserParameters" yaml:"superuserParameters"`
}

