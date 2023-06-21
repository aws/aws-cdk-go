package awsquicksight


// The configuration of a word cloud visual.
//
// Example:
//
//
type CfnTemplate_WordCloudChartConfigurationProperty struct {
	// The label options (label text, label visibility, and sort icon visibility) for the word cloud category.
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The sort configuration of a word cloud visual.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The options for a word cloud visual.
	WordCloudOptions interface{} `field:"optional" json:"wordCloudOptions" yaml:"wordCloudOptions"`
}

