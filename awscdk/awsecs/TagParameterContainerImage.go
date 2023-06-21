package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/constructs-go/constructs/v10"
)

// A special type of `ContainerImage` that uses an ECR repository for the image, but a CloudFormation Parameter for the tag of the image in that repository.
//
// This allows providing this tag through the Parameter at deploy time,
// for example in a CodePipeline that pushes a new tag of the image to the repository during a build step,
// and then provides that new tag through the CloudFormation Parameter in the deploy step.
//
// Example:
//   /**
//    * These are the construction properties for `EcsAppStack`.
//    * They extend the standard Stack properties,
//    * but also require providing the ContainerImage that the service will use.
//    * That Image will be provided from the Stack containing the CodePipeline.
//    */
//   type ecsAppStackProps struct {
//   	stackProps
//   	image containerImage
//   }
//
//   /**
//    * This is the Stack containing a simple ECS Service that uses the provided ContainerImage.
//    */
//   type EcsAppStack struct {
//   	stack
//   }
//
//   func NewEcsAppStack(scope construct, id *string, props ecsAppStackProps) *EcsAppStack {
//   	this := &EcsAppStack{}
//   	cdk.NewStack_Override(this, scope, id, props)
//
//   	taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TaskDefinition"), &TaskDefinitionProps{
//   		Compatibility: ecs.Compatibility_FARGATE,
//   		Cpu: jsii.String("1024"),
//   		MemoryMiB: jsii.String("2048"),
//   	})
//   	taskDefinition.AddContainer(jsii.String("AppContainer"), &ContainerDefinitionOptions{
//   		Image: props.image,
//   	})
//   	ecs.NewFargateService(this, jsii.String("EcsService"), &FargateServiceProps{
//   		TaskDefinition: TaskDefinition,
//   		Cluster: ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   			Vpc: ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   				MaxAzs: jsii.Number(1),
//   			}),
//   		}),
//   	})
//   	return this
//   }
//
//   /**
//    * This is the Stack containing the CodePipeline definition that deploys an ECS Service.
//    */
//   type PipelineStack struct {
//   	stack
//   	tagParameterContainerImage tagParameterContainerImage
//   }tagParameterContainerImage tagParameterContainerImage
//
//   func NewPipelineStack(scope construct, id *string, props stackProps) *PipelineStack {
//   	this := &PipelineStack{}
//   	cdk.NewStack_Override(this, scope, id, props)
//
//   	/* ********** ECS part **************** */
//
//   	// this is the ECR repository where the built Docker image will be pushed
//   	appEcrRepo := ecr.NewRepository(this, jsii.String("EcsDeployRepository"))
//   	// the build that creates the Docker image, and pushes it to the ECR repo
//   	appCodeDockerBuild := codebuild.NewPipelineProject(this, jsii.String("AppCodeDockerImageBuildAndPushProject"), &PipelineProjectProps{
//   		Environment: &BuildEnvironment{
//   			// we need to run Docker
//   			Privileged: jsii.Boolean(true),
//   		},
//   		BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   			"version": jsii.String("0.2"),
//   			"phases": map[string]map[string][]*string{
//   				"build": map[string][]*string{
//   					"commands": []*string{
//   						jsii.String("$(aws ecr get-login --region $AWS_DEFAULT_REGION --no-include-email)"),
//   						jsii.String("docker build -t $REPOSITORY_URI:$CODEBUILD_RESOLVED_SOURCE_VERSION ."),
//   					},
//   				},
//   				"post_build": map[string][]*string{
//   					"commands": []*string{
//   						jsii.String("docker push $REPOSITORY_URI:$CODEBUILD_RESOLVED_SOURCE_VERSION"),
//   						jsii.String("export imageTag=$CODEBUILD_RESOLVED_SOURCE_VERSION"),
//   					},
//   				},
//   			},
//   			"env": map[string][]*string{
//   				// save the imageTag environment variable as a CodePipeline Variable
//   				"exported-variables": []*string{
//   					jsii.String("imageTag"),
//   				},
//   			},
//   		}),
//   		EnvironmentVariables: map[string]buildEnvironmentVariable{
//   			"REPOSITORY_URI": &buildEnvironmentVariable{
//   				"value": appEcrRepo.repositoryUri,
//   			},
//   		},
//   	})
//   	// needed for `docker push`
//   	appEcrRepo.GrantPullPush(appCodeDockerBuild)
//   	// create the ContainerImage used for the ECS application Stack
//   	this.tagParameterContainerImage = ecs.NewTagParameterContainerImage(appEcrRepo)
//
//   	cdkCodeBuild := codebuild.NewPipelineProject(this, jsii.String("CdkCodeBuildProject"), &PipelineProjectProps{
//   		BuildSpec: codebuild.BuildSpec_*FromObject(map[string]interface{}{
//   			"version": jsii.String("0.2"),
//   			"phases": map[string]map[string][]*string{
//   				"install": map[string][]*string{
//   					"commands": []*string{
//   						jsii.String("npm install"),
//   					},
//   				},
//   				"build": map[string][]*string{
//   					"commands": []*string{
//   						jsii.String("npx cdk synth --verbose"),
//   					},
//   				},
//   			},
//   			"artifacts": map[string]*string{
//   				// store the entire Cloud Assembly as the output artifact
//   				"base-directory": jsii.String("cdk.out"),
//   				"files": jsii.String("**/*"),
//   			},
//   		}),
//   	})
//
//   	/* ********** Pipeline part **************** */
//
//   	appCodeSourceOutput := codepipeline.NewArtifact()
//   	cdkCodeSourceOutput := codepipeline.NewArtifact()
//   	cdkCodeBuildOutput := codepipeline.NewArtifact()
//   	appCodeBuildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   		ActionName: jsii.String("AppCodeDockerImageBuildAndPush"),
//   		Project: appCodeDockerBuild,
//   		Input: appCodeSourceOutput,
//   	})
//   	codepipeline.NewPipeline(this, jsii.String("CodePipelineDeployingEcsApplication"), &PipelineProps{
//   		ArtifactBucket: s3.NewBucket(this, jsii.String("ArtifactBucket"), &BucketProps{
//   			RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   		}),
//   		Stages: []stageProps{
//   			&stageProps{
//   				StageName: jsii.String("Source"),
//   				Actions: []iAction{
//   					// this is the Action that takes the source of your application code
//   					codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
//   						ActionName: jsii.String("AppCodeSource"),
//   						Repository: codecommit.NewRepository(this, jsii.String("AppCodeSourceRepository"), &RepositoryProps{
//   							RepositoryName: jsii.String("AppCodeSourceRepository"),
//   						}),
//   						Output: appCodeSourceOutput,
//   					}),
//   					// this is the Action that takes the source of your CDK code
//   					// (which would probably include this Pipeline code as well)
//   					codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
//   						ActionName: jsii.String("CdkCodeSource"),
//   						Repository: codecommit.NewRepository(this, jsii.String("CdkCodeSourceRepository"), &RepositoryProps{
//   							RepositoryName: jsii.String("CdkCodeSourceRepository"),
//   						}),
//   						Output: cdkCodeSourceOutput,
//   					}),
//   				},
//   			},
//   			&stageProps{
//   				StageName: jsii.String("Build"),
//   				Actions: []*iAction{
//   					appCodeBuildAction,
//   					codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   						ActionName: jsii.String("CdkCodeBuildAndSynth"),
//   						Project: cdkCodeBuild,
//   						Input: cdkCodeSourceOutput,
//   						Outputs: []artifact{
//   							cdkCodeBuildOutput,
//   						},
//   					}),
//   				},
//   			},
//   			&stageProps{
//   				StageName: jsii.String("Deploy"),
//   				Actions: []*iAction{
//   					codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&CloudFormationCreateUpdateStackActionProps{
//   						ActionName: jsii.String("CFN_Deploy"),
//   						StackName: jsii.String("SampleEcsStackDeployedFromCodePipeline"),
//   						// this name has to be the same name as used below in the CDK code for the application Stack
//   						TemplatePath: cdkCodeBuildOutput.AtPath(jsii.String("EcsStackDeployedInPipeline.template.json")),
//   						AdminPermissions: jsii.Boolean(true),
//   						ParameterOverrides: map[string]interface{}{
//   							// read the tag pushed to the ECR repository from the CodePipeline Variable saved by the application build step,
//   							// and pass it as the CloudFormation Parameter for the tag
//   							this.tagParameterContainerImage.tagParameterName: appCodeBuildAction.variable(jsii.String("imageTag")),
//   						},
//   					}),
//   				},
//   			},
//   		},
//   	})
//   	return this
//   }
//
//   app := cdk.NewApp()
//
//   // the CodePipeline Stack needs to be created first
//   pipelineStack := NewPipelineStack(app, jsii.String("aws-cdk-pipeline-ecs-separate-sources"))
//   // we supply the image to the ECS application Stack from the CodePipeline Stack
//   // we supply the image to the ECS application Stack from the CodePipeline Stack
//   NewEcsAppStack(app, jsii.String("EcsStackDeployedInPipeline"), &ecsAppStackProps{
//   	image: pipelineStack.tagParameterContainerImage,
//   })
//
// See:  #tagParameterName.
//
type TagParameterContainerImage interface {
	ContainerImage
	// Returns the name of the CloudFormation Parameter that represents the tag of the image in the ECR repository.
	TagParameterName() *string
	// Returns the value of the CloudFormation Parameter that represents the tag of the image in the ECR repository.
	TagParameterValue() *string
	// Called when the image is used by a ContainerDefinition.
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for TagParameterContainerImage
type jsiiProxy_TagParameterContainerImage struct {
	jsiiProxy_ContainerImage
}

func (j *jsiiProxy_TagParameterContainerImage) TagParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagParameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TagParameterContainerImage) TagParameterValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagParameterValue",
		&returns,
	)
	return returns
}


func NewTagParameterContainerImage(repository awsecr.IRepository) TagParameterContainerImage {
	_init_.Initialize()

	if err := validateNewTagParameterContainerImageParameters(repository); err != nil {
		panic(err)
	}
	j := jsiiProxy_TagParameterContainerImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		[]interface{}{repository},
		&j,
	)

	return &j
}

func NewTagParameterContainerImage_Override(t TagParameterContainerImage, repository awsecr.IRepository) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		[]interface{}{repository},
		t,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
func TagParameterContainerImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	if err := validateTagParameterContainerImage_FromAssetParameters(directory, props); err != nil {
		panic(err)
	}
	var returns AssetImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
func TagParameterContainerImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	if err := validateTagParameterContainerImage_FromDockerImageAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
func TagParameterContainerImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	if err := validateTagParameterContainerImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns EcrImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func TagParameterContainerImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	if err := validateTagParameterContainerImage_FromRegistryParameters(name, props); err != nil {
		panic(err)
	}
	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		"fromRegistry",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Use an existing tarball for this container image.
//
// Use this method if the container image has already been created by another process (e.g. jib)
// and you want to add it as a container image asset.
func TagParameterContainerImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	if err := validateTagParameterContainerImage_FromTarballParameters(tarballFile); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.TagParameterContainerImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TagParameterContainerImage) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	if err := t.validateBindParameters(scope, containerDefinition); err != nil {
		panic(err)
	}
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

