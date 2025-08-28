package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Cross-region inference enables you to seamlessly manage unplanned traffic bursts by utilizing compute across different AWS Regions.
//
// With cross-region
// inference, you can distribute traffic across multiple AWS Regions, enabling
// higher throughput and enhanced resilience during periods of peak demands.
//
// This construct represents a system-defined inference profile that routes
// requests across multiple regions based on availability and demand.
//
// Example:
//   // Create a cross-region inference profile
//   crossRegionProfile := bedrock.CrossRegionInferenceProfile_FromConfig(&CrossRegionInferenceProfileProps{
//   	GeoRegion: bedrock.CrossRegionInferenceProfileRegion_US,
//   	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V2_0(),
//   })
//
//   // Create an application inference profile across regions
//   appProfile := bedrock.NewApplicationInferenceProfile(this, jsii.String("MyMultiRegionProfile"), &ApplicationInferenceProfileProps{
//   	ApplicationInferenceProfileName: jsii.String("claude-35-sonnet-v2-multi-region"),
//   	ModelSource: crossRegionProfile,
//   	Description: jsii.String("Multi-region application profile for cost tracking"),
//   })
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html
//
// Experimental.
type CrossRegionInferenceProfile interface {
	IBedrockInvokable
	IInferenceProfile
	// The ARN of the inference profile.
	// Experimental.
	InferenceProfileArn() *string
	// The unique identifier of the inference profile.
	//
	// Format: {geoRegion}.{modelId}
	// Experimental.
	InferenceProfileId() *string
	// The underlying foundation model supporting cross-region inference.
	// Experimental.
	InferenceProfileModel() BedrockFoundationModel
	// The ARN used for invoking this inference profile.
	//
	// This equals to the inferenceProfileArn property, useful for implementing IBedrockInvokable interface.
	// Experimental.
	InvokableArn() *string
	// The type of inference profile.
	//
	// Always SYSTEM_DEFINED for cross-region profiles.
	// Experimental.
	Type() InferenceProfileType
	// Gives the appropriate policies to invoke and use the Foundation Model.
	//
	// For cross-region inference profiles, this method grants permissions to:
	// - Invoke the model in all regions where the inference profile can route requests
	// - Use the inference profile itself.
	//
	// Returns: An IAM Grant object representing the granted permissions.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Grants appropriate permissions to use the cross-region inference profile.
	//
	// This method adds the necessary IAM permissions to allow the grantee to:
	// - Get inference profile details (bedrock:GetInferenceProfile)
	// - Invoke the model through the inference profile (bedrock:InvokeModel*)
	//
	// Note: This does not grant permissions to use the underlying model directly.
	// For comprehensive permissions, use grantInvoke() instead.
	//
	// Returns: An IAM Grant object representing the granted permissions.
	// Experimental.
	GrantProfileUsage(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for CrossRegionInferenceProfile
type jsiiProxy_CrossRegionInferenceProfile struct {
	jsiiProxy_IBedrockInvokable
	jsiiProxy_IInferenceProfile
}

func (j *jsiiProxy_CrossRegionInferenceProfile) InferenceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossRegionInferenceProfile) InferenceProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossRegionInferenceProfile) InferenceProfileModel() BedrockFoundationModel {
	var returns BedrockFoundationModel
	_jsii_.Get(
		j,
		"inferenceProfileModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossRegionInferenceProfile) InvokableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossRegionInferenceProfile) Type() InferenceProfileType {
	var returns InferenceProfileType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates a Cross-Region Inference Profile from the provided configuration.
//
// Returns: A new CrossRegionInferenceProfile instance.
// Experimental.
func CrossRegionInferenceProfile_FromConfig(config *CrossRegionInferenceProfileProps) CrossRegionInferenceProfile {
	_init_.Initialize()

	if err := validateCrossRegionInferenceProfile_FromConfigParameters(config); err != nil {
		panic(err)
	}
	var returns CrossRegionInferenceProfile

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.CrossRegionInferenceProfile",
		"fromConfig",
		[]interface{}{config},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossRegionInferenceProfile) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossRegionInferenceProfile) GrantProfileUsage(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateGrantProfileUsageParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantProfileUsage",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

