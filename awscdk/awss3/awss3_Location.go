package awss3


// An interface that represents the location of a specific object in an S3 Bucket.
//
// Example:
//   lambdaStack := cdk.NewStack(app, jsii.String("LambdaStack"))
//   lambdaCode := lambda.code.fromCfnParameters()
//   lambda.NewFunction(lambdaStack, jsii.String("Lambda"), &functionProps{
//   	code: lambdaCode,
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   })
//   // other resources that your Lambda needs, added to the lambdaStack...
//
//   pipelineStack := cdk.NewStack(app, jsii.String("PipelineStack"))
//   pipeline := codepipeline.NewPipeline(pipelineStack, jsii.String("Pipeline"))
//
//   // add the source code repository containing this code to your Pipeline,
//   // and the source code of the Lambda Function, if they're separate
//   cdkSourceOutput := codepipeline.NewArtifact()
//   cdkSourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	repository: codecommit.NewRepository(pipelineStack, jsii.String("CdkCodeRepo"), &repositoryProps{
//   		repositoryName: jsii.String("CdkCodeRepo"),
//   	}),
//   	actionName: jsii.String("CdkCode_Source"),
//   	output: cdkSourceOutput,
//   })
//   lambdaSourceOutput := codepipeline.NewArtifact()
//   lambdaSourceAction := codepipeline_actions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	repository: codecommit.NewRepository(pipelineStack, jsii.String("LambdaCodeRepo"), &repositoryProps{
//   		repositoryName: jsii.String("LambdaCodeRepo"),
//   	}),
//   	actionName: jsii.String("LambdaCode_Source"),
//   	output: lambdaSourceOutput,
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		cdkSourceAction,
//   		lambdaSourceAction,
//   	},
//   })
//
//   // synthesize the Lambda CDK template, using CodeBuild
//   // the below values are just examples, assuming your CDK code is in TypeScript/JavaScript -
//   // adjust the build environment and/or commands accordingly
//   cdkBuildProject := codebuild.NewProject(pipelineStack, jsii.String("CdkBuildProject"), &projectProps{
//   	environment: &buildEnvironment{
//   		buildImage: codebuild.linuxBuildImage_UBUNTU_14_04_NODEJS_10_1_0(),
//   	},
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
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
//   cdkBuildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("CDK_Build"),
//   	project: cdkBuildProject,
//   	input: cdkSourceOutput,
//   	outputs: []artifact{
//   		cdkBuildOutput,
//   	},
//   })
//
//   // build your Lambda code, using CodeBuild
//   // again, this example assumes your Lambda is written in TypeScript/JavaScript -
//   // make sure to adjust the build environment and/or commands if they don't match your specific situation
//   lambdaBuildProject := codebuild.NewProject(pipelineStack, jsii.String("LambdaBuildProject"), &projectProps{
//   	environment: &buildEnvironment{
//   		buildImage: codebuild.*linuxBuildImage_UBUNTU_14_04_NODEJS_10_1_0(),
//   	},
//   	buildSpec: codebuild.*buildSpec.fromObject(map[string]interface{}{
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
//   lambdaBuildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("Lambda_Build"),
//   	project: lambdaBuildProject,
//   	input: lambdaSourceOutput,
//   	outputs: []*artifact{
//   		lambdaBuildOutput,
//   	},
//   })
//
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Build"),
//   	actions: []*iAction{
//   		cdkBuildAction,
//   		lambdaBuildAction,
//   	},
//   })
//
//   // finally, deploy your Lambda Stack
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   	actions: []*iAction{
//   		codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
//   			actionName: jsii.String("Lambda_CFN_Deploy"),
//   			templatePath: cdkBuildOutput.atPath(jsii.String("LambdaStack.template.yaml")),
//   			stackName: jsii.String("LambdaStackDeployedName"),
//   			adminPermissions: jsii.Boolean(true),
//   			parameterOverrides: lambdaCode.assign(lambdaBuildOutput.s3Location),
//   			extraInputs: []*artifact{
//   				lambdaBuildOutput,
//   			},
//   		}),
//   	},
//   })
//
type Location struct {
	// The name of the S3 Bucket the object is in.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The path inside the Bucket where the object is located at.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// The S3 object version.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

