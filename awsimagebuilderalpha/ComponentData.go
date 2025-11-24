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
	// Indicates that the provided component data is an S3 reference.
	// Experimental.
	IsS3Reference() *bool
	// The resulting inline string or S3 URL which references the component data.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ComponentData
type jsiiProxy_ComponentData struct {
	_ byte // padding
}

func (j *jsiiProxy_ComponentData) IsS3Reference() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isS3Reference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComponentData) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewComponentData_Override(c ComponentData, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentData",
		[]interface{}{value},
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

