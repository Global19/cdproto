package layertree

// AUTOGENERATED. DO NOT EDIT.

import (
	"errors"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/cdp/dom"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// LayerID unique Layer identifier.
type LayerID string

// String returns the LayerID as string value.
func (t LayerID) String() string {
	return string(t)
}

// SnapshotID unique snapshot identifier.
type SnapshotID string

// String returns the SnapshotID as string value.
func (t SnapshotID) String() string {
	return string(t)
}

// ScrollRect rectangle where scrolling happens on the main thread.
type ScrollRect struct {
	Rect *dom.Rect      `json:"rect"` // Rectangle itself.
	Type ScrollRectType `json:"type"` // Reason for rectangle to force scrolling on the main thread
}

// PictureTile serialized fragment of layer picture along with its offset
// within the layer.
type PictureTile struct {
	X       float64 `json:"x"`       // Offset from owning layer left boundary
	Y       float64 `json:"y"`       // Offset from owning layer top boundary
	Picture string  `json:"picture"` // Base64-encoded snapshot data.
}

// Layer information about a compositing layer.
type Layer struct {
	LayerID       LayerID           `json:"layerId"`                 // The unique id for this layer.
	ParentLayerID LayerID           `json:"parentLayerId,omitempty"` // The id of parent (not present for root).
	BackendNodeID cdp.BackendNodeID `json:"backendNodeId,omitempty"` // The backend id for the node associated with this layer.
	OffsetX       float64           `json:"offsetX"`                 // Offset from parent layer, X coordinate.
	OffsetY       float64           `json:"offsetY"`                 // Offset from parent layer, Y coordinate.
	Width         float64           `json:"width"`                   // Layer width.
	Height        float64           `json:"height"`                  // Layer height.
	Transform     []float64         `json:"transform,omitempty"`     // Transformation matrix for layer, default is identity matrix
	AnchorX       float64           `json:"anchorX,omitempty"`       // Transform anchor point X, absent if no transform specified
	AnchorY       float64           `json:"anchorY,omitempty"`       // Transform anchor point Y, absent if no transform specified
	AnchorZ       float64           `json:"anchorZ,omitempty"`       // Transform anchor point Z, absent if no transform specified
	PaintCount    int64             `json:"paintCount"`              // Indicates how many time this layer has painted.
	DrawsContent  bool              `json:"drawsContent"`            // Indicates whether this layer hosts any content, rather than being used for transform/scrolling purposes only.
	Invisible     bool              `json:"invisible,omitempty"`     // Set if layer is not visible.
	ScrollRects   []*ScrollRect     `json:"scrollRects,omitempty"`   // Rectangles scrolling on main thread only.
}

// PaintProfile array of timings, one per paint step.
type PaintProfile []float64

// ScrollRectType reason for rectangle to force scrolling on the main thread.
type ScrollRectType string

// String returns the ScrollRectType as string value.
func (t ScrollRectType) String() string {
	return string(t)
}

// ScrollRectType values.
const (
	ScrollRectTypeRepaintsOnScroll  ScrollRectType = "RepaintsOnScroll"
	ScrollRectTypeTouchEventHandler ScrollRectType = "TouchEventHandler"
	ScrollRectTypeWheelEventHandler ScrollRectType = "WheelEventHandler"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ScrollRectType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ScrollRectType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ScrollRectType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ScrollRectType(in.String()) {
	case ScrollRectTypeRepaintsOnScroll:
		*t = ScrollRectTypeRepaintsOnScroll
	case ScrollRectTypeTouchEventHandler:
		*t = ScrollRectTypeTouchEventHandler
	case ScrollRectTypeWheelEventHandler:
		*t = ScrollRectTypeWheelEventHandler

	default:
		in.AddError(errors.New("unknown ScrollRectType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ScrollRectType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
