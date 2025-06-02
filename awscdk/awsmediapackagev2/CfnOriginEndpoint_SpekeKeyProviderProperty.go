package awsmediapackagev2


// The parameters for the SPEKE key provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spekeKeyProviderProperty := &SpekeKeyProviderProperty{
//   	DrmSystems: []*string{
//   		jsii.String("drmSystems"),
//   	},
//   	EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   		PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   		PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   	},
//   	ResourceId: jsii.String("resourceId"),
//   	RoleArn: jsii.String("roleArn"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-spekekeyprovider.html
//
type CfnOriginEndpoint_SpekeKeyProviderProperty struct {
	// The DRM solution provider you're using to protect your content during distribution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-spekekeyprovider.html#cfn-mediapackagev2-originendpoint-spekekeyprovider-drmsystems
	//
	DrmSystems *[]*string `field:"required" json:"drmSystems" yaml:"drmSystems"`
	// The encryption contract configuration associated with the SPEKE key provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-spekekeyprovider.html#cfn-mediapackagev2-originendpoint-spekekeyprovider-encryptioncontractconfiguration
	//
	EncryptionContractConfiguration interface{} `field:"required" json:"encryptionContractConfiguration" yaml:"encryptionContractConfiguration"`
	// The unique identifier for the content.
	//
	// The service sends this identifier to the key server to identify the current endpoint. How unique you make this identifier depends on how fine-grained you want access controls to be. The service does not permit you to use the same ID for two simultaneous encryption processes. The resource ID is also known as the content ID.
	//
	// The following example shows a resource ID: `MovieNight20171126093045`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-spekekeyprovider.html#cfn-mediapackagev2-originendpoint-spekekeyprovider-resourceid
	//
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The ARN for the IAM role granted by the key provider that provides access to the key provider API.
	//
	// This role must have a trust policy that allows MediaPackage to assume the role, and it must have a sufficient permissions policy to allow access to the specific key retrieval URL. Get this from your DRM solution provider.
	//
	// Valid format: `arn:aws:iam::{accountID}:role/{name}` . The following example shows a role ARN: `arn:aws:iam::444455556666:role/SpekeAccess`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-spekekeyprovider.html#cfn-mediapackagev2-originendpoint-spekekeyprovider-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The URL of the SPEKE key provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-spekekeyprovider.html#cfn-mediapackagev2-originendpoint-spekekeyprovider-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
}

