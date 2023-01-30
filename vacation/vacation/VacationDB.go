package vacation
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type VacationDB struct{
	db *gorm.DB
}

var vacationrepository *VacationDB

func VacationDBInit() {
	var err error
	vacationrepository = &VacationDB{}
	vacationrepository.db, err = gorm.Open(sqlite.Open("Vacation_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	vacationrepository.db.AutoMigrate(&Vacation{})

}

func VacationRepository() *VacationDB {
	return vacationrepository
}

func (self *VacationDB)save(entity interface{}) error {
	
	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *VacationDB)GetList() []Vacation{
	
	entities := []Vacation{}
	self.db.Find(&entities)

	return entities
}

func (self *VacationDB)FindById(id int) (*Vacation, error){
	entity := &Vacation{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		return entity, nil
	}
}

func (self *VacationDB) Delete(entity *Vacation) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *VacationDB) Update(id int, params map[string]string) (*Vacation, error){
	entity := &Vacation{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		update := &Vacation{}
		err := ObjectMapping(update, params)
		if err != nil {
			return nil, err
		}
		self.db.Model(&entity).Updates(update)

		return entity, nil
	}
}