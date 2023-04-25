// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsSettings provides settings for githubreceiver metrics.
type MetricsSettings struct {
	GhRepoBranchCommitsCount MetricSettings `mapstructure:"gh.repo.branch.commits.count"`
	GhRepoBranchesCount      MetricSettings `mapstructure:"gh.repo.branches.count"`
	GhRepoContributorsCount  MetricSettings `mapstructure:"gh.repo.contributors.count"`
	GhRepoCount              MetricSettings `mapstructure:"gh.repo.count"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		GhRepoBranchCommitsCount: MetricSettings{
			Enabled: true,
		},
		GhRepoBranchesCount: MetricSettings{
			Enabled: true,
		},
		GhRepoContributorsCount: MetricSettings{
			Enabled: true,
		},
		GhRepoCount: MetricSettings{
			Enabled: true,
		},
	}
}

// ResourceAttributeSettings provides common settings for a particular metric.
type ResourceAttributeSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesSettings provides settings for githubreceiver metrics.
type ResourceAttributesSettings struct {
}

func DefaultResourceAttributesSettings() ResourceAttributesSettings {
	return ResourceAttributesSettings{}
}

type metricGhRepoBranchCommitsCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills gh.repo.branch.commits.count metric with initial data.
func (m *metricGhRepoBranchCommitsCount) init() {
	m.data.SetName("gh.repo.branch.commits.count")
	m.data.SetDescription("The total number of commits on a given branch")
	m.data.SetUnit("1")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricGhRepoBranchCommitsCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, ghRepoNameAttributeValue string, ghOrgAttributeValue string, ghRepoBranchNameAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("gh.repo.name", ghRepoNameAttributeValue)
	dp.Attributes().PutStr("gh.org", ghOrgAttributeValue)
	dp.Attributes().PutStr("gh.repo.branch.name", ghRepoBranchNameAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricGhRepoBranchCommitsCount) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricGhRepoBranchCommitsCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricGhRepoBranchCommitsCount(settings MetricSettings) metricGhRepoBranchCommitsCount {
	m := metricGhRepoBranchCommitsCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricGhRepoBranchesCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills gh.repo.branches.count metric with initial data.
func (m *metricGhRepoBranchesCount) init() {
	m.data.SetName("gh.repo.branches.count")
	m.data.SetDescription("Number of branches that exist in the repository")
	m.data.SetUnit("ratio")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricGhRepoBranchesCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, ghRepoNameAttributeValue string, ghOrgAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("gh.repo.name", ghRepoNameAttributeValue)
	dp.Attributes().PutStr("gh.org", ghOrgAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricGhRepoBranchesCount) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricGhRepoBranchesCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricGhRepoBranchesCount(settings MetricSettings) metricGhRepoBranchesCount {
	m := metricGhRepoBranchesCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricGhRepoContributorsCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills gh.repo.contributors.count metric with initial data.
func (m *metricGhRepoContributorsCount) init() {
	m.data.SetName("gh.repo.contributors.count")
	m.data.SetDescription("Total number of unique contributors to this repository")
	m.data.SetUnit("ratio")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricGhRepoContributorsCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, ghRepoNameAttributeValue string, ghOrgAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("gh.repo.name", ghRepoNameAttributeValue)
	dp.Attributes().PutStr("gh.org", ghOrgAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricGhRepoContributorsCount) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricGhRepoContributorsCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricGhRepoContributorsCount(settings MetricSettings) metricGhRepoContributorsCount {
	m := metricGhRepoContributorsCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricGhRepoCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills gh.repo.count metric with initial data.
func (m *metricGhRepoCount) init() {
	m.data.SetName("gh.repo.count")
	m.data.SetDescription("Number of repositories that exist in an organization")
	m.data.SetUnit("ratio")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricGhRepoCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, ghOrgAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("gh.org", ghOrgAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricGhRepoCount) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricGhRepoCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricGhRepoCount(settings MetricSettings) metricGhRepoCount {
	m := metricGhRepoCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilderConfig is a structural subset of an otherwise 1-1 copy of metadata.yaml
type MetricsBuilderConfig struct {
	Metrics            MetricsSettings            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesSettings `mapstructure:"resource_attributes"`
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                      pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                int                 // maximum observed number of metrics per resource.
	resourceCapacity               int                 // maximum observed number of resource attributes.
	metricsBuffer                  pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                      component.BuildInfo // contains version information
	resourceAttributesSettings     ResourceAttributesSettings
	metricGhRepoBranchCommitsCount metricGhRepoBranchCommitsCount
	metricGhRepoBranchesCount      metricGhRepoBranchesCount
	metricGhRepoContributorsCount  metricGhRepoContributorsCount
	metricGhRepoCount              metricGhRepoCount
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsSettings(),
		ResourceAttributes: DefaultResourceAttributesSettings(),
	}
}

func NewMetricsBuilderConfig(ms MetricsSettings, ras ResourceAttributesSettings) MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            ms,
		ResourceAttributes: ras,
	}
}

func NewMetricsBuilder(mbc MetricsBuilderConfig, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                      pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                  pmetric.NewMetrics(),
		buildInfo:                      settings.BuildInfo,
		resourceAttributesSettings:     mbc.ResourceAttributes,
		metricGhRepoBranchCommitsCount: newMetricGhRepoBranchCommitsCount(mbc.Metrics.GhRepoBranchCommitsCount),
		metricGhRepoBranchesCount:      newMetricGhRepoBranchesCount(mbc.Metrics.GhRepoBranchesCount),
		metricGhRepoContributorsCount:  newMetricGhRepoContributorsCount(mbc.Metrics.GhRepoContributorsCount),
		metricGhRepoCount:              newMetricGhRepoCount(mbc.Metrics.GhRepoCount),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(ResourceAttributesSettings, pmetric.ResourceMetrics)

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(ras ResourceAttributesSettings, rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/githubreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricGhRepoBranchCommitsCount.emit(ils.Metrics())
	mb.metricGhRepoBranchesCount.emit(ils.Metrics())
	mb.metricGhRepoContributorsCount.emit(ils.Metrics())
	mb.metricGhRepoCount.emit(ils.Metrics())

	for _, op := range rmo {
		op(mb.resourceAttributesSettings, rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := mb.metricsBuffer
	mb.metricsBuffer = pmetric.NewMetrics()
	return metrics
}

// RecordGhRepoBranchCommitsCountDataPoint adds a data point to gh.repo.branch.commits.count metric.
func (mb *MetricsBuilder) RecordGhRepoBranchCommitsCountDataPoint(ts pcommon.Timestamp, val int64, ghRepoNameAttributeValue string, ghOrgAttributeValue string, ghRepoBranchNameAttributeValue string) {
	mb.metricGhRepoBranchCommitsCount.recordDataPoint(mb.startTime, ts, val, ghRepoNameAttributeValue, ghOrgAttributeValue, ghRepoBranchNameAttributeValue)
}

// RecordGhRepoBranchesCountDataPoint adds a data point to gh.repo.branches.count metric.
func (mb *MetricsBuilder) RecordGhRepoBranchesCountDataPoint(ts pcommon.Timestamp, val int64, ghRepoNameAttributeValue string, ghOrgAttributeValue string) {
	mb.metricGhRepoBranchesCount.recordDataPoint(mb.startTime, ts, val, ghRepoNameAttributeValue, ghOrgAttributeValue)
}

// RecordGhRepoContributorsCountDataPoint adds a data point to gh.repo.contributors.count metric.
func (mb *MetricsBuilder) RecordGhRepoContributorsCountDataPoint(ts pcommon.Timestamp, val int64, ghRepoNameAttributeValue string, ghOrgAttributeValue string) {
	mb.metricGhRepoContributorsCount.recordDataPoint(mb.startTime, ts, val, ghRepoNameAttributeValue, ghOrgAttributeValue)
}

// RecordGhRepoCountDataPoint adds a data point to gh.repo.count metric.
func (mb *MetricsBuilder) RecordGhRepoCountDataPoint(ts pcommon.Timestamp, val int64, ghOrgAttributeValue string) {
	mb.metricGhRepoCount.recordDataPoint(mb.startTime, ts, val, ghOrgAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}