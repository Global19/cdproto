package debugger

import (
	"errors"

	"github.com/knq/chromedp/cdp/runtime"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// AUTOGENERATED. DO NOT EDIT.

// BreakpointID breakpoint identifier.
type BreakpointID string

// String returns the BreakpointID as string value.
func (t BreakpointID) String() string {
	return string(t)
}

// CallFrameID call frame identifier.
type CallFrameID string

// String returns the CallFrameID as string value.
func (t CallFrameID) String() string {
	return string(t)
}

// Location location in the source code.
type Location struct {
	ScriptID     runtime.ScriptID `json:"scriptId"`               // Script identifier as reported in the Debugger.scriptParsed.
	LineNumber   int64            `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber int64            `json:"columnNumber,omitempty"` // Column number in the script (0-based).
}

// ScriptPosition location in the source code.
type ScriptPosition struct {
	LineNumber   int64 `json:"lineNumber"`
	ColumnNumber int64 `json:"columnNumber"`
}

// CallFrame javaScript call frame. Array of call frames form the call stack.
type CallFrame struct {
	CallFrameID      CallFrameID           `json:"callFrameId"`                // Call frame identifier. This identifier is only valid while the virtual machine is paused.
	FunctionName     string                `json:"functionName"`               // Name of the JavaScript function called on this call frame.
	FunctionLocation *Location             `json:"functionLocation,omitempty"` // Location in the source code.
	Location         *Location             `json:"location"`                   // Location in the source code.
	ScopeChain       []*Scope              `json:"scopeChain"`                 // Scope chain for this call frame.
	This             *runtime.RemoteObject `json:"this"`                       // this object for this call frame.
	ReturnValue      *runtime.RemoteObject `json:"returnValue,omitempty"`      // The value being returned, if the function is at return point.
}

// Scope scope description.
type Scope struct {
	Type          ScopeType             `json:"type"`   // Scope type.
	Object        *runtime.RemoteObject `json:"object"` // Object representing the scope. For global and with scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
	Name          string                `json:"name,omitempty"`
	StartLocation *Location             `json:"startLocation,omitempty"` // Location in the source code where scope starts
	EndLocation   *Location             `json:"endLocation,omitempty"`   // Location in the source code where scope ends
}

// SearchMatch search match for resource.
type SearchMatch struct {
	LineNumber  float64 `json:"lineNumber"`  // Line number in resource content.
	LineContent string  `json:"lineContent"` // Line with match content.
}

// BreakLocation [no description].
type BreakLocation struct {
	ScriptID     runtime.ScriptID  `json:"scriptId"`               // Script identifier as reported in the Debugger.scriptParsed.
	LineNumber   int64             `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber int64             `json:"columnNumber,omitempty"` // Column number in the script (0-based).
	Type         BreakLocationType `json:"type,omitempty"`
}

// ScopeType scope type.
type ScopeType string

// String returns the ScopeType as string value.
func (t ScopeType) String() string {
	return string(t)
}

// ScopeType values.
const (
	ScopeTypeGlobal  ScopeType = "global"
	ScopeTypeLocal   ScopeType = "local"
	ScopeTypeWith    ScopeType = "with"
	ScopeTypeClosure ScopeType = "closure"
	ScopeTypeCatch   ScopeType = "catch"
	ScopeTypeBlock   ScopeType = "block"
	ScopeTypeScript  ScopeType = "script"
	ScopeTypeEval    ScopeType = "eval"
	ScopeTypeModule  ScopeType = "module"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ScopeType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ScopeType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ScopeType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ScopeType(in.String()) {
	case ScopeTypeGlobal:
		*t = ScopeTypeGlobal
	case ScopeTypeLocal:
		*t = ScopeTypeLocal
	case ScopeTypeWith:
		*t = ScopeTypeWith
	case ScopeTypeClosure:
		*t = ScopeTypeClosure
	case ScopeTypeCatch:
		*t = ScopeTypeCatch
	case ScopeTypeBlock:
		*t = ScopeTypeBlock
	case ScopeTypeScript:
		*t = ScopeTypeScript
	case ScopeTypeEval:
		*t = ScopeTypeEval
	case ScopeTypeModule:
		*t = ScopeTypeModule

	default:
		in.AddError(errors.New("unknown ScopeType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ScopeType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// BreakLocationType [no description].
type BreakLocationType string

// String returns the BreakLocationType as string value.
func (t BreakLocationType) String() string {
	return string(t)
}

// BreakLocationType values.
const (
	BreakLocationTypeDebuggerStatement BreakLocationType = "debuggerStatement"
	BreakLocationTypeCall              BreakLocationType = "call"
	BreakLocationTypeReturn            BreakLocationType = "return"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t BreakLocationType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t BreakLocationType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *BreakLocationType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch BreakLocationType(in.String()) {
	case BreakLocationTypeDebuggerStatement:
		*t = BreakLocationTypeDebuggerStatement
	case BreakLocationTypeCall:
		*t = BreakLocationTypeCall
	case BreakLocationTypeReturn:
		*t = BreakLocationTypeReturn

	default:
		in.AddError(errors.New("unknown BreakLocationType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *BreakLocationType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// PausedReason pause reason.
type PausedReason string

// String returns the PausedReason as string value.
func (t PausedReason) String() string {
	return string(t)
}

// PausedReason values.
const (
	PausedReasonXHR              PausedReason = "XHR"
	PausedReasonDOM              PausedReason = "DOM"
	PausedReasonEventListener    PausedReason = "EventListener"
	PausedReasonException        PausedReason = "exception"
	PausedReasonAssert           PausedReason = "assert"
	PausedReasonDebugCommand     PausedReason = "debugCommand"
	PausedReasonPromiseRejection PausedReason = "promiseRejection"
	PausedReasonOOM              PausedReason = "OOM"
	PausedReasonOther            PausedReason = "other"
	PausedReasonAmbiguous        PausedReason = "ambiguous"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PausedReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PausedReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PausedReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch PausedReason(in.String()) {
	case PausedReasonXHR:
		*t = PausedReasonXHR
	case PausedReasonDOM:
		*t = PausedReasonDOM
	case PausedReasonEventListener:
		*t = PausedReasonEventListener
	case PausedReasonException:
		*t = PausedReasonException
	case PausedReasonAssert:
		*t = PausedReasonAssert
	case PausedReasonDebugCommand:
		*t = PausedReasonDebugCommand
	case PausedReasonPromiseRejection:
		*t = PausedReasonPromiseRejection
	case PausedReasonOOM:
		*t = PausedReasonOOM
	case PausedReasonOther:
		*t = PausedReasonOther
	case PausedReasonAmbiguous:
		*t = PausedReasonAmbiguous

	default:
		in.AddError(errors.New("unknown PausedReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PausedReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ContinueToLocationTargetCallFrames [no description].
type ContinueToLocationTargetCallFrames string

// String returns the ContinueToLocationTargetCallFrames as string value.
func (t ContinueToLocationTargetCallFrames) String() string {
	return string(t)
}

// ContinueToLocationTargetCallFrames values.
const (
	ContinueToLocationTargetCallFramesAny     ContinueToLocationTargetCallFrames = "any"
	ContinueToLocationTargetCallFramesCurrent ContinueToLocationTargetCallFrames = "current"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ContinueToLocationTargetCallFrames) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ContinueToLocationTargetCallFrames) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ContinueToLocationTargetCallFrames) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ContinueToLocationTargetCallFrames(in.String()) {
	case ContinueToLocationTargetCallFramesAny:
		*t = ContinueToLocationTargetCallFramesAny
	case ContinueToLocationTargetCallFramesCurrent:
		*t = ContinueToLocationTargetCallFramesCurrent

	default:
		in.AddError(errors.New("unknown ContinueToLocationTargetCallFrames value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ContinueToLocationTargetCallFrames) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ExceptionsState pause on exceptions mode.
type ExceptionsState string

// String returns the ExceptionsState as string value.
func (t ExceptionsState) String() string {
	return string(t)
}

// ExceptionsState values.
const (
	ExceptionsStateNone     ExceptionsState = "none"
	ExceptionsStateUncaught ExceptionsState = "uncaught"
	ExceptionsStateAll      ExceptionsState = "all"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ExceptionsState) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ExceptionsState) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ExceptionsState) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ExceptionsState(in.String()) {
	case ExceptionsStateNone:
		*t = ExceptionsStateNone
	case ExceptionsStateUncaught:
		*t = ExceptionsStateUncaught
	case ExceptionsStateAll:
		*t = ExceptionsStateAll

	default:
		in.AddError(errors.New("unknown ExceptionsState value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ExceptionsState) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
