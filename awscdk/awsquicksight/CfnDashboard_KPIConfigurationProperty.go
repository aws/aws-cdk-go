package awsquicksight


// The configuration of a KPI visual.
//
// Example:
//
//
type CfnDashboard_KPIConfigurationProperty struct {
	// The field well configuration of a KPI visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The options that determine the presentation of a KPI visual.
	KpiOptions interface{} `field:"optional" json:"kpiOptions" yaml:"kpiOptions"`
	// The sort configuration of a KPI visual.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
}

