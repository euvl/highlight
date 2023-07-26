package clickhouse

import (
	"fmt"
	"time"

	"go.opentelemetry.io/collector/pdata/ptrace"
)

type TraceRow struct {
	ProjectId          uint32
	SecureSessionId    string
	Timestamp          time.Time
	TraceId            string
	SpanId             string
	ParentSpanId       string
	TraceState         string
	SpanName           string
	SpanKind           string
	ServiceName        string
	ResourceAttributes map[string]string
	SpanAttributes     map[string]string
	Duration           int64
	StatusCode         string
	StatusMessage      string
	Events             []Event
	Links              []Link
}

type Event struct {
	Timestamp  time.Time         `json:"timestamp"`
	Name       string            `json:"name"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type Link struct {
	TraceId    string            `json:"trace_id"`
	SpanId     string            `json:"span_id"`
	TraceState string            `json:"trace_state"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

// TODO: See if we can use types from extracted.go. They can't be imported right
// now for some reason :thinking:
type Fields struct {
	projectIDInt uint32
}

func NewTraceRow(span ptrace.Span, fields Fields) *TraceRow {
	attributes := attributesToMap(span.Attributes().AsRaw())

	var traceEvents []Event
	spanEvents := span.Events()
	for i := 0; i < spanEvents.Len(); i++ {
		event := spanEvents.At(i)
		traceEvents = append(traceEvents, Event{
			Timestamp:  event.Timestamp().AsTime(),
			Name:       event.Name(),
			Attributes: attributesToMap(event.Attributes().AsRaw()),
		})
	}

	var traceLinks []Link
	spanLinks := span.Links()
	for l := 0; l < spanLinks.Len(); l++ {
		link := spanLinks.At(l)
		traceLinks = append(traceLinks, Link{
			TraceId:    link.TraceID().String(),
			SpanId:     link.SpanID().String(),
			TraceState: link.TraceState().AsRaw(),
			Attributes: attributesToMap(link.Attributes().AsRaw()),
		})
	}

	traceRow := &TraceRow{
		ProjectId:          fields.projectIDInt,
		SecureSessionId:    attributes["highlight.session_id"],
		Timestamp:          span.StartTimestamp().AsTime(),
		TraceId:            span.TraceID().String(),
		SpanId:             span.SpanID().String(),
		ParentSpanId:       span.ParentSpanID().String(),
		TraceState:         span.TraceState().AsRaw(),
		SpanName:           span.Name(),
		SpanKind:           span.Kind().String(),
		Duration:           int64(span.EndTimestamp().AsTime().Sub(span.StartTimestamp().AsTime()).Nanoseconds()),
		ServiceName:        "",
		ResourceAttributes: attributes,
		SpanAttributes:     attributes,
		StatusCode:         span.Status().Code().String(),
		StatusMessage:      span.Status().Message(),
		Events:             traceEvents,
		Links:              traceLinks,
	}

	return traceRow
}

func attributesToMap(attributes map[string]any) map[string]string {
	newAttrMap := make(map[string]string)
	for k, v := range attributes {
		newAttrMap[k] = fmt.Sprintf("%v", v)
	}
	return newAttrMap
}
