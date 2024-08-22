package awscdksagemakeralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdksagemakeralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Model data represents the source of model artifacts, which will ultimately be loaded from an S3 location.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var image containerImage
//   var modelData modelData
//
//
//   model := sagemaker.NewModel(this, jsii.String("ContainerModel"), &ModelProps{
//   	Containers: []containerDefinition{
//   		&containerDefinition{
//   			Image: *Image,
//   			ModelData: *ModelData,
//   		},
//   	},
//   	NetworkIsolation: jsii.Boolean(true),
//   })
//
// Experimental.
type ModelData interface {
	// This method is invoked by the SageMaker Model construct when it needs to resolve the model data to a URI.
	// Experimental.
	Bind(scope constructs.Construct, model IModel) *ModelDataConfig
}

// The jsii proxy struct for ModelData
type jsiiProxy_ModelData struct {
	_ byte // padding
}

// Experimental.
func NewModelData_Override(m ModelData) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-sagemaker-alpha.ModelData",
		nil, // no parameters
		m,
	)
}

// Constructs model data that will be uploaded to S3 as part of the CDK app deployment.
// Experimental.
func ModelData_FromAsset(path *string, options *awss3assets.AssetOptions) ModelData {
	_init_.Initialize()

	if err := validateModelData_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns ModelData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.ModelData",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Constructs model data which is already available within S3.
// Experimental.
func ModelData_FromBucket(bucket awss3.IBucket, objectKey *string) ModelData {
	_init_.Initialize()

	if err := validateModelData_FromBucketParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns ModelData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.ModelData",
		"fromBucket",
		[]interface{}{bucket, objectKey},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelData) Bind(scope constructs.Construct, model IModel) *ModelDataConfig {
	if err := m.validateBindParameters(scope, model); err != nil {
		panic(err)
	}
	var returns *ModelDataConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{scope, model},
		&returns,
	)

	return returns
}

