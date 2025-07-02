package awsbedrock

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Bedrock provisioned model.
//
// Note: CloudFormation does not currently support creating Bedrock Provisioned Throughput
// resources outside of a custom resource. You can import provisioned models created by
// provisioning throughput in Bedrock outside the CDK or via a custom resource with
// {@link ProvisionedModel#fromProvisionedModelArn }.
//
// Example:
//   import bedrock "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bedrock.ProvisionedModel_FromProvisionedModelArn(this, jsii.String("Model"), jsii.String("arn:aws:bedrock:us-east-2:123456789012:provisioned-model/abc-123"))
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
//
type ProvisionedModel interface {
	IModel
	// The ARN of the provisioned model.
	ModelArn() *string
}

// The jsii proxy struct for ProvisionedModel
type jsiiProxy_ProvisionedModel struct {
	jsiiProxy_IModel
}

func (j *jsiiProxy_ProvisionedModel) ModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelArn",
		&returns,
	)
	return returns
}


// Import an provisioned model given an ARN.
func ProvisionedModel_FromProvisionedModelArn(_scope constructs.Construct, _id *string, provisionedModelArn *string) IModel {
	_init_.Initialize()

	if err := validateProvisionedModel_FromProvisionedModelArnParameters(_scope, _id, provisionedModelArn); err != nil {
		panic(err)
	}
	var returns IModel

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrock.ProvisionedModel",
		"fromProvisionedModelArn",
		[]interface{}{_scope, _id, provisionedModelArn},
		&returns,
	)

	return returns
}

