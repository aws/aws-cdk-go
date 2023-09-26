package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Canonical names of the IANA time zones, derived from the IANA Time Zone Database.
//
// For more information, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.
//
// Example:
//   var target lambdaInvoke
//
//
//   rateBasedSchedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
//   	Target: Target,
//   	Description: jsii.String("This is a test rate-based schedule"),
//   })
//
//   cronBasedSchedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Cron(&CronOptionsWithTimezone{
//   		Minute: jsii.String("0"),
//   		Hour: jsii.String("23"),
//   		Day: jsii.String("20"),
//   		Month: jsii.String("11"),
//   		TimeZone: awscdk.TimeZone_AMERICA_NEW_YORK(),
//   	}),
//   	Target: Target,
//   	Description: jsii.String("This is a test cron-based schedule that will run at 11:00 PM, on day 20 of the month, only in November in New York timezone"),
//   })
//
type TimeZone interface {
	// The name of the timezone.
	TimezoneName() *string
}

// The jsii proxy struct for TimeZone
type jsiiProxy_TimeZone struct {
	_ byte // padding
}

func (j *jsiiProxy_TimeZone) TimezoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneName",
		&returns,
	)
	return returns
}


// Use this to add a timezone not in this class.
//
// Returns: a new Timezone.
func TimeZone_Of(timezoneName *string) TimeZone {
	_init_.Initialize()

	if err := validateTimeZone_OfParameters(timezoneName); err != nil {
		panic(err)
	}
	var returns TimeZone

	_jsii_.StaticInvoke(
		"aws-cdk-lib.TimeZone",
		"of",
		[]interface{}{timezoneName},
		&returns,
	)

	return returns
}

func TimeZone_AFRICA_ABIDJAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_ABIDJAN",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_ALGIERS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_ALGIERS",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_BISSAU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_BISSAU",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_CAIRO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_CAIRO",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_CASABLANCA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_CASABLANCA",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_CEUTA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_CEUTA",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_EL_AAIUN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_EL_AAIUN",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_JOHANNESBURG() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_JOHANNESBURG",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_JUBA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_JUBA",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_KHARTOUM() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_KHARTOUM",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_LAGOS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_LAGOS",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_MAPUTO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_MAPUTO",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_MONROVIA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_MONROVIA",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_NAIROBI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_NAIROBI",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_NDJAMENA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_NDJAMENA",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_SAO_TOME() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_SAO_TOME",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_TRIPOLI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_TRIPOLI",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_TUNIS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_TUNIS",
		&returns,
	)
	return returns
}

func TimeZone_AFRICA_WINDHOEK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AFRICA_WINDHOEK",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ADAK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ADAK",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ANCHORAGE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ANCHORAGE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARAGUAINA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARAGUAINA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_BUENOS_AIRES() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_BUENOS_AIRES",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_CATAMARCA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_CATAMARCA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_CORDOBA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_CORDOBA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_JUJUY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_JUJUY",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_LA_RIOJA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_LA_RIOJA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_MENDOZA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_MENDOZA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_RIO_GALLEGOS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_RIO_GALLEGOS",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_SALTA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_SALTA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_SAN_JUAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_SAN_JUAN",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_SAN_LUIS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_SAN_LUIS",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_TUCUMAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_TUCUMAN",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ARGENTINA_USHUAIA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ARGENTINA_USHUAIA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ASUNCION() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ASUNCION",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_BAHIA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_BAHIA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_BAHIA_BANDERAS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_BAHIA_BANDERAS",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_BARBADOS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_BARBADOS",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_BELEM() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_BELEM",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_BELIZE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_BELIZE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_BOA_VISTA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_BOA_VISTA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_BOGOTA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_BOGOTA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_BOISE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_BOISE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_CAMBRIDGE_BAY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_CAMBRIDGE_BAY",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_CAMPO_GRANDE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_CAMPO_GRANDE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_CANCUN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_CANCUN",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_CARACAS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_CARACAS",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_CAYENNE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_CAYENNE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_CHICAGO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_CHICAGO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_CHIHUAHUA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_CHIHUAHUA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_CIUDAD_JUAREZ() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_CIUDAD_JUAREZ",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_COSTA_RICA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_COSTA_RICA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_CUIABA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_CUIABA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_DANMARKSHAVN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_DANMARKSHAVN",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_DAWSON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_DAWSON",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_DAWSON_CREEK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_DAWSON_CREEK",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_DENVER() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_DENVER",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_DETROIT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_DETROIT",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_EDMONTON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_EDMONTON",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_EIRUNEPE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_EIRUNEPE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_EL_SALVADOR() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_EL_SALVADOR",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_FORT_NELSON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_FORT_NELSON",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_FORTALEZA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_FORTALEZA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_GLACE_BAY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_GLACE_BAY",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_GOOSE_BAY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_GOOSE_BAY",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_GRAND_TURK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_GRAND_TURK",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_GUATEMALA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_GUATEMALA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_GUAYAQUIL() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_GUAYAQUIL",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_GUYANA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_GUYANA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_HALIFAX() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_HALIFAX",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_HAVANA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_HAVANA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_HERMOSILLO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_HERMOSILLO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_INDIANA_INDIANAPOLIS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_INDIANA_INDIANAPOLIS",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_INDIANA_KNOX() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_INDIANA_KNOX",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_INDIANA_MARENGO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_INDIANA_MARENGO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_INDIANA_PETERSBURG() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_INDIANA_PETERSBURG",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_INDIANA_TELL_CITY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_INDIANA_TELL_CITY",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_INDIANA_VEVAY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_INDIANA_VEVAY",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_INDIANA_VINCENNES() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_INDIANA_VINCENNES",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_INDIANA_WINAMAC() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_INDIANA_WINAMAC",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_INUVIK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_INUVIK",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_IQALUIT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_IQALUIT",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_JAMAICA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_JAMAICA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_JUNEAU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_JUNEAU",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_KENTUCKY_LOUISVILLE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_KENTUCKY_LOUISVILLE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_KENTUCKY_MONTICELLO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_KENTUCKY_MONTICELLO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_LA_PAZ() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_LA_PAZ",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_LIMA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_LIMA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_LOS_ANGELES() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_LOS_ANGELES",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MACEIO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MACEIO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MANAGUA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MANAGUA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MANAUS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MANAUS",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MARTINIQUE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MARTINIQUE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MATAMOROS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MATAMOROS",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MAZATLAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MAZATLAN",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MENOMINEE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MENOMINEE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MERIDA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MERIDA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_METLAKATLA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_METLAKATLA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MEXICO_CITY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MEXICO_CITY",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MIQUELON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MIQUELON",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MONCTON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MONCTON",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MONTERREY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MONTERREY",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_MONTEVIDEO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_MONTEVIDEO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_NEW_YORK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_NEW_YORK",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_NOME() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_NOME",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_NORONHA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_NORONHA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_NORTH_DAKOTA_BEULAH() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_NORTH_DAKOTA_BEULAH",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_NORTH_DAKOTA_CENTER() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_NORTH_DAKOTA_CENTER",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_NORTH_DAKOTA_NEW_SALEM() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_NORTH_DAKOTA_NEW_SALEM",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_NUUK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_NUUK",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_OJINAGA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_OJINAGA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_PANAMA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_PANAMA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_PARAMARIBO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_PARAMARIBO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_PHOENIX() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_PHOENIX",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_PORT_MINUS_AU_MINUS_PRINCE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_PORT_MINUS_AU_MINUS_PRINCE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_PORTO_VELHO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_PORTO_VELHO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_PUERTO_RICO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_PUERTO_RICO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_PUNTA_ARENAS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_PUNTA_ARENAS",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_RANKIN_INLET() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_RANKIN_INLET",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_RECIFE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_RECIFE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_REGINA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_REGINA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_RESOLUTE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_RESOLUTE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_RIO_BRANCO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_RIO_BRANCO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_SANTAREM() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_SANTAREM",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_SANTIAGO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_SANTIAGO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_SANTO_DOMINGO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_SANTO_DOMINGO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_SAO_PAULO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_SAO_PAULO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_SCORESBYSUND() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_SCORESBYSUND",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_SITKA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_SITKA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_ST_JOHNS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_ST_JOHNS",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_SWIFT_CURRENT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_SWIFT_CURRENT",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_TEGUCIGALPA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_TEGUCIGALPA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_THULE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_THULE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_TIJUANA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_TIJUANA",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_TORONTO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_TORONTO",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_VANCOUVER() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_VANCOUVER",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_WHITEHORSE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_WHITEHORSE",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_WINNIPEG() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_WINNIPEG",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_YAKUTAT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_YAKUTAT",
		&returns,
	)
	return returns
}

func TimeZone_AMERICA_YELLOWKNIFE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AMERICA_YELLOWKNIFE",
		&returns,
	)
	return returns
}

func TimeZone_ANTARCTICA_CASEY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ANTARCTICA_CASEY",
		&returns,
	)
	return returns
}

func TimeZone_ANTARCTICA_DAVIS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ANTARCTICA_DAVIS",
		&returns,
	)
	return returns
}

func TimeZone_ANTARCTICA_MACQUARIE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ANTARCTICA_MACQUARIE",
		&returns,
	)
	return returns
}

func TimeZone_ANTARCTICA_MAWSON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ANTARCTICA_MAWSON",
		&returns,
	)
	return returns
}

func TimeZone_ANTARCTICA_PALMER() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ANTARCTICA_PALMER",
		&returns,
	)
	return returns
}

func TimeZone_ANTARCTICA_ROTHERA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ANTARCTICA_ROTHERA",
		&returns,
	)
	return returns
}

func TimeZone_ANTARCTICA_TROLL() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ANTARCTICA_TROLL",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_ALMATY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_ALMATY",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_AMMAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_AMMAN",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_ANADYR() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_ANADYR",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_AQTAU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_AQTAU",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_AQTOBE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_AQTOBE",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_ASHGABAT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_ASHGABAT",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_ATYRAU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_ATYRAU",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_BAGHDAD() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_BAGHDAD",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_BAKU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_BAKU",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_BANGKOK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_BANGKOK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_BARNAUL() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_BARNAUL",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_BEIRUT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_BEIRUT",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_BISHKEK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_BISHKEK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_CHITA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_CHITA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_CHOIBALSAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_CHOIBALSAN",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_COLOMBO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_COLOMBO",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_DAMASCUS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_DAMASCUS",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_DHAKA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_DHAKA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_DILI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_DILI",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_DUBAI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_DUBAI",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_DUSHANBE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_DUSHANBE",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_FAMAGUSTA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_FAMAGUSTA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_GAZA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_GAZA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_HEBRON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_HEBRON",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_HO_CHI_MINH() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_HO_CHI_MINH",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_HONG_KONG() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_HONG_KONG",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_HOVD() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_HOVD",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_IRKUTSK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_IRKUTSK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_JAKARTA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_JAKARTA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_JAYAPURA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_JAYAPURA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_JERUSALEM() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_JERUSALEM",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_KABUL() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_KABUL",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_KAMCHATKA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_KAMCHATKA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_KARACHI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_KARACHI",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_KATHMANDU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_KATHMANDU",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_KHANDYGA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_KHANDYGA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_KOLKATA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_KOLKATA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_KRASNOYARSK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_KRASNOYARSK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_KUCHING() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_KUCHING",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_MACAU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_MACAU",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_MAGADAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_MAGADAN",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_MAKASSAR() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_MAKASSAR",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_MANILA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_MANILA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_NICOSIA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_NICOSIA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_NOVOKUZNETSK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_NOVOKUZNETSK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_NOVOSIBIRSK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_NOVOSIBIRSK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_OMSK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_OMSK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_ORAL() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_ORAL",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_PONTIANAK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_PONTIANAK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_PYONGYANG() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_PYONGYANG",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_QATAR() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_QATAR",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_QOSTANAY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_QOSTANAY",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_QYZYLORDA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_QYZYLORDA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_RIYADH() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_RIYADH",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_SAKHALIN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_SAKHALIN",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_SAMARKAND() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_SAMARKAND",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_SEOUL() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_SEOUL",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_SHANGHAI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_SHANGHAI",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_SINGAPORE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_SINGAPORE",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_SREDNEKOLYMSK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_SREDNEKOLYMSK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_TAIPEI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_TAIPEI",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_TASHKENT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_TASHKENT",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_TBILISI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_TBILISI",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_TEHRAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_TEHRAN",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_THIMPHU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_THIMPHU",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_TOKYO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_TOKYO",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_TOMSK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_TOMSK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_ULAANBAATAR() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_ULAANBAATAR",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_URUMQI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_URUMQI",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_UST_MINUS_NERA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_UST_MINUS_NERA",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_VLADIVOSTOK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_VLADIVOSTOK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_YAKUTSK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_YAKUTSK",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_YANGON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_YANGON",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_YEKATERINBURG() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_YEKATERINBURG",
		&returns,
	)
	return returns
}

func TimeZone_ASIA_YEREVAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ASIA_YEREVAN",
		&returns,
	)
	return returns
}

func TimeZone_ATLANTIC_AZORES() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ATLANTIC_AZORES",
		&returns,
	)
	return returns
}

func TimeZone_ATLANTIC_BERMUDA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ATLANTIC_BERMUDA",
		&returns,
	)
	return returns
}

func TimeZone_ATLANTIC_CANARY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ATLANTIC_CANARY",
		&returns,
	)
	return returns
}

func TimeZone_ATLANTIC_CAPE_VERDE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ATLANTIC_CAPE_VERDE",
		&returns,
	)
	return returns
}

func TimeZone_ATLANTIC_FAROE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ATLANTIC_FAROE",
		&returns,
	)
	return returns
}

func TimeZone_ATLANTIC_MADEIRA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ATLANTIC_MADEIRA",
		&returns,
	)
	return returns
}

func TimeZone_ATLANTIC_SOUTH_GEORGIA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ATLANTIC_SOUTH_GEORGIA",
		&returns,
	)
	return returns
}

func TimeZone_ATLANTIC_STANLEY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ATLANTIC_STANLEY",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_ADELAIDE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_ADELAIDE",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_BRISBANE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_BRISBANE",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_BROKEN_HILL() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_BROKEN_HILL",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_DARWIN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_DARWIN",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_EUCLA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_EUCLA",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_HOBART() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_HOBART",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_LINDEMAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_LINDEMAN",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_LORD_HOWE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_LORD_HOWE",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_MELBOURNE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_MELBOURNE",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_PERTH() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_PERTH",
		&returns,
	)
	return returns
}

func TimeZone_AUSTRALIA_SYDNEY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"AUSTRALIA_SYDNEY",
		&returns,
	)
	return returns
}

func TimeZone_CET() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"CET",
		&returns,
	)
	return returns
}

func TimeZone_CST6CDT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"CST6CDT",
		&returns,
	)
	return returns
}

func TimeZone_EET() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EET",
		&returns,
	)
	return returns
}

func TimeZone_EST() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EST",
		&returns,
	)
	return returns
}

func TimeZone_EST5EDT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EST5EDT",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_1() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_1",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_10() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_10",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_11() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_11",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_12() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_12",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_13() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_13",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_14() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_14",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_2() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_2",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_3() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_3",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_4() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_4",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_5() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_5",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_6() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_6",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_7() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_7",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_8() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_8",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_MINUS_9() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_MINUS_9",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_1() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_1",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_10() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_10",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_11() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_11",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_12() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_12",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_2() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_2",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_3() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_3",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_4() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_4",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_5() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_5",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_6() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_6",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_7() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_7",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_8() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_8",
		&returns,
	)
	return returns
}

func TimeZone_ETC_GMT_PLUS_9() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_GMT_PLUS_9",
		&returns,
	)
	return returns
}

func TimeZone_ETC_UTC() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"ETC_UTC",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_ANDORRA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_ANDORRA",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_ASTRAKHAN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_ASTRAKHAN",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_ATHENS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_ATHENS",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_BELGRADE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_BELGRADE",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_BERLIN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_BERLIN",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_BRUSSELS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_BRUSSELS",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_BUCHAREST() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_BUCHAREST",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_BUDAPEST() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_BUDAPEST",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_CHISINAU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_CHISINAU",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_DUBLIN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_DUBLIN",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_GIBRALTAR() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_GIBRALTAR",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_HELSINKI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_HELSINKI",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_ISTANBUL() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_ISTANBUL",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_KALININGRAD() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_KALININGRAD",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_KIROV() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_KIROV",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_KYIV() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_KYIV",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_LISBON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_LISBON",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_LONDON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_LONDON",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_MADRID() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_MADRID",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_MALTA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_MALTA",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_MINSK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_MINSK",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_MOSCOW() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_MOSCOW",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_PARIS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_PARIS",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_PRAGUE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_PRAGUE",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_RIGA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_RIGA",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_ROME() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_ROME",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_SAMARA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_SAMARA",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_SARATOV() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_SARATOV",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_SIMFEROPOL() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_SIMFEROPOL",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_SOFIA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_SOFIA",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_TALLINN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_TALLINN",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_TIRANE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_TIRANE",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_ULYANOVSK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_ULYANOVSK",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_VIENNA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_VIENNA",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_VILNIUS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_VILNIUS",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_VOLGOGRAD() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_VOLGOGRAD",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_WARSAW() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_WARSAW",
		&returns,
	)
	return returns
}

func TimeZone_EUROPE_ZURICH() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"EUROPE_ZURICH",
		&returns,
	)
	return returns
}

func TimeZone_FACTORY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"FACTORY",
		&returns,
	)
	return returns
}

func TimeZone_HST() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"HST",
		&returns,
	)
	return returns
}

func TimeZone_INDIAN_CHAGOS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"INDIAN_CHAGOS",
		&returns,
	)
	return returns
}

func TimeZone_INDIAN_MALDIVES() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"INDIAN_MALDIVES",
		&returns,
	)
	return returns
}

func TimeZone_INDIAN_MAURITIUS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"INDIAN_MAURITIUS",
		&returns,
	)
	return returns
}

func TimeZone_MET() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"MET",
		&returns,
	)
	return returns
}

func TimeZone_MST() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"MST",
		&returns,
	)
	return returns
}

func TimeZone_MST7MDT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"MST7MDT",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_APIA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_APIA",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_AUCKLAND() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_AUCKLAND",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_BOUGAINVILLE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_BOUGAINVILLE",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_CHATHAM() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_CHATHAM",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_EASTER() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_EASTER",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_EFATE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_EFATE",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_FAKAOFO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_FAKAOFO",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_FIJI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_FIJI",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_GALAPAGOS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_GALAPAGOS",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_GAMBIER() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_GAMBIER",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_GUADALCANAL() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_GUADALCANAL",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_GUAM() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_GUAM",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_HONOLULU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_HONOLULU",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_KANTON() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_KANTON",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_KIRITIMATI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_KIRITIMATI",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_KOSRAE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_KOSRAE",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_KWAJALEIN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_KWAJALEIN",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_MARQUESAS() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_MARQUESAS",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_NAURU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_NAURU",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_NIUE() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_NIUE",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_NORFOLK() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_NORFOLK",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_NOUMEA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_NOUMEA",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_PAGO_PAGO() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_PAGO_PAGO",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_PALAU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_PALAU",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_PITCAIRN() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_PITCAIRN",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_PORT_MORESBY() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_PORT_MORESBY",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_RAROTONGA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_RAROTONGA",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_TAHITI() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_TAHITI",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_TARAWA() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_TARAWA",
		&returns,
	)
	return returns
}

func TimeZone_PACIFIC_TONGATAPU() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PACIFIC_TONGATAPU",
		&returns,
	)
	return returns
}

func TimeZone_PST8PDT() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"PST8PDT",
		&returns,
	)
	return returns
}

func TimeZone_WET() TimeZone {
	_init_.Initialize()
	var returns TimeZone
	_jsii_.StaticGet(
		"aws-cdk-lib.TimeZone",
		"WET",
		&returns,
	)
	return returns
}

