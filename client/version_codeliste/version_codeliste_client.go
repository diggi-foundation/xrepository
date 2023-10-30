// Code generated by go-swagger; DO NOT EDIT.

package version_codeliste

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new version codeliste API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for version codeliste API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Abbestellen1(params *Abbestellen1Params, opts ...ClientOption) error

	Abonnieren1(params *Abonnieren1Params, opts ...ClientOption) error

	AddVersionCodeliste(params *AddVersionCodelisteParams, opts ...ClientOption) error

	DownloadGenericodeFachmodell(params *DownloadGenericodeFachmodellParams, opts ...ClientOption) error

	ExportFormat(params *ExportFormatParams, opts ...ClientOption) error

	GetAbonnenten1(params *GetAbonnenten1Params, opts ...ClientOption) error

	GetAsExcel(params *GetAsExcelParams, opts ...ClientOption) error

	GetAsJSON(params *GetAsJSONParams, opts ...ClientOption) error

	GetAsMagicDraw(params *GetAsMagicDrawParams, opts ...ClientOption) error

	GetAsMarkdown(params *GetAsMarkdownParams, opts ...ClientOption) error

	GetBeziehungen1(params *GetBeziehungen1Params, opts ...ClientOption) (*GetBeziehungen1OK, error)

	GetDebugInfo1(params *GetDebugInfo1Params, opts ...ClientOption) (*GetDebugInfo1OK, error)

	GetDokument(params *GetDokumentParams, opts ...ClientOption) error

	GetDokumente1(params *GetDokumente1Params, opts ...ClientOption) error

	GetGenericode(params *GetGenericodeParams, opts ...ClientOption) error

	GetGenericodeDaten(params *GetGenericodeDatenParams, opts ...ClientOption) error

	GetGenericodeDaten1(params *GetGenericodeDaten1Params, opts ...ClientOption) error

	GetGenericode1(params *GetGenericode1Params, opts ...ClientOption) error

	GetJSON1(params *GetJSON1Params, opts ...ClientOption) error

	GetModelPlain1(params *GetModelPlain1Params, opts ...ClientOption) error

	GetModel1(params *GetModel1Params, opts ...ClientOption) error

	ImportGenericode(params *ImportGenericodeParams, opts ...ClientOption) error

	ImportGenericode1(params *ImportGenericode1Params, opts ...ClientOption) error

	Metadaten1(params *Metadaten1Params, opts ...ClientOption) error

	RemoveDokumentFromVersionCodeliste(params *RemoveDokumentFromVersionCodelisteParams, opts ...ClientOption) error

	Save(params *SaveParams, opts ...ClientOption) error

	Save1(params *Save1Params, opts ...ClientOption) error

	TransferOwnership1(params *TransferOwnership1Params, opts ...ClientOption) error

	UpdateStatus1(params *UpdateStatus1Params, opts ...ClientOption) error

	Update1(params *Update1Params, opts ...ClientOption) error

	UploadDokument1(params *UploadDokument1Params, opts ...ClientOption) error

	Validate1(params *Validate1Params, opts ...ClientOption) error

	Validate2(params *Validate2Params, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
Abbestellen1 abbestellen 1 API
*/
func (a *Client) Abbestellen1(params *Abbestellen1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAbbestellen1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "abbestellen_1",
		Method:             "POST",
		PathPattern:        "/version_codeliste/{urn}/abbestellen",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Abbestellen1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Abonnieren1 abonnieren 1 API
*/
func (a *Client) Abonnieren1(params *Abonnieren1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAbonnieren1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "abonnieren_1",
		Method:             "POST",
		PathPattern:        "/version_codeliste/{urn}/abonnieren",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Abonnieren1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
AddVersionCodeliste add version codeliste API
*/
func (a *Client) AddVersionCodeliste(params *AddVersionCodelisteParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddVersionCodelisteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addVersionCodeliste",
		Method:             "POST",
		PathPattern:        "/version_codeliste/hochladen",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddVersionCodelisteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
DownloadGenericodeFachmodell lädts das genericode fachmodell herunter
*/
func (a *Client) DownloadGenericodeFachmodell(params *DownloadGenericodeFachmodellParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadGenericodeFachmodellParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadGenericodeFachmodell",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/technischerBestandteilGenericode",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DownloadGenericodeFachmodellReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
ExportFormat stellts ein exportformat der codeliste bereit
*/
func (a *Client) ExportFormat(params *ExportFormatParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportFormatParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportFormat",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/{format}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportFormatReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetAbonnenten1 get abonnenten 1 API
*/
func (a *Client) GetAbonnenten1(params *GetAbonnenten1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAbonnenten1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAbonnenten_1",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{urn}/abonnenten",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAbonnenten1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetAsExcel stellts das exportformat excel der codeliste bereit
*/
func (a *Client) GetAsExcel(params *GetAsExcelParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAsExcelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAsExcel",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/download/{name}.xlsx",
		ProducesMediaTypes: []string{"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAsExcelReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetAsJSON stellts das exportformat JSON der codeliste bereit
*/
func (a *Client) GetAsJSON(params *GetAsJSONParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAsJSONParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAsJson",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/download/{name}.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAsJSONReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetAsMagicDraw stellts das exportformat magic draw der codeliste bereit
*/
func (a *Client) GetAsMagicDraw(params *GetAsMagicDrawParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAsMagicDrawParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAsMagicDraw",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/download/{name}.mdxml",
		ProducesMediaTypes: []string{"application/x-uml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAsMagicDrawReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetAsMarkdown stellts das exportformat markdown der codeliste bereit
*/
func (a *Client) GetAsMarkdown(params *GetAsMarkdownParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAsMarkdownParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAsMarkdown",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/download/{name}.md",
		ProducesMediaTypes: []string{"text/markdown"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAsMarkdownReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetBeziehungen1 get beziehungen 1 API
*/
func (a *Client) GetBeziehungen1(params *GetBeziehungen1Params, opts ...ClientOption) (*GetBeziehungen1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBeziehungen1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBeziehungen_1",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{urn}/beziehungen",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBeziehungen1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBeziehungen1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBeziehungen_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDebugInfo1 get debug info 1 API
*/
func (a *Client) GetDebugInfo1(params *GetDebugInfo1Params, opts ...ClientOption) (*GetDebugInfo1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDebugInfo1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDebugInfo_1",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{urn}/debugInfo",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDebugInfo1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDebugInfo1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDebugInfo_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDokument lädts ein dokument herunter
*/
func (a *Client) GetDokument(params *GetDokumentParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDokumentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDokument",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/dokument/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDokumentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetDokumente1 lädts alle dokumente des objekts zu der kennung als z IP herunter
*/
func (a *Client) GetDokumente1(params *GetDokumente1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDokumente1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDokumente_1",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/dokumente",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDokumente1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetGenericode lieferts die genericode datei
*/
func (a *Client) GetGenericode(params *GetGenericodeParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGenericodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGenericode",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/genericode",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGenericodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetGenericodeDaten lieferts die genericode daten zur darstellung in der g UI tabelle
*/
func (a *Client) GetGenericodeDaten(params *GetGenericodeDatenParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGenericodeDatenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGenericodeDaten",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/genericode/daten",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGenericodeDatenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetGenericodeDaten1 lieferts die genericode daten zur darstellung in der g UI tabelle
*/
func (a *Client) GetGenericodeDaten1(params *GetGenericodeDaten1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGenericodeDaten1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGenericodeDaten_1",
		Method:             "GET",
		PathPattern:        "/version_codeliste/genericode/daten",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGenericodeDaten1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetGenericode1 lieferts die genericode datei
*/
func (a *Client) GetGenericode1(params *GetGenericode1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGenericode1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGenericode_1",
		Method:             "GET",
		PathPattern:        "/version_codeliste/genericode",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGenericode1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetJSON1 get Json 1 API
*/
func (a *Client) GetJSON1(params *GetJSON1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJSON1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getJson_1",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{urn}/model",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetJSON1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetModelPlain1 herunterladens des modells daten dieses objekts
*/
func (a *Client) GetModelPlain1(params *GetModelPlain1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetModelPlain1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getModelPlain_1",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetModelPlain1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetModel1 get model 1 API
*/
func (a *Client) GetModel1(params *GetModel1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetModel1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getModel_1",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{kennung}/model",
		ProducesMediaTypes: []string{"application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetModel1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
ImportGenericode lädts genericode zur anzeige im genericode editor
*/
func (a *Client) ImportGenericode(params *ImportGenericodeParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportGenericodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "importGenericode",
		Method:             "POST",
		PathPattern:        "/version_codeliste/{kennung}/genericode/load",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImportGenericodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
ImportGenericode1 lädts genericode zur anzeige im genericode editor
*/
func (a *Client) ImportGenericode1(params *ImportGenericode1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportGenericode1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "importGenericode_1",
		Method:             "POST",
		PathPattern:        "/version_codeliste/genericode/load",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImportGenericode1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Metadaten1 metadaten 1 API
*/
func (a *Client) Metadaten1(params *Metadaten1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMetadaten1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "metadaten_1",
		Method:             "GET",
		PathPattern:        "/version_codeliste/{urn}/metadaten",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Metadaten1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
RemoveDokumentFromVersionCodeliste löschens eines dokuments
*/
func (a *Client) RemoveDokumentFromVersionCodeliste(params *RemoveDokumentFromVersionCodelisteParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveDokumentFromVersionCodelisteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "removeDokumentFromVersionCodeliste",
		Method:             "DELETE",
		PathPattern:        "/version_codeliste/{kennung}/dokument/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveDokumentFromVersionCodelisteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Save speicherts das g c editor ergebnis in dem ein extraktion auf die erfassten daten durchgeführt wird
*/
func (a *Client) Save(params *SaveParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "save",
		Method:             "POST",
		PathPattern:        "/version_codeliste/{kennung}/genericode/daten",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SaveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Save1 speicherts das g c editor ergebnis in dem ein extraktion auf die erfassten daten durchgeführt wird
*/
func (a *Client) Save1(params *Save1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSave1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "save_1",
		Method:             "POST",
		PathPattern:        "/version_codeliste/genericode/daten",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Save1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
TransferOwnership1 transfer ownership 1 API
*/
func (a *Client) TransferOwnership1(params *TransferOwnership1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransferOwnership1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "transferOwnership_1",
		Method:             "POST",
		PathPattern:        "/version_codeliste/{urn}/transfer-owner",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TransferOwnership1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
UpdateStatus1 update status 1 API
*/
func (a *Client) UpdateStatus1(params *UpdateStatus1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStatus1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateStatus_1",
		Method:             "POST",
		PathPattern:        "/version_codeliste/{urn}/updateStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateStatus1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Update1 update 1 API
*/
func (a *Client) Update1(params *Update1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdate1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "update_1",
		Method:             "POST",
		PathPattern:        "/version_codeliste/{urn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Update1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
UploadDokument1 uploads eines neuen dokuments
*/
func (a *Client) UploadDokument1(params *UploadDokument1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadDokument1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadDokument_1",
		Method:             "POST",
		PathPattern:        "/version_codeliste/{urn}/add-document",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadDokument1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Validate1 validierts eine genericode datei
*/
func (a *Client) Validate1(params *Validate1Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidate1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "validate_1",
		Method:             "POST",
		PathPattern:        "/version_codeliste/{kennung}/genericode/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Validate1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
Validate2 validierts eine genericode datei
*/
func (a *Client) Validate2(params *Validate2Params, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidate2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "validate_2",
		Method:             "POST",
		PathPattern:        "/version_codeliste/genericode/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Validate2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}