package vacation
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type VacationDaysStatusDB struct{
	db *gorm.DB
}

var vacationDaysStatusrepository *VacationDaysStatusDB

func VacationDaysStatusDBInit() {
	var err error
	vacationDaysStatusrepository = &VacationDaysStatusDB{}
	vacationDaysStatusrepository.db, err = gorm.Open(sqlite.Open("VacationDaysStatus_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	vacationDaysStatusrepository.db.AutoMigrate(&VacationDaysStatus{})

}

func VacationDaysStatusRepository() *VacationDaysStatusDB {
	return vacationDaysStatusrepository
}


func (self *VacationDaysStatusDB)save(entity interface{}) error {

	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *VacationDaysStatusDB)GetList() []VacationDaysStatus{

	entities := []VacationDaysStatus{}
	self.db.Find(&entities)

	return entities
}

func (self *VacationDaysStatusDB)FindById(id int) (*VacationDaysStatus, error){
	entity := &VacationDaysStatus{}
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

func (self *VacationDaysStatusDB) Delete(entity *VacationDaysStatus) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *VacationDaysStatusDB) Update(id int, params map[string]string) (*VacationDaysStatus, error){
	entity := &VacationDaysStatus{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		update := &VacationDaysStatus{}
		err := ObjectMapping(update, params)
		if err != nil {
			return nil, err
		}
		self.db.Model(&entity).Updates(update)

		return entity, nil
	}
}