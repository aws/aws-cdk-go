package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Abstract base class for mixins that provides default implementations.
//
// Example:
//   type myEncryptionAtRest struct {
//   	Mixin
//   }
//
//   func newMyEncryptionAtRest(props myEncryptionAtRestProps) *myEncryptionAtRest {
//   	if props == nil {
//   		props = &myEncryptionAtRestProps{
//   		}
//   	}
//   	this := &myEncryptionAtRest{}
//   	newMixin_Override(this, )
//   	// Validate Mixin props at construction time
//   	if *props.bucketKey && *props.algorithm == "aws:kms:dsse" {throw new Error("Cannot use S3 Bucket Key and DSSE together");
//   	}
//   	return this
//   }
//
//   func (this *myEncryptionAtRest) supports(construct interface{}) *bool {
//   	return s3.CfnBucket_IsCfnBucket(*construct)
//   }
//
//   func (this *myEncryptionAtRest) applyTo(target CfnBucket) CfnBucket {
//   	// Validate pre-conditions on the target, throw if error is unrecoverable
//   	if !*target.BucketEncryption {throw new Error("Bucket encryption not configured");
//   	}
//
//   	// Validate properties are met after app execution
//   	*target.Node.AddValidation(map[string]validate{
//   		"validate": () => isKmsEncrypted(target)
//   		        ? ['This bucket must use aws:kms encryption.']
//   		        : [],
//   	})
//
//   	*target.BucketEncryption = &BucketEncryptionProperty{
//   		ServerSideEncryptionConfiguration: []interface{}{
//   			&ServerSideEncryptionRuleProperty{
//   				BucketKeyEnabled: jsii.Boolean(true),
//   				ServerSideEncryptionByDefault: &ServerSideEncryptionByDefaultProperty{
//   					SseAlgorithm: jsii.String("aws:kms"),
//   				},
//   			},
//   		},
//   	}
//   	return *target
//   }
//
type Mixin interface {
	constructs.IMixin
	// Applies the mixin functionality to the target construct.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for Mixin
type jsiiProxy_Mixin struct {
	internal.Type__constructsIMixin
}

func NewMixin_Override(m Mixin) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Mixin",
		nil, // no parameters
		m,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func Mixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Mixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Mixin) ApplyTo(construct constructs.IConstruct) {
	if err := m.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"applyTo",
		[]interface{}{construct},
	)
}

func (m *jsiiProxy_Mixin) Supports(construct constructs.IConstruct) *bool {
	if err := m.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		m,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

