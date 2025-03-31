package pipelines


// Properties for a `StackDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackDeploymentProps := &StackDeploymentProps{
//   	AbsoluteTemplatePath: jsii.String("absoluteTemplatePath"),
//   	ConstructPath: jsii.String("constructPath"),
//   	StackArtifactId: jsii.String("stackArtifactId"),
//   	StackName: jsii.String("stackName"),
//
//   	// the properties below are optional
//   	Account: jsii.String("account"),
//   	Assets: []stackAsset{
//   		&stackAsset{
//   			AssetId: jsii.String("assetId"),
//   			AssetManifestPath: jsii.String("assetManifestPath"),
//   			AssetSelector: jsii.String("assetSelector"),
//   			AssetType: awscdk.Pipelines.AssetType_FILE,
//   			IsTemplate: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AssetPublishingRoleArn: jsii.String("assetPublishingRoleArn"),
//   			DisplayName: jsii.String("displayName"),
//   		},
//   	},
//   	AssumeRoleArn: jsii.String("assumeRoleArn"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Region: jsii.String("region"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TemplateS3Uri: jsii.String("templateS3Uri"),
//   }
//
type StackDeploymentProps struct {
	// Template path on disk to cloud assembly (cdk.out).
	AbsoluteTemplatePath *string `field:"required" json:"absoluteTemplatePath" yaml:"absoluteTemplatePath"`
	// Construct path for this stack.
	ConstructPath *string `field:"required" json:"constructPath" yaml:"constructPath"`
	// Artifact ID for this stack.
	StackArtifactId *string `field:"required" json:"stackArtifactId" yaml:"stackArtifactId"`
	// Name for this stack.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Account where the stack should be deployed.
	// Default: - Pipeline account.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Assets referenced by this stack.
	// Default: - No assets.
	//
	Assets *[]*StackAsset `field:"optional" json:"assets" yaml:"assets"`
	// Role to assume before deploying this stack.
	// Default: - Don't assume any role.
	//
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// Execution role to pass to CloudFormation.
	// Default: - No execution role.
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Region where the stack should be deployed.
	// Default: - Pipeline region.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Tags to apply to the stack.
	// Default: - No tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The S3 URL which points to the template asset location in the publishing bucket.
	// Default: - Stack template is not published.
	//
	TemplateS3Uri *string `field:"optional" json:"templateS3Uri" yaml:"templateS3Uri"`
}

