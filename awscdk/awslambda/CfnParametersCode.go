package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Lambda code defined using 2 CloudFormation parameters.
//
// Useful when you don't have access to the code of your Lambda from your CDK code, so you can't use Assets,
// and you want to deploy the Lambda in a CodePipeline, using CloudFormation Actions -
// you can fill the parameters using the `#assign` method.
//
// Example:
//   lambdaStack := cdk.NewStack(app, jsii.String("LambdaStack"))
//   lambdaCode := lambda.Code_FromCfnParameters()
//   lambda.NewFunction(lambdaStack, jsii.String("Lambda"), &FunctionProps{
//   	Code: lambdaCode,
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   })
//   // other resources that your Lambda needs, added to the lambdaStack...
//
//   pipelineStack := cdk.NewStack(app, jsii.String("PipelineStack"))
//   pipeline := codepipeline.NewPipeline(pipelineStack, jsii.String("Pipeline"), &PipelineProps{
//   	CrossAccountKeys: jsii.Boolean(true),
//   })
//
//   // add the source code repository containing this code to your Pipeline,
//   // and the source code of the Lambda Function, if they're separate
//   cdkSourceOutput := codepipeline.NewArtifact()
//   cdkSourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
//   	Repository: codecommit.NewRepository(pipelineStack, jsii.String("CdkCodeRepo"), &RepositoryProps{
//   		RepositoryName: jsii.String("CdkCodeRepo"),
//   	}),
//   	ActionName: jsii.String("CdkCode_Source"),
//   	Output: cdkSourceOutput,
//   })
//   lambdaSourceOutput := codepipeline.NewArtifact()
//   lambdaSourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
//   	Repository: codecommit.NewRepository(pipelineStack, jsii.String("LambdaCodeRepo"), &RepositoryProps{
//   		RepositoryName: jsii.String("LambdaCodeRepo"),
//   	}),
//   	ActionName: jsii.String("LambdaCode_Source"),
//   	Output: lambdaSourceOutput,
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Source"),
//   	Actions: []iAction{
//   		cdkSourceAction,
//   		lambdaSourceAction,
//   	},
//   })
//
//   // synthesize the Lambda CDK template, using CodeBuild
//   // the below values are just examples, assuming your CDK code is in TypeScript/JavaScript -
//   // adjust the build environment and/or commands accordingly
//   cdkBuildProject := codebuild.NewProject(pipelineStack, jsii.String("CdkBuildProject"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		BuildImage: codebuild.LinuxBuildImage_STANDARD_7_0(),
//   	},
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string]*string{
//   			"install": map[string]*string{
//   				"commands": jsii.String("npm install"),
//   			},
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("npm run build"),
//   					jsii.String("npm run cdk synth LambdaStack -- -o ."),
//   				},
//   			},
//   		},
//   		"artifacts": map[string]*string{
//   			"files": jsii.String("LambdaStack.template.yaml"),
//   		},
//   	}),
//   })
//   cdkBuildOutput := codepipeline.NewArtifact()
//   cdkBuildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("CDK_Build"),
//   	Project: cdkBuildProject,
//   	Input: cdkSourceOutput,
//   	Outputs: []artifact{
//   		cdkBuildOutput,
//   	},
//   })
//
//   // build your Lambda code, using CodeBuild
//   // again, this example assumes your Lambda is written in TypeScript/JavaScript -
//   // make sure to adjust the build environment and/or commands if they don't match your specific situation
//   lambdaBuildProject := codebuild.NewProject(pipelineStack, jsii.String("LambdaBuildProject"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		BuildImage: codebuild.LinuxBuildImage_STANDARD_7_0(),
//   	},
//   	BuildSpec: codebuild.BuildSpec_*FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string]*string{
//   			"install": map[string]*string{
//   				"commands": jsii.String("npm install"),
//   			},
//   			"build": map[string]*string{
//   				"commands": jsii.String("npm run build"),
//   			},
//   		},
//   		"artifacts": map[string][]*string{
//   			"files": []*string{
//   				jsii.String("index.js"),
//   				jsii.String("node_modules/**/*"),
//   			},
//   		},
//   	}),
//   })
//   lambdaBuildOutput := codepipeline.NewArtifact()
//   lambdaBuildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("Lambda_Build"),
//   	Project: lambdaBuildProject,
//   	Input: lambdaSourceOutput,
//   	Outputs: []*artifact{
//   		lambdaBuildOutput,
//   	},
//   })
//
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Build"),
//   	Actions: []*iAction{
//   		cdkBuildAction,
//   		lambdaBuildAction,
//   	},
//   })
//
//   // finally, deploy your Lambda Stack
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   	Actions: []*iAction{
//   		codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
//   			ActionName: jsii.String("Lambda_CFN_Deploy"),
//   			TemplatePath: cdkBuildOutput.AtPath(jsii.String("LambdaStack.template.yaml")),
//   			StackName: jsii.String("LambdaStackDeployedName"),
//   			AdminPermissions: jsii.Boolean(true),
//   			ParameterOverrides: lambdaCode.Assign(lambdaBuildOutput.s3Location),
//   			ExtraInputs: []*artifact{
//   				lambdaBuildOutput,
//   			},
//   		}),
//   	},
//   })
//
type CfnParametersCode interface {
	Code
	BucketNameParam() *string
	// Determines whether this Code is inline code or not.
	IsInline() *bool
	ObjectKeyParam() *string
	// Create a parameters map from this instance's CloudFormation parameters.
	//
	// It returns a map with 2 keys that correspond to the names of the parameters defined in this Lambda code,
	// and as values it contains the appropriate expressions pointing at the provided S3 location
	// (most likely, obtained from a CodePipeline Artifact by calling the `artifact.s3Location` method).
	// The result should be provided to the CloudFormation Action
	// that is deploying the Stack that the Lambda with this code is part of,
	// in the `parameterOverrides` property.
	Assign(location *awss3.Location) *map[string]interface{}
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions)
}

// The jsii proxy struct for CfnParametersCode
type jsiiProxy_CfnParametersCode struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_CfnParametersCode) BucketNameParam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParametersCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParametersCode) ObjectKeyParam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectKeyParam",
		&returns,
	)
	return returns
}


func NewCfnParametersCode(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	if err := validateNewCfnParametersCodeParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnParametersCode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCfnParametersCode_Override(c CfnParametersCode, props *CfnParametersCodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		[]interface{}{props},
		c,
	)
}

// Loads the function code from a local disk path.
func CfnParametersCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateCfnParametersCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func CfnParametersCode_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	if err := validateCfnParametersCode_FromAssetImageParameters(directory, props); err != nil {
		panic(err)
	}
	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func CfnParametersCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateCfnParametersCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`.
func CfnParametersCode_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	if err := validateCfnParametersCode_FromCfnParametersParameters(props); err != nil {
		panic(err)
	}
	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func CfnParametersCode_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateCfnParametersCode_FromDockerBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func CfnParametersCode_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	if err := validateCfnParametersCode_FromEcrImageParameters(repository, props); err != nil {
		panic(err)
	}
	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func CfnParametersCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateCfnParametersCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParametersCode) Assign(location *awss3.Location) *map[string]interface{} {
	if err := c.validateAssignParameters(location); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"assign",
		[]interface{}{location},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParametersCode) Bind(scope constructs.Construct) *CodeConfig {
	if err := c.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnParametersCode) BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions) {
	if err := c.validateBindToResourceParameters(_resource, _options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}

