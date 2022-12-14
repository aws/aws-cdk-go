package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validationProfileProperty := &validationProfileProperty{
//   	profileName: jsii.String("profileName"),
//   	transformJobDefinition: &transformJobDefinitionProperty{
//   		transformInput: &transformInputProperty{
//   			dataSource: &dataSourceProperty{
//   				s3DataSource: &s3DataSourceProperty{
//   					s3DataType: jsii.String("s3DataType"),
//   					s3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			compressionType: jsii.String("compressionType"),
//   			contentType: jsii.String("contentType"),
//   			splitType: jsii.String("splitType"),
//   		},
//   		transformOutput: &transformOutputProperty{
//   			s3OutputPath: jsii.String("s3OutputPath"),
//
//   			// the properties below are optional
//   			accept: jsii.String("accept"),
//   			assembleWith: jsii.String("assembleWith"),
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   		},
//   		transformResources: &transformResourcesProperty{
//   			instanceCount: jsii.Number(123),
//   			instanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//
//   		// the properties below are optional
//   		batchStrategy: jsii.String("batchStrategy"),
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		maxConcurrentTransforms: jsii.Number(123),
//   		maxPayloadInMb: jsii.Number(123),
//   	},
//   }
//
type CfnModelPackage_ValidationProfileProperty struct {
	// `CfnModelPackage.ValidationProfileProperty.ProfileName`.
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// `CfnModelPackage.ValidationProfileProperty.TransformJobDefinition`.
	TransformJobDefinition interface{} `field:"required" json:"transformJobDefinition" yaml:"transformJobDefinition"`
}

