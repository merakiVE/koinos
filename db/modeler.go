package db

import (
	"errors"
	"github.com/hostelix/aranGO"
)

type modelerDBStruct struct {
	modeler aranGO.Modeler
	db      *aranGO.Database
}

func Model(model aranGO.Modeler, db *aranGO.Database) *modelerDBStruct {
	return &modelerDBStruct{
		modeler: model,
		db:      db,
	}
}

func (this *modelerDBStruct) Create(m aranGO.Modeler) (error) {
	ctx, err := aranGO.NewContext(this.db)
	if err != nil {
		return err
	}
	if e := ctx.Save(m); len(e) >= 1 {
		return errors.New("Error save model in database")
	}
	return nil
}

func (this *modelerDBStruct) FindOne(out interface{}, filter ...interface{}) (error) {
	aql := aranGO.NewAqlStruct()
	aql.For("v", this.modeler.GetCollection())
	aql.Filter(filter...)
	aql.Return("v")

	c, err := aql.Execute(this.db)
	if err != nil {
		return err
	}

	if ok := c.FetchOne(out); !ok {
		return errors.New("Error get data")
	}
	return nil
}

func (this *modelerDBStruct) Find(out interface{}, filter ...interface{}) (error) {
	aql := aranGO.NewAqlStruct()
	aql.For("v", this.modeler.GetCollection())
	aql.Filter(filter...)
	aql.Return("v")

	c, err := aql.Execute(this.db)
	if err != nil {
		return err
	}
	return c.FetchBatch(out)
}
