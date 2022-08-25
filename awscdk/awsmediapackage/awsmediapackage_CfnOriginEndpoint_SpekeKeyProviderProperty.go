package awsmediapackage


// Keyprovider settings for DRM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spekeKeyProviderProperty := &spekeKeyProviderProperty{
//   	resourceId: jsii.String("resourceId"),
//   	roleArn: jsii.String("roleArn"),
//   	systemIds: []*string{
//   		jsii.String("systemIds"),
//   	},
//   	url: jsii.String("url"),
//
//   	// the properties below are optional
//   	certificateArn: jsii.String("certificateArn"),
//   	encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   		presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   		presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   	},
//   }
//
type CfnOriginEndpoint_SpekeKeyProviderProperty struct {
	// Unique identifier for this endpoint, as it is configured in the key provider service.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The ARN for the IAM role that's granted by the key provider to provide access to the key provider API.
	//
	// This role must have a trust policy that allows AWS Elemental MediaPackage to assume the role, and it must have a sufficient permissions policy to allow access to the specific key retrieval URL. Valid format: arn:aws:iam::{accountID}:role/{name}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// List of unique identifiers for the DRM systems to use, as defined in the CPIX specification.
	SystemIds *[]*string `field:"required" json:"systemIds" yaml:"systemIds"`
	// URL for the key provider’s key retrieval API endpoint.
	//
	// Must start with https://.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The Amazon Resource Name (ARN) for the certificate that you imported to AWS Certificate Manager to add content key encryption to this endpoint.
	//
	// For this feature to work, your DRM key provider must support content key encryption.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// `CfnOriginEndpoint.SpekeKeyProviderProperty.EncryptionContractConfiguration`.
	EncryptionContractConfiguration interface{} `field:"optional" json:"encryptionContractConfiguration" yaml:"encryptionContractConfiguration"`
}

