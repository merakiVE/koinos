package db

import (
	"errors"
	"github.com/hostelix/aranGO"
	"encoding/json"
)

type ormStruct struct {
	model aranGO.Modeler
	db      *aranGO.Database
}

func Model(model aranGO.Modeler, db *aranGO.Database) *ormStruct {
	return &ormStruct{
		model: model,
		db:      db,
	}
}

func (this *ormStruct) Create(m aranGO.Modeler) (error) {
	ctx, err := aranGO.NewContext(this.db)
	if err != nil {
		return err
	}
	if e := ctx.Save(m); len(e) >= 1 {
		error_str, _ := json.Marshal(e)
		return errors.New("Error save model in database " + string(error_str))
	}
	return nil
}

func (this *ormStruct) FindOne(out interface{}, filter ...interface{}) (error) {
	aql := aranGO.NewAqlStruct()
	aql.For("v", this.model.GetCollection())
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

func (this *ormStruct) Find(out interface{}, filter ...interface{}) (error) {
	aql := aranGO.NewAqlStruct()
	aql.For("v", this.model.GetCollection())
	aql.Filter(filter...)
	aql.Return("v")

	c, err := aql.Execute(this.db)
	if err != nil {
		return err
	}
	return c.FetchBatch(out)
}
