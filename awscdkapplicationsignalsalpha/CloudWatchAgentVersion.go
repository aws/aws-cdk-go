package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Provides version information and image selection for CloudWatch Agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   cloudWatchAgentVersion := applicationsignals_alpha.NewCloudWatchAgentVersion()
//
// Experimental.
type CloudWatchAgentVersion interface {
}

// The jsii proxy struct for CloudWatchAgentVersion
type jsiiProxy_CloudWatchAgentVersion struct {
	_ byte // padding
}

// Experimental.
func NewCloudWatchAgentVersion() CloudWatchAgentVersion {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchAgentVersion{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentVersion",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchAgentVersion_Override(c CloudWatchAgentVersion) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentVersion",
		nil, // no parameters
		c,
	)
}

// Gets the appropriate CloudWatch Agent image based on the operating system.
//
// Returns: The CloudWatch Agent image URI.
// Experimental.
func CloudWatchAgentVersion_GetCloudWatchAgentImage(operatingSystemFamily awsecs.OperatingSystemFamily) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentVersion",
		"getCloudWatchAgentImage",
		[]interface{}{operatingSystemFamily},
		&returns,
	)

	return returns
}

func CloudWatchAgentVersion_CLOUDWATCH_AGENT_IMAGE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentVersion",
		"CLOUDWATCH_AGENT_IMAGE",
		&returns,
	)
	return returns
}

func CloudWatchAgentVersion_CLOUDWATCH_AGENT_IMAGE_WIN2019() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentVersion",
		"CLOUDWATCH_AGENT_IMAGE_WIN2019",
		&returns,
	)
	return returns
}

func CloudWatchAgentVersion_CLOUDWATCH_AGENT_IMAGE_WIN2022() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.CloudWatchAgentVersion",
		"CLOUDWATCH_AGENT_IMAGE_WIN2022",
		&returns,
	)
	return returns
}

