package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		collection := &models.Collection{
			Name:       "webhooks",
			Type:       models.CollectionTypeBase,
			ListRule:   nil,
			ViewRule:   types.Pointer("@request.auth.id != ''"),
			CreateRule: types.Pointer(""),
			UpdateRule: types.Pointer("@request.auth.id != ''"),
			DeleteRule: nil,
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:     "paymentez_id",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "status",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "status_detail",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "order_description",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "authorization_code",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "date",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "dev_reference",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "amount",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "paid_date",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "stoken",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "application_code",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "pay_method",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "user_id",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "user_email",
					Type:     schema.FieldTypeText,
					Required: true,
				},
			),
		}

		return dao.SaveCollection(collection)

	}, func(db dbx.Builder) error {
		return nil
	})
}
