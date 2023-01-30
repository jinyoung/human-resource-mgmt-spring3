package vacation
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type VacationDaysLeftDB struct{
	db *gorm.DB
}

var vacationDaysLeftrepository *VacationDaysLeftDB

func VacationDaysLeftDBInit() {
	var err error
	vacationDaysLeftrepository = &VacationDaysLeftDB{}
	vacationDaysLeftrepository.db, err = gorm.Open(sqlite.Open("VacationDaysLeft_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	vacationDaysLeftrepository.db.AutoMigrate(&VacationDaysLeft{})

}

func VacationDaysLeftRepository() *VacationDaysLeftDB {
	return vacationDaysLeftrepository
}

func (self *VacationDaysLeftDB)save(entity interface{}) error {
	
	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *VacationDaysLeftDB)GetList() []VacationDaysLeft{
	
	entities := []VacationDaysLeft{}
	self.db.Find(&entities)

	return entities
}

func (self *VacationDaysLeftDB)FindById(id int) (*VacationDaysLeft, error){
	entity := &VacationDaysLeft{}
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

func (self *VacationDaysLeftDB) Delete(entity *VacationDaysLeft) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *VacationDaysLeftDB) Update(id int, params map[string]string) (*VacationDaysLeft, error){
	entity := &VacationDaysLeft{}
	txDb := self.db.Where("id = ?", id)
	if txDb.Error != nil {
		return nil, txDb.Error
	} else {
		txDbRow := txDb.First(entity)
		if txDbRow.Error != nil {
			return nil, txDbRow.Error
		}
		update := &VacationDaysLeft{}
		err := ObjectMapping(update, params)
		if err != nil {
			return nil, err
		}
		self.db.Model(&entity).Updates(update)

		return entity, nil
	}
}