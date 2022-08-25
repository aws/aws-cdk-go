package awsimagebuilder


// Properties for defining a `CfnComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnComponentProps := &cfnComponentProps{
//   	name: jsii.String("name"),
//   	platform: jsii.String("platform"),
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	changeDescription: jsii.String("changeDescription"),
//   	data: jsii.String("data"),
//   	description: jsii.String("description"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	supportedOsVersions: []*string{
//   		jsii.String("supportedOsVersions"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	uri: jsii.String("uri"),
//   }
//
type CfnComponentProps struct {
	// The name of the component.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The platform of the component.
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// The component version.
	//
	// For example, `1.0.0` .
	Version *string `field:"required" json:"version" yaml:"version"`
	// The change description of the component.
	//
	// Describes what change has been made in this version, or what makes this version different from other versions of this component.
	ChangeDescription *string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// Component `data` contains inline YAML document content for the component.
	//
	// Alternatively, you can specify the `uri` of a YAML document file stored in Amazon S3. However, you cannot specify both properties.
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The description of the component.
	//
	// Describes the contents of the component.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the KMS key that should be used to encrypt this component.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The operating system (OS) version supported by the component.
	//
	// If the OS information is available, a prefix match is performed against the base image OS version during image recipe creation.
	SupportedOsVersions *[]*string `field:"optional" json:"supportedOsVersions" yaml:"supportedOsVersions"`
	// The tags of the component.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The `uri` of a YAML component document file.
	//
	// This must be an S3 URL ( `s3://bucket/key` ), and the requester must have permission to access the S3 bucket it points to. If you use Amazon S3, you can specify component content up to your service quota.
	//
	// Alternatively, you can specify the YAML document inline, using the component `data` property. You cannot specify both properties.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

