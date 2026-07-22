package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   modelPackageContainerDefinitionProperty := &ModelPackageContainerDefinitionProperty{
//   	ContainerHostname: jsii.String("containerHostname"),
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	Framework: jsii.String("framework"),
//   	FrameworkVersion: jsii.String("frameworkVersion"),
//   	Image: jsii.String("image"),
//   	ImageDigest: jsii.String("imageDigest"),
//   	IsCheckpoint: jsii.Boolean(false),
//   	ModelInput: &ModelInputProperty{
//   		DataInputConfig: jsii.String("dataInputConfig"),
//   	},
//   	NearestModelName: jsii.String("nearestModelName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelpackagecontainerdefinition.html
//
type CfnAlgorithmPropsMixin_ModelPackageContainerDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelpackagecontainerdefinition.html#cfn-sagemaker-algorithm-modelpackagecontainerdefinition-containerhostname
	//
	ContainerHostname *string `field:"optional" json:"containerHostname" yaml:"containerHostname"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelpackagecontainerdefinition.html#cfn-sagemaker-algorithm-modelpackagecontainerdefinition-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelpackagecontainerdefinition.html#cfn-sagemaker-algorithm-modelpackagecontainerdefinition-framework
	//
	Framework *string `field:"optional" json:"framework" yaml:"framework"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelpackagecontainerdefinition.html#cfn-sagemaker-algorithm-modelpackagecontainerdefinition-frameworkversion
	//
	FrameworkVersion *string `field:"optional" json:"frameworkVersion" yaml:"frameworkVersion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelpackagecontainerdefinition.html#cfn-sagemaker-algorithm-modelpackagecontainerdefinition-image
	//
	Image *string `field:"optional" json:"image" yaml:"image"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelpackagecontainerdefinition.html#cfn-sagemaker-algorithm-modelpackagecontainerdefinition-imagedigest
	//
	ImageDigest *string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelpackagecontainerdefinition.html#cfn-sagemaker-algorithm-modelpackagecontainerdefinition-ischeckpoint
	//
	IsCheckpoint interface{} `field:"optional" json:"isCheckpoint" yaml:"isCheckpoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelpackagecontainerdefinition.html#cfn-sagemaker-algorithm-modelpackagecontainerdefinition-modelinput
	//
	ModelInput interface{} `field:"optional" json:"modelInput" yaml:"modelInput"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-algorithm-modelpackagecontainerdefinition.html#cfn-sagemaker-algorithm-modelpackagecontainerdefinition-nearestmodelname
	//
	NearestModelName *string `field:"optional" json:"nearestModelName" yaml:"nearestModelName"`
}

