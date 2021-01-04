package audits

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// AffectedCookie information about a cookie that is affected by an inspector
// issue.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-AffectedCookie
type AffectedCookie struct {
	Name   string `json:"name"` // The following three properties uniquely identify a cookie
	Path   string `json:"path"`
	Domain string `json:"domain"`
}

// AffectedRequest information about a request that is affected by an
// inspector issue.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-AffectedRequest
type AffectedRequest struct {
	RequestID network.RequestID `json:"requestId"` // The unique request id.
	URL       string            `json:"url,omitempty"`
}

// AffectedFrame information about the frame affected by an inspector issue.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-AffectedFrame
type AffectedFrame struct {
	FrameID cdp.FrameID `json:"frameId"`
}

// SameSiteCookieExclusionReason [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-SameSiteCookieExclusionReason
type SameSiteCookieExclusionReason string

// String returns the SameSiteCookieExclusionReason as string value.
func (t SameSiteCookieExclusionReason) String() string {
	return string(t)
}

// SameSiteCookieExclusionReason values.
const (
	SameSiteCookieExclusionReasonExcludeSameSiteUnspecifiedTreatedAsLax SameSiteCookieExclusionReason = "ExcludeSameSiteUnspecifiedTreatedAsLax"
	SameSiteCookieExclusionReasonExcludeSameSiteNoneInsecure            SameSiteCookieExclusionReason = "ExcludeSameSiteNoneInsecure"
	SameSiteCookieExclusionReasonExcludeSameSiteLax                     SameSiteCookieExclusionReason = "ExcludeSameSiteLax"
	SameSiteCookieExclusionReasonExcludeSameSiteStrict                  SameSiteCookieExclusionReason = "ExcludeSameSiteStrict"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SameSiteCookieExclusionReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SameSiteCookieExclusionReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SameSiteCookieExclusionReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SameSiteCookieExclusionReason(in.String()) {
	case SameSiteCookieExclusionReasonExcludeSameSiteUnspecifiedTreatedAsLax:
		*t = SameSiteCookieExclusionReasonExcludeSameSiteUnspecifiedTreatedAsLax
	case SameSiteCookieExclusionReasonExcludeSameSiteNoneInsecure:
		*t = SameSiteCookieExclusionReasonExcludeSameSiteNoneInsecure
	case SameSiteCookieExclusionReasonExcludeSameSiteLax:
		*t = SameSiteCookieExclusionReasonExcludeSameSiteLax
	case SameSiteCookieExclusionReasonExcludeSameSiteStrict:
		*t = SameSiteCookieExclusionReasonExcludeSameSiteStrict

	default:
		in.AddError(errors.New("unknown SameSiteCookieExclusionReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SameSiteCookieExclusionReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SameSiteCookieWarningReason [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-SameSiteCookieWarningReason
type SameSiteCookieWarningReason string

// String returns the SameSiteCookieWarningReason as string value.
func (t SameSiteCookieWarningReason) String() string {
	return string(t)
}

// SameSiteCookieWarningReason values.
const (
	SameSiteCookieWarningReasonWarnSameSiteUnspecifiedCrossSiteContext SameSiteCookieWarningReason = "WarnSameSiteUnspecifiedCrossSiteContext"
	SameSiteCookieWarningReasonWarnSameSiteNoneInsecure                SameSiteCookieWarningReason = "WarnSameSiteNoneInsecure"
	SameSiteCookieWarningReasonWarnSameSiteUnspecifiedLaxAllowUnsafe   SameSiteCookieWarningReason = "WarnSameSiteUnspecifiedLaxAllowUnsafe"
	SameSiteCookieWarningReasonWarnSameSiteStrictLaxDowngradeStrict    SameSiteCookieWarningReason = "WarnSameSiteStrictLaxDowngradeStrict"
	SameSiteCookieWarningReasonWarnSameSiteStrictCrossDowngradeStrict  SameSiteCookieWarningReason = "WarnSameSiteStrictCrossDowngradeStrict"
	SameSiteCookieWarningReasonWarnSameSiteStrictCrossDowngradeLax     SameSiteCookieWarningReason = "WarnSameSiteStrictCrossDowngradeLax"
	SameSiteCookieWarningReasonWarnSameSiteLaxCrossDowngradeStrict     SameSiteCookieWarningReason = "WarnSameSiteLaxCrossDowngradeStrict"
	SameSiteCookieWarningReasonWarnSameSiteLaxCrossDowngradeLax        SameSiteCookieWarningReason = "WarnSameSiteLaxCrossDowngradeLax"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SameSiteCookieWarningReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SameSiteCookieWarningReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SameSiteCookieWarningReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SameSiteCookieWarningReason(in.String()) {
	case SameSiteCookieWarningReasonWarnSameSiteUnspecifiedCrossSiteContext:
		*t = SameSiteCookieWarningReasonWarnSameSiteUnspecifiedCrossSiteContext
	case SameSiteCookieWarningReasonWarnSameSiteNoneInsecure:
		*t = SameSiteCookieWarningReasonWarnSameSiteNoneInsecure
	case SameSiteCookieWarningReasonWarnSameSiteUnspecifiedLaxAllowUnsafe:
		*t = SameSiteCookieWarningReasonWarnSameSiteUnspecifiedLaxAllowUnsafe
	case SameSiteCookieWarningReasonWarnSameSiteStrictLaxDowngradeStrict:
		*t = SameSiteCookieWarningReasonWarnSameSiteStrictLaxDowngradeStrict
	case SameSiteCookieWarningReasonWarnSameSiteStrictCrossDowngradeStrict:
		*t = SameSiteCookieWarningReasonWarnSameSiteStrictCrossDowngradeStrict
	case SameSiteCookieWarningReasonWarnSameSiteStrictCrossDowngradeLax:
		*t = SameSiteCookieWarningReasonWarnSameSiteStrictCrossDowngradeLax
	case SameSiteCookieWarningReasonWarnSameSiteLaxCrossDowngradeStrict:
		*t = SameSiteCookieWarningReasonWarnSameSiteLaxCrossDowngradeStrict
	case SameSiteCookieWarningReasonWarnSameSiteLaxCrossDowngradeLax:
		*t = SameSiteCookieWarningReasonWarnSameSiteLaxCrossDowngradeLax

	default:
		in.AddError(errors.New("unknown SameSiteCookieWarningReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SameSiteCookieWarningReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SameSiteCookieOperation [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-SameSiteCookieOperation
type SameSiteCookieOperation string

// String returns the SameSiteCookieOperation as string value.
func (t SameSiteCookieOperation) String() string {
	return string(t)
}

// SameSiteCookieOperation values.
const (
	SameSiteCookieOperationSetCookie  SameSiteCookieOperation = "SetCookie"
	SameSiteCookieOperationReadCookie SameSiteCookieOperation = "ReadCookie"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SameSiteCookieOperation) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SameSiteCookieOperation) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SameSiteCookieOperation) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SameSiteCookieOperation(in.String()) {
	case SameSiteCookieOperationSetCookie:
		*t = SameSiteCookieOperationSetCookie
	case SameSiteCookieOperationReadCookie:
		*t = SameSiteCookieOperationReadCookie

	default:
		in.AddError(errors.New("unknown SameSiteCookieOperation value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SameSiteCookieOperation) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SameSiteCookieIssueDetails this information is currently necessary, as the
// front-end has a difficult time finding a specific cookie. With this, we can
// convey specific error information without the cookie.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-SameSiteCookieIssueDetails
type SameSiteCookieIssueDetails struct {
	Cookie                 *AffectedCookie                 `json:"cookie"`
	CookieWarningReasons   []SameSiteCookieWarningReason   `json:"cookieWarningReasons"`
	CookieExclusionReasons []SameSiteCookieExclusionReason `json:"cookieExclusionReasons"`
	Operation              SameSiteCookieOperation         `json:"operation"` // Optionally identifies the site-for-cookies and the cookie url, which may be used by the front-end as additional context.
	SiteForCookies         string                          `json:"siteForCookies,omitempty"`
	CookieURL              string                          `json:"cookieUrl,omitempty"`
	Request                *AffectedRequest                `json:"request,omitempty"`
}

// MixedContentResolutionStatus [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-MixedContentResolutionStatus
type MixedContentResolutionStatus string

// String returns the MixedContentResolutionStatus as string value.
func (t MixedContentResolutionStatus) String() string {
	return string(t)
}

// MixedContentResolutionStatus values.
const (
	MixedContentResolutionStatusMixedContentBlocked               MixedContentResolutionStatus = "MixedContentBlocked"
	MixedContentResolutionStatusMixedContentAutomaticallyUpgraded MixedContentResolutionStatus = "MixedContentAutomaticallyUpgraded"
	MixedContentResolutionStatusMixedContentWarning               MixedContentResolutionStatus = "MixedContentWarning"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t MixedContentResolutionStatus) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t MixedContentResolutionStatus) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *MixedContentResolutionStatus) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch MixedContentResolutionStatus(in.String()) {
	case MixedContentResolutionStatusMixedContentBlocked:
		*t = MixedContentResolutionStatusMixedContentBlocked
	case MixedContentResolutionStatusMixedContentAutomaticallyUpgraded:
		*t = MixedContentResolutionStatusMixedContentAutomaticallyUpgraded
	case MixedContentResolutionStatusMixedContentWarning:
		*t = MixedContentResolutionStatusMixedContentWarning

	default:
		in.AddError(errors.New("unknown MixedContentResolutionStatus value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *MixedContentResolutionStatus) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// MixedContentResourceType [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-MixedContentResourceType
type MixedContentResourceType string

// String returns the MixedContentResourceType as string value.
func (t MixedContentResourceType) String() string {
	return string(t)
}

// MixedContentResourceType values.
const (
	MixedContentResourceTypeAudio          MixedContentResourceType = "Audio"
	MixedContentResourceTypeBeacon         MixedContentResourceType = "Beacon"
	MixedContentResourceTypeCSPReport      MixedContentResourceType = "CSPReport"
	MixedContentResourceTypeDownload       MixedContentResourceType = "Download"
	MixedContentResourceTypeEventSource    MixedContentResourceType = "EventSource"
	MixedContentResourceTypeFavicon        MixedContentResourceType = "Favicon"
	MixedContentResourceTypeFont           MixedContentResourceType = "Font"
	MixedContentResourceTypeForm           MixedContentResourceType = "Form"
	MixedContentResourceTypeFrame          MixedContentResourceType = "Frame"
	MixedContentResourceTypeImage          MixedContentResourceType = "Image"
	MixedContentResourceTypeImport         MixedContentResourceType = "Import"
	MixedContentResourceTypeManifest       MixedContentResourceType = "Manifest"
	MixedContentResourceTypePing           MixedContentResourceType = "Ping"
	MixedContentResourceTypePluginData     MixedContentResourceType = "PluginData"
	MixedContentResourceTypePluginResource MixedContentResourceType = "PluginResource"
	MixedContentResourceTypePrefetch       MixedContentResourceType = "Prefetch"
	MixedContentResourceTypeResource       MixedContentResourceType = "Resource"
	MixedContentResourceTypeScript         MixedContentResourceType = "Script"
	MixedContentResourceTypeServiceWorker  MixedContentResourceType = "ServiceWorker"
	MixedContentResourceTypeSharedWorker   MixedContentResourceType = "SharedWorker"
	MixedContentResourceTypeStylesheet     MixedContentResourceType = "Stylesheet"
	MixedContentResourceTypeTrack          MixedContentResourceType = "Track"
	MixedContentResourceTypeVideo          MixedContentResourceType = "Video"
	MixedContentResourceTypeWorker         MixedContentResourceType = "Worker"
	MixedContentResourceTypeXMLHTTPRequest MixedContentResourceType = "XMLHttpRequest"
	MixedContentResourceTypeXSLT           MixedContentResourceType = "XSLT"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t MixedContentResourceType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t MixedContentResourceType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *MixedContentResourceType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch MixedContentResourceType(in.String()) {
	case MixedContentResourceTypeAudio:
		*t = MixedContentResourceTypeAudio
	case MixedContentResourceTypeBeacon:
		*t = MixedContentResourceTypeBeacon
	case MixedContentResourceTypeCSPReport:
		*t = MixedContentResourceTypeCSPReport
	case MixedContentResourceTypeDownload:
		*t = MixedContentResourceTypeDownload
	case MixedContentResourceTypeEventSource:
		*t = MixedContentResourceTypeEventSource
	case MixedContentResourceTypeFavicon:
		*t = MixedContentResourceTypeFavicon
	case MixedContentResourceTypeFont:
		*t = MixedContentResourceTypeFont
	case MixedContentResourceTypeForm:
		*t = MixedContentResourceTypeForm
	case MixedContentResourceTypeFrame:
		*t = MixedContentResourceTypeFrame
	case MixedContentResourceTypeImage:
		*t = MixedContentResourceTypeImage
	case MixedContentResourceTypeImport:
		*t = MixedContentResourceTypeImport
	case MixedContentResourceTypeManifest:
		*t = MixedContentResourceTypeManifest
	case MixedContentResourceTypePing:
		*t = MixedContentResourceTypePing
	case MixedContentResourceTypePluginData:
		*t = MixedContentResourceTypePluginData
	case MixedContentResourceTypePluginResource:
		*t = MixedContentResourceTypePluginResource
	case MixedContentResourceTypePrefetch:
		*t = MixedContentResourceTypePrefetch
	case MixedContentResourceTypeResource:
		*t = MixedContentResourceTypeResource
	case MixedContentResourceTypeScript:
		*t = MixedContentResourceTypeScript
	case MixedContentResourceTypeServiceWorker:
		*t = MixedContentResourceTypeServiceWorker
	case MixedContentResourceTypeSharedWorker:
		*t = MixedContentResourceTypeSharedWorker
	case MixedContentResourceTypeStylesheet:
		*t = MixedContentResourceTypeStylesheet
	case MixedContentResourceTypeTrack:
		*t = MixedContentResourceTypeTrack
	case MixedContentResourceTypeVideo:
		*t = MixedContentResourceTypeVideo
	case MixedContentResourceTypeWorker:
		*t = MixedContentResourceTypeWorker
	case MixedContentResourceTypeXMLHTTPRequest:
		*t = MixedContentResourceTypeXMLHTTPRequest
	case MixedContentResourceTypeXSLT:
		*t = MixedContentResourceTypeXSLT

	default:
		in.AddError(errors.New("unknown MixedContentResourceType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *MixedContentResourceType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// MixedContentIssueDetails [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-MixedContentIssueDetails
type MixedContentIssueDetails struct {
	ResourceType     MixedContentResourceType     `json:"resourceType,omitempty"` // The type of resource causing the mixed content issue (css, js, iframe, form,...). Marked as optional because it is mapped to from blink::mojom::RequestContextType, which will be replaced by network::mojom::RequestDestination
	ResolutionStatus MixedContentResolutionStatus `json:"resolutionStatus"`       // The way the mixed content issue is being resolved.
	InsecureURL      string                       `json:"insecureURL"`            // The unsafe http url causing the mixed content issue.
	MainResourceURL  string                       `json:"mainResourceURL"`        // The url responsible for the call to an unsafe url.
	Request          *AffectedRequest             `json:"request,omitempty"`      // The mixed content request. Does not always exist (e.g. for unsafe form submission urls).
	Frame            *AffectedFrame               `json:"frame,omitempty"`        // Optional because not every mixed content issue is necessarily linked to a frame.
}

// BlockedByResponseReason enum indicating the reason a response has been
// blocked. These reasons are refinements of the net error BLOCKED_BY_RESPONSE.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-BlockedByResponseReason
type BlockedByResponseReason string

// String returns the BlockedByResponseReason as string value.
func (t BlockedByResponseReason) String() string {
	return string(t)
}

// BlockedByResponseReason values.
const (
	BlockedByResponseReasonCoepFrameResourceNeedsCoepHeader                  BlockedByResponseReason = "CoepFrameResourceNeedsCoepHeader"
	BlockedByResponseReasonCoopSandboxedIFrameCannotNavigateToCoopPage       BlockedByResponseReason = "CoopSandboxedIFrameCannotNavigateToCoopPage"
	BlockedByResponseReasonCorpNotSameOrigin                                 BlockedByResponseReason = "CorpNotSameOrigin"
	BlockedByResponseReasonCorpNotSameOriginAfterDefaultedToSameOriginByCoep BlockedByResponseReason = "CorpNotSameOriginAfterDefaultedToSameOriginByCoep"
	BlockedByResponseReasonCorpNotSameSite                                   BlockedByResponseReason = "CorpNotSameSite"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t BlockedByResponseReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t BlockedByResponseReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *BlockedByResponseReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch BlockedByResponseReason(in.String()) {
	case BlockedByResponseReasonCoepFrameResourceNeedsCoepHeader:
		*t = BlockedByResponseReasonCoepFrameResourceNeedsCoepHeader
	case BlockedByResponseReasonCoopSandboxedIFrameCannotNavigateToCoopPage:
		*t = BlockedByResponseReasonCoopSandboxedIFrameCannotNavigateToCoopPage
	case BlockedByResponseReasonCorpNotSameOrigin:
		*t = BlockedByResponseReasonCorpNotSameOrigin
	case BlockedByResponseReasonCorpNotSameOriginAfterDefaultedToSameOriginByCoep:
		*t = BlockedByResponseReasonCorpNotSameOriginAfterDefaultedToSameOriginByCoep
	case BlockedByResponseReasonCorpNotSameSite:
		*t = BlockedByResponseReasonCorpNotSameSite

	default:
		in.AddError(errors.New("unknown BlockedByResponseReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *BlockedByResponseReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// BlockedByResponseIssueDetails details for a request that has been blocked
// with the BLOCKED_BY_RESPONSE code. Currently only used for COEP/COOP, but may
// be extended to include some CSP errors in the future.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-BlockedByResponseIssueDetails
type BlockedByResponseIssueDetails struct {
	Request      *AffectedRequest        `json:"request"`
	ParentFrame  *AffectedFrame          `json:"parentFrame,omitempty"`
	BlockedFrame *AffectedFrame          `json:"blockedFrame,omitempty"`
	Reason       BlockedByResponseReason `json:"reason"`
}

// HeavyAdResolutionStatus [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-HeavyAdResolutionStatus
type HeavyAdResolutionStatus string

// String returns the HeavyAdResolutionStatus as string value.
func (t HeavyAdResolutionStatus) String() string {
	return string(t)
}

// HeavyAdResolutionStatus values.
const (
	HeavyAdResolutionStatusHeavyAdBlocked HeavyAdResolutionStatus = "HeavyAdBlocked"
	HeavyAdResolutionStatusHeavyAdWarning HeavyAdResolutionStatus = "HeavyAdWarning"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t HeavyAdResolutionStatus) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t HeavyAdResolutionStatus) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *HeavyAdResolutionStatus) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch HeavyAdResolutionStatus(in.String()) {
	case HeavyAdResolutionStatusHeavyAdBlocked:
		*t = HeavyAdResolutionStatusHeavyAdBlocked
	case HeavyAdResolutionStatusHeavyAdWarning:
		*t = HeavyAdResolutionStatusHeavyAdWarning

	default:
		in.AddError(errors.New("unknown HeavyAdResolutionStatus value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *HeavyAdResolutionStatus) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// HeavyAdReason [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-HeavyAdReason
type HeavyAdReason string

// String returns the HeavyAdReason as string value.
func (t HeavyAdReason) String() string {
	return string(t)
}

// HeavyAdReason values.
const (
	HeavyAdReasonNetworkTotalLimit HeavyAdReason = "NetworkTotalLimit"
	HeavyAdReasonCPUTotalLimit     HeavyAdReason = "CpuTotalLimit"
	HeavyAdReasonCPUPeakLimit      HeavyAdReason = "CpuPeakLimit"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t HeavyAdReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t HeavyAdReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *HeavyAdReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch HeavyAdReason(in.String()) {
	case HeavyAdReasonNetworkTotalLimit:
		*t = HeavyAdReasonNetworkTotalLimit
	case HeavyAdReasonCPUTotalLimit:
		*t = HeavyAdReasonCPUTotalLimit
	case HeavyAdReasonCPUPeakLimit:
		*t = HeavyAdReasonCPUPeakLimit

	default:
		in.AddError(errors.New("unknown HeavyAdReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *HeavyAdReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// HeavyAdIssueDetails [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-HeavyAdIssueDetails
type HeavyAdIssueDetails struct {
	Resolution HeavyAdResolutionStatus `json:"resolution"` // The resolution status, either blocking the content or warning.
	Reason     HeavyAdReason           `json:"reason"`     // The reason the ad was blocked, total network or cpu or peak cpu.
	Frame      *AffectedFrame          `json:"frame"`      // The frame that was blocked.
}

// ContentSecurityPolicyViolationType [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-ContentSecurityPolicyViolationType
type ContentSecurityPolicyViolationType string

// String returns the ContentSecurityPolicyViolationType as string value.
func (t ContentSecurityPolicyViolationType) String() string {
	return string(t)
}

// ContentSecurityPolicyViolationType values.
const (
	ContentSecurityPolicyViolationTypeKInlineViolation             ContentSecurityPolicyViolationType = "kInlineViolation"
	ContentSecurityPolicyViolationTypeKEvalViolation               ContentSecurityPolicyViolationType = "kEvalViolation"
	ContentSecurityPolicyViolationTypeKURLViolation                ContentSecurityPolicyViolationType = "kURLViolation"
	ContentSecurityPolicyViolationTypeKTrustedTypesSinkViolation   ContentSecurityPolicyViolationType = "kTrustedTypesSinkViolation"
	ContentSecurityPolicyViolationTypeKTrustedTypesPolicyViolation ContentSecurityPolicyViolationType = "kTrustedTypesPolicyViolation"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ContentSecurityPolicyViolationType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ContentSecurityPolicyViolationType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ContentSecurityPolicyViolationType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ContentSecurityPolicyViolationType(in.String()) {
	case ContentSecurityPolicyViolationTypeKInlineViolation:
		*t = ContentSecurityPolicyViolationTypeKInlineViolation
	case ContentSecurityPolicyViolationTypeKEvalViolation:
		*t = ContentSecurityPolicyViolationTypeKEvalViolation
	case ContentSecurityPolicyViolationTypeKURLViolation:
		*t = ContentSecurityPolicyViolationTypeKURLViolation
	case ContentSecurityPolicyViolationTypeKTrustedTypesSinkViolation:
		*t = ContentSecurityPolicyViolationTypeKTrustedTypesSinkViolation
	case ContentSecurityPolicyViolationTypeKTrustedTypesPolicyViolation:
		*t = ContentSecurityPolicyViolationTypeKTrustedTypesPolicyViolation

	default:
		in.AddError(errors.New("unknown ContentSecurityPolicyViolationType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ContentSecurityPolicyViolationType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SourceCodeLocation [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-SourceCodeLocation
type SourceCodeLocation struct {
	URL          string `json:"url"`
	LineNumber   int64  `json:"lineNumber"`
	ColumnNumber int64  `json:"columnNumber"`
}

// ContentSecurityPolicyIssueDetails [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-ContentSecurityPolicyIssueDetails
type ContentSecurityPolicyIssueDetails struct {
	BlockedURL                         string                             `json:"blockedURL,omitempty"` // The url not included in allowed sources.
	ViolatedDirective                  string                             `json:"violatedDirective"`    // Specific directive that is violated, causing the CSP issue.
	IsReportOnly                       bool                               `json:"isReportOnly"`
	ContentSecurityPolicyViolationType ContentSecurityPolicyViolationType `json:"contentSecurityPolicyViolationType"`
	FrameAncestor                      *AffectedFrame                     `json:"frameAncestor,omitempty"`
	SourceCodeLocation                 *SourceCodeLocation                `json:"sourceCodeLocation,omitempty"`
	ViolatingNodeID                    cdp.BackendNodeID                  `json:"violatingNodeId,omitempty"`
}

// SharedArrayBufferTransferIssueDetails details for a request that has been
// blocked with the BLOCKED_BY_RESPONSE code. Currently only used for COEP/COOP,
// but may be extended to include some CSP errors in the future.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-SharedArrayBufferTransferIssueDetails
type SharedArrayBufferTransferIssueDetails struct {
	SourceCodeLocation *SourceCodeLocation `json:"sourceCodeLocation"`
	IsWarning          bool                `json:"isWarning"`
}

// InspectorIssueCode a unique identifier for the type of issue. Each type
// may use one of the optional fields in InspectorIssueDetails to convey more
// specific information about the kind of issue.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-InspectorIssueCode
type InspectorIssueCode string

// String returns the InspectorIssueCode as string value.
func (t InspectorIssueCode) String() string {
	return string(t)
}

// InspectorIssueCode values.
const (
	InspectorIssueCodeSameSiteCookieIssue            InspectorIssueCode = "SameSiteCookieIssue"
	InspectorIssueCodeMixedContentIssue              InspectorIssueCode = "MixedContentIssue"
	InspectorIssueCodeBlockedByResponseIssue         InspectorIssueCode = "BlockedByResponseIssue"
	InspectorIssueCodeHeavyAdIssue                   InspectorIssueCode = "HeavyAdIssue"
	InspectorIssueCodeContentSecurityPolicyIssue     InspectorIssueCode = "ContentSecurityPolicyIssue"
	InspectorIssueCodeSharedArrayBufferTransferIssue InspectorIssueCode = "SharedArrayBufferTransferIssue"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t InspectorIssueCode) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t InspectorIssueCode) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *InspectorIssueCode) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch InspectorIssueCode(in.String()) {
	case InspectorIssueCodeSameSiteCookieIssue:
		*t = InspectorIssueCodeSameSiteCookieIssue
	case InspectorIssueCodeMixedContentIssue:
		*t = InspectorIssueCodeMixedContentIssue
	case InspectorIssueCodeBlockedByResponseIssue:
		*t = InspectorIssueCodeBlockedByResponseIssue
	case InspectorIssueCodeHeavyAdIssue:
		*t = InspectorIssueCodeHeavyAdIssue
	case InspectorIssueCodeContentSecurityPolicyIssue:
		*t = InspectorIssueCodeContentSecurityPolicyIssue
	case InspectorIssueCodeSharedArrayBufferTransferIssue:
		*t = InspectorIssueCodeSharedArrayBufferTransferIssue

	default:
		in.AddError(errors.New("unknown InspectorIssueCode value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *InspectorIssueCode) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// InspectorIssueDetails this struct holds a list of optional fields with
// additional information specific to the kind of issue. When adding a new issue
// code, please also add a new optional field to this type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-InspectorIssueDetails
type InspectorIssueDetails struct {
	SameSiteCookieIssueDetails            *SameSiteCookieIssueDetails            `json:"sameSiteCookieIssueDetails,omitempty"`
	MixedContentIssueDetails              *MixedContentIssueDetails              `json:"mixedContentIssueDetails,omitempty"`
	BlockedByResponseIssueDetails         *BlockedByResponseIssueDetails         `json:"blockedByResponseIssueDetails,omitempty"`
	HeavyAdIssueDetails                   *HeavyAdIssueDetails                   `json:"heavyAdIssueDetails,omitempty"`
	ContentSecurityPolicyIssueDetails     *ContentSecurityPolicyIssueDetails     `json:"contentSecurityPolicyIssueDetails,omitempty"`
	SharedArrayBufferTransferIssueDetails *SharedArrayBufferTransferIssueDetails `json:"sharedArrayBufferTransferIssueDetails,omitempty"`
}

// InspectorIssue an inspector issue reported from the back-end.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-InspectorIssue
type InspectorIssue struct {
	Code    InspectorIssueCode     `json:"code"`
	Details *InspectorIssueDetails `json:"details"`
}

// GetEncodedResponseEncoding the encoding to use.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#method-getEncodedResponse
type GetEncodedResponseEncoding string

// String returns the GetEncodedResponseEncoding as string value.
func (t GetEncodedResponseEncoding) String() string {
	return string(t)
}

// GetEncodedResponseEncoding values.
const (
	GetEncodedResponseEncodingWebp GetEncodedResponseEncoding = "webp"
	GetEncodedResponseEncodingJpeg GetEncodedResponseEncoding = "jpeg"
	GetEncodedResponseEncodingPng  GetEncodedResponseEncoding = "png"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t GetEncodedResponseEncoding) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t GetEncodedResponseEncoding) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *GetEncodedResponseEncoding) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch GetEncodedResponseEncoding(in.String()) {
	case GetEncodedResponseEncodingWebp:
		*t = GetEncodedResponseEncodingWebp
	case GetEncodedResponseEncodingJpeg:
		*t = GetEncodedResponseEncodingJpeg
	case GetEncodedResponseEncodingPng:
		*t = GetEncodedResponseEncodingPng

	default:
		in.AddError(errors.New("unknown GetEncodedResponseEncoding value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *GetEncodedResponseEncoding) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
