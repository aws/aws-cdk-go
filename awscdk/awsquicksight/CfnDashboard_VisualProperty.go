package awsquicksight


// A visual displayed on a sheet in an analysis, dashboard, or template.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html
//
type CfnDashboard_VisualProperty struct {
	// A bar chart.
	//
	// For more information, see [Using bar charts](https://docs.aws.amazon.com/quicksight/latest/user/bar-charts.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-barchartvisual
	//
	BarChartVisual interface{} `field:"optional" json:"barChartVisual" yaml:"barChartVisual"`
	// A box plot.
	//
	// For more information, see [Using box plots](https://docs.aws.amazon.com/quicksight/latest/user/box-plots.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-boxplotvisual
	//
	BoxPlotVisual interface{} `field:"optional" json:"boxPlotVisual" yaml:"boxPlotVisual"`
	// A combo chart.
	//
	// For more information, see [Using combo charts](https://docs.aws.amazon.com/quicksight/latest/user/combo-charts.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-combochartvisual
	//
	ComboChartVisual interface{} `field:"optional" json:"comboChartVisual" yaml:"comboChartVisual"`
	// A visual that contains custom content.
	//
	// For more information, see [Using custom visual content](https://docs.aws.amazon.com/quicksight/latest/user/custom-visual-content.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-customcontentvisual
	//
	CustomContentVisual interface{} `field:"optional" json:"customContentVisual" yaml:"customContentVisual"`
	// An empty visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-emptyvisual
	//
	EmptyVisual interface{} `field:"optional" json:"emptyVisual" yaml:"emptyVisual"`
	// A filled map.
	//
	// For more information, see [Creating filled maps](https://docs.aws.amazon.com/quicksight/latest/user/filled-maps.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-filledmapvisual
	//
	FilledMapVisual interface{} `field:"optional" json:"filledMapVisual" yaml:"filledMapVisual"`
	// A funnel chart.
	//
	// For more information, see [Using funnel charts](https://docs.aws.amazon.com/quicksight/latest/user/funnel-visual-content.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-funnelchartvisual
	//
	FunnelChartVisual interface{} `field:"optional" json:"funnelChartVisual" yaml:"funnelChartVisual"`
	// A gauge chart.
	//
	// For more information, see [Using gauge charts](https://docs.aws.amazon.com/quicksight/latest/user/gauge-chart.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-gaugechartvisual
	//
	GaugeChartVisual interface{} `field:"optional" json:"gaugeChartVisual" yaml:"gaugeChartVisual"`
	// A geospatial map or a points on map visual.
	//
	// For more information, see [Creating point maps](https://docs.aws.amazon.com/quicksight/latest/user/point-maps.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-geospatialmapvisual
	//
	GeospatialMapVisual interface{} `field:"optional" json:"geospatialMapVisual" yaml:"geospatialMapVisual"`
	// A heat map.
	//
	// For more information, see [Using heat maps](https://docs.aws.amazon.com/quicksight/latest/user/heat-map.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-heatmapvisual
	//
	HeatMapVisual interface{} `field:"optional" json:"heatMapVisual" yaml:"heatMapVisual"`
	// A histogram.
	//
	// For more information, see [Using histograms](https://docs.aws.amazon.com/quicksight/latest/user/histogram-charts.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-histogramvisual
	//
	HistogramVisual interface{} `field:"optional" json:"histogramVisual" yaml:"histogramVisual"`
	// An insight visual.
	//
	// For more information, see [Working with insights](https://docs.aws.amazon.com/quicksight/latest/user/computational-insights.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-insightvisual
	//
	InsightVisual interface{} `field:"optional" json:"insightVisual" yaml:"insightVisual"`
	// A key performance indicator (KPI).
	//
	// For more information, see [Using KPIs](https://docs.aws.amazon.com/quicksight/latest/user/kpi.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-kpivisual
	//
	KpiVisual interface{} `field:"optional" json:"kpiVisual" yaml:"kpiVisual"`
	// A line chart.
	//
	// For more information, see [Using line charts](https://docs.aws.amazon.com/quicksight/latest/user/line-charts.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-linechartvisual
	//
	LineChartVisual interface{} `field:"optional" json:"lineChartVisual" yaml:"lineChartVisual"`
	// A pie or donut chart.
	//
	// For more information, see [Using pie charts](https://docs.aws.amazon.com/quicksight/latest/user/pie-chart.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-piechartvisual
	//
	PieChartVisual interface{} `field:"optional" json:"pieChartVisual" yaml:"pieChartVisual"`
	// A pivot table.
	//
	// For more information, see [Using pivot tables](https://docs.aws.amazon.com/quicksight/latest/user/pivot-table.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-pivottablevisual
	//
	PivotTableVisual interface{} `field:"optional" json:"pivotTableVisual" yaml:"pivotTableVisual"`
	// A radar chart visual.
	//
	// For more information, see [Using radar charts](https://docs.aws.amazon.com/quicksight/latest/user/radar-chart.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-radarchartvisual
	//
	RadarChartVisual interface{} `field:"optional" json:"radarChartVisual" yaml:"radarChartVisual"`
	// A sankey diagram.
	//
	// For more information, see [Using Sankey diagrams](https://docs.aws.amazon.com/quicksight/latest/user/sankey-diagram.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-sankeydiagramvisual
	//
	SankeyDiagramVisual interface{} `field:"optional" json:"sankeyDiagramVisual" yaml:"sankeyDiagramVisual"`
	// A scatter plot.
	//
	// For more information, see [Using scatter plots](https://docs.aws.amazon.com/quicksight/latest/user/scatter-plot.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-scatterplotvisual
	//
	ScatterPlotVisual interface{} `field:"optional" json:"scatterPlotVisual" yaml:"scatterPlotVisual"`
	// A table visual.
	//
	// For more information, see [Using tables as visuals](https://docs.aws.amazon.com/quicksight/latest/user/tabular.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-tablevisual
	//
	TableVisual interface{} `field:"optional" json:"tableVisual" yaml:"tableVisual"`
	// A tree map.
	//
	// For more information, see [Using tree maps](https://docs.aws.amazon.com/quicksight/latest/user/tree-map.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-treemapvisual
	//
	TreeMapVisual interface{} `field:"optional" json:"treeMapVisual" yaml:"treeMapVisual"`
	// A waterfall chart.
	//
	// For more information, see [Using waterfall charts](https://docs.aws.amazon.com/quicksight/latest/user/waterfall-chart.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-waterfallvisual
	//
	WaterfallVisual interface{} `field:"optional" json:"waterfallVisual" yaml:"waterfallVisual"`
	// A word cloud.
	//
	// For more information, see [Using word clouds](https://docs.aws.amazon.com/quicksight/latest/user/word-cloud.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-wordcloudvisual
	//
	WordCloudVisual interface{} `field:"optional" json:"wordCloudVisual" yaml:"wordCloudVisual"`
}

