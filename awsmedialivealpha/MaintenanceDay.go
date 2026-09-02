package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Day of the week for maintenance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   maintenanceDay := medialive_alpha.MaintenanceDay_FRIDAY()
//
// Experimental.
type MaintenanceDay interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MaintenanceDay
type jsiiProxy_MaintenanceDay struct {
	_ byte // padding
}

func (j *jsiiProxy_MaintenanceDay) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func MaintenanceDay_Of(value *string) MaintenanceDay {
	_init_.Initialize()

	if err := validateMaintenanceDay_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MaintenanceDay

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MaintenanceDay",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MaintenanceDay_FRIDAY() MaintenanceDay {
	_init_.Initialize()
	var returns MaintenanceDay
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MaintenanceDay",
		"FRIDAY",
		&returns,
	)
	return returns
}

func MaintenanceDay_MONDAY() MaintenanceDay {
	_init_.Initialize()
	var returns MaintenanceDay
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MaintenanceDay",
		"MONDAY",
		&returns,
	)
	return returns
}

func MaintenanceDay_SATURDAY() MaintenanceDay {
	_init_.Initialize()
	var returns MaintenanceDay
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MaintenanceDay",
		"SATURDAY",
		&returns,
	)
	return returns
}

func MaintenanceDay_SUNDAY() MaintenanceDay {
	_init_.Initialize()
	var returns MaintenanceDay
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MaintenanceDay",
		"SUNDAY",
		&returns,
	)
	return returns
}

func MaintenanceDay_THURSDAY() MaintenanceDay {
	_init_.Initialize()
	var returns MaintenanceDay
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MaintenanceDay",
		"THURSDAY",
		&returns,
	)
	return returns
}

func MaintenanceDay_TUESDAY() MaintenanceDay {
	_init_.Initialize()
	var returns MaintenanceDay
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MaintenanceDay",
		"TUESDAY",
		&returns,
	)
	return returns
}

func MaintenanceDay_WEDNESDAY() MaintenanceDay {
	_init_.Initialize()
	var returns MaintenanceDay
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MaintenanceDay",
		"WEDNESDAY",
		&returns,
	)
	return returns
}

