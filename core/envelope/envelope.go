package envelope

import (
	"context"
	"io/fs"

	"github.com/ItsNotGoodName/smtpbridge/core"
	"github.com/ItsNotGoodName/smtpbridge/core/paginate"
)

type (
	Envelope struct {
		Message     Message      `json:"message"`
		Attachments []Attachment `json:"attachments"`
	}

	CreateEnvelopeRequest struct {
		From       string
		To         []string
		Subject    string
		Text       string
		HTML       string
		Attachment []CreateAttachmentRequest
	}

	CreateAttachmentRequest struct {
		Name string
		Data []byte
	}

	Service interface {
		ListEnvelope(ctx context.Context, page *paginate.Page) ([]Envelope, error)
		GetEnvelope(ctx context.Context, msgID int64) (*Envelope, error)
		CreateEnvelope(ctx context.Context, req *CreateEnvelopeRequest) (int64, error)
		DeleteEnvelope(ctx context.Context, msgID int64) error
	}

	Store interface {
		ListEnvelope(ctx context.Context, offset, limit int, ascending bool) ([]Envelope, int, error)
		CreateEnvelope(ctx context.Context, msg *Message, atts []Attachment) (int64, error)
		GetEnvelope(ctx context.Context, msgID int64) (*Envelope, error)
		GetAndDeleteEnvelope(ctx context.Context, msgID int64) (*Envelope, error)
	}

	DataStore interface {
		CreateData(ctx context.Context, att *Attachment, data []byte) error
		GetData(ctx context.Context, att *Attachment) ([]byte, error)
		DeleteData(ctx context.Context, att *Attachment) error
	}

	LocalDataStore interface {
		DataFS() (fs.FS, error)
	}
)

type EnvelopeService struct {
	store     Store
	dataStore DataStore
}

func NewEnvelopeService(store Store, dataStore DataStore) *EnvelopeService {
	return &EnvelopeService{
		store:     store,
		dataStore: dataStore,
	}
}

func (es *EnvelopeService) ListEnvelope(ctx context.Context, page *paginate.Page) ([]Envelope, error) {
	envs, count, err := es.store.ListEnvelope(ctx, page.Offset(), page.Limit, page.Ascending)
	if err != nil {
		return nil, err
	}

	page.SetCount(count)

	return envs, nil
}

func (es *EnvelopeService) CreateEnvelope(ctx context.Context, req *CreateEnvelopeRequest) (int64, error) {
	// Create message and attachments
	msg := NewMessage(req.From, req.To, req.Subject, req.Text, req.HTML)
	atts := make([]Attachment, 0, len(req.Attachment))
	for _, attReq := range req.Attachment {
		atts = append(atts, *NewAttachment(msg.ID, attReq.Name, attReq.Data))
	}

	// Save envelope
	msgID, err := es.store.CreateEnvelope(ctx, msg, atts)
	if err != nil {
		return 0, nil
	}

	// Save attachments' data
	for i, att := range atts {
		if err := es.dataStore.CreateData(ctx, &att, req.Attachment[i].Data); err != nil {
			return 0, err
		}
	}

	return msgID, nil
}

func (es *EnvelopeService) GetEnvelope(ctx context.Context, msgID int64) (*Envelope, error) {
	return es.store.GetEnvelope(ctx, msgID)
}

func (es *EnvelopeService) DeleteEnvelope(ctx context.Context, msgID int64) error {
	// Delete envelope
	env, err := es.store.GetAndDeleteEnvelope(ctx, msgID)
	if err != nil {
		return err
	}

	// Delete attachments' data
	for _, att := range env.Attachments {
		if err := es.dataStore.DeleteData(ctx, &att); err != nil && err != core.ErrDataNotFound {
			return err
		}
	}

	return nil
}