package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// Fargate profiles allows an administrator to declare which pods run on Fargate.
//
// This declaration is done through the profile’s selectors. Each
// profile can have up to five selectors that contain a namespace and optional
// labels. You must define a namespace for every selector. The label field
// consists of multiple optional key-value pairs. Pods that match a selector (by
// matching a namespace for the selector and all of the labels specified in the
// selector) are scheduled on Fargate. If a namespace selector is defined
// without any labels, Amazon EKS will attempt to schedule all pods that run in
// that namespace onto Fargate using the profile. If a to-be-scheduled pod
// matches any of the selectors in the Fargate profile, then that pod is
// scheduled on Fargate.
//
// If a pod matches multiple Fargate profiles, Amazon EKS picks one of the
// matches at random. In this case, you can specify which profile a pod should
// use by adding the following Kubernetes label to the pod specification:
// eks.amazonaws.com/fargate-profile: profile_name. However, the pod must still
// match a selector in that profile in order to be scheduled onto Fargate.
//
// Example:
//   var cluster cluster
//
//   eks.NewFargateProfile(this, jsii.String("MyProfile"), &fargateProfileProps{
//   	cluster: cluster,
//   	selectors: []selector{
//   		&selector{
//   			namespace: jsii.String("default"),
//   		},
//   	},
//   })
//
// Experimental.
type FargateProfile interface {
	awscdk.Construct
	awscdk.ITaggable
	// The full Amazon Resource Name (ARN) of the Fargate profile.
	// Experimental.
	FargateProfileArn() *string
	// The name of the Fargate profile.
	// Experimental.
	FargateProfileName() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The pod execution role to use for pods that match the selectors in the Fargate profile.
	//
	// The pod execution role allows Fargate infrastructure to
	// register with your cluster as a node, and it provides read access to Amazon
	// ECR image repositories.
	// Experimental.
	PodExecutionRole() awsiam.IRole
	// Resource tags.
	// Experimental.
	Tags() awscdk.TagManager
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for FargateProfile
type jsiiProxy_FargateProfile struct {
	internal.Type__awscdkConstruct
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_FargateProfile) FargateProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fargateProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateProfile) FargateProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fargateProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateProfile) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateProfile) PodExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"podExecutionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateProfile) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}


// Experimental.
func NewFargateProfile(scope constructs.Construct, id *string, props *FargateProfileProps) FargateProfile {
	_init_.Initialize()

	j := jsiiProxy_FargateProfile{}

	_jsii_.Create(
		"monocdk.aws_eks.FargateProfile",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFargateProfile_Override(f FargateProfile, scope constructs.Construct, id *string, props *FargateProfileProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.FargateProfile",
		[]interface{}{scope, id, props},
		f,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func FargateProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.FargateProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateProfile) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FargateProfile) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FargateProfile) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateProfile) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FargateProfile) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FargateProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateProfile) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

