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
            Name:       "duplicate_webhooks",
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
                    Name:     "data",
                    Type:     schema.FieldTypeJson,
                    Required: true,
                },

            ),
        }

        return dao.SaveCollection(collection)

    }, func(db dbx.Builder) error {
        return nil
    })
}
