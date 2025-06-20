package awsgreengrass


// Settings for an Amazon S3 machine learning resource.
//
// For more information, see [Perform Machine Learning Inference](https://docs.aws.amazon.com/greengrass/v1/developerguide/ml-inference.html) in the *Developer Guide* .
//
// In an AWS CloudFormation template, `S3MachineLearningModelResourceData` can be used in the [`ResourceDataContainer`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourcedatacontainer.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3MachineLearningModelResourceDataProperty := &S3MachineLearningModelResourceDataProperty{
//   	DestinationPath: jsii.String("destinationPath"),
//   	S3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	OwnerSetting: &ResourceDownloadOwnerSettingProperty{
//   		GroupOwner: jsii.String("groupOwner"),
//   		GroupPermission: jsii.String("groupPermission"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-s3machinelearningmodelresourcedata.html
//
type CfnResourceDefinition_S3MachineLearningModelResourceDataProperty struct {
	// The absolute local path of the resource inside the Lambda environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-s3machinelearningmodelresourcedata.html#cfn-greengrass-resourcedefinition-s3machinelearningmodelresourcedata-destinationpath
	//
	DestinationPath *string `field:"required" json:"destinationPath" yaml:"destinationPath"`
	// The URI of the source model in an Amazon S3 bucket.
	//
	// The model package must be in `tar.gz` or `.zip` format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-s3machinelearningmodelresourcedata.html#cfn-greengrass-resourcedefinition-s3machinelearningmodelresourcedata-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The owner setting for the downloaded machine learning resource.
	//
	// For more information, see [Access Machine Learning Resources from Lambda Functions](https://docs.aws.amazon.com/greengrass/v1/developerguide/access-ml-resources.html) in the *Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-s3machinelearningmodelresourcedata.html#cfn-greengrass-resourcedefinition-s3machinelearningmodelresourcedata-ownersetting
	//
	OwnerSetting interface{} `field:"optional" json:"ownerSetting" yaml:"ownerSetting"`
}

