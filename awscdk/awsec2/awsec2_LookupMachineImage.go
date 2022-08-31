package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// A machine image whose AMI ID will be searched using DescribeImages.
//
// The most recent, available, launchable image matching the given filter
// criteria will be used. Looking up AMIs may take a long time; specify
// as many filter criteria as possible to narrow down the search.
//
// The AMI selected will be cached in `cdk.context.json` and the same value
// will be used on future runs. To refresh the AMI lookup, you will have to
// evict the value from the cache using the `cdk context` command. See
// https://docs.aws.amazon.com/cdk/latest/guide/context.html for more information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userData userData
//
//   lookupMachineImage := awscdk.Aws_ec2.NewLookupMachineImage(&lookupMachineImageProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	filters: map[string][]*string{
//   		"filtersKey": []*string{
//   			jsii.String("filters"),
//   		},
//   	},
//   	owners: []*string{
//   		jsii.String("owners"),
//   	},
//   	userData: userData,
//   	windows: jsii.Boolean(false),
//   })
//
// Experimental.
type LookupMachineImage interface {
	IMachineImage
	// Return the image to use in the given context.
	// Experimental.
	GetImage(scope awscdk.Construct) *MachineImageConfig
}

// The jsii proxy struct for LookupMachineImage
type jsiiProxy_LookupMachineImage struct {
	jsiiProxy_IMachineImage
}

// Experimental.
func NewLookupMachineImage(props *LookupMachineImageProps) LookupMachineImage {
	_init_.Initialize()

	j := jsiiProxy_LookupMachineImage{}

	_jsii_.Create(
		"monocdk.aws_ec2.LookupMachineImage",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewLookupMachineImage_Override(l LookupMachineImage, props *LookupMachineImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.LookupMachineImage",
		[]interface{}{props},
		l,
	)
}

func (l *jsiiProxy_LookupMachineImage) GetImage(scope awscdk.Construct) *MachineImageConfig {
	var returns *MachineImageConfig

	_jsii_.Invoke(
		l,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

