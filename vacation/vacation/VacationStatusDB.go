package vacation
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type VacationStatusDB struct{
	db *gorm.DB
}

var vacationStatusrepository *VacationStatusDB

func VacationStatusDBInit() {
	var err error
	vacationStatusrepository = &VacationStatusDB{}
	vacationStatusrepository.db, err = gorm.Open(sqlite.Open("VacationStatus_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	vacationStatusrepository.db.AutoMigrate(&VacationStatus{})

}

func VacationStatusRepository() *VacationStatusDB {
	return vacationStatusrepository
}


func (self *VacationStatusDB)save(entity interface{}) error {

	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *VacationStatusDB)GetList() []VacationStatus{

	entities := []VacationStatus{}
	self.db.Find(&entities)

	return entities
}

func (self *VacationStatusDB)FindById(id int) (*VacationStatus, error){
	entity := &VacationStatus{}
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

func (self *VacationStatusDB) Delete(entity *VacationStatus) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *VacationStatusDB) Update(id int, params map[string]string) (*VacationStatus, error){
	entity := &VacationStatus{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		update := &VacationStatus{}
		err := ObjectMapping(update, params)
		if err != nil {
			return nil, err
		}
		self.db.Model(&entity).Updates(update)

		return entity, nil
	}
}