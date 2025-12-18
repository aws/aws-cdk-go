package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class for referencing and uploading component data.
//
// Example:
//   component := imagebuilder.NewComponent(this, jsii.String("EncryptedComponent"), &ComponentProps{
//   	Platform: imagebuilder.Platform_LINUX,
//   	KmsKey: kms.NewKey(this, jsii.String("ComponentKey")),
//   	Data: imagebuilder.ComponentData_FromJsonObject(map[string]interface{}{
//   		"schemaVersion": imagebuilder.ComponentSchemaVersion_V1_0,
//   		"phases": []interface{}{
//   			map[string]interface{}{
//   				"name": imagebuilder.ComponentPhaseName_BUILD,
//   				"steps": []map[string]interface{}{
//   					map[string]interface{}{
//   						"name": jsii.String("secure-setup"),
//   						"action": imagebuilder.ComponentAction_EXECUTE_BASH,
//   						"inputs": map[string][]*string{
//   							"commands": []*string{
//   								jsii.String("echo \"This component data is encrypted with KMS\""),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type ComponentData interface {
	// The rendered component data value, for use in CloudFormation.
	//
	// - For inline components, data is the component text
	// - For S3-backed components, uri is the S3 URL.
	// Experimental.
	Render() *ComponentDataConfig
}

// The jsii proxy struct for ComponentData
type jsiiProxy_ComponentData struct {
	_ byte // padding
}

// Experimental.
func NewComponentData_Override(c ComponentData) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentData",
		nil, // no parameters
		c,
	)
}

// Uploads component data from a local file to S3 to use as the component data.
// Experimental.
func ComponentData_FromAsset(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) S3ComponentData {
	_init_.Initialize()

	if err := validateComponentData_FromAssetParameters(scope, id, path, options); err != nil {
		panic(err)
	}
	var returns S3ComponentData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentData",
		"fromAsset",
		[]interface{}{scope, id, path, options},
		&returns,
	)

	return returns
}

// Uses an inline JSON object as the component data, using the ComponentDocument interface.
// Experimental.
func ComponentData_FromComponentDocumentJsonObject(data *ComponentDocument) ComponentData {
	_init_.Initialize()

	if err := validateComponentData_FromComponentDocumentJsonObjectParameters(data); err != nil {
		panic(err)
	}
	var returns ComponentData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentData",
		"fromComponentDocumentJsonObject",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// Uses an inline JSON/YAML string as the component data.
// Experimental.
func ComponentData_FromInline(data *string) ComponentData {
	_init_.Initialize()

	if err := validateComponentData_FromInlineParameters(data); err != nil {
		panic(err)
	}
	var returns ComponentData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentData",
		"fromInline",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// Uses an inline JSON object as the component data.
// Experimental.
func ComponentData_FromJsonObject(data *map[string]interface{}) ComponentData {
	_init_.Initialize()

	if err := validateComponentData_FromJsonObjectParameters(data); err != nil {
		panic(err)
	}
	var returns ComponentData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentData",
		"fromJsonObject",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// References component data from a pre-existing S3 object.
// Experimental.
func ComponentData_FromS3(bucket awss3.IBucket, key *string) S3ComponentData {
	_init_.Initialize()

	if err := validateComponentData_FromS3Parameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3ComponentData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentData",
		"fromS3",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComponentData) Render() *ComponentDataConfig {
	var returns *ComponentDataConfig

	_jsii_.Invoke(
		c,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

