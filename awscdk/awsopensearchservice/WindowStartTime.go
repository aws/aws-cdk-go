package awsopensearchservice


// Example:
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_3(),
//   	OffPeakWindowEnabled: jsii.Boolean(true),
//   	 // can be omitted if offPeakWindowStart is set
//   	OffPeakWindowStart: &WindowStartTime{
//   		Hours: jsii.Number(20),
//   		Minutes: jsii.Number(0),
//   	},
//   })
//
type WindowStartTime struct {
	// The start hour of the window in Coordinated Universal Time (UTC), using 24-hour time.
	//
	// For example, 17 refers to 5:00 P.M. UTC.
	// Default: - 22.
	//
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// The start minute of the window, in UTC.
	// Default: - 0.
	//
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

