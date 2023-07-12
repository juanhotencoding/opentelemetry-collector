// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package telemetry

import (
	"encoding/json"
	"fmt"
)

type Attributes struct {
	// ServiceName corresponds to the JSON schema field "service.name".
	ServiceName *string `mapstructure:"service.name,omitempty"`
}

type BatchLogRecordProcessor struct {
	// ExportTimeout corresponds to the JSON schema field "export_timeout".
	ExportTimeout *int `mapstructure:"export_timeout,omitempty"`

	// Exporter corresponds to the JSON schema field "exporter".
	Exporter *LogRecordExporter `mapstructure:"exporter,omitempty"`

	// MaxExportBatchSize corresponds to the JSON schema field
	// "max_export_batch_size".
	MaxExportBatchSize *int `mapstructure:"max_export_batch_size,omitempty"`

	// MaxQueueSize corresponds to the JSON schema field "max_queue_size".
	MaxQueueSize *int `mapstructure:"max_queue_size,omitempty"`

	// ScheduleDelay corresponds to the JSON schema field "schedule_delay".
	ScheduleDelay *int `mapstructure:"schedule_delay,omitempty"`
}

type BatchSpanProcessor struct {
	// ExportTimeout corresponds to the JSON schema field "export_timeout".
	ExportTimeout *int `mapstructure:"export_timeout,omitempty"`

	// Exporter corresponds to the JSON schema field "exporter".
	Exporter *SpanExporter `mapstructure:"exporter,omitempty"`

	// MaxExportBatchSize corresponds to the JSON schema field
	// "max_export_batch_size".
	MaxExportBatchSize *int `mapstructure:"max_export_batch_size,omitempty"`

	// MaxQueueSize corresponds to the JSON schema field "max_queue_size".
	MaxQueueSize *int `mapstructure:"max_queue_size,omitempty"`

	// ScheduleDelay corresponds to the JSON schema field "schedule_delay".
	ScheduleDelay *int `mapstructure:"schedule_delay,omitempty"`
}

type CommonJson map[string]interface{}

type Console map[string]interface{}

type Headers map[string]interface{}

type LogRecordExporter struct {
	// Otlp corresponds to the JSON schema field "otlp".
	Otlp *Otlp `mapstructure:"otlp,omitempty"`
}

type LogRecordLimits struct {
	// AttributeCountLimit corresponds to the JSON schema field
	// "attribute_count_limit".
	AttributeCountLimit *int `mapstructure:"attribute_count_limit,omitempty"`

	// AttributeValueLengthLimit corresponds to the JSON schema field
	// "attribute_value_length_limit".
	AttributeValueLengthLimit *int `mapstructure:"attribute_value_length_limit,omitempty"`
}

type LogRecordProcessor struct {
	// Batch corresponds to the JSON schema field "batch".
	Batch *BatchLogRecordProcessor `mapstructure:"batch,omitempty"`

	// Simple corresponds to the JSON schema field "simple".
	Simple *SimpleLogRecordProcessor `mapstructure:"simple,omitempty"`
}

type LoggerProviderJson struct {
	// Limits corresponds to the JSON schema field "limits".
	Limits *LogRecordLimits `mapstructure:"limits,omitempty"`

	// Processors corresponds to the JSON schema field "processors".
	Processors []LogRecordProcessor `mapstructure:"processors,omitempty"`
}

type MeterProviderJson struct {
	// Readers corresponds to the JSON schema field "readers".
	Readers []MetricReader `mapstructure:"readers,omitempty"`

	// Views corresponds to the JSON schema field "views".
	Views []View `mapstructure:"views,omitempty"`
}

type MetricExporter struct {
	// Console corresponds to the JSON schema field "console".
	Console Console `mapstructure:"console,omitempty"`

	// Otlp corresponds to the JSON schema field "otlp".
	Otlp *OtlpMetric `mapstructure:"otlp,omitempty"`

	// Prometheus corresponds to the JSON schema field "prometheus".
	Prometheus *Prometheus `mapstructure:"prometheus,omitempty"`
}

type MetricReader struct {
	// Periodic corresponds to the JSON schema field "periodic".
	Periodic *PeriodicMetricReader `mapstructure:"periodic,omitempty"`

	// Pull corresponds to the JSON schema field "pull".
	Pull *PullMetricReader `mapstructure:"pull,omitempty"`
}

type Otlp struct {
	// Certificate corresponds to the JSON schema field "certificate".
	Certificate *string `mapstructure:"certificate,omitempty"`

	// ClientCertificate corresponds to the JSON schema field "client_certificate".
	ClientCertificate *string `mapstructure:"client_certificate,omitempty"`

	// ClientKey corresponds to the JSON schema field "client_key".
	ClientKey *string `mapstructure:"client_key,omitempty"`

	// Compression corresponds to the JSON schema field "compression".
	Compression *string `mapstructure:"compression,omitempty"`

	// Endpoint corresponds to the JSON schema field "endpoint".
	Endpoint string `mapstructure:"endpoint"`

	// Headers corresponds to the JSON schema field "headers".
	Headers Headers `mapstructure:"headers,omitempty"`

	// Protocol corresponds to the JSON schema field "protocol".
	Protocol string `mapstructure:"protocol"`

	// Timeout corresponds to the JSON schema field "timeout".
	Timeout *int `mapstructure:"timeout,omitempty"`
}

type OtlpMetric struct {
	// Certificate corresponds to the JSON schema field "certificate".
	Certificate *string `mapstructure:"certificate,omitempty"`

	// ClientCertificate corresponds to the JSON schema field "client_certificate".
	ClientCertificate *string `mapstructure:"client_certificate,omitempty"`

	// ClientKey corresponds to the JSON schema field "client_key".
	ClientKey *string `mapstructure:"client_key,omitempty"`

	// Compression corresponds to the JSON schema field "compression".
	Compression *string `mapstructure:"compression,omitempty"`

	// DefaultHistogramAggregation corresponds to the JSON schema field
	// "default_histogram_aggregation".
	DefaultHistogramAggregation *string `mapstructure:"default_histogram_aggregation,omitempty"`

	// Endpoint corresponds to the JSON schema field "endpoint".
	Endpoint string `mapstructure:"endpoint"`

	// Headers corresponds to the JSON schema field "headers".
	Headers Headers `mapstructure:"headers,omitempty"`

	// Protocol corresponds to the JSON schema field "protocol".
	Protocol string `mapstructure:"protocol"`

	// TemporalityPreference corresponds to the JSON schema field
	// "temporality_preference".
	TemporalityPreference *string `mapstructure:"temporality_preference,omitempty"`

	// Timeout corresponds to the JSON schema field "timeout".
	Timeout *int `mapstructure:"timeout,omitempty"`
}

type PeriodicMetricReader struct {
	// Exporter corresponds to the JSON schema field "exporter".
	Exporter MetricExporter `mapstructure:"exporter"`

	// Interval corresponds to the JSON schema field "interval".
	Interval *int `mapstructure:"interval,omitempty"`

	// Timeout corresponds to the JSON schema field "timeout".
	Timeout *int `mapstructure:"timeout,omitempty"`
}

type Prometheus struct {
	// Host corresponds to the JSON schema field "host".
	Host *string `mapstructure:"host,omitempty"`

	// Port corresponds to the JSON schema field "port".
	Port *int `mapstructure:"port,omitempty"`
}

type PullMetricReader struct {
	// Exporter corresponds to the JSON schema field "exporter".
	Exporter MetricExporter `mapstructure:"exporter"`
}

type ResourceJson struct {
	// Attributes corresponds to the JSON schema field "attributes".
	Attributes *Attributes `mapstructure:"attributes,omitempty"`
}

type Sampler struct {
	// AlwaysOff corresponds to the JSON schema field "always_off".
	AlwaysOff SamplerAlwaysOff `mapstructure:"always_off,omitempty"`

	// AlwaysOn corresponds to the JSON schema field "always_on".
	AlwaysOn SamplerAlwaysOn `mapstructure:"always_on,omitempty"`

	// JaegerRemote corresponds to the JSON schema field "jaeger_remote".
	JaegerRemote *SamplerJaegerRemote `mapstructure:"jaeger_remote,omitempty"`

	// ParentBased corresponds to the JSON schema field "parent_based".
	ParentBased *SamplerParentBased `mapstructure:"parent_based,omitempty"`

	// TraceIdRatioBased corresponds to the JSON schema field "trace_id_ratio_based".
	TraceIdRatioBased *SamplerTraceIdRatioBased `mapstructure:"trace_id_ratio_based,omitempty"`
}

type SamplerAlwaysOff map[string]interface{}

type SamplerAlwaysOn map[string]interface{}

type SamplerJaegerRemote struct {
	// Endpoint corresponds to the JSON schema field "endpoint".
	Endpoint *string `mapstructure:"endpoint,omitempty"`

	// InitialSampler corresponds to the JSON schema field "initial_sampler".
	InitialSampler *Sampler `mapstructure:"initial_sampler,omitempty"`

	// Interval corresponds to the JSON schema field "interval".
	Interval *int `mapstructure:"interval,omitempty"`
}

type SamplerParentBased struct {
	// LocalParentNotSampled corresponds to the JSON schema field
	// "local_parent_not_sampled".
	LocalParentNotSampled *Sampler `mapstructure:"local_parent_not_sampled,omitempty"`

	// LocalParentSampled corresponds to the JSON schema field "local_parent_sampled".
	LocalParentSampled *Sampler `mapstructure:"local_parent_sampled,omitempty"`

	// RemoteParentNotSampled corresponds to the JSON schema field
	// "remote_parent_not_sampled".
	RemoteParentNotSampled *Sampler `mapstructure:"remote_parent_not_sampled,omitempty"`

	// RemoteParentSampled corresponds to the JSON schema field
	// "remote_parent_sampled".
	RemoteParentSampled *Sampler `mapstructure:"remote_parent_sampled,omitempty"`

	// Root corresponds to the JSON schema field "root".
	Root *Sampler `mapstructure:"root,omitempty"`
}

type SamplerTraceIdRatioBased struct {
	// Ratio corresponds to the JSON schema field "ratio".
	Ratio *float64 `mapstructure:"ratio,omitempty"`
}

type SimpleLogRecordProcessor struct {
	// Exporter corresponds to the JSON schema field "exporter".
	Exporter *LogRecordExporter `mapstructure:"exporter,omitempty"`
}

type SimpleSpanProcessor struct {
	// Exporter corresponds to the JSON schema field "exporter".
	Exporter *SpanExporter `mapstructure:"exporter,omitempty"`
}

type SpanExporter struct {
	// Console corresponds to the JSON schema field "console".
	Console Console `mapstructure:"console,omitempty"`

	// Otlp corresponds to the JSON schema field "otlp".
	Otlp *Otlp `mapstructure:"otlp,omitempty"`
}

type SpanLimits struct {
	// AttributeCountLimit corresponds to the JSON schema field
	// "attribute_count_limit".
	AttributeCountLimit *int `mapstructure:"attribute_count_limit,omitempty"`

	// AttributeValueLengthLimit corresponds to the JSON schema field
	// "attribute_value_length_limit".
	AttributeValueLengthLimit *int `mapstructure:"attribute_value_length_limit,omitempty"`

	// EventAttributeCountLimit corresponds to the JSON schema field
	// "event_attribute_count_limit".
	EventAttributeCountLimit *int `mapstructure:"event_attribute_count_limit,omitempty"`

	// EventCountLimit corresponds to the JSON schema field "event_count_limit".
	EventCountLimit *int `mapstructure:"event_count_limit,omitempty"`

	// LinkAttributeCountLimit corresponds to the JSON schema field
	// "link_attribute_count_limit".
	LinkAttributeCountLimit *int `mapstructure:"link_attribute_count_limit,omitempty"`

	// LinkCountLimit corresponds to the JSON schema field "link_count_limit".
	LinkCountLimit *int `mapstructure:"link_count_limit,omitempty"`
}

type SpanProcessor struct {
	// Batch corresponds to the JSON schema field "batch".
	Batch *BatchSpanProcessor `mapstructure:"batch,omitempty"`

	// Simple corresponds to the JSON schema field "simple".
	Simple *SimpleSpanProcessor `mapstructure:"simple,omitempty"`
}

type TracerProviderJson struct {
	// Limits corresponds to the JSON schema field "limits".
	Limits *SpanLimits `mapstructure:"limits,omitempty"`

	// Processors corresponds to the JSON schema field "processors".
	Processors []SpanProcessor `mapstructure:"processors,omitempty"`

	// Sampler corresponds to the JSON schema field "sampler".
	Sampler *Sampler `mapstructure:"sampler,omitempty"`
}

type View struct {
	// Selector corresponds to the JSON schema field "selector".
	Selector *ViewSelector `mapstructure:"selector,omitempty"`

	// Stream corresponds to the JSON schema field "stream".
	Stream *ViewStream `mapstructure:"stream,omitempty"`
}

type ViewSelector struct {
	// InstrumentName corresponds to the JSON schema field "instrument_name".
	InstrumentName *string `mapstructure:"instrument_name,omitempty"`

	// InstrumentType corresponds to the JSON schema field "instrument_type".
	InstrumentType *string `mapstructure:"instrument_type,omitempty"`

	// MeterName corresponds to the JSON schema field "meter_name".
	MeterName *string `mapstructure:"meter_name,omitempty"`

	// MeterSchemaUrl corresponds to the JSON schema field "meter_schema_url".
	MeterSchemaUrl *string `mapstructure:"meter_schema_url,omitempty"`

	// MeterVersion corresponds to the JSON schema field "meter_version".
	MeterVersion *string `mapstructure:"meter_version,omitempty"`
}

type ViewStream struct {
	// Aggregation corresponds to the JSON schema field "aggregation".
	Aggregation *ViewStreamAggregation `mapstructure:"aggregation,omitempty"`

	// AttributeKeys corresponds to the JSON schema field "attribute_keys".
	AttributeKeys []string `mapstructure:"attribute_keys,omitempty"`

	// Description corresponds to the JSON schema field "description".
	Description *string `mapstructure:"description,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name *string `mapstructure:"name,omitempty"`
}

type ViewStreamAggregation struct {
	// Default corresponds to the JSON schema field "default".
	Default interface{} `mapstructure:"default,omitempty"`

	// Drop corresponds to the JSON schema field "drop".
	Drop interface{} `mapstructure:"drop,omitempty"`

	// ExplicitBucketHistogram corresponds to the JSON schema field
	// "explicit_bucket_histogram".
	ExplicitBucketHistogram *ViewStreamAggregationExplicitBucketHistogram `mapstructure:"explicit_bucket_histogram,omitempty"`

	// ExponentialBucketHistogram corresponds to the JSON schema field
	// "exponential_bucket_histogram".
	ExponentialBucketHistogram *ViewStreamAggregationExponentialBucketHistogram `mapstructure:"exponential_bucket_histogram,omitempty"`

	// LastValue corresponds to the JSON schema field "last_value".
	LastValue interface{} `mapstructure:"last_value,omitempty"`

	// Sum corresponds to the JSON schema field "sum".
	Sum interface{} `mapstructure:"sum,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PullMetricReader) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["exporter"]; !ok || v == nil {
		return fmt.Errorf("field exporter in PullMetricReader: required")
	}
	type Plain PullMetricReader
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PullMetricReader(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PeriodicMetricReader) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["exporter"]; !ok || v == nil {
		return fmt.Errorf("field exporter in PeriodicMetricReader: required")
	}
	type Plain PeriodicMetricReader
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PeriodicMetricReader(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OtlpMetric) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["endpoint"]; !ok || v == nil {
		return fmt.Errorf("field endpoint in OtlpMetric: required")
	}
	if v, ok := raw["protocol"]; !ok || v == nil {
		return fmt.Errorf("field protocol in OtlpMetric: required")
	}
	type Plain OtlpMetric
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = OtlpMetric(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Otlp) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["endpoint"]; !ok || v == nil {
		return fmt.Errorf("field endpoint in Otlp: required")
	}
	if v, ok := raw["protocol"]; !ok || v == nil {
		return fmt.Errorf("field protocol in Otlp: required")
	}
	type Plain Otlp
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Otlp(plain)
	return nil
}

type ViewStreamAggregationExplicitBucketHistogram struct {
	// Boundaries corresponds to the JSON schema field "boundaries".
	Boundaries []float64 `mapstructure:"boundaries,omitempty"`

	// RecordMinMax corresponds to the JSON schema field "record_min_max".
	RecordMinMax *bool `mapstructure:"record_min_max,omitempty"`
}

type ViewStreamAggregationExponentialBucketHistogram struct {
	// MaxScale corresponds to the JSON schema field "max_scale".
	MaxScale *int `mapstructure:"max_scale,omitempty"`

	// MaxSize corresponds to the JSON schema field "max_size".
	MaxSize *int `mapstructure:"max_size,omitempty"`

	// RecordMinMax corresponds to the JSON schema field "record_min_max".
	RecordMinMax *bool `mapstructure:"record_min_max,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ViewStreamAggregation) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain ViewStreamAggregation
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if plain.Default != nil {
		return fmt.Errorf("field %s: must be null", "default")
	}
	if plain.Drop != nil {
		return fmt.Errorf("field %s: must be null", "drop")
	}
	if plain.LastValue != nil {
		return fmt.Errorf("field %s: must be null", "last_value")
	}
	if plain.Sum != nil {
		return fmt.Errorf("field %s: must be null", "sum")
	}
	*j = ViewStreamAggregation(plain)
	return nil
}
