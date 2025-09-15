package awsimagebuilder


// Properties for defining a `CfnComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnComponentProps := &CfnComponentProps{
//   	Name: jsii.String("name"),
//   	Platform: jsii.String("platform"),
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	ChangeDescription: jsii.String("changeDescription"),
//   	Data: jsii.String("data"),
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SupportedOsVersions: []*string{
//   		jsii.String("supportedOsVersions"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html
//
type CfnComponentProps struct {
	// The name of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The operating system platform of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-platform
	//
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// The component version.
	//
	// For example, `1.0.0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-version
	//
	Version *string `field:"required" json:"version" yaml:"version"`
	// The change description of the component.
	//
	// Describes what change has been made in this version, or what makes this version different from other versions of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-changedescription
	//
	ChangeDescription *string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// Component `data` contains inline YAML document content for the component.
	//
	// Alternatively, you can specify the `uri` of a YAML document file stored in Amazon S3. However, you cannot specify both properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-data
	//
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Describes the contents of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) that uniquely identifies the KMS key used to encrypt this component.
	//
	// This can be either the Key ARN or the Alias ARN. For more information, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The operating system (OS) version supported by the component.
	//
	// If the OS information is available, a prefix match is performed against the base image OS version during image recipe creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-supportedosversions
	//
	SupportedOsVersions *[]*string `field:"optional" json:"supportedOsVersions" yaml:"supportedOsVersions"`
	// The tags that apply to the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The `uri` of a YAML component document file.
	//
	// This must be an S3 URL ( `s3://bucket/key` ), and the requester must have permission to access the S3 bucket it points to. If you use Amazon S3, you can specify component content up to your service quota.
	//
	// Alternatively, you can specify the YAML document inline, using the component `data` property. You cannot specify both properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-component.html#cfn-imagebuilder-component-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

