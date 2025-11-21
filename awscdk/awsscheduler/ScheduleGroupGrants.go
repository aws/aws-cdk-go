package awsscheduler

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsscheduler"
)

// Collection of grant methods for a IScheduleGroupRef.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var scheduleGroupRef IScheduleGroupRef
//
//   scheduleGroupGrants := awscdk.Aws_scheduler.ScheduleGroupGrants_FromScheduleGroup(scheduleGroupRef)
//
type ScheduleGroupGrants interface {
	Resource() interfacesawsscheduler.IScheduleGroupRef
	// Grant delete schedule permission for schedules in this group to the given principal.
	DeleteSchedules(grantee awsiam.IGrantable) awsiam.Grant
	// Grant list and get schedule permissions for schedules in this group to the given principal.
	ReadSchedules(grantee awsiam.IGrantable) awsiam.Grant
	// Grant create and update schedule permissions for schedules in this group to the given principal.
	WriteSchedules(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for ScheduleGroupGrants
type jsiiProxy_ScheduleGroupGrants struct {
	_ byte // padding
}

func (j *jsiiProxy_ScheduleGroupGrants) Resource() interfacesawsscheduler.IScheduleGroupRef {
	var returns interfacesawsscheduler.IScheduleGroupRef
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


// Creates grants for ScheduleGroupGrants.
func ScheduleGroupGrants_FromScheduleGroup(resource interfacesawsscheduler.IScheduleGroupRef) ScheduleGroupGrants {
	_init_.Initialize()

	if err := validateScheduleGroupGrants_FromScheduleGroupParameters(resource); err != nil {
		panic(err)
	}
	var returns ScheduleGroupGrants

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_scheduler.ScheduleGroupGrants",
		"fromScheduleGroup",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleGroupGrants) DeleteSchedules(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateDeleteSchedulesParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"deleteSchedules",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleGroupGrants) ReadSchedules(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateReadSchedulesParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"readSchedules",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleGroupGrants) WriteSchedules(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateWriteSchedulesParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"writeSchedules",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

