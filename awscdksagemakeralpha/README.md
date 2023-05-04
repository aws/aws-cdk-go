# Amazon SageMaker Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

Amazon SageMaker provides every developer and data scientist with the ability to build, train, and
deploy machine learning models quickly. Amazon SageMaker is a fully-managed service that covers the
entire machine learning workflow to label and prepare your data, choose an algorithm, train the
model, tune and optimize it for deployment, make predictions, and take action. Your models get to
production faster with much less effort and lower cost.

## Model

To create a machine learning model with Amazon Sagemaker, use the `Model` construct. This construct
includes properties that can be configured to define model components, including the model inference
code as a Docker image and an optional set of separate model data artifacts. See the [AWS
documentation](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-marketplace-develop.html)
to learn more about SageMaker models.

### Single Container Model

In the event that a single container is sufficient for your inference use-case, you can define a
single-container model:

```go
import "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
import "github.com/aws-samples/dummy/path"


image := sagemaker.ContainerImage_FromAsset(path.join(jsii.String("path"), jsii.String("to"), jsii.String("Dockerfile"), jsii.String("directory")))
modelData := sagemaker.ModelData_FromAsset(path.join(jsii.String("path"), jsii.String("to"), jsii.String("artifact"), jsii.String("file.tar.gz")))

model := sagemaker.NewModel(this, jsii.String("PrimaryContainerModel"), &ModelProps{
	Containers: []containerDefinition{
		&containerDefinition{
			Image: image,
			ModelData: modelData,
		},
	},
})
```

### Inference Pipeline Model

An inference pipeline is an Amazon SageMaker model that is composed of a linear sequence of multiple
containers that process requests for inferences on data. See the [AWS
documentation](https://docs.aws.amazon.com/sagemaker/latest/dg/inference-pipelines.html) to learn
more about SageMaker inference pipelines. To define an inference pipeline, you can provide
additional containers for your model:

```go
import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"

var image1 containerImage
var modelData1 modelData
var image2 containerImage
var modelData2 modelData
var image3 containerImage
var modelData3 modelData


model := sagemaker.NewModel(this, jsii.String("InferencePipelineModel"), &ModelProps{
	Containers: []containerDefinition{
		&containerDefinition{
			Image: image1,
			ModelData: modelData1,
		},
		&containerDefinition{
			Image: image2,
			ModelData: modelData2,
		},
		&containerDefinition{
			Image: image3,
			ModelData: modelData3,
		},
	},
})
```

### Container Images

Inference code can be stored in the Amazon EC2 Container Registry (Amazon ECR), which is specified
via `ContainerDefinition`'s `image` property which accepts a class that extends the `ContainerImage`
abstract base class.

#### Asset Image

Reference a local directory containing a Dockerfile:

```go
import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
import path "github.com/aws-samples/dummy/path"


image := sagemaker.ContainerImage_FromAsset(path.join(jsii.String("path"), jsii.String("to"), jsii.String("Dockerfile"), jsii.String("directory")))
```

#### ECR Image

Reference an image available within ECR:

```go
import ecr "github.com/aws/aws-cdk-go/awscdk"
import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"


repository := ecr.Repository_FromRepositoryName(this, jsii.String("Repository"), jsii.String("repo"))
image := sagemaker.ContainerImage_FromEcrRepository(repository, jsii.String("tag"))
```

#### DLC Image

Reference a deep learning container image:

```go
import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"


repositoryName := "huggingface-pytorch-training"
tag := "1.13.1-transformers4.26.0-gpu-py39-cu117-ubuntu20.04"

image := sagemaker.ContainerImage_FromDlc(repositoryName, tag)
```

### Model Artifacts

If you choose to decouple your model artifacts from your inference code (as is natural given
different rates of change between inference code and model artifacts), the artifacts can be
specified via the `modelData` property which accepts a class that extends the `ModelData` abstract
base class. The default is to have no model artifacts associated with a model.

#### Asset Model Data

Reference local model data:

```go
import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
import path "github.com/aws-samples/dummy/path"


modelData := sagemaker.ModelData_FromAsset(path.join(jsii.String("path"), jsii.String("to"), jsii.String("artifact"), jsii.String("file.tar.gz")))
```

#### S3 Model Data

Reference an S3 bucket and object key as the artifacts for a model:

```go
import s3 "github.com/aws/aws-cdk-go/awscdk"
import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"


bucket := s3.NewBucket(this, jsii.String("MyBucket"))
modelData := sagemaker.ModelData_FromBucket(bucket, jsii.String("path/to/artifact/file.tar.gz"))
```

## Model Hosting

Amazon SageMaker provides model hosting services for model deployment. Amazon SageMaker provides an
HTTPS endpoint where your machine learning model is available to provide inferences.

### Endpoint Configuration

By using the `EndpointConfig` construct, you can define a set of endpoint configuration which can be
used to provision one or more endpoints. In this configuration, you identify one or more models to
deploy and the resources that you want Amazon SageMaker to provision. You define one or more
production variants, each of which identifies a model. Each production variant also describes the
resources that you want Amazon SageMaker to provision. If you are hosting multiple models, you also
assign a variant weight to specify how much traffic you want to allocate to each model. For example,
suppose that you want to host two models, A and B, and you assign traffic weight 2 for model A and 1
for model B. Amazon SageMaker distributes two-thirds of the traffic to Model A, and one-third to
model B:

```go
import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"

var modelA model
var modelB model


endpointConfig := sagemaker.NewEndpointConfig(this, jsii.String("EndpointConfig"), &EndpointConfigProps{
	InstanceProductionVariants: []instanceProductionVariantProps{
		&instanceProductionVariantProps{
			Model: modelA,
			VariantName: jsii.String("modelA"),
			InitialVariantWeight: jsii.Number(2),
		},
		&instanceProductionVariantProps{
			Model: modelB,
			VariantName: jsii.String("variantB"),
			InitialVariantWeight: jsii.Number(1),
		},
	},
})
```

### Endpoint

When you create an endpoint from an `EndpointConfig`, Amazon SageMaker launches the ML compute
instances and deploys the model or models as specified in the configuration. To get inferences from
the model, client applications send requests to the Amazon SageMaker Runtime HTTPS endpoint. For
more information about the API, see the
[InvokeEndpoint](https://docs.aws.amazon.com/sagemaker/latest/dg/API_runtime_InvokeEndpoint.html)
API. Defining an endpoint requires at minimum the associated endpoint configuration:

```go
import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"

var endpointConfig endpointConfig


endpoint := sagemaker.NewEndpoint(this, jsii.String("Endpoint"), &EndpointProps{
	EndpointConfig: EndpointConfig,
})
```

### AutoScaling

To enable autoscaling on the production variant, use the `autoScaleInstanceCount` method:

```go
import "github.com/aws/aws-cdk-go/awscdksagemakeralpha"

var model model


variantName := "my-variant"
endpointConfig := sagemaker.NewEndpointConfig(this, jsii.String("EndpointConfig"), &EndpointConfigProps{
	InstanceProductionVariants: []instanceProductionVariantProps{
		&instanceProductionVariantProps{
			Model: model,
			VariantName: variantName,
		},
	},
})

endpoint := sagemaker.NewEndpoint(this, jsii.String("Endpoint"), &EndpointProps{
	EndpointConfig: EndpointConfig,
})
productionVariant := endpoint.FindInstanceProductionVariant(variantName)
instanceCount := productionVariant.AutoScaleInstanceCount(&EnableScalingProps{
	MaxCapacity: jsii.Number(3),
})
instanceCount.ScaleOnInvocations(jsii.String("LimitRPS"), &InvocationsScalingProps{
	MaxRequestsPerSecond: jsii.Number(30),
})
```

For load testing guidance on determining the maximum requests per second per instance, please see
this [documentation](https://docs.aws.amazon.com/sagemaker/latest/dg/endpoint-scaling-loadtest.html).

### Metrics

To monitor CloudWatch metrics for a production variant, use one or more of the metric convenience
methods:

```go
import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"

var endpointConfig endpointConfig


endpoint := sagemaker.NewEndpoint(this, jsii.String("Endpoint"), &EndpointProps{
	EndpointConfig: EndpointConfig,
})
productionVariant := endpoint.FindInstanceProductionVariant(jsii.String("my-variant"))
productionVariant.MetricModelLatency().CreateAlarm(this, jsii.String("ModelLatencyAlarm"), &CreateAlarmOptions{
	Threshold: jsii.Number(100000),
	EvaluationPeriods: jsii.Number(3),
})
```
