package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class for referencing and uploading workflow data.
//
// Example:
//   workflow := imagebuilder.NewWorkflow(this, jsii.String("EncryptedWorkflow"), &WorkflowProps{
//   	WorkflowType: imagebuilder.WorkflowType_BUILD,
//   	KmsKey: kms.NewKey(this, jsii.String("WorkflowKey")),
//   	Data: imagebuilder.WorkflowData_FromJsonObject(map[string]interface{}{
//   		"schemaVersion": imagebuilder.WorkflowSchemaVersion_V1_0,
//   		"steps": []interface{}{
//   			map[string]interface{}{
//   				"name": jsii.String("LaunchBuildInstance"),
//   				"action": imagebuilder.WorkflowAction_LAUNCH_INSTANCE,
//   				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
//   				"inputs": map[string]*string{
//   					"waitFor": jsii.String("ssmAgent"),
//   				},
//   			},
//   			map[string]interface{}{
//   				"name": jsii.String("CreateImage"),
//   				"action": imagebuilder.WorkflowAction_CREATE_IMAGE,
//   				"onFailure": imagebuilder.WorkflowOnFailure_ABORT,
//   				"inputs": map[string]*string{
//   					"instanceId": jsii.String("i-123"),
//   				},
//   			},
//   			map[string]interface{}{
//   				"name": jsii.String("TerminateInstance"),
//   				"action": imagebuilder.WorkflowAction_TERMINATE_INSTANCE,
//   				"onFailure": imagebuilder.WorkflowOnFailure_CONTINUE,
//   				"inputs": map[string]*string{
//   					"instanceId": jsii.String("i-123"),
//   				},
//   			},
//   		},
//   		"outputs": []interface{}{
//   			map[string]*string{
//   				"name": jsii.String("ImageId"),
//   				"value": jsii.String("$.stepOutputs.CreateImage.imageId"),
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type WorkflowData interface {
	// The rendered workflow data value, for use in CloudFormation.
	//
	// - For inline workflows, data is the workflow text
	// - For S3-backed workflows, uri is the S3 URL.
	// Experimental.
	Render() *WorkflowDataConfig
}

// The jsii proxy struct for WorkflowData
type jsiiProxy_WorkflowData struct {
	_ byte // padding
}

// Experimental.
func NewWorkflowData_Override(w WorkflowData) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowData",
		nil, // no parameters
		w,
	)
}

// Uploads workflow data from a local file to S3 to use as the workflow data.
// Experimental.
func WorkflowData_FromAsset(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) S3WorkflowData {
	_init_.Initialize()

	if err := validateWorkflowData_FromAssetParameters(scope, id, path, options); err != nil {
		panic(err)
	}
	var returns S3WorkflowData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowData",
		"fromAsset",
		[]interface{}{scope, id, path, options},
		&returns,
	)

	return returns
}

// Uses an inline JSON or YAML string as the workflow data.
// Experimental.
func WorkflowData_FromInline(data *string) WorkflowData {
	_init_.Initialize()

	if err := validateWorkflowData_FromInlineParameters(data); err != nil {
		panic(err)
	}
	var returns WorkflowData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowData",
		"fromInline",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// Uses an inline JSON object as the workflow data.
// Experimental.
func WorkflowData_FromJsonObject(data *map[string]interface{}) WorkflowData {
	_init_.Initialize()

	if err := validateWorkflowData_FromJsonObjectParameters(data); err != nil {
		panic(err)
	}
	var returns WorkflowData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowData",
		"fromJsonObject",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// References workflow data from a pre-existing S3 object.
// Experimental.
func WorkflowData_FromS3(bucket awss3.IBucket, key *string) S3WorkflowData {
	_init_.Initialize()

	if err := validateWorkflowData_FromS3Parameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3WorkflowData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowData",
		"fromS3",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkflowData) Render() *WorkflowDataConfig {
	var returns *WorkflowDataConfig

	_jsii_.Invoke(
		w,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

